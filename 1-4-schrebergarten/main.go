package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"
)

type garden struct {
	X       int // left top
	Y       int // left top
	Width   int
	Height  int
	Rotated bool
}

type field struct {
	width  int
	height int
}

var area [][]int              // resulting area with indeces of garden on x and y
var areaWidth, areaHeight int // width and height of area
var lowestArea [][]int        // result of calculations (multiple results)
var currHeight int            // height, which is currently used
var gardens []garden          // all gardens
var tries int                 // counts calculations

func main() {
	// Command Line Arguments
	var datasetchoice = flag.String("dataset", "", "Select Dataset Path for calculation (beispieldaten/set1.json)")
	flag.Parse()

	if *datasetchoice != "" {
		gardens = parseJSONSampleSets(*datasetchoice)

	} else {
		fmt.Println("You musst select a dataset, select (beispieldaten/set1.json) with --dataset=")
		os.Exit(1)
	}

	startFreshAndClean()

	fmt.Println("calculating...")

	gcdScale := calcGCD(gardens)              // gcd (greatest common divisor) of every side length to get optimal scale value
	gardens = scaleGardens(gardens, gcdScale) // scaling all side lengths by gcd

	areaWidth = getMaxWidth(gardens) // the maximum of possible width --> width and height
	areaHeight = areaWidth

	lowestArea = make([][]int, 1) // [0] = area size, [1] = x, [2] = y, [3] = rotated (1/0)  --> every garden
	lowestArea[0] = make([]int, len(gardens)*3+1)
	lowestArea[0][0] = 4 * areaWidth * areaHeight // should not be the result (just for avoiding errors, cp.: Algo() 2nd-for-loop-condition)

	// calculates lowest possible area size and all possibilities for that
	Algo()
	// does stuff with results
	Results(gcdScale)
	fmt.Println()
	fmt.Println("Finished... you can now see the results in the files field*.png in the current folder")
}

func startFreshAndClean() {
	// We want to start clean so we delete all generated field*.pngs. So theres no confusion
	// We dont handle errors here because its not relevant to the user
	os.Remove("field0.png")
	os.Remove("field1.png")
	os.Remove("field2.png")
	os.Remove("field3.png")
}

func parseJSONSampleSets(filePath string) []garden {
	fmt.Printf("// reading sampleset %s\n", filePath)
	file, err1 := ioutil.ReadFile(filePath)
	if err1 != nil {
		fmt.Printf("// error while reading file %s\n", filePath)
		fmt.Printf("File error: %v\n", err1)
		os.Exit(1)
	}

	var jsongardenset []garden

	err2 := json.Unmarshal(file, &jsongardenset)
	if err2 != nil {
		fmt.Println("error:", err2)
		os.Exit(1)
	}
	return jsongardenset
}

func scaleGardens(gardens []garden, value int) []garden {
	for i := 0; i < len(gardens); i++ {
		gardens[i].Width /= value  // scale Width
		gardens[i].Height /= value // scale height
	}
	return gardens
}

//=========================== Algorithm ================================0

func Algo() {
	/**
	This Algorithm calculates lowest possible area size and all possibilities for that
	**/
	tries = 0                       // nothing is calculated yet
	area = make([][]int, areaWidth) // area width / x
	for i := 0; i < len(area); i++ {
		area[i] = make([]int, areaHeight) // area height / y
	}
	minHeight := calcMinHeight(gardens) // no lower height possible

	clearArea() // no garden on area --> space for new ones
	for currHeight = minHeight; currHeight < lowestArea[0][0]/minHeight; currHeight++ {
		if lowestArea[0][0]/currHeight < currHeight {
			break
		}
		runField(0, currHeight, nil) // first step of calculation
		clearArea()                  // no garden on area --> space for new ones
	}
}

func runField(x int, height int, unused []bool) { // x = column at x, height = maxHeight of column, unused = all unused indeces of gardens are true
	nextHeight := 0 // height of next empty space in y-direction
	nextY := -1     // starting y at x (first empty field in column)
	if x == 0 {     // if nothing is placed (x == 0), than make new unused array and all is unused
		unused = make([]bool, len(gardens))
		for i := 0; i < len(unused); i++ {
			unused[i] = true
		}
	}

	if !contains(unused, true) {
		return // if every garden is used --> return (do not calculate further from this point)
	}
	for i := 0; i < len(unused); i++ {
		if unused[i] {
			if gardens[i].Height > lowestArea[0][0]/height+x &&
				gardens[i].Width > lowestArea[0][0]/height+x {
				return // if garden's position + its width/height is too far, next possible Area can only be greater than current lowest Area --> return (do not calculate further from this point)
			}
		}
	}

	if x >= areaWidth {
		return // if x greater than width of area --> return
	}
	for i := 0; i < height; i++ {
		if nextY < 0 { // if has no starting yet
			if area[x][i] == -1 { // field is empty
				nextY = i // starting y = first empty field
				nextHeight++
			}
		} else {
			if area[x][i] == -1 { // field is empty
				nextHeight++
			} else {
				break // starting y set and this field not empty --> nextHeight calculated
			}
		}
	}

	// get all possibilities of combinations in this column with further information (x, unused gardens and the above calculalted height)
	poss := findPossibilities(x, unused, nextHeight)

	// go through every possibility
	for i := 0; i < len(poss); i++ {
		buffY := nextY // for not positioning every garden on same y location

		// positions every garden
		for _, index := range poss[i] {
			if index < 0 {
				index *= (-1) // if index = 0, rotated or not has no difference --> all indeces int poss[] += 1
				index--       // if index = 0, rotated or not has no difference --> all indeces in poss[] += 1
				gardens[index].Rotated = true
			} else {
				index-- // if index = 0, rotated or not has no difference --> all indeces in poss[] += 1
				gardens[index].Rotated = false
			}

			gardens[index].X = x     // position garden
			gardens[index].Y = buffY // position garden

			if !gardens[index].Rotated {
				buffY += gardens[index].Height // adds height of garden in area to buffY (next garden's y)
			} else {
				buffY += gardens[index].Width // adds height of garden in area to buffY (next garden's y)
			}
			unused[index] = false // this garden is no longer unused
		}
		updateArea(unused)           // updates Area with all used gardens
		if !contains(unused, true) { //if there is no unused garden anymore
			tries++
			calcArea()
		} else {
			runField(x+1, height, unused) // calculate next column of area with the now unused gardens
		}
		// every garden in this possibility-go through is now available again
		for _, v := range poss[i] {
			if v < 0 {
				v *= (-1)
			}
			v--
			unused[v] = true
		}
	}
	// if this column has no empty spaces or too small spaces, try next column
	if len(poss) < 1 && contains(unused, true) {
		runField(x+1, height, unused)
	}
}

// finds possibilities at column x, and a specific height with unused gardens
func findPossibilities(x int, unused []bool, height int) [][]int {

	poss := make([][]int, 0) // all possibilities of lowest Area size

	for h := height; h > 0; h-- { // every height equal or lower than the entered height is tested
		for i := h; i > 0; i-- { // every height equal or lower than h tested
			for j := 0; j < len(gardens); j++ { //every garden is tested
				if unused[j] {
					if gardens[j].Height == i && gardens[j].Width+x <= lowestArea[0][0]/currHeight { // if garden's height is equal to i and garden fits in the remaining width (width-x)
						if gardens[j].Height == h { // if this is true, no other garden fits in remaining space --> only possibility
							newA := []int{j + 1}      // if index = 0, rotated or not has no difference --> all indeces += 1
							poss = append(poss, newA) // add to possibilities
						} else {
							newA := []int{j + 1}                                             // if index = 0, rotated or not has no difference --> all indeces += 1
							unused[j] = false                                                // garden will not be used in next step
							lIPoss := findPossibilities(x, unused, height-gardens[j].Height) // get possibilities with remaining height and remaining gardens
							unused[j] = true
							if len(lIPoss) > 0 {
								for k := 0; k < len(lIPoss); k++ {
									if !has2DArrayValues(poss, addToArray(newA, lIPoss[k])) { // is this combination stored already?
										poss = append(poss, addToArray(newA, lIPoss[k])) // add combination to possibilities
									}
								}
							} else {
								if !has2DArrayValues(poss, newA) { // is this combination stored already?
									poss = append(poss, newA) // add to possibilities
								}
							}
						}
					} else if gardens[j].Width == i && gardens[j].Height+x <= lowestArea[0][0]/currHeight { // if garden's width is equal to i and garden fits in the remaining width (width-x)
						if gardens[j].Width == h { // if this is true, no other garden fits in remaining space --> only possibility
							newA := []int{-(j + 1)}   // negative --> rotated. if index = 0, rotated or not has no difference --> all indeces += 1
							poss = append(poss, newA) // add to possibilities
						} else {
							newA := []int{-(j + 1)}                                         // negative --> rotated. if index = 0, rotated or not has no difference --> all indeces += 1
							unused[j] = false                                               // garden will not be used in next step
							lIPoss := findPossibilities(x, unused, height-gardens[j].Width) // get possibilities with remaining height and remaining gardens
							unused[j] = true
							if len(lIPoss) > 0 {
								for k := 0; k < len(lIPoss); k++ {
									if !has2DArrayValues(poss, addToArray(newA, lIPoss[k])) { // is this combination stored already?
										poss = append(poss, addToArray(newA, lIPoss[k])) // add combination to possibilities
									}
								}
							} else {
								if !has2DArrayValues(poss, newA) { // is this combination stored already?
									poss = append(poss, newA) // add combination to possibilities
								}
							}
						}
					}
				}
			}
		}
	}
	return poss
}

func updateArea(unused []bool) {
	clearArea() // --> no garden in area

	for i := 0; i < len(gardens); i++ { // for every garden
		if unused == nil || !unused[i] { // if garden is used
			if !gardens[i].Rotated { // if not rotated --> width = width and height = height
				for j := 0; j < gardens[i].Width; j++ {
					for k := 0; k < gardens[i].Height; k++ {
						area[gardens[i].X+j][gardens[i].Y+k] = i // area value on position = index of garden
					}
				}
			} else { // if rotated --> width = height and height = width
				for j := 0; j < gardens[i].Height; j++ {
					for k := 0; k < gardens[i].Width; k++ {
						area[gardens[i].X+j][gardens[i].Y+k] = i // area value on position = index of garden
					}
				}
			}
		}
	}
}

func clearArea() {
	// every position in area is -1 --> no gardens
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[i]); j++ {
			area[i][j] = -1
		}
	}
}

func Results(gcdScaleResult int) {
	fmt.Println()
	fmt.Println("lowest:", lowestArea[0][0]*gcdScaleResult*gcdScaleResult) // scale lowest area output back to origin
	//	fmt.Println("tries:", tries)
	lastHeight := 0                        // cp currHeight
	outputCounter := 0                     // counts number of images
	scale := 750 / len(area)               // image's height and width about 500
	for c := 0; c < len(lowestArea); c++ { // every calculated combination with lowest area size
		for i := 1; i <= len(gardens); i++ { // get information of every garden in combination c
			gardens[i-1].X = lowestArea[c][i*3-2]
			gardens[i-1].Y = lowestArea[c][i*3-1]
			if lowestArea[c][i*3] == 0 {
				gardens[i-1].Rotated = false
			} else {
				gardens[i-1].Rotated = true
			}
		}
		updateArea(nil) // every garden is used and on area
		currHeight := 0 // for no excessive amount of images
		for i := 0; i < len(area); i++ {
			for j := 0; j < len(area[i]); j++ {
				if area[i][j] != -1 && j > currHeight {
					currHeight = j
				}
			}
		}
		// if currHeight = lastHeight, this combination won't be used, because there was something similar before
		if currHeight != lastHeight {
			width, height := getAreaBounds(area)
			width *= gcdScaleResult
			height *= gcdScaleResult
			fmt.Println(width, "x", height) // prints areaBounds
			// Create Image and draw the Gardens/Field
			var img = image.NewRGBA(image.Rect(0, 0, len(area)*scale, len(area[0])*scale))
			drawField(area, len(gardens), scale, img)
			imgName := "field" + strconv.Itoa(outputCounter) + ".png"
			writeImageToFile(imgName, img)
			outputCounter++

			lastHeight = currHeight
		}
	}
}

package main

/**
In this File all the high level functions
for actually outputing a image are located
**/

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

// Generates a random Color in a colorrange
func randomColor() color.RGBA {
	col := color.RGBA{uint8(rand.Intn(100)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
	return col
}

// returns an array of wanted length with different colors
func getRandomColors(count int) []color.RGBA {
	multCol := make([]color.RGBA, count) //return array
	for count > 0 {                      //because of subtraction the line below the lowest index will be 0
		count--
		multCol[count] = randomColor()
	}
	return multCol
}

// paints every Pixel with values of field
func drawField(field [][]int, gardenCount int, scale int, img *image.RGBA) {
	white := color.RGBA{255, 255, 255, 255} // WHITE
	colors := getRandomColors(gardenCount)  // random color for every value
	for i := 0; i < len(field); i++ {       // x
		for s1 := 0; s1 < scale; s1++ { // scale of x
			for j := 0; j < len(field[i]); j++ { // y
				for s2 := 0; s2 < scale; s2++ { //scale of y
					if field[i][j] >= 0 {
						img.Set(i*scale+s1, j*scale+s2, colors[field[i][j]])
					} else {
						img.Set(i*scale+s1, j*scale+s2, white) //pixel is white if value of index == -1 (no garden on position)
					}
				}
			}
		}
	}
}

func writeImageToFile(filename string, img *image.RGBA) {
	// Create a File out of it.
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

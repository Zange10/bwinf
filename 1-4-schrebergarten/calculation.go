package main

func calcGCD(gardens []garden) int {
	smallest := gardens[0].Width        // smallest side length
	for i := 0; i < len(gardens); i++ { // get all side length of every garden
		if gardens[i].Width < smallest {
			smallest = gardens[i].Width
		}
		if gardens[i].Height < smallest {
			smallest = gardens[i].Height
		}
	}

	gcd := 1                           // greatest common divisor; every natural can be divided by 1
	for c := 1; c <= smallest/2; c++ { // c = every try for the gcd
		mayGCD := true // becomes false if even one side length cannot be divided by c
		for i := 0; i < len(gardens); i++ {
			if gardens[i].Width%c != 0 {
				mayGCD = false
				break
			}
			if gardens[i].Height%c != 0 {
				mayGCD = false
				break
			}
		}
		if mayGCD {
			gcd = c
		}
	}
	return gcd
}

// returns minimum of Height for every possible combinations
func calcMinHeight(tmpGardens []garden) int {
	height := 0
	for _, currGarden := range tmpGardens {
		// if garden does not fit in current lowest height --> lowest height = smallest side length
		if currGarden.Width > height && currGarden.Height > height {
			if currGarden.Height < currGarden.Width {
				height = currGarden.Height
			} else {
				height = currGarden.Width
			}
		}
	}
	return height
}

// calculates area and stores it, if it fulfills further conditions
func calcArea() {
	// area calculation
	width, height := getAreaBounds(area)
	size := width * height

	if size < lowestArea[0][0] { // size is lower than current lowest size
		lowestArea = make([][]int, 1) // all possibilities of lowest Area size
		for i := 0; i < len(lowestArea); i++ {
			lowestArea[i] = make([]int, len(gardens)*3+1) // [0] = area size, [1] = x, [2] = y, [3] = rotated (1/0)  --> every garden
		}

		lowestArea[0][0] = size
		for i := 1; i <= len(gardens); i++ {
			lowestArea[0][i*3-2] = gardens[i-1].X
			lowestArea[0][i*3-1] = gardens[i-1].Y
			if !gardens[i-1].Rotated {
				lowestArea[0][i*3] = 0
			} else {
				lowestArea[0][i*3] = 1
			}
		}
	} else if size == lowestArea[0][0] { // size is equal to current lowest size
		addLowestArea := make([][]int, 1) // all possibilities of lowest Area size, which can be added to the current combinations
		for i := 0; i < len(addLowestArea); i++ {
			addLowestArea[i] = make([]int, len(gardens)*3+1) // [0] = area size, [1] = x, [2] = y, [3] = rotated (1/0)  --> every garden
		}
		addLowestArea[0][0] = size
		for i := 1; i <= len(gardens); i++ { // storing data of every garden
			addLowestArea[0][i*3-2] = gardens[i-1].X
			addLowestArea[0][i*3-1] = gardens[i-1].Y
			if !gardens[i-1].Rotated {
				addLowestArea[0][i*3] = 0
			} else {
				addLowestArea[0][i*3] = 1
			}
		}

		if !hasValues(lowestArea, addLowestArea) { // if this combination does not exist yet
			lowestArea = addTo2DArray(lowestArea, addLowestArea, len(gardens)*3+1)
		}
	}
}

// returns sidelengths
func getAreaBounds(area [][]int) (int, int) {
	width := 0
	height := 0
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[i]); j++ {
			if area[i][j] >= 0 && i > width { // if a garden is on position and its x > current highest width --> heightest width = x
				width = i
			}
			if area[i][j] >= 0 && j > height { // if a garden is on position and its x > current highest width --> heightest width = x
				height = j
			}
		}
	}
	width++  // it started with 0
	height++ // it started with 0
	return width, height
}

// calculates maximum width of area (returns the greatest side length of all gardens --> all gardens can be placed next to each other without further problems)
func getMaxWidth(tmpGardens []garden) int {
	mw := 0 //return value
	for i := 0; i < len(tmpGardens); i++ {
		if tmpGardens[i].Width > tmpGardens[i].Height { // adds greatest side length to mw
			mw += tmpGardens[i].Width
		} else {
			mw += tmpGardens[i].Height
		}
	}
	return mw
}

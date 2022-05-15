package main

// if array contains b --> return true
func contains(array []bool, b bool) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == b {
			return true
		}
	}
	return false
}

// adds new indeces for array in newArray
func addToArray(array []int, addArray []int) []int {
	newArray := make([]int, len(array)+len(addArray))
	for i := 0; i < len(array); i++ {
		newArray[i] = array[i]
	}
	for i := 0; i < len(addArray); i++ {
		newArray[len(array)+i] = addArray[i]
	}
	return newArray
}

// adds new indeces for array in newArray
func addTo2DArray(array [][]int, addArray [][]int, sndSize int) [][]int {
	newArray := make([][]int, len(array)+len(addArray)) // all possibilities of lowest Area size
	for i := 0; i < len(newArray); i++ {
		newArray[i] = make([]int, sndSize) // [0] = area size, [1] = x, [2] = y, [3] = rotated (1/0)  --> every garden
	}
	for i := 0; i < len(array); i++ {
		newArray[i] = array[i]
	}
	for i := 0; i < len(addArray); i++ {
		newArray[len(array)+i] = addArray[i]
	}
	return newArray
}

// if any index of array2d == array --> return true
func has2DArrayValues(array2d [][]int, array []int) bool {
	thisIndex := true
	for i := 0; i < len(array2d); i++ {
		array1 := array2d[i]
		for j := 0; j < len(array1); j++ {
			if array1[0] != array[0] {
				thisIndex = false
			}
		}
		if thisIndex {
			return true
		}
		thisIndex = true
	}
	return false
}

// if any index of array1 == any index of array2 --> return true
func hasValues(array1 [][]int, array2 [][]int) bool {
	thisIndex := true
	if len(array1[0]) == len(array2[0]) {
		for _, a1 := range array1 {
			for _, a2 := range array2 {
				for i := 0; i < len(a1); i++ {
					if a1[i] != a2[i] {
						thisIndex = false
						break
					}
				}
				if thisIndex {
					return true
				}
				thisIndex = true
			}
		}
	}
	return false
}

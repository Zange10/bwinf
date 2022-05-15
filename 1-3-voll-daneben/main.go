package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Declare Flag --dataset to set SampleData path
	var datasetpath = flag.String("dataset", "", "Input the Path to your Dataset")
	flag.Parse()
	if *datasetpath == "" {
		// We have no dataset path so we exit
		fmt.Println("Please specify dataset Path with flag --dataset")
		os.Exit(1)
	}
	// Parse all Numbers from the Players into an integer slice and sort them
	playerNumbers := parseSampleData(*datasetpath)
	sort.Ints(playerNumbers)
	// Find Als Numbers and print them out
	luckynumbers := findAlsNumbers(playerNumbers)
	fmt.Println("These are Als Numbers for the Dataset")
	fmt.Println(luckynumbers)
}

func findAlsNumbers(tmpSlice []int) []int {
	// Returns ALs ten Numbers out of an sorted Array
	alsNumbers := make([]int, 10)
	// Divide into 10 Parts and the caluclate Median of every part
	playerParts := divideIntoParts(tmpSlice)
	for index := 0; index < len(playerParts); index++ {
		alsNumbers[index] = medianOfSlice(playerParts[index])
	}
	return alsNumbers
}

func medianOfSlice(tmpSlice []int) int {
	// Calculates the Median of multiple Numbers in a Slice
	l := len(tmpSlice)
	if l == 0 {
		// No math is needed if there are no numbers
		return 0
	} else if l%2 == 0 {
		// For even numbers we add the two middle numbers
		// and divide by two
		middlenumbers := tmpSlice[l/2-1 : l/2+1]
		return (middlenumbers[0] + middlenumbers[1]) / 2
	} else {
		// For odd numbers we just use the middle number
		return tmpSlice[l/2]
	}
}

func divideIntoParts(tmpSlice []int) [][]int {
	// Divides an slice into 10 even Parts. If the Slice is not even, some numbers just get kicked out. Could be fixed late.
	x := 0
	partSize := len(tmpSlice) / 10

	parts := [][]int{tmpSlice[x : x+partSize]}
	x += partSize
	for index := 1; index < 10; index++ {
		parts = append(parts, tmpSlice[x:x+partSize])
		x += partSize
	}
	return parts
}

func parseSampleData(filename string) []int {
	// Opens the selected Sample File .txt
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer file.Close()

	// Parses the File into a string array
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// string to int
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		lines = append(lines, i)
	}

	return lines
}

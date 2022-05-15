package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	var origText string               //the original text
	var dictionary []string           //all words from the dictionary
	var twistedText string            //the twisted text
	var detwistedText string          //the detwisted text (hopefully equals origText)
	var wordsAndChars []string        //all words and chatacters from origText
	var twistedWordsAndChars []string //all words and chatacters from twistedText

	// Declare Flags for Parameters
	var datasetpath = flag.String("dataset", "", "Input the Path to your Dataset")
	var dictionarypath = flag.String("dictionary", "beispieldaten/woerterliste.txt", "Path of your Dictionary")
	flag.Parse()
	if *datasetpath == "" {
		// We have no dataset path so we exit
		fmt.Println("Please specify dataset Path with flag --dataset")
		os.Exit(1)
	}
	// Load the Dictionary and Dataset
	origText = parseSampleData(*datasetpath)
	dictionary = parseData(*dictionarypath)

	// Split to Slice
	wordsAndChars = transformTextToSlice(origText)

	//Forming "twisted" text by using wordsAndChars
	for _, wordOrChar := range wordsAndChars {
		if len(wordOrChar) > 3 {
			twistedText += twistWord(wordOrChar)
		} else {
			twistedText += wordOrChar
		}
	}

	fmt.Println(twistedText)

	//Twisting --> Detwisting --------------------------

	twistedWordsAndChars = transformTextToSlice(twistedText)

	for _, wordOrChar := range twistedWordsAndChars {
		if len(wordOrChar) > 3 {
			detwistedText += detwistWord(wordOrChar, dictionary)
		} else {
			detwistedText += wordOrChar
		}
	}

	fmt.Println(detwistedText)
}

func transformTextToSlice(text string) []string {
	var words []string     //stores words and Characters
	buffer := 0            //stores beginning of a word
	lastWasLetter := false //Is true, if the character before the charakter, which is checked, was a letter

	//Transforms text to words und characters
	for i, r := range text {
		if !unicode.IsLetter(r) {
			if lastWasLetter {
				words = append(words, text[buffer:i]) //adding word to words
				buffer = i
			}
			words = append(words, text[i:i+1]) //adding character, which is no letter to words
			buffer++
		} else {
			lastWasLetter = true
		}

		if i == (len(text) - 1) {
			words = append(words, text[buffer:i+1])
		} //Last letter in text
	}
	return words
}

//Twisting -------------------------------------------------------------------------------------------------

func twistWord(word string) string {
	var chars []string     //Later: All characters in word
	var twistedWord string //Result
	var randNum int

	indexesOfUmlauts := make([]string, len(word)) //indexes of german umlauts in word

	//find german umlauts with runes
	for i, r := range word {
		switch r {
		case 223:
			indexesOfUmlauts[i] = "ß"
		case 228:
			indexesOfUmlauts[i] = "ä"
		case 246:
			indexesOfUmlauts[i] = "ö"
		case 252:
			indexesOfUmlauts[i] = "ü"
		case 196:
			indexesOfUmlauts[i] = "Ä"
		case 214:
			indexesOfUmlauts[i] = "Ö"
		case 220:
			indexesOfUmlauts[i] = "Ü"
		default:
		}
	}

	//chars[] gets all characters in word
	for i := 0; i < len(word); i++ {
		//if german umlaut at this point
		if indexesOfUmlauts[i] != "" {
			chars = append(chars, indexesOfUmlauts[i])
			i++ //skipping character (added after every german umlaut)
		} else {
			chars = append(chars, word[i:i+1])
		}
	}

	twistedWord = chars[0] //First letter in word
	maxChanges := len(chars) - 2

	//Creating random number and adding chars[random Number] to twistedWord
	for i := 0; i < len(chars)-2; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		randNum = rand.Intn(maxChanges-0) + 1
		for chars[randNum] == "" {
			rand.Seed(time.Now().UTC().UnixNano())
			randNum = rand.Intn(maxChanges-0) + 1
		}
		twistedWord += chars[randNum]
		chars[randNum] = "" //For no duplications
	}
	twistedWord += chars[len(chars)-1] //Last letter in word
	return twistedWord
}

//Detwisting ----------------------------------------------------------------------------------------------

func detwistWord(twistedWord string, dictionary []string) string {
	//Finds out all possible words out of dictionary
	possibleWords := findPossibleWords(twistedWord, dictionary)
	if len(possibleWords) == 0 { //if there is no possible word --> no result
		return "(" + twistedWord + "(???))"
	} else if len(possibleWords) == 1 { //if there is only one possible word --> result
		return possibleWords[0]
	} else { //if there are more than one possible words --> multiple results
		if len(possibleWords) == 2 && possibleWords[0] == possibleWords[1] {
			return possibleWords[0]
		}
		morePosWords := "("
		for i, r := range possibleWords {
			morePosWords += r
			//last possible word
			if i < len(possibleWords)-1 {
				morePosWords += ", "
			}
		}
		morePosWords += ")"
		return morePosWords // e.g. "(posWord1, posWord2, posWord3)"
	}
}

func findPossibleWords(twistedWord string, dictionary []string) []string {
	var possibleWords []string   //all possible words out of dictionary
	var sortedTwistedWord string //all letters sorted (exept first and last)
	var sortedWord string        //all letters sorted (exept first and last)

	for _, word := range dictionary {
		//if first letter, length and last letter are the same
		if strings.ToLower(twistedWord[:1]) == strings.ToLower(word[:1]) &&
			len(twistedWord) == len(word) &&
			twistedWord[len(twistedWord)-1:len(twistedWord)] == word[len(word)-1:len(word)] {

			sortedTwistedWord = sortLetters(twistedWord) //e.g. "k-abcdef-z" (first and last not sorted)
			sortedWord = sortLetters(word)               //e.g. "k-abcdef-z" (first and last not sorted)

			if strings.ToLower(sortedWord) == strings.ToLower(sortedTwistedWord) {
				//if has all the same letters --> append
				possibleWords = append(possibleWords, word)
			}
		}
	}
	return possibleWords
}

func sortLetters(word string) string {
	var smallestInd int
	var letters []string
	var sortedWord string

	//stores every character of word as string
	for i := 0; i < len(word); i++ {
		letters = append(letters, word[i:i+1])
	}

	//Selectionsort
	for i := 1; i < len(letters)-1; i++ {
		smallestInd = i //starting point
		for j := i + 1; j < len(letters)-1; j++ {
			if letters[j] < letters[smallestInd] {
				smallestInd = j //the index of the smallest letter (a less than b)
			}
		}
		//smallest letter and letter of index i are swapped
		buffer := letters[i]
		letters[i] = letters[smallestInd]
		letters[smallestInd] = buffer
	}

	for _, r := range letters {
		sortedWord += r
	}
	return sortedWord
}

//Parsing Data -------------------------------------------------------------------------------------------

func parseSampleData(filename string) string {
	lines := parseData(filename)
	// add a wordwrap at the end of every line
	for i := 0; i < len(lines); i++ {
		lines[i] += "\n"
	}

	// Convert string array to string
	var stringOut string
	for _, text := range lines {
		stringOut += text
	}
	return stringOut
}

func parseData(filename string) []string {
	// Opens the selected Sample File .txt
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		var er []string
		//returns null string array (should never come to this point)
		return er
	}
	defer file.Close()

	// Parses the File into a string array
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

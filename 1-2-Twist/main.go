package main

import (
    "time"
	"fmt"
    "math/rand"
	"unicode"
)

func main() {
	var origText string
	var twistedText string
	var wordsAndChars []string

	origText = "Das ist ein Beispieltext zum Austesten des Twisten!"

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
	
}

func transformTextToSlice(text string) []string {
	var words []string //stores words and Characters
	buffer := 0 //stores beginning of a word
	lastWasLetter := false //Is true, if the charakter before the charakter, which is checked, was a letter

	//Transforms text to words und charakters
	for i, r := range text {
        if !unicode.IsLetter(r) {
			if lastWasLetter {
				words = append(words,text[buffer:i])	//adding word to words
				buffer = i
			}
			words = append(words,text[i:i+1])	//adding charakter, which is no letter to words
			buffer++
		} else {
			lastWasLetter = true
		}

		if i == (len(text)-1) {words = append(words,text[buffer:i+1])}	//Last letter in text
    }
	return words
}

func twistWord(word string) string {
	chars := make([]string, len(word))	//Later: All charakters in word
	var twistedWord string	//Result
	var randNum int

	//chars[] gets all charakters in word
	for i := 0; i < len(word); i++ {
		chars[i] = word[i:i+1]
	}

	twistedWord = chars[0]	//First letter in word
	maxChanges := len(word) - 2

	//Creating random number and adding chars[random Number] to twistedWord
	for i := 0; i < len(word)-2; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		randNum = rand.Intn(maxChanges-0) + 1
		for chars[randNum] == "" {
			time.Sleep(2)		//Doesn't work faster without it --> better performance
			rand.Seed(time.Now().UTC().UnixNano())
			randNum = rand.Intn(maxChanges-0) + 1
		}
		twistedWord += chars[randNum]
		chars[randNum] = ""		//For no duplications
	}
		
	twistedWord += chars[len(word)-1]	//Last letter in word
	return twistedWord
}  
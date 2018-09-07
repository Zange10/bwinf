package main

import (
	"fmt"
)

// Basic Structure to define an user
type user struct {
	name, group string
	followers   []string
	following   []string
}

func main() {

	// ====== Define some sample Data =========
	justin := user{
		name:      "Justin",
		group:     "Kids",
		followers: []string{"Hailey", "Selena"},
	}

	selena := user{
		name:      "Selena",
		group:     "Kids",
		followers: []string{"Hailey"},
		following: []string{"Justin"},
	}

	hailey := user{
		name:      "Hailey",
		group:     "Kids",
		following: []string{"Justin", "Selena"},
	}
	// Print out Sample Data
	fmt.Println(justin)
	fmt.Println(selena)
	fmt.Println(hailey)
}

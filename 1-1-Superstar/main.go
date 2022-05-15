package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Basic Structure to define a user
type user struct {
	name         string
	followers    []string
	following    []string
	maySuperstar bool
}

// Counter for counting the requests
var isFollowingCounter int

// Our TeeniGram Group
var teenigramgroup []user

func main() {
	// Declare Flag --dataset to set SampleData path
	var datasetpath = flag.String("dataset", "", "Input the Path to your Dataset")
	flag.Parse()
	if *datasetpath == "" {
		// We have no dataset path so we exit
		fmt.Println("Please specify dataset Path with flag --dataset")
		os.Exit(1)
	}
	// We parse our sample Data into a group object
	teenigramgroup = parseSampleData(*datasetpath)
	fmt.Println("We have", len(teenigramgroup), "Users in our Group")
	// Checks if theres a Superstar in group[] and stores his/her name
	superstar := findSuperStar(teenigramgroup)
	if superstar != "" {
		fmt.Println(superstar, "is a superstar")
	} else {
		fmt.Println("No superstar in this group")
	}

	fmt.Println("We used", isFollowingCounter, "Requests")
}

func findSuperStar(searchgroup []user) string {
	index := 0
	// Loops over the group and checks if the user is a superstar
	for index = 0; index < len(searchgroup); index++ {
		if searchgroup[index].maySuperstar {
			if isUserSuperstar(index, searchgroup) {
				break
			}
		}
	}
	if index < len(searchgroup) {
		return searchgroup[index].name
	}
	return ""
}

func isUserSuperstar(userindex int, tmpGroup []user) bool {
	testUser := tmpGroup[userindex]
	// Deletes testUser from tmpGroup that he wont be checked
	testGroup := removeUserFromGroup(userindex, tmpGroup)

	// User cant be a Superstar if hes following someone in the Group
	for index := 0; index < len(testGroup); index++ {
		if isFollowing(testUser, testGroup[index]) {
			return false // Abort
		}
		// If testUser isnt following testGroup[index] user, he cant be a Superstar
		teenigramgroup[findUserIndexFromGroup(testGroup[index].name)].maySuperstar = false
	}
	// Tests if all Users in the Group are following testUser
	for index := 0; index < len(testGroup); index++ {
		if !isFollowing(testGroup[index], testUser) {
			return false
		}
	}
	// If these two pass, we have got a Superstar
	return true
}

func isFollowing(person1, person2 user) bool {
	// Increment the RequestCounter
	isFollowingCounter++
	// Iterate over the followers of person2 and check if person1's name is in there
	for _, follower := range person2.followers {
		if follower == person1.name {
			// We have a match
			return true
		}
	}
	return false
}

func removeUserFromGroup(userindex int, tmpGroup []user) []user {
	removedGroup := make([]user, len(tmpGroup))
	copy(removedGroup, tmpGroup)
	removedGroup = append(removedGroup[:userindex], removedGroup[userindex+1:]...)
	return removedGroup
}

func findUserIndexFromGroup(finduser string) int {
	for index, follower := range teenigramgroup {
		if follower.name == finduser {
			return index
		}
	}
	// Error
	return -1
}

func parseSampleData(filename string) []user {
	// Opens the selected Sample File .txt
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	// Parses the File into a string array
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Creates a SampleGroup with all Users in it.
	datapoints := strings.Fields(lines[0])
	var samplegroup = make([]user, len(datapoints))
	for index, username := range datapoints {
		samplegroup[index] = user{
			name:         username,
			maySuperstar: true,
			// Switches to false if its not possible anymore for him/her to be a superstar
		}
	}

	// Loops over every line and parses it into an user object
	for index := 1; index < len(lines)-1; index++ {
		followingdata := strings.Fields(lines[index])
		for index, searchuser := range samplegroup {
			if searchuser.name == followingdata[0] {
				// Add
				samplegroup[index].following = append(samplegroup[index].following, followingdata[1])
			}
			if searchuser.name == followingdata[1] {
				samplegroup[index].followers = append(samplegroup[index].followers, followingdata[0])
			}
		}
	}
	return samplegroup
}

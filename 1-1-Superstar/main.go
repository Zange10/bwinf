package main

import (
	"fmt"
)

// Basic Structure to define an user
type user struct {
	name      string
	followers []string
	following []string
}

// Counter to know how much Requests we are making
var isFollowingCounter int

func main() {

	// ====== Define a sample Group with Users in it =========
	var samplegroup = []user{
		user{
			name:      "Selena",
			followers: []string{"Hailey"},
			following: []string{"Justin"},
		},
		user{
			name:      "Justin",
			followers: []string{"Hailey", "Selena"},
		},
		user{
			name:      "Hailey",
			following: []string{"Justin", "Selena"},
		},
	}
	// Check if theres a Superstar in group[]
	getSuperstar(samplegroup)
}

func getSuperstar(searchgroup []user) {
	for index := 0; index < len(searchgroup); index++ {
		// fmt.Println("Testing this Group: ", searchgroup)
		isUserSuperstar(index, searchgroup)
	}
	fmt.Println("We used", isFollowingCounter, "Requests")
}

func isUserSuperstar(userindex int, tmpGroup []user) {
	testUser := tmpGroup[userindex]
	fmt.Println("Testing:", testUser.name)
	// Deletes testUser from tmpGroup
	testGroup := removeUserFromGroup(userindex, tmpGroup)

	// User cant be an Superstar if hes following someone in the Group
	for index := 0; index < len(testGroup); index++ {
		if isFollowing(testUser, testGroup[index]) {
			fmt.Println(testUser.name, "is following someone in the group, he's/she's not a Superstar")
			return // Abort
		}
	}
	// Test if all Users in the Group are following testUser
	for index := 0; index < len(testGroup); index++ {
		if !isFollowing(testGroup[index], testUser) {
			fmt.Println("Not all Users are following", testUser.name, "he's/she's not a Superstar")
			return
		}
	}
	// If these two pass we have a Superstar
	fmt.Println(testUser.name, "is an SuperStar")
}

func isFollowing(person1, person2 user) bool {
	// Increment the FollowingCounter
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

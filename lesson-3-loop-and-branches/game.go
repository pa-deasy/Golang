package main

import (
	"fmt"
)

func main() {
	checkRoom("cave")
}

func checkRoom(room string) {
	if room == "cave" {
		fmt.Println("You enter a dark dank room")
	} else if room == "entrance" {
		fmt.Println("You enter a huge open cavern with water dripping from the roof")
	} else if room == "mountain" {
		fmt.Println("You stand afoot the highest mountain in the land")
	} else {
		println("You stand still as a breeze flutters by")
	}
}

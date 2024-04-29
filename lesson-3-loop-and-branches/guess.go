package main

import (
	"fmt"
	"math/rand"
)

func main() {
	guess()
}

func guess() {
	const myNumber = 9
	var guessedCorrectly = false

	for !guessedCorrectly {
		var guess = rand.Intn(21)

		if guess == myNumber {
			fmt.Printf("The guess %v is correct \n", guess)
			guessedCorrectly = true
		} else if guess > myNumber {
			fmt.Printf("The guess %v is too high \n", guess)
		} else if guess < myNumber {
			fmt.Printf("The guess %v is too low \n", guess)
		}
	}
}

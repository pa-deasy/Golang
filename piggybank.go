package main

import (
	"fmt"
	"math/rand"
)

const nickle = 0.05
const dime = 0.10
const quater = 0.25

func main() {
	saveTill(20)
}

func saveTill(target float64) {
	for piggyTotal := 0.0; piggyTotal <= target; {
		fmt.Println("Balance is: ", piggyTotal)

		randomCoin := randomCoin()
		piggyTotal += randomCoin

		fmt.Println("Added: ", randomCoin)
		fmt.Println("New balance is: ", piggyTotal)
	}
}

func randomCoin() float64 {
	switch coinNumber := rand.Intn(3); coinNumber {
	case 0:
		return nickle
	case 1:
		return dime
	case 2:
		return quater
	default:
		return 0
	}
}

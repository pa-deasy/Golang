package main

import (
	"fmt"
	"math/rand"
)

const nickle = 5
const dime = 10
const quater = 25

func main() {
	saveTill(20)
}

func saveTill(targetDollars int) {
	for piggyTotalCents := 0; piggyTotalCents <= targetDollars*100; {
		fmt.Println("\nBalance in cents: ", piggyTotalCents)

		randomCoin := randomCoin()
		piggyTotalCents += randomCoin

		totalDollars := piggyTotalCents / 100
		remainingCents := piggyTotalCents % 100

		fmt.Println("Added cents: ", randomCoin)
		fmt.Printf("New balance is: $%v.%v \n", totalDollars, remainingCents)
	}
}

func randomCoin() int {
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

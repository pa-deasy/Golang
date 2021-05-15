package main

import (
	"fmt"
	"math/rand"
)

func main() {
	AgeAndWeight()
	TravelDays()
}

func AgeAndWeight() {
	fmt.Printf("My age on Mars would be %v \n", 32*365/687)
	fmt.Printf("My weight on Mars would be %v kg\n", 70*0.3783)
}

func TravelDays() {
	const hoursPerDay = 24
	var distance = 96300000
	var speed = 100800

	fmt.Printf("It would take %v days to reach Mars\n", distance/(hoursPerDay*speed))

	distance = rand.Intn(345000001) + 56000000

	fmt.Printf("Days to mars could also be %v\n", distance/(hoursPerDay*speed))
}

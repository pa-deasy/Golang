package main

import (
	"fmt"
	"math/rand"
)

const speedRange = 16

func main() {
	printTickets()
}

func printTickets() {
	println("Spaceline         Days  Triptype    Price")
	println("========================================")
	for index := 0; index < 10; index++ {
		printRandomTicket()
	}
}

func printRandomTicket() {
	spaceline := randomAirline()
	speedKm := randomSpeed()
	days := calculateTravelDays(speedKm)
	tripType := randomTripType()
	price := calculateCost(speedKm, tripType)

	fmt.Printf("%-17v %-5v %-11v $%4v \n", spaceline, days, tripType, price)
}

func randomAirline() string {
	switch airlineNumber := rand.Intn(3); airlineNumber {
	case 0:
		return "Space Adventures"
	case 1:
		return "SpaceX"
	case 2:
		return "Virgin Galactic"
	default:
		return "Airline not found"
	}
}

func randomSpeed() int {
	const baseSpeedKm = 14
	speedKm := rand.Intn(baseSpeedKm+1) + speedRange

	return speedKm
}

func calculateTravelDays(speedKm int) int {
	const distanceToMarsKm = 62100000
	const seconds = 60
	const minutes = 60
	const hours = 24

	speedDay := speedKm * seconds * minutes * hours

	return distanceToMarsKm / speedDay
}

func randomTripType() string {
	switch tripType := rand.Intn(2); tripType {
	case 0:
		return "One-way"
	case 1:
		return "Round-trip"
	default:
		return "Invalid trip"
	}
}

func calculateCost(speedKm int, tripType string) int {
	const basePrice = 36
	speedDiff := speedKm - speedRange

	cost := basePrice + speedDiff

	if tripType == "Round-trip" {
		return cost * 2
	} else {
		return cost
	}
}

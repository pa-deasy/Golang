package main

import "fmt"

func main() {
	isLeapYear(2021)
}

func isLeapYear(year int) {
	var isLeapYear = (year%4 == 0 && year%100 != 0) || (year % 400 == 0	)

	if isLeapYear {
		fmt.Printf("The year %v is a leap year", year)
	} else {
		fmt.Printf("The year %v is not a leap year", year)
	}
}

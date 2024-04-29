package main

import "fmt"

type report struct {
	sol int
	temperature
	location
}

type temperature struct {
	high celcius
	low  celcius
}

type celcius float64

type location struct {
	lat  float64
	long float64
}

func main() {
	report := report{
		15,
		temperature{-1.0, -78.0},
		location{-4.5895, 137.4417},
	}

	fmt.Printf("%+v\n", report)
	fmt.Printf("average %v °C\n", report.average())
	fmt.Printf("average %v °C\n", report.temperature.average())
	fmt.Printf("high %v °C\n", report.high)
	fmt.Printf("high %v °C\n", report.temperature.high)
}

func (t temperature) average() celcius {
	return (t.high + t.low) / 2
}

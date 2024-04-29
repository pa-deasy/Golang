package main

import "fmt"

type report struct {
	sol         int
	temperature temperature
	location    location
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
	bradbury := location{-4.5895, 137.4417}
	t := temperature{-1.0, -78.0}
	report := report{15, t, bradbury}

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v °C\n", report.temperature.high)
}

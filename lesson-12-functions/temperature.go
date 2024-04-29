package main

import "fmt"

func main() {
	kelvin := 290
	fmt.Println("Kevlin of", kelvin, "converts to celcius of", kelvinToCelcius(float64(kelvin)))

	celcius := 16.850000000000023
	fmt.Println("Cecius of", celcius, "converts to fahrenheit of", celciusToFahrenheit(celcius))

	fmt.Println("Kelvin of", kelvin, "converts to fahrenheit of", kelvinToFahrenheit(float64(kelvin)))
}

func kelvinToCelcius(kelvin float64) float64 {
	celcius := kelvin - 273.15
	return celcius
}

func celciusToFahrenheit(celcius float64) float64 {
	fahrenheit := (celcius * 9.0 / 5.0) + 32.0
	return fahrenheit
}

func kelvinToFahrenheit(kelvin float64) float64 {
	return celciusToFahrenheit(kelvinToCelcius(kelvin))
}

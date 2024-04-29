package main

import "fmt"

type kelvin float64
type celcius float64
type fahrenheit float64

func main() {
	var kelvin kelvin = 290
	fmt.Println("Kevlin of", kelvin, "converts to celcius of", kelvinToCelcius(kelvin))

	var celcius celcius = 16.850000000000023
	fmt.Println("Cecius of", celcius, "converts to fahrenheit of", celciusToFahrenheit(celcius))

	fmt.Println("Kelvin of", kelvin, "converts to fahrenheit of", kelvinToFahrenheit(kelvin))
}

func kelvinToCelcius(k kelvin) celcius {
	return celcius(k - 273.15)
}

func kelvinToFahrenheit(k kelvin) fahrenheit {
	return fahrenheit(k - 459.67)
}

func celciusToKelvin(c celcius) kelvin {
	return kelvin(c + 273.15)
}

func celciusToFahrenheit(c celcius) fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func fahrenheitToKelvin(f fahrenheit) kelvin {
	return kelvin(f + 459.67)
}

func fahrenheitToCelcius(f fahrenheit) celcius {
	return celcius((f - 32.0) * 5.0 / 9.0)
}

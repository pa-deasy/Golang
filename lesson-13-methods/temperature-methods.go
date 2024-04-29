package main

import "fmt"

type kelvin float64
type celcius float64
type fahrenheit float64

func main() {
	var kelvin kelvin = 290
	fmt.Println("Kevlin of", kelvin, "converts to celcius of", kelvin.toCelcius())

	var celcius celcius = 16.850000000000023
	fmt.Println("Cecius of", celcius, "converts to fahrenheit of", celcius.toFahrenheit())

	fmt.Println("Kelvin of", kelvin, "converts to fahrenheit of", kelvin.toFahrenheit())
}

func (k kelvin) toCelcius() celcius {
	return celcius(k - 273.15)
}

func (k kelvin) toFahrenheit() fahrenheit {
	return fahrenheit(k - 459.67)
}

func (c celcius) toKelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celcius) toFahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) toKelvin() kelvin {
	return kelvin(f + 459.67)
}

func (f fahrenheit) toCelcius() celcius {
	return celcius((f - 32.0) * 5.0 / 9.0)
}

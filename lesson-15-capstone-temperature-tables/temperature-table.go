package main

import (
	"fmt"
)

type celcius float64
type fahrenheit float64

func main() {
	drawTable("c", "f", drawCelciusRows)
	drawTable("f", "c", drawFahrenheitRows)
}

func drawTable(c1 string, c2 string, rowFn func()) {
	fmt.Println("=======================")
	fmt.Printf("| °%v       | °%v       |", c1, c2)
	fmt.Println("=======================")
	rowFn()
	fmt.Println("=======================")
}

func drawCelciusRows() {
	for c := celcius(-40); c <= 100; c += 5 {
		drawTableRow(fmt.Sprint(c), fmt.Sprint(c.toFahrenheit()))
	}
}

func drawFahrenheitRows() {
	for f := fahrenheit(-40); f <= 100; f += 5 {
		drawTableRow(fmt.Sprint(f), fmt.Sprint(f.toCelcius()))
	}
}

func drawTableRow(v1 string, v2 string) {
	fmt.Printf("| %-9v| %-9v|\n", v1, v2)
}

func (c celcius) toFahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) toCelcius() celcius {
	return celcius((f - 32.0) * 5.0 / 9.0)
}

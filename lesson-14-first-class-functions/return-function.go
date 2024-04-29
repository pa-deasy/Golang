package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

type sensor func() kelvin

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())

	fakeSensor := calibrate(fakeSensor, 10)
	fmt.Println(fakeSensor())
	fmt.Println(fakeSensor())
	fmt.Println(fakeSensor())
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return kelvin(0)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

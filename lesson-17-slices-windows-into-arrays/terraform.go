package main

import "fmt"

type planets []string

func main() {
	planets := planets{
		"Mercury", "Venus", "Earth", "Marsbar",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	planets[3:4].terraform()
	planets[6:].terraform()

	fmt.Println(planets)
}

func (planets planets) terraform() {
	for index := range planets {
		planets[index] = "New " + planets[index]
	}
}

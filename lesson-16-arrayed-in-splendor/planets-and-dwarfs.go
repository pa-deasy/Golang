package main

import "fmt"

func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumera", "Makemake", "Eris"}
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	for i, d := range dwarfs {
		fmt.Println(i+1, d)
	}

	fmt.Println(planets)
}

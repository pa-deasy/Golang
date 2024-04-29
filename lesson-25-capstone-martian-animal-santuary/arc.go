package main

import (
	"fmt"
	"math/rand"
)

type animal interface {
	move() string
	eat() string
}

type cow struct {
	name string
}

type dog struct {
	name string
}

type pig struct {
	name string
}

type gecko struct {
	name string
}

func main() {
	animals := []animal{
		cow{"Princess"},
		dog{"Rover"},
		pig{"Spider pig"},
		gecko{"Kiki"},
	}

	const sunrise, sunset int = 06, 21

	for hour := 0; hour <= 72; hour++ {
		time := hour % 24
		fmt.Println(time, ":00")
		if sunset < time || time < sunrise {
			fmt.Printf("The animals are asleep")
		} else {
			random := rand.Intn(4)

			description(animals[random])
		}
		fmt.Println()
	}
}

func description(a animal) {
	fmt.Printf("%v the %T %v and %v", a, a, a.move(), a.eat())
}

func (c cow) String() string {
	return c.name
}

func (c cow) move() string {
	return "plods along"
}

func (c cow) eat() string {
	return "chews the grass"
}

func (d dog) String() string {
	return d.name
}

func (d dog) move() string {
	return "trots about"
}

func (d dog) eat() string {
	return "bites the bone"
}

func (p pig) String() string {
	return p.name
}

func (d pig) move() string {
	return "trys to get up but fails"
}

func (d pig) eat() string {
	return "swallows anything in the vacinity"
}

func (g gecko) String() string {
	return g.name
}

func (g gecko) move() string {
	return "springs swiftly forward"
}

func (g gecko) eat() string {
	return "swallows a cricket whole"
}

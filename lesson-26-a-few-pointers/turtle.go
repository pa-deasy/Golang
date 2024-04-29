package main

import "fmt"

type location struct {
	x int
	y int
}

type turtle struct {
	name     string
	location location
}

func main() {
	turtle := turtle{
		name:     "Franky",
		location: location{x: 0, y: 0},
	}

	turtle.moveRight()
	turtle.moveDown()
	turtle.moveDown()
	turtle.moveUp()
	turtle.moveLeft()
	turtle.moveRight()
	turtle.moveDown()

	fmt.Printf("%+v", turtle)
}

func (t *turtle) moveLeft() {
	t.location.x--
}

func (t *turtle) moveRight() {
	t.location.x++
}

func (t *turtle) moveUp() {
	t.location.y--
}

func (t *turtle) moveDown() {
	t.location.y++
}

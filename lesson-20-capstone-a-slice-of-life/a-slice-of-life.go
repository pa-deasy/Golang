package main

import (
	"fmt"
	"math/rand"
	"time"
)

type universe [][]bool

const (
	width  = 80
	height = 15
)

func main() {
	u1 := newUniverse()
	u2 := newUniverse()
	u1.seed()

	for index := 0; index < 30; index++ {
		step(u1, u2)
		u1.show()
		time.Sleep(time.Second / 30)
		u1, u2 = u2, u1
	}

}

func newUniverse() universe {
	u := make(universe, height)
	for index := range u {
		u[index] = make([]bool, width)
	}

	return u
}

func (u universe) seed() {
	for index := 0; index < (width*height)/4; index++ {
		u[rand.Intn(height)][rand.Intn(width)] = true
	}
}

func (u universe) alive(y, x int) bool {
	y = (y + height) % height
	x = (x + width) % width
	return u[y][x]
}

func (u universe) neighbors(y, x int) int {
	n := 0

	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.alive(y+v, x+h) {
				n++
			}
		}
	}

	return n
}

func (u universe) next(y, x int) bool {
	neighbors := u.neighbors(y, x)
	isAlive := u.alive(y, x)

	return isAlive && neighbors == 2 || neighbors == 3
}

func step(u1, u2 universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			u2[y][x] = u1.next(y, x)
		}
	}
}

func (u universe) show() {
	const line = "------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------"

	fmt.Println("\x0c")

	fmt.Println(line)
	for _, row := range u {
		for _, column := range row {
			if column {
				fmt.Printf("|*|")
			} else {
				fmt.Printf("| |")
			}
		}
		fmt.Println()
		fmt.Println(line)
	}
}

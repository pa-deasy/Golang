package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

type grid [rows][columns]int8

func main() {
	var g grid
	err := g.set(10, 0, 5)
	if err != nil {
		fmt.Printf("An error occurred: %v.\n", err)
		os.Exit(1)
	}
}

func (g *grid) set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return errors.New("out of bounds")
	}
	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}

	if column < 0 || column >= columns {
		return false
	}

	return true
}

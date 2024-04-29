package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrBounds     = errors.New("out of bounds")
	ErrDigit      = errors.New("invalid digit")
	ErrRow        = errors.New("Row already contains digit")
	ErrColumn     = errors.New("Column already contains digit")
	ErrSubregion  = errors.New("Subregion already contains digit")
	ErrFixedDigit = errors.New("Digit is fixed")
)

const rows, columns = 9, 9

type grid [rows][columns]cell

type cell struct {
	value int8
	fixed bool
}

func main() {
	g := newSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	fmt.Println(g)
	fmt.Println()

	//setErr := g.set(10, 0, 1)
	//setErr := g.set(0, 0, 11)
	//setErr := g.set(0, 0, 7)
	//setErr := g.set(0, 0, 4)
	//setErr := g.set(0, 0, 9)
	//setErr := g.set(0, 0, 1)
	setErr := g.set(0, 2, 1)

	if setErr != nil {
		fmt.Println(setErr)
		os.Exit(1)
	}

	fmt.Println(g)
	fmt.Println()

	//clearErr := g.clear(10, 0)
	//clearErr := g.clear(0, 0)
	clearErr := g.clear(0, 2)

	if clearErr != nil {
		fmt.Println(clearErr)
		os.Exit(1)
	}

	fmt.Println(g)
	fmt.Println()
}

func newSudoku(layout [rows][columns]int8) grid {
	var cells grid

	for rIndex, r := range layout {
		for cIndex, c := range r {
			cells[rIndex][cIndex] = cell{value: c, fixed: c != 0}
		}
	}

	return cells
}

func (g *grid) set(row, column int, digit int8) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	case g.rowContainsDigit(row, digit):
		return ErrRow
	case g.columnContainsDigit(column, digit):
		return ErrColumn
	case g.subregionContainsDigit(row, column, digit):
		return ErrSubregion
	case g.isFixedDigit(row, column):
		return ErrFixedDigit
	}

	g[row][column] = cell{value: digit, fixed: false}
	return nil
}

func (g *grid) clear(row, column int) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case g.isFixedDigit(row, column):
		return ErrFixedDigit
	}

	g[row][column].value = 0
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

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (g *grid) rowContainsDigit(row int, digit int8) bool {
	for _, c := range g[row] {
		if c.value == digit {
			return true
		}
	}
	return false
}

func (g *grid) columnContainsDigit(column int, digit int8) bool {
	for _, c := range g {
		if c[column].value == digit {
			return true
		}
	}
	return false
}

func (g *grid) subregionContainsDigit(row, column int, digit int8) bool {
	startRow := (row / 3) * 3
	startColumn := (column / 3) * 3

	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].value == digit {
				return true
			}
		}
	}
	return false
}

func (g *grid) isFixedDigit(row, column int) bool {
	return g[row][column].fixed
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type error interface {
	Error() string
}

type sudokuError []error

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

const rows, columns = 9, 9

type grid [rows][columns]int8

func main() {
	var g grid
	err := g.set(10, 0, 11)
	if err != nil {
		if errs, ok := err.(sudokuError); ok {
			fmt.Printf("%d errors occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("-%v\n", e)
			}
		}

		os.Exit(1)
	}
}

func (g *grid) set(row, column int, digit int8) error {
	var errs sudokuError

	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}

	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}

	if len(errs) > 0 {
		return errs
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

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (se sudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}

	return strings.Join(s, ", ")
}

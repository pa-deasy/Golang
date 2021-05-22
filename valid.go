package main

import "fmt"

type number struct {
	isValid bool
	value   int
}

func main() {
	n := newNumber(42)
	fmt.Println(n)

	e := number{}
	fmt.Println(e)
}

func newNumber(v int) number {
	return number{true, v}
}

func (n number) String() string {
	if !n.isValid {
		return "not set"
	}

	return fmt.Sprintf("%d", n.value)
}

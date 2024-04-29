package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)

	for _, temp := range temperatures {
		set[temp] = true
	}

	if set[-28.0] {
		fmt.Println("set member")
	}

	fmt.Println(set)

	unique := make([]float64, 0, len(set))

	for temp := range set {
		unique = append(unique, temp)
	}

	sort.Float64s(unique)

	fmt.Println(unique)
}

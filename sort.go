package main

import (
	"fmt"
	"sort"
)

func main() {
	food := []string{"onion", "carrot", "celery", "pea"}
	sortStrings(food, nil)
	fmt.Println(food)

	sortStrings(food, func(i, j int) bool { return len(food[i]) < len(food[j]) })
	fmt.Println()
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}

	sort.Slice(s, less)
}

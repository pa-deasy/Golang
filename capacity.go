package main

import "fmt"

func main() {
	s := []string{}

	for i := 0; i < 100; i++ {
		s := append(s, "An element")
		dump(s)
	}
}

func dump(slice []string) {
	fmt.Printf("%v capacity: %v, length: %v\n", slice, cap(slice), len(slice))
}

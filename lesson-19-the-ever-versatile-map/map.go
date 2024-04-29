package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	fmt.Printf("Average temperature on Earth is: %v\n", temperature["Earth"])

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	fmt.Printf("Default value for missing Moon is: %v\n", temperature["Moon"])

	if moon, exists := temperature["Moon"]; exists {
		fmt.Printf("Average temperature on the Moon is: %v\n", moon)
	} else {
		fmt.Println("Moon does not exist")
	}
}

package main

import "fmt"

func main() {
	fn := func(message string) {
		fmt.Println(message)
	}

	fn("Go to the party")
}

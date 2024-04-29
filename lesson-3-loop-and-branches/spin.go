package main

import "fmt"

func main() {
	spin()
}

func spin() {
	var degrees = 0

	for {
		if degrees >= 360 {
			break
		}

		fmt.Println(degrees)

		degrees++
	}
}

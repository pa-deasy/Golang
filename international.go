package main

import "fmt"

func main() {
	message := "Hola Estación Espacial International"

	println("Message: ", message)

	cipher(message)
}

func cipher(message string) {
	println("Ciphered to: ")

	for index := 0; index < len(message); index++ {
		c := message[index]

		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}

		fmt.Printf("%c", c)
	}
}

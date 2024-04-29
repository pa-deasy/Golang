package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "Wedig you luv the gophers"
	key := "golang"
	cipher(message, key)
}

func cipher(message string, key string) {
	formattedMessage := strings.ToUpper(strings.ReplaceAll(message, " ", ""))
	formattedKey := strings.ToUpper(strings.ReplaceAll(key, " ", ""))
	ciphered := ""

	for index, keyIndex := 0, 0; index < len(formattedMessage); index++ {
		c := formattedMessage[index] - 'A'
		k := formattedKey[keyIndex] - 'A'

		c = (c+k)%26 + 'A'
		ciphered += string(c)

		keyIndex++
		keyIndex %= len(formattedKey)
	}

	fmt.Println("Ciphered message is: ", ciphered)
}

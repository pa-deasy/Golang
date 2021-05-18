package main

import "fmt"

func main() {
	message := "CSOITEUIWUIZNSROCNKFD"
	key := "GOLANG"
	decipher(message, key)
}

func decipher(message string, key string) {
	deciphered := ""

	for index, keyIndex := 0, 0; index < len(message); index++ {
		c := message[index] - 'A'
		k := key[keyIndex] - 'A'

		c = (c-k+26)%26 + 'A'
		deciphered += string(c)

		keyIndex++
		keyIndex %= len(key)
	}

	fmt.Println("Deciphered message is: ", deciphered)
}

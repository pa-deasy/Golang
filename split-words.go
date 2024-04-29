package main

import (
	"fmt"
	"strings"
)

func main() {

	sourceChannel := make(chan string)
	splitChannel := make(chan string)

	go sourceSentences(sourceChannel)
	go splitSentences(sourceChannel, splitChannel)
	print(splitChannel)
}

func sourceSentences(downstream chan string) {
	sentences := []string{"I am Patrick", "I like pizza", "I'm Irish"}
	for _, sentence := range sentences {
		downstream <- sentence
	}
	close(downstream)
}

func splitSentences(upstream, downstream chan string) {
	for sentence := range upstream {
		words := strings.Fields(sentence)
		for _, word := range words {
			downstream <- word
		}
	}
	close(downstream)
}

func print(upstream chan string) {
	for item := range upstream {
		fmt.Println(item)
	}
}

package main

import "strings"

func main() {
	passage := "The man, was walking away down the road. The road was long and windy!"
	wordsCount := countWordsIn(passage)

	for word, count := range wordsCount {
		println(word, count)
	}

}

func countWordsIn(passage string) map[string]int {
	words := strings.Fields(passage)
	wordsCount := make(map[string]int, len(words))

	for _, word := range words {
		formattedWord := strings.ToLower(strings.Trim(word, `,.!`))
		wordsCount[formattedWord]++
	}

	return wordsCount
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	sourceChannel := make(chan string)
	filterChannel := make(chan string)

	go sourceGopher(sourceChannel)
	go filterGopher(sourceChannel, filterChannel)
	printGopher(filterChannel)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"Hello world", "a bad apple", "goodbye"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for {
		item, opened := <-upstream
		if !opened {
			close(downstream)
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

func printGopher(upstream chan string) {
	for {
		item, opened := <-upstream
		if !opened {
			return
		}

		fmt.Println(item)
	}
}

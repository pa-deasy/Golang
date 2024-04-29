package main

import (
	"fmt"
	"strings"
)

func main() {
	sourceChannel := make(chan string)
	filterChannel := make(chan string)
	duplicateChannel := make(chan string)

	go sourceGopher(sourceChannel)
	go filterGopher(sourceChannel, filterChannel)
	go duplicatesGopher(filterChannel, duplicateChannel)
	printGopher(duplicateChannel)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"Hello world", "a bad apple", "goodbye", "goodbye"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func duplicatesGopher(upstream, downstream chan string) {
	sent := make([]string, len(upstream))
	for item := range upstream {
		if !contains(sent, item) {
			downstream <- item
			sent = append(sent, item)
		}
	}
	close(downstream)
}

func contains(s []string, v string) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}
	return false
}

func printGopher(upstream chan string) {
	for item := range upstream {
		fmt.Println(item)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func main() {
	v := visited{visited: make(map[string]int)}

	go v.visitLink("http://www.google.com")
	go v.visitLink("http://www.yahoo.com")
	go v.visitLink("http://www.bing.com")
	go v.visitLink("http://www.google.com")

	time.Sleep(1 * time.Second)

	for url, count := range v.visited {
		fmt.Println(url, "occurs", count, "times")
	}
}

func (v *visited) visitLink(url string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.visited[url]++
}

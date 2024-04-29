package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, channel)
	}

	for i := 0; i < 5; i++ {
		gopherId := <-channel
		fmt.Println("gopher ", gopherId, " has finished sleeping")
	}
}

func sleepyGopher(id int, channel chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, " snore...")
	channel <- id
}

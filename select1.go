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

	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-channel:
			fmt.Println("gopher ", gopherId, " has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}

func sleepyGopher(id int, channel chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, " snore...")
	channel <- id
}

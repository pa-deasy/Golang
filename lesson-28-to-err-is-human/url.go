package main

import (
	"fmt"
	"net/url"
	"os"
)

type address string

func main() {
	address := address("https://a b.com")

	address.validate()
}

func (a address) validate() {
	parsed, err := url.Parse(string(a))

	if err != nil {
		fmt.Println(err)
		fmt.Printf("Not a valid address: %#v\n", err)

		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op:", e.Op)
			fmt.Println("URL:", e.URL)
			fmt.Println("Err:", e.Err)
		}

		os.Exit(1)
	}

	fmt.Println(parsed)
}

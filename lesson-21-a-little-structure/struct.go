package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Latitude  float64
	Longitude float64
}

func main() {
	opportunity := location{Latitude: -1.9462, Longitude: 354.4734}
	insight := location{-1.7839, 412.8792}

	insight.Latitude = 1

	fmt.Printf("%v\n", opportunity)
	fmt.Printf("%+v\n", insight)

	opportunity.serialize()
}

func (l location) serialize() {
	bytes, err := json.Marshal(l)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes))
}

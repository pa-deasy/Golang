package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Latitude  float64 `json:"lati-tude"`
	Longitude float64 `json:"longi-tude"`
}

func main() {
	opportunity := location{Latitude: -1.9462, Longitude: 354.4734}

	fmt.Printf("%v\n", opportunity)

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

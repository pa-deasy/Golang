package main

import (
	"encoding/json"
	"fmt"
)

type location struct {
	Name string
	Lat  float64
	Long float64
}

func main() {
	locations := []location{
		{Name: "Site A", Lat: -1, Long: -2},
		{Name: "Site B", Lat: -1, Long: -2},
		{Name: "Site C", Lat: -1, Long: -2},
	}

	serialize(locations)
}

func serialize(locations []location) {
	bytes, err := json.MarshalIndent(locations, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bytes))
}

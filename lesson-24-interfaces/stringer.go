package main

import "fmt"

type location struct {
	lat  float64
	long float64
}

func main() {
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity)
}

func (l location) String() string {
	return fmt.Sprintf("%v°, %v°", l.lat, l.long)
}

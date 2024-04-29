package main

import (
	"fmt"
	"math"
)

type gps struct {
	current     location
	destination location
	world       world
}

type location struct {
	name string
	long float64
	lat  float64
}

type world struct {
	radius float64
}

type rover struct {
	gps
}

func main() {
	mars := gps{
		location{"Bradbury landing", -4.5895, 137.4417},
		location{"Elysium Planitia", 4.5, 135.9},
		world{radius: 3389.5},
	}

	curiosity := rover{mars}

	fmt.Println(curiosity.message())
}

func (l location) description() string {
	return fmt.Sprintf("%v (%.1f°, %.1f°)", l.name, l.long, l.lat)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func (gps gps) distance() float64 {
	return gps.world.distance(gps.current, gps.destination)
}

func (gps gps) message() string {
	return fmt.Sprintf("%.1f kms remaining till %v", gps.distance(), gps.destination.description())
}

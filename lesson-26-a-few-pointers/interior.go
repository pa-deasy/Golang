package main

import "fmt"

type stats struct {
	level     int
	endurance int
	health    int
}

type character struct {
	name  string
	stats stats
}

func main() {
	player := character{name: "Patrick"}

	levelUp(&player.stats)
	fmt.Printf("%+v\n", player.stats)

	levelUp(&player.stats)
	fmt.Printf("%+v\n", player.stats)
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

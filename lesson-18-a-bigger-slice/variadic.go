package main

import "fmt"

func main() {
	twoWorlds := terraform("new", "Venus", "Mars")

	fmt.Println(twoWorlds)

	planets := []string{"Venus", "Mars", "Jupiter"}
	destroyedPlanets := terraform("destroyed", planets...)

	fmt.Println(destroyedPlanets)
}

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for index := range newWorlds {
		newWorlds[index] = prefix + worlds[index]
	}

	return newWorlds
}

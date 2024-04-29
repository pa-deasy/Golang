package main

import "fmt"

type character struct {
	name     string
	leftHand *item
}

type item struct {
	description string
}

func main() {
	arthur := character{name: "Arthur"}
	knight := character{name: "Knight"}

	arthur.give(&knight)

	sword := item{description: "Majestic glowing sword"}
	arthur.pickup(&sword)
	arthur.give(&knight)
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}

	if i == nil {
		fmt.Printf("Item not found")
	} else {
		c.leftHand = i
		fmt.Printf("%v picks up %v", c.name, i.description)
	}
	fmt.Println()
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}

	if c.leftHand == nil {
		fmt.Printf("%v cannot give to %v what they do not have", c.name, to.name)
	} else {
		to.leftHand = c.leftHand
		c.leftHand = nil

		fmt.Printf("%v gives %v to %v", c.name, to.leftHand.description, to.name)
	}

	fmt.Println()
}

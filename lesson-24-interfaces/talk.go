package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct{}

type laser int

func main() {
	martian := martian{}
	laser := laser(2)

	fmt.Println(martian.talk())
	fmt.Println(shout(martian))

	fmt.Println(laser.talk())
	fmt.Println(shout(laser))
}

func (m martian) talk() string {
	return "nack nack"
}

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func shout(t talker) string {
	return strings.ToUpper(t.talk())
}

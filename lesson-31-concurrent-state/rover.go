package main

import (
	"fmt"
	"image"
	"time"
)

type roverDriver struct {
	commandc chan command
}

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

func main() {
	r := newRoverDriver()
	time.Sleep(3 * time.Second)
	r.start()
	time.Sleep(3 * time.Second)
	r.left()
	time.Sleep(3 * time.Second)
	r.stop()
	time.Sleep(3 * time.Second)
	r.start()
	time.Sleep(3 * time.Second)
	r.right()
	time.Sleep(3 * time.Second)
}

func newRoverDriver() *roverDriver {
	r := &roverDriver{commandc: make(chan command)}
	go r.drive()
	return r
}

func (r *roverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	isMoving := false
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{X: -direction.Y, Y: direction.X}
				fmt.Println("heading right")
			case left:
				direction = image.Point{X: direction.Y, Y: -direction.X}
				fmt.Println("heading left")
			case start:
				isMoving = true
				fmt.Println("starting rover")
			case stop:
				isMoving = false
				fmt.Println("stopping rover")
			}
		case <-nextMove:
			if isMoving {
				pos = pos.Add(direction)
				fmt.Printf("moved to %v\n", pos)
			} else {
				fmt.Println("rover is not moving")
			}
			nextMove = time.After(updateInterval)
		}
	}
}

func (r *roverDriver) left() {
	r.commandc <- left
}

func (r *roverDriver) right() {
	r.commandc <- right
}

func (r *roverDriver) start() {
	r.commandc <- start
}

func (r *roverDriver) stop() {
	r.commandc <- stop
}

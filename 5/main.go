package main

import (
	"fmt"
	"math"
)

var data = 277678

const RIGHT = 0
const UP = 1
const LEFT = 2
const DOWN = 3

var currentDirection = RIGHT
var currentSteps = 0

var maxSteps = 1
var directionChanges = 0

var locX float64 = 0
var locY float64 = 0

func main() {
	for i := 1; i <= data; i++ {
		switch currentDirection {
		case UP:
			locY += 1
			break
		case DOWN:
			locY -= 1
			break
		case LEFT:
			locX -= 1
			break
		case RIGHT:
			locX += 1
			break
		}

		currentSteps += 1

		if currentSteps == maxSteps {
			directionChanges += 1
			currentDirection = (currentDirection + 1) % 4
			currentSteps = 0

			if directionChanges % 2 == 0 {
				maxSteps += 1
			}
		}
	}

	fmt.Println(math.Abs(locX) + math.Abs(locY) - 1)
}

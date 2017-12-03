package main

import (
	"fmt"
)

var data = 277678

var tableLength = 20

const RIGHT = 0
const UP = 1
const LEFT = 2
const DOWN = 3

var currentDirection = RIGHT
var currentSteps = 0

var maxSteps = 1
var directionChanges = 0

var locX = 0
var locY = 0

func main() {
	var table = make([][]int, 0)

	for i := 0; i < tableLength; i++ {
		table = append(table, make([]int, tableLength))
	}

	locY = int(tableLength/2)
	locX = locY+1

	table[locY][locX - 1] = 1

	for locX <= tableLength && locY <= tableLength {
		if locY > 0 && locX > 0 {
			table[locY][locX] += table[locY - 1][locX - 1]
		}

		if locY > 0 {
			table[locY][locX] += table[locY - 1][locX]
		}

		if locY > 0 && locX < tableLength {
			table[locY][locX] += table[locY - 1][locX + 1]
		}

		if locX > 0 {
			table[locY][locX] += table[locY][locX - 1]
		}

		if locX < tableLength {
			table[locY][locX] += table[locY][locX + 1]
		}

		if locY < tableLength && locX > 0 {
			table[locY][locX] += table[locY + 1][locX - 1]
		}

		if locY < tableLength {
			table[locY][locX] += table[locY + 1][locX]
		}

		if locY < tableLength && locX < tableLength {
			table[locY][locX] += table[locY + 1][locX + 1]
		}

		if table[locY][locX] > data {
			fmt.Println(table[locY][locX])
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
	}
}

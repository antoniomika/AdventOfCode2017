package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func main() {
	file, err := os.Open("./data.txt");
	if err != nil {
		log.Fatal(err)
	}

	jumps := make([]int, 0)
	currentJump := 0
	prevNextJump := 0
	nextJump := 0
	steps := 0

	buf := bufio.NewScanner(file)

	for buf.Scan() {
		line := buf.Text()
		i, _ := strconv.Atoi(line)

		jumps = append(jumps, i)
	}

	for !exited(jumps, currentJump + nextJump) {
		prevNextJump = nextJump
		nextJump = jumps[currentJump + nextJump]
		inc := 1
		jumps[currentJump + prevNextJump] += inc
		steps += 1
		currentJump = currentJump + prevNextJump
	}

	fmt.Println(len(jumps), steps)
}

func exited(jumps []int, loc int) bool {
	if loc >= len(jumps) {
		return true
	}

	return false
}
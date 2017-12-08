package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
)

func main() {
	file, err := os.Open("./data.txt");
	if err != nil {
		log.Fatal(err)
	}

	instr := make(map[string]int, 0)

	buf := bufio.NewScanner(file)

	for buf.Scan() {
		line := buf.Text()

		fields := strings.Fields(line)

		variable := fields[0]
		function := fields[1]
		value, _ := strconv.Atoi(fields[2])
		variable2 := fields[4]
		condition := fields[5]
		value2, _ := strconv.Atoi(fields[6])

		if _, ok := instr[variable]; !ok {
			instr[variable] = 0
		}

		if _, ok := instr[variable2]; !ok {
			instr[variable2] = 0
		}

		switch(condition) {
		case "==":
			if instr[variable2] == value2 {
				instr = doFunc(function, value, variable, instr)
			}
			break
		case ">=":
			if instr[variable2] >= value2 {
				instr = doFunc(function, value, variable, instr)
			}
			break
		case "<=":
			if instr[variable2] <= value2 {
				instr = doFunc(function, value, variable, instr)
			}
			break
		case "!=":
			if instr[variable2] != value2 {
				instr = doFunc(function, value, variable, instr)
			}
			break
		case "<":
			if instr[variable2] < value2 {
				instr = doFunc(function, value, variable, instr)
			}
			break
		case ">":
			if instr[variable2] > value2 {
				instr = doFunc(function, value, variable, instr)
			}
			break
		}
	}

	m := ""
	b := 0

	for k, v := range instr {
		if v > b {
			b = v
			m = k
		}
	}

	fmt.Println(m, b)
}

func doFunc(function string, value int, variable string, instr map[string]int) map[string]int {
	switch(function) {
	case "inc":
		instr[variable] += value
		break
	case "dec":
		instr[variable] -= value
		break
	}

	return instr
}

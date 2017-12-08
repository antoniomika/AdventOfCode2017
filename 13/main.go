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

	progs := make(map[string]float64, 0)
	childToParent := make(map[string]string, 0)
	parentToChild := make(map[string][]string, 0)

	buf := bufio.NewScanner(file)

	for buf.Scan() {
		line := buf.Text()

		fields := strings.Fields(line)

		program := fields[0]
		weight := strings.Replace(fields[1], "(", "", -1)
		weight = strings.Replace(weight, ")", "", -1)

		weightInt, _ := strconv.Atoi(weight)

		progs[program] = float64(weightInt)

		childs := strings.Join(fields[2:], " ")
		childs = strings.Replace(childs, "-> ", "", -1)
		childsSlice := strings.Fields(childs)

		parentToChild[program] = childsSlice

		for _, v := range childsSlice {
			childToParent[v] = program
		}
	}

	new := make(map[string]float64, 0)

	for k, _ := range parentToChild {
		if _, ok := childToParent[k]; !ok {
			new[k] = progs[k]
		}
	}

	printMax(new)
}

func printMax(progs map[string]float64) {
	m := ""
	b := math.Inf(1)

	for k, v := range progs {
		if v < b {
			m = k
			b = v
		}
	}

	fmt.Println(m, b)
}
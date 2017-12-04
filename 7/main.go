package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

func main() {
	file, err := os.Open("./data.txt");
	if err != nil {
		log.Fatal(err)
	}

	var viable = make([]string, 0)

	buf := bufio.NewScanner(file)

	for buf.Scan() {
		line := buf.Text()
		stringSlice := strings.Split(line, " ")
		isViable := true

		for _, val := range stringSlice {
			if stringInSlice(val, stringSlice) > 1 {
				isViable = false
				break
			}
		}

		if isViable {
			viable = append(viable, line)
		}
	}

	fmt.Println(len(viable))
}

func stringInSlice(a string, list []string) int {
	count := 0
	for _, b := range list {
		if b == a {
			count += 1
		}
	}
	return count
}


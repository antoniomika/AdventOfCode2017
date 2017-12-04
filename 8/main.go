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

			for _, val2 := range stringSlice {
				if val != val2 && checkAnagrams(val, val2) {
					isViable = false
					break
				}
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

func checkAnagrams(a, b string) bool {
	letters1 := make(map[string]int, 0)
	letters2 := make(map[string]int, 0)

	for _, val := range a {
		c := string(val)
		if _, ok := letters1[c]; ok {
			letters1[c] += 1
		} else {
			letters1[c] = 1
		}
	}

	for _, val := range b {
		c := string(val)
		if _, ok := letters2[c]; ok {
			letters2[c] += 1
		} else {
			letters2[c] = 1
		}
	}

	if len(letters1) != len(letters2) {
		return false
	}

	for k, v := range letters1 {
		if val, ok := letters2[k]; ok {
			if letters1[k] == val {
				if v == len(letters1) - 1 {
					continue
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}


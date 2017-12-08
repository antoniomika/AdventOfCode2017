package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	file, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		log.Fatal(err)
	}

	cache := make([][]int, 0)
	fields := strings.Fields(string(file))

	data := make([]int, 0)

	i := 0

	for _, v := range fields {
		strint, _ := strconv.Atoi(v)
		data = append(data, strint)
	}

	fmt.Println(data)

	for !incache(cache, data) {
		newdata := make([]int, len(data))
		copy(newdata, data)
		cache = append(cache, newdata)

		maxIndex := max(data)
		maxValue := data[maxIndex]
		moveIndex := maxIndex

		for i := maxValue; i > 0; i-- {
			moveIndex = (moveIndex + 1) % len(data)

			data[moveIndex]++
			data[maxIndex]--
		}

		i++
	}

	fmt.Println(i)
}

func max(data []int) int {
	i := 0
	m := 0

	for k, v := range data {
		if v > m {
			i = k
			m = v
		}
	}

	return i
}

func incache(cache [][]int, data []int) bool {
	count := 0
	//fmt.Println(cache, data)
	for _, v := range cache {
		if equals(v, data) {
			count++
		}
	}

	if count > 0 {
		return true
	}

	return false
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func convert(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic("not a number?")
	}
	return res
}

func sliding_window(xs []int, size int) [][]int {
	var part []int
	var result [][]int
	for index, _ := range xs {
		for i := 0; i < size && i+index < len(xs); i++ {
			part = append(part, xs[i+index])
		}
		if len(part) == size {
			result = append(result, part)
		}
		part = make([]int, 0)
	}
	return result
}

func sum(xs []int) int {
	var result int
	for _, x := range xs {
		result += x
	}
	return result
}

// sliding_window([1, 2, 3, 4, 5], 3)
// -> [[1, 2, 3], [2, 3, 4], [3, 4, 5]]

func main() {
	sample, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic("bad file")
	}
	lines := strings.Split(string(sample), "\n")
	var numbers []int
	for i := 0; i < len(lines)-1; i++ {
		numbers = append(numbers, convert(lines[i]))
	}

	var count_increase int
	for index, _ := range numbers {
		if index == 0 {
			count_increase++
		} else {
			if numbers[index] > numbers[index-1] {
				count_increase++
			}
		}
	}

	count_increase = 0
	windows := sliding_window(numbers, 3)
	for index, _ := range windows {
		if index > 0 && sum(windows[index]) > sum(windows[index-1]) {
			count_increase++
		}
	}
	fmt.Println("Solution part2: ", count_increase)
}

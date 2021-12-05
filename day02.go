package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parse_move(move string) (string, int) {
	var current_move string
	var depth_value int
	move_params := strings.Fields(move)
	current_move = string(move_params[0])
	depth_value, err := strconv.Atoi(string(move_params[1]))
	if err != nil {
		panic("not a number?")
	}
	return current_move, depth_value
}

func absolute_value(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}

func main() {
	sample, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic("bad file")
	}
	lines := strings.Split(string(sample), "\n")

	var lines_tab []string
	for i := 0; i < len(lines)-1; i++ {
		lines_tab = append(lines, lines[i])
	}

	var depth int
	var depth_v2 int
	var possition int
	var possition_v2 int
	var new_depth int
	var aim int
	for i := 0; i < len(lines_tab)-2; i++ {
		move, move_value := parse_move(lines_tab[i])
		if move == "down" {
			depth += -move_value
			if depth_v2 == 0 {
				depth_v2 += absolute_value(move_value)
			}
			aim += move_value
			fmt.Printf("Moving down, aim is %v and depth is %v and possition: %v \n", aim, depth_v2, possition_v2)
		} else if move == "up" {
			depth += move_value
			aim -= absolute_value(move_value)
			fmt.Printf("Moving up, aim is %v and depth is %v and possition: %v \n", aim, depth_v2, possition_v2)
		} else if move == "forward" {
			possition += move_value
			possition_v2 += absolute_value(move_value)
			fmt.Printf("Moving forward, aim is %v and depth is %v and possition: %v \n", aim, depth_v2, possition_v2)
			if depth_v2 > 0 {
				if aim > 0 {
					depth_v2 = move_value * aim
					new_depth += depth_v2
					fmt.Printf("CHANGING AIM, now aim is %v and depth is %v and possition %v \n", aim, depth_v2, possition_v2)
					fmt.Println("The new depth is", new_depth)
				} else {
					depth_v2 = move_value * aim
					fmt.Println("Aim is 0, so will skipp depth calculations")
				}
			}
		}
	}
	exercise1_answer := absolute_value(depth * possition)
	fmt.Println("Possition is ", possition)
	fmt.Println("Depth is", depth)
	fmt.Println("AIM is", aim)
	fmt.Println("Answer for first excercise is: ", exercise1_answer)
	exercise2_answer := absolute_value(new_depth * possition_v2)
	fmt.Println("Answer for second is", exercise2_answer)
}

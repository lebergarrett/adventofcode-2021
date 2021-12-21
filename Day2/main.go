package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	commands := readInput("input.txt")

	// Create vars for storing coords
	var pos1 int   // X coordinate for part 1
	var pos2 int   // X coordinate for part 2
	var depth1 int // Y coord for part 1
	var depth2 int // Y coord for part 2
	var aim int
	// Main logic loop
	for _, command := range commands {
		words := strings.Fields(command)
		direction := words[0]
		value, _ := strconv.Atoi(words[1])

		pos1, depth1 = pt1Movement(direction, value, pos1, depth1)
		pos2, depth2, aim = pt2Movement(direction, value, pos2, depth2, aim)
	}
	fmt.Println("Part 1: What do you get if you multiply your final horizontal position by your final depth?", pos1*depth1)
	fmt.Println("Part 2: What do you get if you multiply your final horizontal position by your final depth?", pos2*depth2)
}

// opening the inputfile and transposing it to a slice
func readInput(inputfile string) []string {
	file, err := os.Open(inputfile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	file.Close()
	return output
}

func pt1Movement(direction string, value int, pos int, depth int) (int, int) {
	if direction == "forward" {
		pos += value
	} else if direction == "up" {
		depth -= value
	} else {
		depth += value
	}
	return pos, depth
}

func pt2Movement(direction string, value int, pos int, depth int, aim int) (int, int, int) {
	if direction == "forward" {
		pos += value
		depth += aim * value
	} else if direction == "up" {
		aim -= value
	} else {
		aim += value
	}
	return pos, depth, aim
}

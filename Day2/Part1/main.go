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
	var pos int
	var depth int
	// Main logic loop
	for _, command := range commands {
		words := strings.Fields(command)
		direction := words[0]
		value, _ := strconv.Atoi(words[1])

		pos, depth = movement(direction, value, pos, depth)
	}
	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?", x*y)
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

func movement(direction string, value int, pos int, depth int) (int, int) {
	if direction == "forward" {
		pos += value
	} else if direction == "up" {
		depth -= value
	} else {
		depth += value
	}
	return pos, depth
}

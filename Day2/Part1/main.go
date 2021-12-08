package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// Opening the inputfile and transposing it to a slice
	inputfile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(inputfile)
	scanner.Split(bufio.ScanLines)
	var commands []string

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	inputfile.Close()

	// Create vars for storing coords
	var x int
	var y int
	// Main logic loop
	for _, command := range commands {
		words := strings.Fields(command)
		direction := words[0]
		value, _ := strconv.Atoi(words[1])

		if direction == "forward" {
			x += value
			fmt.Print(command)
			color.New(color.FgGreen).Add(color.Bold).Print(" (Increased X Coord ", x, ")\n")
		} else if direction == "up" {
			y -= value
			fmt.Print(command)
			color.New(color.FgRed).Add(color.Bold).Print(" (Decreased Y Coord ", x, ")\n")
		} else {
			y += value
			fmt.Print(command)
			color.New(color.FgYellow).Add(color.Bold).Print(" (Increased Y Coord ", x, ")\n")
		}
	}
	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?", x*y)
}

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
	var aim int
	// Main logic loop
	for _, command := range commands {
		words := strings.Fields(command)
		direction := words[0]
		value, _ := strconv.Atoi(words[1])

		if direction == "forward" {
			x += value
			y += aim * value
			fmt.Print(command)
			color.New(color.FgYellow).Add(color.Bold).Print(" (X Coord ", x, ")")
			color.New(color.FgBlue).Add(color.Bold).Print("      (Y Coord ", y, ")\n")
		} else if direction == "up" {
			aim -= value
			fmt.Print(command)
			color.New(color.FgRed).Add(color.Bold).Print("      (Decreased aim ", aim, ")")
			color.New(color.FgMagenta).Add(color.Bold).Print(" (Aim: ", aim, ")\n")
		} else {
			aim += value
			fmt.Print(command)
			color.New(color.FgGreen).Add(color.Bold).Print("    (Increased aim ", aim, ")")
			color.New(color.FgMagenta).Add(color.Bold).Print(" (Aim: ", aim, ")\n")
		}
	}
	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?", x*y)
}

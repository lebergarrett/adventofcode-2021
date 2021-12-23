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

	part1Answer, err := CalcPart1(commands)
	ErrorCheck(err)
	part2Answer, err := CalcPart2(commands)
	ErrorCheck(err)

	fmt.Println("Part 1: What do you get if you multiply your final horizontal position by your final depth?", part1Answer)
	fmt.Println("Part 2: What do you get if you multiply your final horizontal position by your final depth?", part2Answer)
}

/*
Pass in a list of commands which contain a direction and value (i.e. "forward 2")
func returns the multiplied coordinates
*/
func CalcPart1(commands []string) (output int, err error) {
	var pos int
	var depth int

	for _, command := range commands {
		words := strings.Fields(command)
		direction := words[0]

		var value int // must be declared separate because err already created
		value, err = strconv.Atoi(words[1])

		pos, depth = Pt1Movement(direction, value, pos, depth)
	}
	return pos * depth, err
}

/*
Pass in a list of commands which contain a direction and value (i.e. "forward 2")
For this, any "up" or "down" commands only change the aim, and don't affect pos until a "forward" command is given
func returns the multiplied coordinates
*/
func CalcPart2(commands []string) (output int, err error) {
	var pos int
	var depth int
	var aim int

	for _, command := range commands {
		words := strings.Fields(command)
		direction := words[0]

		var value int
		value, err = strconv.Atoi(words[1])

		pos, depth, aim = Pt2Movement(direction, value, pos, depth, aim)
	}
	return pos * depth, err
}

/*
direction is a string for which direction the movement is in (i.e. forward, up, down)
value is how far in that direction you should move
pos is the X coordinate
depth is the Y coordinate

function returns the new coordinates
*/
func Pt1Movement(direction string, value int, pos int, depth int) (newpos, newdepth int) {
	if direction == "forward" {
		pos += value
	} else if direction == "up" {
		depth -= value
	} else {
		depth += value
	}
	return pos, depth
}

/*
direction is a string for which direction the movement is in (i.e. forward, up, down)
value is how far in that direction you should move
pos is the X coordinate
depth is the Y coordinate
aim is which way you are facing, used to compute depth

function returns the new coordinates and aim
*/
func Pt2Movement(direction string, value int, pos int, depth int, aim int) (newpos, newdepth, newaim int) {
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

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

// open the inputfile and transposing it to a slice
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

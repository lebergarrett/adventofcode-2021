package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	measurements := readInput("input.txt")

	// Vars that need to persist outside of loop
	inputlen := len(measurements) - 1
	var pt1counter int
	var pt2counter int
	var lastwindow int

	for i := range measurements {
		var window int
		var curr int
		var next int
		var nextnext int

		if i+1 > inputlen {
			curr, _ = strconv.Atoi(measurements[i])

			window = curr
			pt2counter += isLarger(window, lastwindow)
		} else if i+2 > inputlen {
			curr, _ = strconv.Atoi(measurements[i])
			next, _ = strconv.Atoi(measurements[i+1])

			window = curr + next
			pt1counter += isLarger(next, curr)
			pt2counter += isLarger(window, lastwindow)
		} else {
			curr, _ = strconv.Atoi(measurements[i])
			next, _ = strconv.Atoi(measurements[i+1])
			nextnext, _ = strconv.Atoi(measurements[i+2])

			window = curr + next + nextnext
			pt1counter += isLarger(next, curr)
			if i != 0 {
				pt2counter += isLarger(window, lastwindow)
			}
		}
		lastwindow = window
	}
	fmt.Println("How many measurements are larger than the previous measurement?", pt1counter)
	fmt.Println("How many sums are larger than the previous sum?", pt2counter)
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

// Opening the inputfile and transposing it to a slice
func readInput(inputfile string) []string {
	file, err := os.Open(inputfile)
	errorCheck(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	file.Close()
	return output
}

// Same func used to compare neighboring measurements as well as windows
func isLarger(first int, second int) (incerement int) {
	if first > second {
		incerement = 1
	} else {
		incerement = 0
	}
	return incerement
}

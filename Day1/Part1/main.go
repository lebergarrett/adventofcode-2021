package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	measurements := readInput("input.txt")

	// Main logic loop
	var counter int
	for i, measurement := range measurements {

		// if it's the first measurement it can't be larger than previous
		if i == 0 {
			fmt.Println(measurement, "(N/A - no previous measurement)")
			continue
		}

		// convert values to ints for comparison
		curr, err := strconv.Atoi(measurements[i])
		errorCheck(err)
		last, err := strconv.Atoi(measurements[i-1])
		errorCheck(err)

		var output string
		output, counter = compareNeighbors(curr, last, counter)
		fmt.Println(measurement, output)
	}
	fmt.Println("How many measurements are larger than the previous measurement?", counter)
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

func compareNeighbors(curr int, last int, counter int) (string, int) {
	var output string

	if curr > last {
		output = color.GreenString("(increased)")
		counter++
	} else if curr == last {
		output = "(no change)"
	} else {
		output = color.YellowString("(decreased)")
	}

	return output, counter
}

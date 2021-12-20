package main

// imports
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
	inputlen := len(measurements) - 1
	var counter int
	var lastwindow int
	for i := range measurements {
		var window int
		var curr int
		var next int
		var nextnext int

		if i == 0 {
			curr, _ = strconv.Atoi(measurements[i])
			next, _ = strconv.Atoi(measurements[i+1])
			nextnext, _ = strconv.Atoi(measurements[i+2])

			window = curr + next + nextnext
			println(window, "(N/A - no previous sum)")
		} else if i+1 > inputlen {
			curr, _ = strconv.Atoi(measurements[i])

			window = curr
			counter += compare(window, lastwindow)
		} else if i+2 > inputlen {
			curr, _ = strconv.Atoi(measurements[i])
			next, _ = strconv.Atoi(measurements[i+1])

			window = curr + next
			counter += compare(window, lastwindow)
		} else {
			curr, _ = strconv.Atoi(measurements[i])
			next, _ = strconv.Atoi(measurements[i+1])
			nextnext, _ = strconv.Atoi(measurements[i+2])

			window = curr + next + nextnext
			counter += compare(window, lastwindow)
		}
		lastwindow = window
	}
	fmt.Println("How many sums are larger than the previous sum?", counter)
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

func compare(window int, lastwindow int) (increment int) {
	if window > lastwindow {
		fmt.Print(window)
		//Green := color.New(color.FgGreen)
		//boldGreen := Green.Add(color.Bold)
		color.New(color.FgGreen).Add(color.Bold).Printf(" (increased)\n")
		increment = 1
	} else if window == lastwindow {
		fmt.Println(window, "(no change)")
		increment = 0
	} else {
		fmt.Println(window, "(decreased)")
		increment = 0
	}
	return increment
}

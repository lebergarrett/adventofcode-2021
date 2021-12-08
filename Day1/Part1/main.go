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
	// Opening the inputfile and transposing it to a slice
	inputfile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(inputfile)
	scanner.Split(bufio.ScanLines)
	var measurements []string

	for scanner.Scan() {
		measurements = append(measurements, scanner.Text())
	}

	inputfile.Close()

	// Main logic loop
	var counter int
	for i, measurement := range measurements {

		// if it's the first measurement it can't be larger than previous
		if i == 0 {
			fmt.Println(measurement, "(N/A - no previous measurement)")
			continue
		}

		curr, _ := strconv.Atoi(measurements[i])
		last, _ := strconv.Atoi(measurements[i-1])
		if curr > last {
			fmt.Print(measurement)
			//Green := color.New(color.FgGreen)
			//boldGreen := Green.Add(color.Bold)
			color.New(color.FgGreen).Add(color.Bold).Printf(" (increased)\n")
			counter++
		} else if curr == last {
			fmt.Println(measurement, "(no change)")
		} else {
			fmt.Println(measurement, "(decreased)")
		}
	}
	fmt.Println("How many measurements are larger than the previous measurement?", counter)
}

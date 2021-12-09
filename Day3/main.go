package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	binarynumbers := readInput("inputfile.txt")

	// Calculate number of zeros and ones in each position
	numzeros := make([]int, len(binarynumbers[0]))
	numones := make([]int, len(binarynumbers[0]))
	for _, numstr := range binarynumbers {
		for i, digit := range numstr {
			// convert ascii to num
			digit -= 48
			if digit == 0 {
				numzeros[i] += 1
			} else {
				numones[i] += 1
			}
		}
	}
	fmt.Println("Num Zeros:", numzeros)
	fmt.Println("Num Ones: ", numones)

	// Build up gamma and epsilon
	gamma := make([]int, len(binarynumbers[0]))
	epsilon := make([]int, len(binarynumbers[0]))
	for i := range numzeros {
		if numzeros[i] > numones[i] {
			gamma[i] = 0
			epsilon[i] = 1
		} else {
			gamma[i] = 1
			epsilon[i] = 0
		}
	}
	fmt.Println("Gamma:    ", gamma)
	fmt.Println("Epsilon:  ", epsilon)

	// turn slices of bits into decimal(base10) value
	var decgamma float64
	// creates 2 iterators, one that counts down and one that counts up
	for i, j := len(gamma)-1, 0; i >= 0; i, j = i-1, j+1 {
		decgamma += float64(gamma[i]) * math.Pow(2, float64(j))
	}
	fmt.Println("decgamma: ", decgamma)

	var decepsilon float64
	for i, j := len(epsilon)-1, 0; i >= 0; i, j = i-1, j+1 {
		decepsilon += float64(epsilon[i]) * math.Pow(2, float64(j))
	}
	fmt.Println("decepsilon: ", decepsilon)

	// Calculate and output power consumption
	powerconsumption := decgamma * decepsilon
	fmt.Print("What is the power consumption of the submarine? ")
	fmt.Printf("%f\n", powerconsumption)
}

// Opening the inputfile and transposing it to a slice
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

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	binarynumbers := readInput("inputfile.txt")

	numzeros, numones := zerosAndOnes(binarynumbers)
	fmt.Println("Num Zeros:", numzeros)
	fmt.Println("Num Ones: ", numones)

	gamma, epsilon := calcGammaAndEpsilon(numzeros, numones)
	fmt.Println("Gamma:    ", gamma)
	fmt.Println("Epsilon:  ", epsilon)

	decgamma := convertToBase10(gamma)
	fmt.Println("decgamma: ", decgamma)
	decepsilon := convertToBase10(epsilon)
	fmt.Println("decepsilon: ", decepsilon)

	// Calculate and trim trailing zeros for power consumption output
	powerconsumption := decgamma * decepsilon
	fmt.Println("What is the power consumption of the submarine?", strconv.FormatFloat(powerconsumption, 'f', -1, 64))
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

// Calculate number of zeros and ones in each position
func zerosAndOnes(slice []string) (numzeros []int, numones []int) {
	numzeros = make([]int, len(slice[0]))
	numones = make([]int, len(slice[0]))

	for _, numstr := range slice {
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

	return numzeros, numones
}

func calcGammaAndEpsilon(numzeros []int, numones []int) (gamma []int, epsilon []int) {
	gamma = make([]int, len(numzeros))
	epsilon = make([]int, len(numzeros))
	for i := range numzeros {
		if numzeros[i] > numones[i] {
			gamma[i], epsilon[i] = 0, 1
		} else {
			gamma[i], epsilon[i] = 1, 0
		}
	}
	return gamma, epsilon
}

// turn slices of bits into decimal(base10) value
func convertToBase10(slice []int) (dec float64) {
	// creates 2 iterators, one that counts down and one that counts up
	for i, j := len(slice)-1, 0; i >= 0; i, j = i-1, j+1 {
		dec += float64(slice[i]) * math.Pow(2, float64(j))
	}
	return dec
}

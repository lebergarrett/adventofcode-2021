package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	binarynumbers := readInput("inputfile.txt")

	numzeros, numones, _ := zerosAndOnes(binarynumbers)
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
func zerosAndOnes(slice []string) (numzeros []int, numones []int, string error) {
	if len(slice) == 0 {
		return nil, nil, errors.New("Error: empty slice passed to zerosAndOnes")
	}

	lenfirstdigit := len(slice[0])
	numzeros = make([]int, lenfirstdigit)
	numones = make([]int, lenfirstdigit)

	for _, numstr := range slice {
		if len(numstr) != lenfirstdigit {
			return nil, nil, errors.New("Error: slice with varying length digits passed to zerosAndOnes")
		}
		for i, digit := range numstr {
			// convert ascii to num
			digit -= '0'
			if digit == 0 {
				numzeros[i] += 1
			} else if digit == 1 {
				numones[i] += 1
			} else {
				return nil, nil, errors.New("Error: digit that is not a zero or one pass to zerosAndOnes")
			}
		}
	}

	return numzeros, numones, nil
}

// Calculate Gamma and Epsilon, which are opposing values
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

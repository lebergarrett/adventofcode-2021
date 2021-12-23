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
	binaryNumbers := ReadInput("inputfile.txt")

	numZeros, numOnes, _ := ZerosAndOnes(binaryNumbers)
	fmt.Println("Num Zeros:", numZeros)
	fmt.Println("Num Ones: ", numOnes)

	gamma, epsilon, _ := CalcGammaAndEpsilon(numZeros, numOnes)
	fmt.Println("Gamma:    ", gamma)
	fmt.Println("Epsilon:  ", epsilon)

	decGamma := ConvertToBase10(gamma)
	fmt.Println("decGamma: ", decGamma)
	decEpsilon := ConvertToBase10(epsilon)
	fmt.Println("decEpsilon: ", decEpsilon)

	// Calculate and trim trailing zeros for power consumption output
	powerConsumption := decGamma * decEpsilon
	fmt.Println("What is the power consumption of the submarine?", strconv.FormatFloat(powerConsumption, 'f', -1, 64))
}

// Opening the inputfile and transposing it to a slice
func ReadInput(inputfile string) []string {
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
func ZerosAndOnes(slice []string) (numZeros []int, numOnes []int, err error) {
	if len(slice) == 0 {
		return nil, nil, errors.New("Error: empty slice passed to ZerosAndOnes")
	}

	lenFirstDigit := len(slice[0])
	numZeros = make([]int, lenFirstDigit)
	numOnes = make([]int, lenFirstDigit)

	for _, numstr := range slice {
		if len(numstr) != lenFirstDigit {
			return nil, nil, errors.New("Error: slice with varying length values passed to ZerosAndOnes")
		}
		for i, digit := range numstr {
			// convert ascii to num
			digit -= '0'
			if digit == 0 {
				numZeros[i] += 1
			} else if digit == 1 {
				numOnes[i] += 1
			} else {
				return nil, nil, errors.New("Error: digit that is not a zero or one pass to ZerosAndOnes")
			}
		}
	}

	return numZeros, numOnes, nil
}

// Calculate Gamma and Epsilon, which are opposing values
func CalcGammaAndEpsilon(numZeros []int, numOnes []int) (gamma []int, epsilon []int, err error) {
	if len(numZeros) == 0 || len(numOnes) == 0 {
		return nil, nil, errors.New("Error: empty slice passed to CalcGammaAndEpsilon")
	}

	gamma = make([]int, len(numZeros))
	epsilon = make([]int, len(numZeros))
	for i := range numZeros {
		if numZeros[i] == numOnes[i] {
			return nil, nil, errors.New("Error: equal amount of ones and zeros passed to CalcGammaAndEpsilon")
		} else if numZeros[i] > numOnes[i] {
			gamma[i], epsilon[i] = 0, 1
		} else {
			gamma[i], epsilon[i] = 1, 0
		}
	}
	return gamma, epsilon, nil
}

// turn slices of bits into decimal(base10) value
func ConvertToBase10(slice []int) (dec float64) {
	// creates 2 iterators, one that counts down and one that counts up
	for i, j := len(slice)-1, 0; i >= 0; i, j = i-1, j+1 {
		dec += float64(slice[i]) * math.Pow(2, float64(j))
	}
	return dec
}

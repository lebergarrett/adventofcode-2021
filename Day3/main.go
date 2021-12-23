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

	numZeros, numOnes, err := ZerosAndOnes(binaryNumbers)
	ErrorCheck(err)
	fmt.Println("Num Zeros:", numZeros)
	fmt.Println("Num Ones: ", numOnes)

	gamma, epsilon, err := CalcGammaAndEpsilon(numZeros, numOnes)
	ErrorCheck(err)
	fmt.Println("Gamma:    ", gamma)
	fmt.Println("Epsilon:  ", epsilon)

	decGamma, err := ConvertToBase10(gamma)
	ErrorCheck(err)
	fmt.Println("decGamma: ", decGamma)

	decEpsilon, err := ConvertToBase10(epsilon)
	ErrorCheck(err)
	fmt.Println("decEpsilon: ", decEpsilon)

	// Calculate and trim trailing zeros for power consumption output
	powerConsumption := decGamma * decEpsilon
	fmt.Println("What is the power consumption of the submarine?", strconv.FormatFloat(powerConsumption, 'f', -1, 64))
}

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

// Opening the inputfile and transposing it to a slice
func ReadInput(inputfile string) []string {
	file, err := os.Open(inputfile)
	ErrorCheck(err)

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
		return nil, nil, errors.New("Empty slice passed to ZerosAndOnes")
	}

	lenFirstDigit := len(slice[0])
	numZeros = make([]int, lenFirstDigit)
	numOnes = make([]int, lenFirstDigit)

	for _, numstr := range slice {
		if len(numstr) != lenFirstDigit {
			return nil, nil, errors.New("Slice with varying length values passed to ZerosAndOnes")
		}
		for i, digit := range numstr {
			// convert ascii to num
			digit -= '0'
			if digit == 0 {
				numZeros[i] += 1
			} else if digit == 1 {
				numOnes[i] += 1
			} else {
				return nil, nil, errors.New("Value that is not a zero or one passed to ZerosAndOnes")
			}
		}
	}

	return numZeros, numOnes, nil
}

// Calculate Gamma and Epsilon, which are opposing values
func CalcGammaAndEpsilon(numZeros []int, numOnes []int) (gamma []int, epsilon []int, err error) {
	if len(numZeros) == 0 || len(numOnes) == 0 {
		return nil, nil, errors.New("Empty slice passed to CalcGammaAndEpsilon")
	}

	gamma = make([]int, len(numZeros))
	epsilon = make([]int, len(numZeros))
	for i := range numZeros {
		if numZeros[i] == numOnes[i] {
			return nil, nil, errors.New("Equal amount of ones and zeros passed to CalcGammaAndEpsilon")
		} else if numZeros[i] < 0 || numOnes[i] < 0 {
			return nil, nil, errors.New("Negative value passed to CalcGammaAndEpsilon")
		}

		if numZeros[i] > numOnes[i] {
			gamma[i], epsilon[i] = 0, 1
		} else {
			gamma[i], epsilon[i] = 1, 0
		}
	}
	return gamma, epsilon, nil
}

// turn slices of bits into decimal(base10) value
func ConvertToBase10(slice []int) (dec float64, err error) {
	if len(slice) == 0 {
		return 0, errors.New("Empty slice passed to ConvertToBase10")
	}

	// creates 2 iterators, one that counts down and one that counts up
	for i, j := len(slice)-1, 0; i >= 0; i, j = i-1, j+1 {
		if slice[i] != 0 && slice[i] != 1 {
			return 0, errors.New("Non-binary value passed to ConvertToBase10")
		}
		dec += float64(slice[i]) * math.Pow(2, float64(j))
	}
	return dec, nil
}

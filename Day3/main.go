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

	powerConsumption, err := CalcPart1(binaryNumbers)
	ErrorCheck(err)
	fmt.Println("What is the power consumption of the submarine?", strconv.FormatFloat(powerConsumption, 'f', -1, 64))

	lifeSupportRating, err := CalcPart2(binaryNumbers)
	ErrorCheck(err)
	fmt.Println("What is the life support rating of the submarine?", strconv.FormatFloat(lifeSupportRating, 'f', -1, 64))
}

func CalcPart1(binaryNumbers []string) (output float64, err error) {
	totalNumZeros := make([]int, len(binaryNumbers[0]))
	totalNumOnes := make([]int, len(binaryNumbers[0]))

	// iterate over the length of the numbers (12)
	for i := 0; i < len(binaryNumbers[0]); i++ {
		// extract all digits at this place (i.e. all first digits, all second digits)
		strDigits := ExtractDigit(binaryNumbers, i)
		numZeros, numOnes, err := ZerosAndOnes(strDigits)
		ErrorCheck(err)

		totalNumZeros[i] = numZeros
		totalNumOnes[i] = numOnes
	}

	gamma, epsilon, err := CalcGammaAndEpsilon(totalNumZeros, totalNumOnes)
	ErrorCheck(err)

	decGamma, err := ConvertToBase10(gamma)
	ErrorCheck(err)

	decEpsilon, err := ConvertToBase10(epsilon)
	ErrorCheck(err)

	powerConsumption := decGamma * decEpsilon
	return powerConsumption, err
}

func CalcPart2(binaryNumbers []string) (output float64, err error) {
	oxygenRating, err := CalcOxygenGenRating(binaryNumbers, 0)
	ErrorCheck(err)

	co2ScrubberRating, err := CalcCO2ScrubberRating(binaryNumbers, 0)
	ErrorCheck(err)

	return oxygenRating * co2ScrubberRating, err
}

func CalcOxygenGenRating(binaryNumbers []string, i int) (rating float64, err error) {
	// establish base case, when only one num is left in slice
	if len(binaryNumbers) == 1 {
		// if base case is hit, convert the string to float
		var sliceOxygenRating []int
		for _, digit := range binaryNumbers[0] {
			// convert ascii to num
			digit -= '0'
			sliceOxygenRating = append(sliceOxygenRating, int(digit))
		}
		decOxygenRating, _ := ConvertToBase10(sliceOxygenRating)
		return decOxygenRating, nil
	}
	// pull the values at the location and tally up zeros and ones
	strDigits := ExtractDigit(binaryNumbers, i)
	numZeros, numOnes, err := ZerosAndOnes(strDigits)
	ErrorCheck(err)

	// remove nums which have the lesser between zeros and ones
	if numZeros > numOnes {
		binaryNumbers = RemoveNumsContaining(binaryNumbers, i, "1")
	} else {
		binaryNumbers = RemoveNumsContaining(binaryNumbers, i, "0")
	}

	// woooo, recursion
	return CalcOxygenGenRating(binaryNumbers, i+1)
}

func CalcCO2ScrubberRating(binaryNumbers []string, i int) (rating float64, err error) {
	// establish base case, when only one num is left in slice
	if len(binaryNumbers) == 1 {
		// if base case is hit, convert the string to float
		var sliceCO2ScrubberRating []int
		for _, digit := range binaryNumbers[0] {
			// convert ascii to num
			digit -= '0'
			sliceCO2ScrubberRating = append(sliceCO2ScrubberRating, int(digit))
		}
		decCO2ScrubberRating, _ := ConvertToBase10(sliceCO2ScrubberRating)
		return decCO2ScrubberRating, nil
	}
	// pull the values at the location and tally up zeros and ones
	strDigits := ExtractDigit(binaryNumbers, i)
	numZeros, numOnes, err := ZerosAndOnes(strDigits)
	ErrorCheck(err)

	// remove nums which have the lesser between zeros and ones
	if numZeros > numOnes {
		binaryNumbers = RemoveNumsContaining(binaryNumbers, i, "0")
	} else {
		binaryNumbers = RemoveNumsContaining(binaryNumbers, i, "1")
	}

	// woooo, recursion
	return CalcCO2ScrubberRating(binaryNumbers, i+1)
}

func RemoveNumsContaining(slice []string, digit int, numStr string) (newSlice []string) {
	for _, value := range slice {
		if digit == len(value) {
			if value[digit-1:] == numStr {
				newSlice = append(newSlice, value)
			}
		} else {
			if value[digit:digit+1] == numStr {
				newSlice = append(newSlice, value)
			}
		}
	}
	return newSlice
}

func ExtractDigit(slice []string, digit int) (str string) {
	for _, value := range slice {
		if digit == len(value) {
			str += value[digit-1:]
		} else {
			str += value[digit : digit+1]
		}
	}
	return str
}

func ZerosAndOnes(binarystr string) (numZeros int, numOnes int, err error) {
	for _, digit := range binarystr {
		// convert ascii to num
		digit -= '0'
		if digit == 0 {
			numZeros += 1
		} else if digit == 1 {
			numOnes += 1
		} else {
			return 0, 0, errors.New("Value that is not a zero or one passed to ZerosAndOnes")
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

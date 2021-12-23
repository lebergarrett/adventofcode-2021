package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	strMeasurements := ReadInput("input.txt")
	intMeasurements, err := MakeInts(strMeasurements)
	ErrorCheck(err)

	Part1Answer, err := CalcPart1(intMeasurements)
	ErrorCheck(err)
	fmt.Println("How many measurements are larger than the previous measurement?", Part1Answer)

	Part2Answer, err := CalcPart2(intMeasurements)
	ErrorCheck(err)
	fmt.Println("How many sums are larger than the previous sum?", Part2Answer)
}

func CalcPart1(ints []int) (count int, err error) {
	for i := 0; i < len(ints)-1; i++ {
		if ints[i] < ints[i+1] {
			count++
		}
	}
	return count, nil
}

func CalcPart2(ints []int) (count int, err error) {
	for i := 0; i < len(ints)-3; i++ {
		window1 := ints[i] + ints[i+1] + ints[i+2]
		window2 := ints[i+1] + ints[i+2] + ints[i+3]

		if window1 < window2 {
			count++
		}
	}
	return count, nil
}

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

// opening the inputfile and transposing it to a slice
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

// turn a slice of strings into ints
func MakeInts(strSlice []string) (intSlice []int, err error) {
	intSlice = make([]int, len(strSlice))
	for i, str := range strSlice {
		intSlice[i], err = strconv.Atoi(str)
	}
	return intSlice, err
}

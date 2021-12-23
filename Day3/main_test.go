package main

import (
	"strings"
	"testing"
)

// ErrorContains checks if the error message in actual contains the text in
// expected.
//
// This is safe when actual is nil. Use an empty string for expected if you want to
// test that err is nil.
func ErrorContains(actual error, expected string) bool {
	if actual == nil {
		return expected == ""
	}
	if expected == "" {
		return false
	}
	return strings.Contains(actual.Error(), expected)
}

// Evaluate if two slices are the same
func SliceIsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestZerosAndOnes(t *testing.T) {
	tables := []struct {
		testCase      []string
		expectedZeros []int
		expectedOnes  []int
		expectedErr   string
	}{
		{[]string{"0", "0", "0", "0"}, []int{4}, []int{0}, ""},
		{[]string{"1", "0", "0", "0"}, []int{3}, []int{1}, ""},
		{[]string{"1", "0", "1", "0", "1", "1", "1", "1"}, []int{2}, []int{6}, ""},
		{[]string{"10", "00", "10", "01", "11", "11", "01", "10"}, []int{3, 4}, []int{5, 4}, ""},
		{[]string{"10", "001", "10", "01"}, nil, nil, "Slice with varying length values passed to ZerosAndOnes"},
		{[]string{}, nil, nil, "Empty slice passed to ZerosAndOnes"},
		{[]string{"string"}, nil, nil, "Value that is not a zero or one passed to ZerosAndOnes"},
		{[]string{"0", "string"}, nil, nil, "Slice with varying length values passed to ZerosAndOnes"},
		{[]string{"0", "s"}, nil, nil, "Value that is not a zero or one passed to ZerosAndOnes"},
		{[]string{"-1", "0"}, nil, nil, "Value that is not a zero or one passed to ZerosAndOnes"},
	}

	for _, table := range tables {
		zeros, ones, err := ZerosAndOnes(table.testCase)
		if !ErrorContains(err, table.expectedErr) {
			t.Errorf("Test Case (%s) was incorrect, got unexpected error: (%v), expected: (%s).", table.testCase, err, table.expectedErr)
		} else if !SliceIsEqual(zeros, table.expectedZeros) {
			t.Errorf("Test Case (%s) was incorrect, got unexpected zeros: (%d), expected: (%d).", table.testCase, zeros, table.expectedZeros)
		} else if !SliceIsEqual(ones, table.expectedOnes) {
			t.Errorf("Test Case (%s) was incorrect, got unexpected ones: (%d), expected: (%d).", table.testCase, ones, table.expectedOnes)
		}
	}
}

func TestCalcGammaAndEpsilon(t *testing.T) {
	tables := []struct {
		testCaseZeros   []int
		testCaseOnes    []int
		expectedGamma   []int
		expectedEpsilon []int
		expectedErr     string
	}{
		{[]int{1}, []int{0}, []int{0}, []int{1}, ""},
		{[]int{0}, []int{1}, []int{1}, []int{0}, ""},
		{[]int{2, 4, 6, 8}, []int{1, 5, 3, 9}, []int{0, 1, 0, 1}, []int{1, 0, 1, 0}, ""},
		{[]int{0}, []int{0}, nil, nil, "Equal amount of ones and zeros passed to CalcGammaAndEpsilon"},
		{[]int{100}, []int{100}, nil, nil, "Equal amount of ones and zeros passed to CalcGammaAndEpsilon"},
		{[]int{5, 30}, []int{10, 30}, nil, nil, "Equal amount of ones and zeros passed to CalcGammaAndEpsilon"},
		{[]int{}, []int{}, nil, nil, "Empty slice passed to CalcGammaAndEpsilon"},
		{[]int{-1}, []int{1}, nil, nil, "Negative value passed to CalcGammaAndEpsilon"},
	}

	for _, table := range tables {
		gamma, epsilon, err := CalcGammaAndEpsilon(table.testCaseZeros, table.testCaseOnes)
		if !ErrorContains(err, table.expectedErr) {
			t.Errorf("Test Case (%v)(%v) was incorrect, got unexpected error: (%v), expected: (%s).", table.testCaseZeros, table.testCaseOnes, err, table.expectedErr)
		} else if !SliceIsEqual(gamma, table.expectedGamma) {
			t.Errorf("Test Case (%v)(%v) was incorrect, got unexpected gamma: (%v), expected: (%v).", table.testCaseZeros, table.testCaseOnes, gamma, table.expectedGamma)
		} else if !SliceIsEqual(epsilon, table.expectedEpsilon) {
			t.Errorf("Test Case (%v)(%v) was incorrect, got unexpected epsilon: (%v), expected: (%v).", table.testCaseZeros, table.testCaseOnes, epsilon, table.expectedEpsilon)
		}
	}
}

func TestConvertToBase10(t *testing.T) {
	tables := []struct {
		testCase    []int
		expectedDec float64
		expectedErr string
	}{
		{[]int{1}, 1, ""},
		{[]int{0, 1}, 1, ""},
		{[]int{1, 0, 0, 0}, 8, ""},
		{[]int{1, 1, 0, 1, 0, 1, 1, 0}, 214, ""},
		{[]int{}, 0, "Empty slice passed to ConvertToBase10"},
		{[]int{2}, 0, "Non-binary value passed to ConvertToBase10"},
	}

	for _, table := range tables {
		dec, err := ConvertToBase10(table.testCase)
		if !ErrorContains(err, table.expectedErr) {
			t.Errorf("Test Case (%v) was incorrect, got unexpected error: (%v), expected: (%s).", table.testCase, err, table.expectedErr)
		} else if dec != table.expectedDec {
			t.Errorf("Test Case (%v) was incorrect, got unexpected decimal: (%v), expected: (%v).", table.testCase, dec, table.expectedDec)
		}
	}
}

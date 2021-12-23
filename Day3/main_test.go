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

// Test if two slices are the same
func testEq(a, b []int) bool {
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
		{[]string{"0", "0", "0", "1"}, []int{3}, []int{1}, ""},
		{[]string{"0", "1", "1", "0"}, []int{2}, []int{2}, ""},
		{[]string{"1", "0", "1", "0"}, []int{2}, []int{2}, ""},
		{[]string{"1", "0", "1", "0", "1", "1", "1", "1"}, []int{2}, []int{6}, ""},
		{[]string{"10", "00", "10", "01", "11", "11", "01", "10"}, []int{3, 4}, []int{5, 4}, ""},
		{[]string{"10", "001", "10", "01"}, nil, nil, "Error: slice with varying length digits passed to zerosAndOnes"},
		{[]string{}, nil, nil, "Error: empty slice passed to zerosAndOnes"},
		{[]string{"string"}, nil, nil, "Error: digit that is not a zero or one pass to zerosAndOnes"},
		{[]string{"0", "string"}, nil, nil, "Error: slice with varying length digits passed to zerosAndOnes"},
		{[]string{"0", "s"}, nil, nil, "Error: digit that is not a zero or one pass to zerosAndOnes"},
	}

	for _, table := range tables {
		zeros, ones, err := zerosAndOnes(table.testCase)
		if !ErrorContains(err, table.expectedErr) {
			t.Errorf("Test Case (%s) was incorrect, got unexpected error: (%v), expected: (%s).", table.testCase, err, table.expectedErr)
		} else if !testEq(zeros, table.expectedZeros) {
			t.Errorf("Test Case (%s) was incorrect, got unexpected zeros: (%d), expected: (%d).", table.testCase, zeros, table.expectedZeros)
		} else if !testEq(ones, table.expectedOnes) {
			t.Errorf("Test Case (%s) was incorrect, got unexpected ones: (%d), expected: (%d).", table.testCase, ones, table.expectedOnes)
		}
	}
}

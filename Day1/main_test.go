package main

import (
	"strings"
	"testing"
)

func TestCalcPart1(t *testing.T) {
	tables := []struct {
		testCase      []int
		expectedCount int
		expectedErr   string
	}{
		{[]int{1, 2, 3, 4}, 3, ""},
		{[]int{}, 0, ""},
		{[]int{1, -2, 3, -4}, 1, ""},
		{[]int{0, 0, 0, 0}, 0, ""},
		{[]int{1, 200, 3000, 40000}, 3, ""},
	}

	for _, table := range tables {
		count, err := CalcPart1(table.testCase)
		if !ErrorContains(err, table.expectedErr) {
			t.Errorf("Test Case (%v) was incorrect, got unexpected error: (%v), expected: (%s).", table.testCase, err, table.expectedErr)
		} else if count != table.expectedCount {
			t.Errorf("Test Case (%v) was incorrect, got unexpected count: (%d), expected: (%d).", table.testCase, count, table.expectedCount)
		}
	}
}

func TestCalcPart2(t *testing.T) {
	tables := []struct {
		testCase      []int
		expectedCount int
		expectedErr   string
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 5, ""},
		{[]int{}, 0, ""},
		{[]int{1, -2, 3, -4, 5, -6, 7, -8}, 2, ""},
		{[]int{0, 0, 0, 0}, 0, ""},
		{[]int{1, 200, 3000, 40000}, 1, ""},
	}

	for _, table := range tables {
		count, err := CalcPart2(table.testCase)
		if !ErrorContains(err, table.expectedErr) {
			t.Errorf("Test Case (%v) was incorrect, got unexpected error: (%v), expected: (%s).", table.testCase, err, table.expectedErr)
		} else if count != table.expectedCount {
			t.Errorf("Test Case (%v) was incorrect, got unexpected count: (%d), expected: (%d).", table.testCase, count, table.expectedCount)
		}
	}
}

func TestMakeInts(t *testing.T) {
	tables := []struct {
		testCase     []string
		expectedInts []int
		expectedErr  string
	}{
		{[]string{"0", "0", "0", "0"}, []int{0, 0, 0, 0}, ""},
		{[]string{"1", "0", "1", "0"}, []int{1, 0, 1, 0}, ""},
		{[]string{"999", "8888", "77", "6"}, []int{999, 8888, 77, 6}, ""},
		{[]string{"0", "string"}, []int{0, 0}, "strconv.Atoi: parsing \"string\": invalid syntax"},
	}

	for _, table := range tables {
		ints, err := MakeInts(table.testCase)
		if !ErrorContains(err, table.expectedErr) {
			t.Errorf("Test Case (%v) was incorrect, got unexpected error: (%v), expected: (%s).", table.testCase, err, table.expectedErr)
		} else if !SliceIsEqual(ints, table.expectedInts) {
			t.Errorf("Test Case (%v) was incorrect, got unexpected ints: (%d), expected: (%d).", table.testCase, ints, table.expectedInts)
		}
	}
}

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

package main

import (
	"strings"
	"testing"
)

func TestCalcPart1(t *testing.T) {
	tables := []struct {
		testCase        []string
		expectedProduct int
		expectedErr     string
	}{
		{[]string{"forward 0", "down 0"}, 0, ""},
		{[]string{"forward 1", "down 2"}, 2, ""},
		{[]string{"forward 15", "up 10"}, -150, ""},
		{[]string{"forward 9", "down 2", "up 1"}, 9, ""},
	}

	for _, table := range tables {
		product, err := CalcPart1(table.testCase)
		if !ErrorContains(err, table.expectedErr) {
			t.Errorf("Test Case (%v) was incorrect, got unexpected error: (%v), expected: (%s).", table.testCase, err, table.expectedErr)
		} else if product != table.expectedProduct {
			t.Errorf("Test Case (%v) was incorrect, got unexpected product: (%d), expected: (%d).", table.testCase, product, table.expectedProduct)
		}
	}
}

func TestCalcPart2(t *testing.T) {
	tables := []struct {
		testCase        []string
		expectedProduct int
		expectedErr     string
	}{
		{[]string{"forward 0", "down 0"}, 0, ""},
		{[]string{"forward 1", "down 2"}, 0, ""},
		{[]string{"forward 15", "up 10"}, 0, ""},
		{[]string{"forward 9", "down 4", "up 1", "forward 3"}, 108, ""},
	}

	for _, table := range tables {
		product, err := CalcPart2(table.testCase)
		if !ErrorContains(err, table.expectedErr) {
			t.Errorf("Test Case (%v) was incorrect, got unexpected error: (%v), expected: (%s).", table.testCase, err, table.expectedErr)
		} else if product != table.expectedProduct {
			t.Errorf("Test Case (%v) was incorrect, got unexpected product: (%d), expected: (%d).", table.testCase, product, table.expectedProduct)
		}
	}
}

func TestPt1Movement(t *testing.T) {
	tables := []struct {
		testDirection string
		testValue     int
		testPos       int
		testDepth     int
		expectedPos   int
		expectedDepth int
	}{
		{"forward", 1, 0, 0, 1, 0},
		{"forward", 5, 3, 2, 8, 2},
		{"down", 3, 0, 0, 0, 3},
		{"down", 5, 3, 2, 3, 7},
		{"up", 3, 0, 0, 0, -3},
		{"up", 5, 3, 2, 3, -3},
	}

	for _, table := range tables {
		pos, depth := Pt1Movement(table.testDirection, table.testValue, table.testPos, table.testDepth)
		if pos != table.expectedPos {
			t.Errorf("Test Case (%s %d)(pos %d)(depth %d) was incorrect, got unexpected pos: (%d), expected: (%d).", table.testDirection, table.testValue, table.testPos, table.testDepth, pos, table.expectedPos)
		} else if depth != table.expectedDepth {
			t.Errorf("Test Case (%s %d)(pos %d)(depth %d) was incorrect, got unexpected depth: (%d), expected: (%d).", table.testDirection, table.testValue, table.testPos, table.testDepth, depth, table.expectedDepth)
		}
	}
}

func TestPt2Movement(t *testing.T) {
	tables := []struct {
		testDirection string
		testValue     int
		testPos       int
		testDepth     int
		testAim       int
		expectedPos   int
		expectedDepth int
		expectedAim   int
	}{
		{"forward", 1, 0, 0, 2, 1, 2, 2},
		{"forward", 5, 3, 2, 1, 8, 7, 1},
		{"down", 3, 0, 0, 3, 0, 0, 6},
		{"down", 5, 3, 2, 0, 3, 2, 5},
		{"up", 3, 0, 0, 2, 0, 0, -1},
		{"up", 5, 3, 2, 5, 3, 2, 0},
	}

	for _, table := range tables {
		pos, depth, aim := Pt2Movement(table.testDirection, table.testValue, table.testPos, table.testDepth, table.testAim)
		if pos != table.expectedPos {
			t.Errorf("Test Case (%s %d)(pos %d)(depth %d)(aim %d) was incorrect, got unexpected pos: (%d), expected: (%d).", table.testDirection, table.testValue, table.testPos, table.testDepth, table.testAim, pos, table.expectedPos)
		} else if depth != table.expectedDepth {
			t.Errorf("Test Case (%s %d)(pos %d)(depth %d)(aim %d) was incorrect, got unexpected depth: (%d), expected: (%d).", table.testDirection, table.testValue, table.testPos, table.testDepth, table.testAim, depth, table.expectedDepth)
		} else if aim != table.expectedAim {
			t.Errorf("Test Case (%s %d)(pos %d)(depth %d)(aim %d) was incorrect, got unexpected aim: (%d), expected: (%d).", table.testDirection, table.testValue, table.testPos, table.testDepth, table.testAim, aim, table.expectedAim)
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

package main

import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatalf("%v != %v", a, b)
}

func TestLineContainsSubset(t *testing.T) {
	assertEqual(t, lineContainsSubset("2-4,6-8"), false)
	assertEqual(t, lineContainsSubset("2-3,4-5"), false)
	assertEqual(t, lineContainsSubset("5-7,7-9"), false)
	assertEqual(t, lineContainsSubset("2-8,3-7"), true)
	assertEqual(t, lineContainsSubset("6-6,4-6"), true)
	assertEqual(t, lineContainsSubset("2-6,4-8"), false)
}

func TestMain(t *testing.T) {
	lines := readFileIntoLines("input_test.txt")
	count := 0
	for _, line := range lines {
		if lineContainsSubset(line) {
			count += 1
		}
	}
	assertEqual(t, count, 2)
}

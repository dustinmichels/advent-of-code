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

// PART 1
func TestCheckLineFullyContained(t *testing.T) {
	assertEqual(t, checkLineFullyContained("2-4,6-8"), false)
	assertEqual(t, checkLineFullyContained("2-3,4-5"), false)
	assertEqual(t, checkLineFullyContained("5-7,7-9"), false)
	assertEqual(t, checkLineFullyContained("2-8,3-7"), true)
	assertEqual(t, checkLineFullyContained("6-6,4-6"), true)
	assertEqual(t, checkLineFullyContained("2-6,4-8"), false)
}

// PART 2
func TestCheckLinePartiallyContained(t *testing.T) {
	assertEqual(t, checkLinePartiallyContained("2-4,6-8"), false)
	assertEqual(t, checkLinePartiallyContained("2-3,4-5"), false)
	assertEqual(t, checkLinePartiallyContained("5-7,7-9"), true)
	assertEqual(t, checkLinePartiallyContained("2-8,3-7"), true)
	assertEqual(t, checkLinePartiallyContained("6-6,4-6"), true)
	assertEqual(t, checkLinePartiallyContained("2-6,4-8"), true)
}

func TestPart1(t *testing.T) {
	lines := readFileIntoLines("input_test.txt")
	count := 0
	for _, line := range lines {
		if checkLineFullyContained(line) {
			count += 1
		}
	}
	assertEqual(t, count, 2)
}

func TestPart2(t *testing.T) {
	lines := readFileIntoLines("input_test.txt")
	count := 0
	for _, line := range lines {
		if checkLinePartiallyContained(line) {
			count += 1
		}
	}
	assertEqual(t, count, 4)
}

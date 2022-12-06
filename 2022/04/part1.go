package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Low  int
	High int
}

// Create range from string. Eg, "2-8" -> Range{2: 8}
func RangeFromString(rangeStr string) Range {
	lowStr, highStr, _ := strings.Cut(rangeStr, "-")
	low, _ := strconv.Atoi(strings.TrimSpace(lowStr))
	high, _ := strconv.Atoi(strings.TrimSpace(highStr))
	return Range{low, high}
}

// Check if range A fully contains range B
func Contains(a Range, b Range) bool {
	if a.Low <= b.Low && a.High >= b.High {
		return true
	}
	return false
}

// Check if a given line features one range
// that fully contains the other range
//
//	Eg, "2-8, 3-7" -> true
func lineContainsSubset(line string) bool {
	first, second, _ := strings.Cut(line, ",")
	r1 := RangeFromString(first)
	r2 := RangeFromString(second)
	return Contains(r1, r2) || Contains(r2, r1)
}

func readFileIntoLines(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Could not read file: %s", filename))
	}
	return strings.Split(string(b), "\n")
}

func main() {
	// Process args
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "input.txt")
		return
	}
	filename := os.Args[1]

	// Read file into slice of lines
	lines := readFileIntoLines(filename)

	// Count the lines containing a complete subset
	count := 0
	for _, line := range lines {
		if lineContainsSubset(line) {
			count += 1
		}
	}
	fmt.Println("Count:", count)
}

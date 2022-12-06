package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------- COMMON ----------

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

func readFileIntoLines(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Could not read file: %s", filename))
	}
	return strings.Split(string(b), "\n")
}

// ---------- PART 1 ----------

// Check if range A fully contains range B
func fullyContained(a Range, b Range) bool {
	if a.Low <= b.Low && a.High >= b.High {
		return true
	}
	return false
}

// Check if a given line features one range
// that fully contains the other range
//
// Eg, "2-8, 3-7" -> true
func checkLineFullyContained(line string) bool {
	first, second, found := strings.Cut(line, ",")
	if !found {
		return false
	}
	r1 := RangeFromString(first)
	r2 := RangeFromString(second)
	return fullyContained(r1, r2) || fullyContained(r2, r1)
}

// ---------- PART 2 ----------

// Check if range A partially contains range B
func partiallyContained(a Range, b Range) bool {
	// is b's low point in range?
	if (b.Low <= a.High) && (b.Low >= a.Low) {
		return true
	}
	// is b's high point in range?
	if (b.High <= a.High) && (b.High >= a.Low) {
		return true
	}
	return false
}

// Check if a given line features on range
// that partially contains the other range
//
// Eg, "5-7, 7-9" -> true
func checkLinePartiallyContained(line string) bool {
	first, second, found := strings.Cut(line, ",")
	if !found {
		return false
	}
	r1 := RangeFromString(first)
	r2 := RangeFromString(second)
	return partiallyContained(r1, r2) || partiallyContained(r2, r1)
}

// ---------- MAIN ----------

func main() {

	// Process args
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "input.txt")
		return
	}
	filename := os.Args[1]

	// Read file into lines
	lines := readFileIntoLines(filename)

	// PART 1
	countPart1 := 0
	for _, line := range lines {
		if checkLineFullyContained(line) {
			countPart1 += 1
		}
	}
	fmt.Println("Pt1 count:", countPart1)

	// PART 2
	countPart2 := 0
	for _, line := range lines {
		if checkLinePartiallyContained(line) {
			countPart2 += 1
		}
	}
	fmt.Println("Pt2 count:", countPart2)
}

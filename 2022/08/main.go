package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func parseInput(inp string) [][]int {
	rows := strings.Split(inp, "\n")
	grid := make([][]int, len(rows))
	for i := range grid {
		ary := parseRow(rows[i])
		grid[i] = ary
	}
	return grid
}

func parseRow(row string) []int {
	strArray := strings.Split(row, "")
	intArray := make([]int, len(strArray))
	for i := range intArray {
		intArray[i], _ = strconv.Atoi(strArray[i])
	}
	return intArray
}

// ----------------------------------------

func makeBoolGrid(grid [][]int) [][]bool {
	boolGrid := make([][]bool, len(grid))
	for i, row := range grid {
		boolGrid[i] = make([]bool, len(row))
	}
	return boolGrid
}

// ----------------------------------------

func checkLeft(grid [][]int) [][]bool {
	boolGrid := makeBoolGrid(grid)
	for i, row := range grid {
		biggestTree := -1
		for j, val := range row {
			if val > biggestTree {
				boolGrid[i][j] = true
				biggestTree = val
			}
		}
	}
	return boolGrid
}

func checkTop(grid [][]int) [][]bool {
	boolGrid := makeBoolGrid(grid)
	// iter cols
	for j := 0; j < len(grid[0]); j++ {
		biggestTree := -1
		for i := 0; i < len(grid); i++ {
			val := grid[i][j]
			if val > biggestTree {
				boolGrid[i][j] = true
				biggestTree = val
			}
		}
	}
	return boolGrid
}

func checkRight(grid [][]int) [][]bool {
	boolGrid := makeBoolGrid(grid)
	width := len(boolGrid[0])
	// fmt.Println(width)
	// height := len(boolGrid)

	for i := range grid {
		biggestTree := -1
		for j := width - 1; j >= 0; j-- {
			val := grid[i][j]
			if val > biggestTree {
				boolGrid[i][j] = true
				biggestTree = val
			}
		}
	}
	return boolGrid
}

func checkBottom(grid [][]int) [][]bool {
	boolGrid := makeBoolGrid(grid)

	width := len(boolGrid[0])
	height := len(boolGrid)

	// iter cols
	for j := 0; j < width; j++ {
		biggestTree := -1
		for i := height - 1; i >= 0; i-- {
			// fmt.Println(i, j)

			val := grid[i][j]
			if val > biggestTree {
				boolGrid[i][j] = true
				biggestTree = val
			}
		}
	}

	return boolGrid
}

// ----------------------------------------

func main() {

	inpB, err := os.ReadFile("input_test.txt")
	// inpB, err := os.ReadFile("input.txt")

	if err != nil {
		panic("Can't find file")
	}

	inp := string(inpB)

	grid := parseInput(inp)

	s.MaxPerLine = 1
	// s.Show("grid", grid)

	// width := len(grid[0])
	// height := len(grid)

	leftGrid := checkLeft(grid)
	topGrid := checkTop(grid)
	rightGrid := checkRight(grid)
	bottomGrid := checkBottom(grid)

	finalBoolGrid := makeBoolGrid(grid)
	count := 0
	for i, row := range finalBoolGrid {
		for j := range row {
			res := leftGrid[i][j] || topGrid[i][j] || rightGrid[i][j] || bottomGrid[i][j]
			finalBoolGrid[i][j] = res
			if res {
				count++
			}
		}
	}

	// s.Show("finalBoolGrid", finalBoolGrid)

	fmt.Println(count)
}

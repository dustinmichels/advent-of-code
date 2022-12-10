package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	X int
	Y int
}

var SEEN = make(map[string]int)

func updateTail(head Pos, tail *Pos) {
	updateTailPos(head, tail)
	updateSeen(*tail)
}

func updateSeen(tail Pos) {
	posStr := fmt.Sprintf("%d%d", tail.X, tail.Y)
	SEEN[posStr] = 1
}

func updateTailPos(head Pos, tail *Pos) {

	// ----- move directly? -----

	// 2 up?
	if head.Y > tail.Y+1 && head.X == tail.X {
		tail.Y++
		return
	}

	// 2 down?
	if head.Y < tail.Y-1 && head.X == tail.X {
		tail.Y--
		return
	}

	// 2 left?
	if head.X < tail.X-1 && head.Y == tail.Y {
		tail.X--
		return
	}

	// 2 right?
	if head.X > tail.X+1 && head.Y == tail.Y {
		tail.X++
		return
	}

	// ----- move diagonally? -----

	// up
	if head.Y > tail.Y+1 {
		tail.Y++
		if head.X > tail.X {
			tail.X++
		}
		if head.X < tail.X {
			tail.X--
		}
		return
	}

	// down
	if head.Y < tail.Y-1 {
		tail.Y--
		if head.X > tail.X {
			tail.X++
		}
		if head.X < tail.X {
			tail.X--
		}
		return
	}

	// left
	if head.X < tail.X-1 {
		tail.X--
		if head.Y > tail.Y {
			tail.Y++
		}
		if head.Y < tail.Y {
			tail.Y--
		}
		return
	}

	// right
	if head.X > tail.X+1 {
		tail.X++
		if head.Y > tail.Y {
			tail.Y++
		}
		if head.Y < tail.Y {
			tail.Y--
		}
		return
	}

}

func main() {

	head := Pos{0, 0}
	tail := Pos{0, 0}

	dat, _ := os.ReadFile("input.txt")
	inp := string(dat)

	lines := strings.Split(inp, "\n")

	for _, line := range lines {

		dir, numS, _ := strings.Cut(line, " ")
		num, _ := strconv.Atoi(numS)

		if dir == "R" {
			for i := 0; i < num; i++ {
				head.X++
				updateTail(head, &tail)
			}
		}
		if dir == "L" {
			for i := 0; i < num; i++ {
				head.X--
				updateTail(head, &tail)
			}
		}
		if dir == "U" {
			for i := 0; i < num; i++ {
				head.Y++
				updateTail(head, &tail)
			}
		}
		if dir == "D" {
			for i := 0; i < num; i++ {
				head.Y--
				updateTail(head, &tail)
			}
		}
	}

	fmt.Println(len(SEEN))

}

package main

import (
	"fmt"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

const (
	right = 'R'
	up    = 'U'
	left  = 'L'
	down  = 'D'
)

type knot struct {
	x, y    int
	visited map[string]bool
}

type rope struct {
	knots []knot
}

func newRope(size int) *rope {
	r := new(rope)
	r.knots = make([]knot, size)
	for i := 0; i < size; i++ {
		r.knots[i].visited = make(map[string]bool)
		r.knots[i].visited["0:0"] = true
	}

	return r
}

func (r *rope) move(direction rune, steps int) {
	for i := 0; i < steps; i++ {
		switch direction {
		case right:
			r.knots[0].x++
		case left:
			r.knots[0].x--
		case up:
			r.knots[0].y++
		case down:
			r.knots[0].y--
		}

		for i := 1; i < len(r.knots); i++ {
			head := &r.knots[i-1]
			tail := &r.knots[i]
			if abs(head.x-tail.x) > 1 || abs(head.y-tail.y) > 1 {
				tail.x += step(head.x, tail.x)
				tail.y += step(head.y, tail.y)
				tail.visited[fmt.Sprintf("%d:%d", tail.x, tail.y)] = true
			}
		}
	}
}

func main() {
	lines := reader.ReadFile(flags.File())

	if flags.Part() == 1 {
		partOne(lines)
	} else {
		partTwo(lines)
	}
}

func partOne(input []string) {
	rope := newRope(2)
	for _, l := range input {
		var direction rune
		var steps int
		fmt.Sscanf(l, "%c %d", &direction, &steps)
		rope.move(direction, steps)
	}

	fmt.Print(len(rope.knots[1].visited))
}

func partTwo(input []string) {
	rope := newRope(10)
	for _, l := range input {
		var direction rune
		var steps int
		fmt.Sscanf(l, "%c %d", &direction, &steps)
		rope.move(direction, steps)
	}

	fmt.Print(len(rope.knots[9].visited))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func step(head, tail int) int {
	if head == tail {
		return 0
	}

	if head < tail {
		return -1
	}

	return 1
}

package main

import (
	"fmt"
	"strings"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

const (
	rock = '#'
	air  = '.'
	sand = 'o'

	start = 500
)

type coordinates struct {
	x, y int
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
	m, lastRock := getMap(input)
	count := 0

	curr := coordinates{start, 0}
	for curr.y <= lastRock {
		curr.y++
		if m[curr.y][curr.x] == air {
			continue
		}
		if m[curr.y][curr.x-1] == air {
			curr.x--
			continue
		}
		if m[curr.y][curr.x+1] == air {
			curr.x++
			continue
		}

		m[curr.y-1][curr.x] = sand
		count++

		curr = coordinates{start, 0}
	}

	fmt.Print(count)
}

func partTwo(input []string) {
	m, lastRock := getMap(input)
	count := 0

	floor := lastRock + 2
	for i := 0; i < 1000; i++ {
		m[floor][i] = rock
	}

	curr := coordinates{start, 0}
	for m[0][start] != sand {
		curr.y++
		if m[curr.y][curr.x] == air {
			continue
		}
		if m[curr.y][curr.x-1] == air {
			curr.x--
			continue
		}
		if m[curr.y][curr.x+1] == air {
			curr.x++
			continue
		}

		m[curr.y-1][curr.x] = sand
		count++

		curr = coordinates{start, 0}
	}

	fmt.Print(count)
}

func getMap(input []string) ([]map[int]rune, int) {
	m := make([]map[int]rune, 200)
	for i := 0; i < 200; i++ {
		m[i] = make(map[int]rune, 1000)
		for j := 0; j < 1000; j++ {
			m[i][j] = air
		}
	}

	drawLine := func(from, to *coordinates) {
		switch {
		case from.x < to.x:
			for i := from.x; i <= to.x; i++ {
				m[from.y][i] = rock
			}
		case from.x > to.x:
			for i := from.x; i >= to.x; i-- {
				m[from.y][i] = rock
			}
		case from.y < to.y:
			for i := from.y; i <= to.y; i++ {
				m[i][from.x] = rock
			}
		case from.y > to.y:
			for i := from.y; i >= to.y; i-- {
				m[i][from.x] = rock
			}
		}
	}

	maxY := 0

	for _, line := range input {
		l := strings.Split(line, " -> ")
		var prev, curr *coordinates
		for _, c := range l {
			curr = &coordinates{}
			fmt.Sscanf(c, "%d,%d", &curr.x, &curr.y)
			if prev != nil {
				drawLine(prev, curr)
			}
			if curr.y > maxY {
				maxY = curr.y
			}

			prev = curr
		}
	}

	return m, maxY
}

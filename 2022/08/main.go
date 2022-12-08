package main

import (
	"fmt"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

func main() {
	fileName := flags.File()
	part := flags.Part()

	lines := reader.ReadFile(fileName)

	if part == 1 {
		partOne(lines)
	} else {
		partTwo(lines)
	}
}

func partOne(input []string) {
	lenI := len(input)
	lenJ := len(input[0])
	count := (lenI + lenJ - 2) * 2

	for i := 1; i < lenI-1; i++ {
		for j := 1; j < lenJ-1; j++ {
			if isVisible(input, i, j) {
				count++
			}
		}
	}

	fmt.Print(count)
}

func partTwo(input []string) {
	lenI := len(input)
	lenJ := len(input[0])
	score := 0

	for i := 1; i < lenI-1; i++ {
		for j := 1; j < lenJ-1; j++ {
			s := getScore(input, i, j)
			if s > score {
				score = s
			}
		}
	}

	fmt.Print(score)
}

func isVisible(grid []string, i, j int) bool {
	checkLeft := func(grid []string, i, j int) bool {
		for k := i - 1; k >= 0; k-- {
			if grid[k][j] >= grid[i][j] {
				return false
			}
		}
		return true
	}

	checkTop := func(grid []string, i, j int) bool {
		for k := j - 1; k >= 0; k-- {
			if grid[i][k] >= grid[i][j] {
				return false
			}
		}
		return true
	}

	checkRight := func(grid []string, i, j int) bool {
		for k := i + 1; k < len(grid[i]); k++ {
			if grid[k][j] >= grid[i][j] {
				return false
			}
		}
		return true
	}

	checkBottom := func(grid []string, i, j int) bool {
		for k := j + 1; k < len(grid); k++ {
			if grid[i][k] >= grid[i][j] {
				return false
			}
		}
		return true
	}

	return checkLeft(grid, i, j) || checkTop(grid, i, j) || checkRight(grid, i, j) || checkBottom(grid, i, j)
}

func getScore(grid []string, i, j int) int {
	left := func(grid []string, i, j int) int {
		count := 0
		for k := i - 1; k >= 0; k-- {
			count++
			if grid[k][j] >= grid[i][j] {
				break
			}
		}
		return count
	}

	top := func(grid []string, i, j int) int {
		count := 0
		for k := j - 1; k >= 0; k-- {
			count++
			if grid[i][k] >= grid[i][j] {
				break
			}
		}
		return count
	}

	right := func(grid []string, i, j int) int {
		count := 0
		for k := i + 1; k < len(grid[i]); k++ {
			count++
			if grid[k][j] >= grid[i][j] {
				break
			}
		}
		return count
	}

	bottom := func(grid []string, i, j int) int {
		count := 0
		for k := j + 1; k < len(grid); k++ {
			count++
			if grid[i][k] >= grid[i][j] {
				break
			}
		}
		return count
	}

	return left(grid, i, j) * top(grid, i, j) * right(grid, i, j) * bottom(grid, i, j)
}

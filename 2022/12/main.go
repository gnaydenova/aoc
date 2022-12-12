package main

import (
	"fmt"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

func main() {
	lines := reader.ReadFile(flags.File())

	if flags.Part() == 1 {
		partOne(lines)
	} else {
		partTwo(lines)
	}
}

func partOne(input []string) {
	start, end := "", ""
	edges := make(map[string][]string)

	for i, line := range input {
		for j, h := range line {
			pos := fmt.Sprintf("%d:%d", i, j)
			edges[pos] = getEdges(input, i, j)
			if h == 'S' {
				start = pos
			} else if h == 'E' {
				end = pos
			}
		}
	}

	fmt.Print(getSteps(start, end, edges))
}

func partTwo(input []string) {
	end := ""
	var possibleStart []string 
	edges := make(map[string][]string)

	for i, line := range input {
		for j, h := range line {
			pos := fmt.Sprintf("%d:%d", i, j)
			edges[pos] = getEdges(input, i, j)
			if h == 'S' || h == 'a' {
				possibleStart = append(possibleStart, pos)
			} else if h == 'E' {
				end = pos
			}
		}
	}

	min := 0

	for _, s := range possibleStart {
		steps := getSteps(s, end, edges) 
		if min == 0 || (steps != -1 && min > steps) {
			min = steps 
		}
	}

	fmt.Print(min)
}

func getEdges(input []string, i, j int) []string {
	var edges []string
	curr := input[i][j]

	if curr == 'E' {
		return edges
	}

	if curr == 'S' {
		curr = 'a'
	}

	canStep := func(next byte) bool {
		if next == 'E' {
			next = 'z'
		}

		return int(next)-int(curr) <= 1
	}

	if i > 0 && canStep(input[i-1][j]) {
		edges = append(edges, fmt.Sprintf("%d:%d", i-1, j))
	}

	if i < len(input)-1 && canStep(input[i+1][j]) {
		edges = append(edges, fmt.Sprintf("%d:%d", i+1, j))
	}

	if j > 0 && canStep(input[i][j-1]) {
		edges = append(edges, fmt.Sprintf("%d:%d", i, j-1))
	}

	if j < len(input[i])-1 && canStep(input[i][j+1]) {
		edges = append(edges, fmt.Sprintf("%d:%d", i, j+1))
	}

	return edges
}

func getSteps(start, end string, edges map[string][]string) int {
	visited := make(map[string]int)
	var q []string
	q = append(q, start)
	visited[start] = 0

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, edge := range edges[curr] {
			if edge == end {
				return visited[curr] + 1
			}
			if _, ok := visited[edge]; !ok {
				visited[edge] = visited[curr] + 1
				q = append(q, edge)
			}
		}
	}

	return -1
}

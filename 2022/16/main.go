package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

type valve struct {
	name      string
	rate      int
	neighbors []string
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
	all := getValves(input)
	graph := floydWarshall(all)

	var start *valve
	var valves []*valve
	for _, v := range all {
		if v.name == "AA" {
			start = v
		}
		if v.rate > 0 {
			valves = append(valves, v)
		}
	}

	const timeLimit = 30

	var maxPressure func(minute int, curr *valve, opened []string, pressure int) int
	maxPressure = func(minute int, curr *valve, opened []string, pressure int) int {
		max := pressure
		for _, v := range valves {
			if v == curr || contains(opened, v.name) {
				continue
			}

			dist := graph[curr][v] + 1
			if minute+dist > timeLimit {
				continue
			}

			p := pressure + (timeLimit-minute-dist)*v.rate
			if next := maxPressure(minute+dist, v, append(opened, v.name), p); next > max {
				max = next
			}
		}

		return max
	}

	fmt.Print(maxPressure(0, start, []string{}, 0))
}

func partTwo(input []string) {
	all := getValves(input)
	graph := floydWarshall(all)

	var start *valve
	var valves []*valve
	for _, v := range all {
		if v.name == "AA" {
			start = v
		}
		if v.rate > 0 {
			valves = append(valves, v)
		}
	}

	const timeLimit = 26

	var maxPressure func(minute, elMin int, curr, elCurr *valve, opened []string, pressure int) int
	maxPressure = func(minute, elMin int, curr, elCurr *valve, opened []string, pressure int) int {
		max := pressure
		for _, v := range valves {
			if contains(opened, v.name) {
				continue
			}
			dist := graph[curr][v] + 1
			if minute+dist > timeLimit {
				continue
			}

			for _, v2 := range valves {
				if v == v2 || contains(opened, v2.name) {
					continue
				}

				elDist := graph[elCurr][v2] + 1
				if elMin+elDist > timeLimit {
					continue
				}
				p2 := pressure + (timeLimit-minute-dist)*v.rate + (timeLimit-elMin-elDist)*v2.rate
				if next := maxPressure(minute+dist, elMin+elDist, v, v2, append(opened, v.name, v2.name), p2); next > max {
					max = next
				}
			}
		}

		return max
	}

	fmt.Print(maxPressure(0, 0, start, start, []string{}, 0))
}

func floydWarshall(valves []*valve) map[*valve]map[*valve]int {
	graph := make(map[*valve]map[*valve]int)

	for _, v1 := range valves {
		graph[v1] = make(map[*valve]int)
		for _, v2 := range valves {
			if v1 == v2 {
				graph[v1][v2] = 0
				continue
			}

			if contains(v1.neighbors, v2.name) {
				graph[v1][v2] = 1
				continue
			}

			graph[v1][v2] = math.MaxInt
		}
	}

	for _, k := range valves {
		for _, i := range valves {
			for _, j := range valves {
				if graph[i][k] == math.MaxInt || graph[k][j] == math.MaxInt {
					continue
				}

				if graph[i][j] > graph[i][k]+graph[k][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
				}
			}
		}
	}

	return graph
}

func getValves(input []string) []*valve {
	valves := make([]*valve, len(input))

	for i, line := range input {
		v := &valve{}

		l := strings.Split(line, ";")
		fmt.Sscanf(l[0], "Valve %s has flow rate=%d", &v.name, &v.rate)

		v.neighbors = strings.Split(strings.TrimLeft(l[1], "tunnels lead to valves"), ", ")

		valves[i] = v
	}

	return valves
}

func contains(valves []string, name string) bool {
	for _, n := range valves {
		if n == name {
			return true
		}
	}
	return false
}

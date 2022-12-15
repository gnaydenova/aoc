package main

import (
	"flag"
	"fmt"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

type coordinates struct {
	x, y int
}

type sensor struct {
	pos      coordinates
	beacon   coordinates
	distance int
}

func (s sensor) getCoverage() []coordinates {
	var coverage []coordinates
	max := s.distance + 1
	for x := 0; x <= max; x++ {
		y := max - x
		coverage = append(
			coverage,
			coordinates{s.pos.x + x, s.pos.y + y},
			coordinates{s.pos.x - y, s.pos.y + x},
			coordinates{s.pos.x + y, s.pos.y - x},
			coordinates{s.pos.x - x, s.pos.y - y},
		)
	}

	return coverage
}

func main() {
	var row int
	flag.IntVar(&row, "row", 10, "count for row") // 2000000
	lines := reader.ReadFile(flags.File())

	if flags.Part() == 1 {
		partOne(lines, row)
	} else {
		partTwo(lines, row*2)
	}
}

func partOne(input []string, row int) {
	sensors := getSensors(input)

	var max, min int
	set := false

	for _, s := range sensors {
		currMin := s.pos.x - s.distance
		if !set || min > currMin {
			min = currMin
			set = true
		}

		currMax := s.pos.x + s.distance
		if max < currMax {
			max = currMax
		}
	}

	count := 0
	for x := min; x <= max; x++ {
		for _, s := range sensors {
			if s.beacon.x == x && s.beacon.y == row {
				break
			}

			if getManhattanDistance(coordinates{x, row}, s.pos) <= s.distance {
				count++
				break
			}
		}
	}

	fmt.Print(count)
}

func partTwo(input []string, max int) {
	const multiplier = 4000000

	sensors := getSensors(input)

	frequency := 0
	for _, s := range sensors {
		for _, possiblePos := range s.getCoverage() {
			if possiblePos.x < 0 || possiblePos.x > max || possiblePos.y < 0 || possiblePos.y > max {
				continue
			}

			isBeacon := true
			for _, s2 := range sensors {
				if getManhattanDistance(possiblePos, s2.pos) <= s2.distance {
					isBeacon = false
					break
				}
			}

			if isBeacon {
				frequency = possiblePos.x*multiplier + possiblePos.y
				break
			}
		}

		if frequency != 0 {
			break
		}
	}

	fmt.Print(frequency)
}

func getSensors(input []string) []sensor {
	var sensors []sensor
	for _, line := range input {
		var s sensor
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.pos.x, &s.pos.y, &s.beacon.x, &s.beacon.y)
		s.distance = getManhattanDistance(s.pos, s.beacon)
		sensors = append(sensors, s)
	}

	return sensors
}

func getManhattanDistance(c1, c2 coordinates) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

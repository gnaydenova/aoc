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
	cycle := 20
	strength, i, x := 0, 1, 1

	tick := func() {
		if i == cycle {
			strength += i * x
			cycle += 40
		}

		i++
	}

	for _, line := range input {
		tick()

		if line == "noop" {
			continue
		}

		val := 0
		fmt.Sscanf(line, "addx %d", &val)
		tick()

		x += val
	}

	fmt.Print(strength)
}

func partTwo(input []string) {
	i, x := 1, 1
	var crt []string
	var row string
	sprite := 0

	tick := func() {
		if x-1 <= sprite && sprite <= x+1 {
			row += "#"
		} else {
			row += "."
		}

		sprite++

		if sprite == 40 {
			crt = append(crt, row)
			sprite = 0
			row = ""
		}

		i++
	}

	for _, line := range input {
		tick()

		if line == "noop" {
			continue
		}

		val := 0
		fmt.Sscanf(line, "addx %d", &val)
		tick()

		x += val
	}

	for _, l := range crt {
		fmt.Println(l)
	}
}

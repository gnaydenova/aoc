package main

import (
	"fmt"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

type elfRange struct {
	start int
	end   int
}

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
	contains := func(elfOne, elfTwo elfRange) bool {
		if elfOne.start <= elfTwo.start && elfTwo.end <= elfOne.end {
			return true
		}

		return false
	}

	count := 0

	for _, line := range input {
		var elfOne, elfTwo elfRange
		fmt.Sscanf(line, "%d-%d,%d-%d", &elfOne.start, &elfOne.end, &elfTwo.start, &elfTwo.end)
		if contains(elfOne, elfTwo) || contains(elfTwo, elfOne) {
			count++
		}
	}

	fmt.Print(count)
}

func partTwo(input []string) {
	count := 0

	for _, line := range input {
		var elfOne, elfTwo elfRange
		fmt.Sscanf(line, "%d-%d,%d-%d", &elfOne.start, &elfOne.end, &elfTwo.start, &elfTwo.end)
		if elfOne.start <= elfTwo.end && elfTwo.start <= elfOne.end {
			count++
		}
	}

	fmt.Print(count)
}

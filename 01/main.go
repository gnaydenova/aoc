package main

import (
	"fmt"
	"strconv"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

func main() {
	fileName := flags.File()
	measurements := reader.ReadFile(fileName)

	partOne(measurements)
	partTwo(measurements)
}

func partOne(measurements []string) {
	count := 0

	var prev *int
	for _, line := range measurements {
		curr, _ := strconv.Atoi(line)
		if prev != nil && curr > *prev {
			count++
		}
		prev = &curr
	}

	fmt.Println(count)
}

func partTwo(measurements []string) {
	count := 0

	var prev *int
	for i := 0; i < len(measurements); i++ {
		curr := sum(measurements[i : i+3])
		if prev != nil && curr > *prev {
			count++
		}
		prev = &curr
	}

	fmt.Println(count)
}

func sum(measurements []string) int {
	sum := 0
	for _, line := range measurements {
		i, _ := strconv.Atoi(line)
		sum += i
	}
	return sum
}

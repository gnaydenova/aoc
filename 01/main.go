package main

import (
	"fmt"
	"log"
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
	prev := 0

	for _, line := range measurements {
		curr, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		if prev > 0 && curr > prev {
			count++
		}

		prev = curr
	}

	fmt.Println(count)
}

func partTwo(measurements []string) {
	count := 0
	prev := 0

	for i := 0; i < len(measurements); i++ {
		curr := sum(measurements[i : i+3])
		if prev > 0 && curr > prev {
			count++
		}
		prev = curr
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

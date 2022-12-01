package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

func main() {
	fileName := flags.File()
	lines := reader.ReadFile(fileName)

	calories := []int{0}

	for i, j := 0, 0; i < len(lines); i++ {
		if lines[i] == "" {
			calories = append(calories, 0)
			j ++
			
			continue
		}

		c, _ := strconv.Atoi(lines[i])
		calories[j] += c
	}

	sort.Ints(calories)	

	partOne(calories)
	partTwo(calories)
}

// In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: 
// they'd like to know how many Calories are being carried by the Elf carrying the most Calories.
// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
func partOne(calories []int) {
	fmt.Printf("Part one: %d\n", calories[len(calories) - 1])
}

// Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
func partTwo(calories []int) {
	sum := 0
	n := len(calories) - 1

	for i := n; i > n - 3; i -- {
		sum += calories[i]
	}
	
	fmt.Printf("Part two: %d\n", sum)
}

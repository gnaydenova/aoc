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
	part := flags.Part()

	lines := reader.ReadFile(fileName)

	if part == 1 {
		partOne(lines)
	} else {
		partTwo(lines)
	}
}

// partOne: In case the Elves get hungry and need extra snacks, they need to know which Elf to ask:
// they'd like to know how many Calories are being carried by the Elf carrying the most Calories.
// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
func partOne(input []string) {
	calories := getSortedCalories(input)

	fmt.Println(calories[len(calories)-1])
}

// partTwo: Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
func partTwo(input []string) {
	calories := getSortedCalories(input)

	sum := 0
	n := len(calories) - 1

	for i := n; i > n-3; i-- {
		sum += calories[i]
	}

	fmt.Println(sum)
}

func getSortedCalories(lines []string) []int {
	calories := []int{0}

	for i, j := 0, 0; i < len(lines); i++ {
		if lines[i] == "" {
			calories = append(calories, 0)
			j++

			continue
		}

		c, _ := strconv.Atoi(lines[i])
		calories[j] += c
	}

	sort.Ints(calories)

	return calories
}

package main

import (
	"fmt"
	"strings"

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
	priorityMap := getPriorityMap()

	total := 0

	for _, line := range input {
		center := len(line) / 2
		compartment1 := line[:center]
		compartment2 := line[center:]

		for _, c := range compartment1 {
			if strings.ContainsRune(compartment2, c) {
				total += priorityMap[c]
				break
			}
		}
	}

	fmt.Print(total)
}

func partTwo(input []string) {
	priorityMap := getPriorityMap()

	total := 0

	for i := 0; i < len(input); i += 3 {
		group := input[i : i+3]
		for _, c := range group[0] {
			if strings.ContainsRune(group[1], c) && strings.ContainsRune(group[2], c) {
				total += priorityMap[c]
				break
			}
		}

	}

	fmt.Print(total)
}

func getPriorityMap() map[rune]int {
	fillPriority := func(priorityMap map[rune]int, asciiStart, priorityStart int) {
		for i := 0; i < 26; i++ {
			p := asciiStart + i
			priorityMap[rune(p)] = priorityStart + i
		}
	}

	priorityMap := make(map[rune]int, 52)
	fillPriority(priorityMap, 97, 1)
	fillPriority(priorityMap, 65, 27)

	return priorityMap
}

package main

import (
	"fmt"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

func main() {
	fileName := flags.File()
	part := flags.Part()

	lines := reader.ReadFile(fileName)

	if part == 1 {
		partOne(lines[0])
	} else {
		partTwo(lines[0])
	}
}

func partOne(input string) {
	fmt.Print(getCharsCount(input, 4))
}

func partTwo(input string) {
	fmt.Print(getCharsCount(input, 14))
}

func getCharsCount(input string, n int) int {
	for i := n; i < len(input); i++ {
		if isUnique(input[i-n : i]) {
			return i
		}
	}

	return 0
}

func isUnique(s string) bool {
	m := make(map[rune]bool)
	for _, c := range s {
		if _, ok := m[c]; ok {
			return false
		}

		m[c] = true
	}

	return true
}

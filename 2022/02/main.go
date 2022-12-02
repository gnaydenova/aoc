package main

import (
	"fmt"
	"strings"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

const (
	rock = iota + 1
	paper
	scissors
)

const (
	win  = 6
	draw = 3
	loss = 0
)

var elfChoiceMap = map[string]int{
	"A": rock,
	"B": paper,
	"C": scissors,
}

var winMap = map[int]int{
	rock:     paper,
	paper:    scissors,
	scissors: rock,
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

func partOne(strategy []string) {
	pcMap := map[string]int{
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	determineScore := func(elfChoice, playerChoice string) int {
		ec, _ := elfChoiceMap[elfChoice]
		pc, _ := pcMap[playerChoice]

		score := pc

		if ec == pc {
			score += draw
		} else if v, ok := winMap[ec]; ok && v == pc {
			score += win
		}

		return score
	}

	score := 0

	for _, turn := range strategy {
		curr := strings.Split(turn, " ")
		score += determineScore(curr[0], curr[1])
	}

	fmt.Println(score)
}

func partTwo(strategy []string) {
	pcMap := map[string]int{
		"X": loss,
		"Y": draw,
		"Z": win,
	}

	determineScore := func(elfChoice, neededEnd string) int {
		ec, _ := elfChoiceMap[elfChoice]
		end, _ := pcMap[neededEnd]

		score := end

		switch end {
		case win:
			score += winMap[ec]
		case draw:
			score += ec
		case loss:
			score += indexOf(winMap, ec)
		}

		return score
	}

	score := 0

	for _, turn := range strategy {
		curr := strings.Split(turn, " ")
		score += determineScore(curr[0], curr[1])
	}

	fmt.Println(score)
}

func indexOf(m map[int]int, n int) int {
	for i, v := range m {
		if v == n {
			return i
		}
	}

	return -1
}

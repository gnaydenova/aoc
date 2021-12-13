package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

func main() {
	file := flags.File()
	instructions := reader.ReadFile(file)

	partOne(instructions)
	partTwo(instructions)
}

func partOne(instructions []string) {
	horizontal := 0
	depth := 0

	for _, line := range instructions {
		direction, step := parseAction(line)

		switch direction {
		case "forward":
			horizontal += step
		case "down":
			depth += step
		case "up":
			depth -= step
		}
	}

	fmt.Printf("Horizontal position: %d\n", horizontal)
	fmt.Printf("Depth: %d\n", depth)
	fmt.Println(horizontal * depth)
}

func partTwo(instructions []string) {
	aim := 0
	horizontal := 0
	depth := 0

	for _, line := range instructions {
		direction, step := parseAction(line)

		switch direction {
		case "forward":
			horizontal += step
			depth += aim * step
		case "down":
			aim += step
		case "up":
			aim -= step
		}
	}

	fmt.Printf("Horizontal position: %d\n", horizontal)
	fmt.Printf("Depth: %d\n", depth)
	fmt.Println(horizontal * depth)
}

func parseAction(action string) (direction string, step int) {
	r := regexp.MustCompile(`(?P<direction>forward|down|up)\s*(?P<step>\d+)`)

	matches := r.FindStringSubmatch(action)
	if len(matches) != 3 {
		log.Fatal(fmt.Sprintf("invalid action %s", action))
	}

	direction = matches[1]
	step, _ = strconv.Atoi(matches[2])

	return
}

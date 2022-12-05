package main

import (
	"fmt"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

type step struct {
	count int
	from  int
	to    int
}

type stack struct {
	items []string
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) push(item string) {
	s.items = append(s.items, item)
}

func (s *stack) pop() string {
	if s.isEmpty() {
		return ""
	}

	i := len(s.items) - 1
	item := s.items[i]
	s.items = s.items[:i]

	return item
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
	procedure, stacks := getProcedureAndStacks(input)

	for _, step := range procedure {
		for i := 0; i < step.count; i++ {
			item := stacks[step.from].pop()
			stacks[step.to].push(item)
		}
	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i].pop())
	}
}

func partTwo(input []string) {
	procedure, stacks := getProcedureAndStacks(input)

	for _, step := range procedure {
		s := &stack{}
		for i := 0; i < step.count; i++ {
			s.push(stacks[step.from].pop())
		}
		for !s.isEmpty() {
			stacks[step.to].push(s.pop())
		}
	}

	for i := 1; i <= len(stacks); i++ {
		fmt.Print(stacks[i].pop())
	}
}

func getProcedureAndStacks(input []string) (procedure []step, stacks map[int]*stack) {
	stackEndIndex := 0

	for i, line := range input {
		if line == "" {
			stackEndIndex = i - 1
		}

		if stackEndIndex > 0 && i > stackEndIndex+1 {
			var s step
			fmt.Sscanf(line, "move %d from %d to %d", &s.count, &s.from, &s.to)
			procedure = append(procedure, s)
		}
	}

	stacks = make(map[int]*stack)
	for i, c := range input[stackEndIndex] {
		if c == ' ' {
			continue
		}

		for j := stackEndIndex - 1; j >= 0; j-- {
			if len(input[j]) < i || input[j][i] == ' ' {
				continue
			}

			index := int(c - '0')
			if _, ok := stacks[index]; !ok {
				stacks[index] = &stack{}
			}
			stacks[index].push(string(input[j][i]))
		}
	}

	return procedure, stacks
}

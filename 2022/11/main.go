package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

type test struct {
	divisor uint64
	ifTrue  int
	ifFalse int
}

func (t test) run(w uint64) int {
	if w%t.divisor == 0 {
		return t.ifTrue
	}

	return t.ifFalse
}

type monkey struct {
	count     int
	items     []uint64
	operation func(curr uint64) uint64
	test      test
}

func (m *monkey) catchItem(item uint64) {
	m.items = append(m.items, item)
}

func (m *monkey) playTurn(monkeys []*monkey, manageWorry func(uint64) uint64) {
	for i := 0; i < len(m.items); i++ {
		w := m.operation(m.items[i])
		m.count++

		w = manageWorry(w)

		throwTo := m.test.run(w)
		monkeys[throwTo].catchItem(w)
	}

	m.items = []uint64{}
}

func main() {
	lines := reader.ReadFile(flags.File())

	if flags.Part() == 1 {
		partOne(lines)
	} else {
		partTwo(lines)
	}
}

func partOne(input []string) {
	monkeys := getMonkeys(input)

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.playTurn(monkeys, func(curr uint64) uint64 { return curr / 3 })
		}
	}

	fmt.Print(getMonkeyBusiness(monkeys))
}

func partTwo(input []string) {
	monkeys := getMonkeys(input)

	lcm := uint64(1)
	for _, monkey := range monkeys {
		lcm *= monkey.test.divisor
	}

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			monkey.playTurn(monkeys, func(curr uint64) uint64 { return curr % lcm })
		}
	}

	fmt.Print(getMonkeyBusiness(monkeys))
}

func getMonkeys(input []string) []*monkey {
	var monkeys []*monkey
	var curr *monkey

	for _, l := range input {
		line := strings.TrimLeft(l, " ")

		switch {
		case strings.HasPrefix(line, "Monkey"):
			curr = new(monkey)
			monkeys = append(monkeys, curr)
		case strings.HasPrefix(line, "Starting items"):
			s := strings.TrimLeft(line, "Starting items: ")
			items := strings.Split(s, ", ")
			curr.items = make([]uint64, len(items))
			for j, item := range items {
				curr.items[j], _ = strconv.ParseUint(item, 0, 64)
			}
		case strings.HasPrefix(line, "Operation"):
			var op, right string
			fmt.Sscanf(line, "Operation: new = old %s %s", &op, &right)
			if right == "old" {
				curr.operation = func(curr uint64) uint64 {
					if op == "*" {
						return curr * curr
					}
					return curr + curr
				}
			} else {
				r, _ := strconv.ParseUint(right, 0, 64)
				curr.operation = func(curr uint64) uint64 {
					if op == "*" {
						return curr * r
					}
					return curr + r
				}
			}
		case strings.HasPrefix(line, "Test"):
			fmt.Sscanf(line, "Test: divisible by %d", &curr.test.divisor)
		case strings.HasPrefix(line, "If true"):
			fmt.Sscanf(line, "If true: throw to monkey %d", &curr.test.ifTrue)
		case strings.HasPrefix(line, "If false"):
			fmt.Sscanf(line, "If false: throw to monkey %d", &curr.test.ifFalse)
		}
	}

	return monkeys
}

func getMonkeyBusiness(monkeys []*monkey) int {
	max1, max2 := 0, 0
	for _, monkey := range monkeys {
		if monkey.count > max1 {
			max2 = max1
			max1 = monkey.count
		} else if monkey.count > max2 {
			max2 = monkey.count
		}
	}

	return max1 * max2
}

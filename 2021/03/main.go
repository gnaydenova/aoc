package main

import (
	"fmt"
	"strconv"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

func main() {
	fileName := flags.File()
	diagnostics := reader.ReadFile(fileName)

	partOne(diagnostics)
	partTwo(diagnostics)
}

func partOne(diagnostics []string) {
	length := len(diagnostics[0])

	gamma := make([]byte, length)
	epsilon := make([]byte, length)

	for i := 0; i < length; i++ {
		gamma[i], epsilon[i] = getMostAndLeastCommon(diagnostics, i)
	}

	fmt.Printf("Gamma: %s\n", gamma)
	fmt.Printf("Epsilon: %s\n", epsilon)

	g, _ := strconv.ParseInt(string(gamma), 2, 64)
	fmt.Printf("Gamma: %d\n", g)

	e, _ := strconv.ParseInt(string(epsilon), 2, 64)
	fmt.Printf("Epsilon: %d\n", e)

	fmt.Printf("Power consumption: %d\n", g*e)
}

func partTwo(diagnostics []string) {
	o2 := getRating(diagnostics, func(s []string, i int) byte {
		most, _ := getMostAndLeastCommon(s, i)
		return most
	})

	fmt.Printf("Oxygen generator rating: %s\n", o2)

	co2 := getRating(diagnostics, func(s []string, i int) byte {
		_, least := getMostAndLeastCommon(s, i)
		return least 
	})

	fmt.Printf("CO2 scrubber rating: %s\n", co2)

	o, _ := strconv.ParseInt(o2, 2, 64)
	c, _ := strconv.ParseInt(co2, 2, 64)

	fmt.Printf("Life support rating: %d\n", o * c)
}

func getMostAndLeastCommon(diagnostics []string, position int) (most, least byte) {
	occ := 0
	for _, line := range diagnostics {
		if line[position] == '1' {
			occ += 1
		} else {
			occ -= 1
		}
	}

	if occ >= 0 {
		return '1', '0'
	}

	return '0', '1'
}

type lookup func([]string, int) byte

func getRating(diagnostics []string, fn lookup) string {
	for i := 0; i < len(diagnostics[0]); i++ {
		diagnostics = filter(diagnostics, i, fn)
		if len(diagnostics) == 1 {
			return diagnostics[0]
		}
	}

	return ""
}

func filter(diagnostics []string, position int, fn lookup) []string {
	l := fn(diagnostics, position)

	var filtered []string
	for _, line := range diagnostics {
		if line[position] == l {
			filtered = append(filtered, line)
		}
	}

	return filtered
}


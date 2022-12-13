package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

func main() {
	lines := reader.ReadFile(flags.File())

	if flags.Part() == 1 {
		partOne(lines)
	} else {
		partTwo(lines)
	}
}

func partOne(input []string) {
	sum := 0

	for i, j := 0, 1; i < len(input)-1; i += 3 {
		var left, right any
		json.Unmarshal([]byte(input[i]), &left)
		json.Unmarshal([]byte(input[i+1]), &right)

		if compare(left, right) >= 0 {
			sum += j
		}

		j++
	}

	fmt.Print(sum)
}

func partTwo(input []string) {
	var packets []any
	input = append(input, "[[2]]", "[[6]]")

	for _, line := range input {
		if line == "" {
			continue
		}

		var packet any
		json.Unmarshal([]byte(line), &packet)

		packets = append(packets, packet)
	}

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) > 0
	})

	decoderKeyStart, decoderKeyEnd := 0, 0

	for i, packet := range packets {
		switch fmt.Sprintf("%v", packet) {
		case "[[2]]":
			decoderKeyStart = i + 1
		case "[[6]]":
			decoderKeyEnd = i + 1
			break
		}
	}

	fmt.Print(decoderKeyStart * decoderKeyEnd)
}

func compare(left, right any) int {
	l, lOk := left.(float64)
	r, rOk := right.(float64)

	if lOk && rOk {
		return int(r - l)
	}

	lList, rList := getList(left), getList(right)

	rListLen, lListLen := len(rList), len(lList)
	max := max(lListLen, rListLen)

	for i := 0; i < max; i++ {
		if i == lListLen {
			return 1
		}

		if i == rListLen {
			return -1
		}

		if result := compare(lList[i], rList[i]); result != 0 {
			return result
		}
	}

	return 0
}

func getList(l any) []any {
	var list []any
	switch l.(type) {
	case []any:
		list = l.([]any)
	case float64:
		list = []any{l}
	}

	return list
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

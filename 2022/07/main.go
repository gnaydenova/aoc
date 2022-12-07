package main

import (
	"fmt"

	"github.com/gnaydenova/aoc/pkg/flags"
	"github.com/gnaydenova/aoc/pkg/reader"
)

const (
	dir = iota + 1
	file
)

type node struct {
	parent   *node
	kind     int
	name     string
	size     int
	children []*node
}

func (n *node) findChild(name string) *node {
	for _, child := range n.children {
		if child.name == name {
			return child
		}
	}

	return nil
}

func (n *node) addChild(c *node) {
	c.parent = n
	n.children = append(n.children, c)
}

func (n *node) getSize() int {
	if n.kind == file {
		return n.size
	}

	s := 0
	for _, c := range n.children {
		s += c.getSize()
	}

	return s
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
	dirSizes := walk(getFileTree(input), []int{})
	total := 0

	for _, s := range dirSizes {
		if s <= 100000 {
			total += s
		}
	}

	fmt.Print(total)
}

func partTwo(input []string) {
	dirSizes := walk(getFileTree(input), []int{})

	max := 0
	for _, s := range dirSizes {
		if s > max {
			max = s
		}
	}
	
	free := 70000000 - max
	needed := 30000000 - free

	toDelete := 0
	for _, s := range dirSizes {
		if s >= needed && (toDelete == 0 || s < toDelete) {
			toDelete = s
		}
	}

	fmt.Print(toDelete)
}

func getFileTree(input []string) *node {
	var curr, top *node

	for i := 0; i < len(input); i++ {
		line := input[i]
		if line[0] == '$' {
			var cmd, arg string
			fmt.Sscanf(line, "$ %s %s", &cmd, &arg)
			if cmd == "cd" {
				if arg != ".." {
					if curr != nil {
						curr = curr.findChild(arg)
						continue
					}
					curr = &node{kind: dir, name: arg}
					top = curr
				} else {
					curr = curr.parent
				}
			}
		} else {
			if line[:3] == "dir" {
				curr.addChild(&node{kind: dir, name: line[4:]})
			} else {
				child := node{kind: file, parent: curr}
				fmt.Sscanf(line, "%d %s", &child.size, &child.name)
				curr.addChild(&child)
			}
		}
	}

	return top
}

func walk(n *node, dirSizes []int) []int {
	if n.kind == file {
		return dirSizes
	}

	dirSizes = append(dirSizes, n.getSize())
	for _, c := range n.children {
		dirSizes = walk(c, dirSizes)
	}

	return dirSizes
}

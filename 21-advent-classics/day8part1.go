package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	name        string
	left, right string
}

func solve(input []string) {
	if len(input) < 3 {
		panic("Invalid input, need at least three lines")
	}
	turns := input[0]
	nodes := []Node{}

	re := regexp.MustCompile(`(\w+)\s*=\s*\((\w+),\s*(\w+)\)`)

	for i := 2; i < len(input); i++ {
		// AAA = (BBB, BBB)
		matches := re.FindStringSubmatch(input[i])
		if matches != nil && len(matches) == 4 {
			nodes = append(nodes, Node{name: matches[1], left: matches[2], right: matches[3]})
		} else {
			fmt.Println("Error: Invalid input, expected format like 'AAA = (BBB, BBB)'", input[i])
		}
	}

	nodesMap := make(map[string]Node)
	for _, node := range nodes {
		nodesMap[node.name] = node
	}

	fmt.Println("Turns", turns)
	fmt.Println("Nodes", nodes)

	steps := 0
	location := "AAA"
	index := 0
	for location != "ZZZ" {
		if index >= len(turns) {
			index = 0
		}
		turn := turns[index]
		index += 1
		steps += 1

		if turn == 'L' {
			location = nodesMap[location].left
		} else if turn == 'R' {
			location = nodesMap[location].right
		} else {
			fmt.Println("Error: Invalid turn", turn)
		}
	}

	fmt.Println("Steps", steps)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	solve(lines)
}

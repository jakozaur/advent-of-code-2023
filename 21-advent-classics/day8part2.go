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

type Solution struct {
	initialSteps int
	cycle        int
	zTurns       []int
}

func findCycle(start string, nodesMap map[string]Node, turns string, index int) Solution {
	steps := 0
	visited := make(map[string]map[int]int)
	location := start
	zLocations := []int{}

	for true {
		if index >= len(turns) {
			index = 0
		}
		turn := turns[index]

		if location[2] == 'Z' {
			zLocations = append(zLocations, steps)
		}

		if visited[location] == nil {
			visited[location] = make(map[int]int)
		} else {
			val, exist := visited[location][index]
			if exist {
				fmt.Println("Cycle found location", location, "index", index, "steps", steps, "visited[location][index]", val, zLocations)
				return Solution{initialSteps: val, cycle: steps - val, zTurns: zLocations}
			}
		}
		visited[location][index] = steps

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

	// No-op
	return Solution{initialSteps: 0, cycle: 0}
}

func moveLocations(locations []string, nodesMap map[string]Node, turns string, index int, bySteps int) ([]string, int) {
	for i := 0; i < bySteps; i++ {
		if index >= len(turns) {
			index = 0
		}
		isSolution := true
		for _, location := range locations {
			if location[2] != 'Z' {
				isSolution = false
				break
			}
		}
		if isSolution {
			fmt.Println("Solution found at step", i)
		}

		turn := turns[index]
		index += 1

		for j := 0; j < len(locations); j++ {
			if turn == 'L' {
				locations[j] = nodesMap[locations[j]].left
			} else if turn == 'R' {
				locations[j] = nodesMap[locations[j]].right
			} else {
				fmt.Println("Error: Invalid turn", turn)
			}
		}
	}

	return locations, index
}

// Function to calculate GCD using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate LCM using GCD
func lcm(a, b int) int {
	return a * b / gcd(a, b)
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

	locations := []string{}

	for _, node := range nodes {
		if node.name[2] == 'A' {
			locations = append(locations, node.name)
		}
	}

	fmt.Println("Turns", turns)
	fmt.Println("Nodes", nodes)
	fmt.Println("Locations", locations)

	steps := 0
	index := 0
	solutions := make([]Solution, len(locations))
	for i, location := range locations {
		solutions[i] = findCycle(location, nodesMap, turns, 0)
	}

	maxInitialSteps := 0
	for _, solution := range solutions {
		if solution.initialSteps > maxInitialSteps {
			maxInitialSteps = solution.initialSteps
		}
	}

	fmt.Println("Move by ", maxInitialSteps)
	locations, index = moveLocations(locations, nodesMap, turns, index, maxInitialSteps)
	steps += maxInitialSteps
	fmt.Println("====")

	solutions2 := make([]Solution, len(locations))
	for i, location := range locations {
		solutions2[i] = findCycle(location, nodesMap, turns, index)
	}

	for _, solution := range solutions2 {
		if len(solution.zTurns) != 1 {
			panic("My specific assumption")
		}
	}

	moveBySteps := solutions2[0].zTurns[0]
	locations, index = moveLocations(locations, nodesMap, turns, index, moveBySteps)
	steps += moveBySteps
	stepSize := solutions2[0].cycle
	fmt.Println("==== steps", steps, "stepSize", stepSize)
	fmt.Println("Locations", locations)

	for solveLocation := 1; solveLocation < 300 && solveLocation < len(locations); solveLocation++ {
		cycle := solutions2[solveLocation].cycle
		zLocation := solutions[solveLocation].zTurns[0]
		fmt.Println("steps", steps, "stepSize", stepSize, "ZLocation", zLocation, "cycle", cycle)

		// TODO: I can lcm here, but I'm lazy
		for steps%cycle != zLocation%cycle {
			steps += stepSize
		}

		stepSize = lcm(stepSize, cycle)
		fmt.Println("==== steps", steps, "stepSize", stepSize)
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

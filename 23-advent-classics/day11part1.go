package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func isRowEmpty(line string) bool {
	for _, c := range line {
		if c != '.' {
			return false
		}
	}
	return true
}

func isColumnEmpty(input []string, column int) bool {
	for _, line := range input {
		if line[column] != '.' {
			return false
		}
	}
	return true
}

func calculateExpandedInput(input []string) []string {
	rowsToDouble := []int{}
	columnsToDouble := []int{}

	for i, line := range input {
		if isRowEmpty(line) {
			rowsToDouble = append(rowsToDouble, i)
		}
	}

	for i := range input[0] {
		if isColumnEmpty(input, i) {
			columnsToDouble = append(columnsToDouble, i)
		}
	}

	expandedInput := make([]string, len(input)+len(rowsToDouble))
	iRowToDouble := 0
	for i := range input {
		iColumnToDouble := 0
		line := ""
		for j := 0; j < len(input[0]); j++ {
			if iColumnToDouble < len(columnsToDouble) && j == columnsToDouble[iColumnToDouble] {
				line += "."
				iColumnToDouble++
			}
			line += string(input[i][j])
		}
		expandedInput[i+iRowToDouble] = line
		if iRowToDouble < len(rowsToDouble) && i == rowsToDouble[iRowToDouble] {
			iRowToDouble++
			expandedInput[i+iRowToDouble] = line
		}
	}
	return expandedInput
}

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func solve(input []string) {
	expandedInput := calculateExpandedInput(input)

	for _, line := range expandedInput {
		fmt.Println(line)
	}

	points := []Point{}
	for i := range expandedInput {
		for j := range expandedInput[i] {
			if expandedInput[i][j] == '#' {
				points = append(points, Point{j, i})
			}
		}
	}

	fmt.Println("Points:", points)
	fmt.Println("Points count:", len(points))

	result := 0
	for i, pA := range points {
		for _, pB := range points[i+1:] {
			dist := absInt(pA.x-pB.x) + absInt(pA.y-pB.y)
			result += dist
		}
	}

	fmt.Println("Result:", result)
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

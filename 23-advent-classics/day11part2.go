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

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func solve(input []string) {
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

	iRowToDouble := 0
	multiplyFactor := 1000000

	points := []Point{}
	for i := range input {
		if iRowToDouble < len(rowsToDouble) && i == rowsToDouble[iRowToDouble] {
			iRowToDouble++
		}
		iColumnToDouble := 0
		for j := range input[i] {
			if iColumnToDouble < len(columnsToDouble) && j == columnsToDouble[iColumnToDouble] {
				iColumnToDouble++
			}
			if input[i][j] == '#' {
				points = append(points, Point{j + iColumnToDouble*(multiplyFactor-1), i + iRowToDouble*(multiplyFactor-1)})
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

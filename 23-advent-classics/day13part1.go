package main

import (
	"bufio"
	"fmt"
	"os"
)

func isVerticalLine(input []string, num int) bool {
	for _, line := range input {
		for i := num; i < len(line) && num-1-(i-num) >= 0; i++ {
			if line[i] != line[num-1-(i-num)] {
				return false
			}
		}
	}
	return true
}

func isHorizontalLine(input []string, num int) bool {
	for i := num; i < len(input) && num-1-(i-num) >= 0; i++ {
		if input[i] != input[num-1-(i-num)] {
			return false
		}
	}
	return true
}

func solvePuzzle(input []string) int {
	for i := 1; i < len(input[0]); i++ {
		if isVerticalLine(input, i) {
			return i
		}
	}
	for i := 1; i < len(input); i++ {
		if isHorizontalLine(input, i) {
			return 100 * i
		}
	}
	fmt.Println("Error: no solution found")
	return 0
}

func solve(input []string) {
	result := 0
	puzzle := []string{}
	for _, line := range input {
		if len(line) > 0 {
			puzzle = append(puzzle, line)
		} else {
			partialResult := solvePuzzle(puzzle)
			puzzle = []string{}
			fmt.Println("Partial result:", partialResult)
			result += partialResult
		}
	}

	// Last one
	partialResult := solvePuzzle(puzzle)
	fmt.Println("Partial result:", partialResult)
	result += partialResult

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

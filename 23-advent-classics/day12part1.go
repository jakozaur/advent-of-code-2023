package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseNumbersByComma(numbers string) []int {
	result := []int{}
	for _, number := range strings.Split(numbers, ",") {
		if len(number) > 0 {
			parsedNumber, err2 := strconv.Atoi(number)
			result = append(result, parsedNumber)
			if err2 != nil {
				fmt.Println("Error: Invalid input, invalid number", err2, number)
			}
		}
	}
	return result
}

func solveLine(line string) int {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		fmt.Println("Error: Invalid input, no ' '", line)
	}
	records := parts[0]
	partsInts := parseNumbersByComma(parts[1])

	var dpFunc func(i, j int) int
	dpFunc = func(i, j int) int {
		if i >= len(records) {
			if j == len(partsInts) {
				return 1
			} else {
				return 0
			}
		}
		switch records[i] {
		case '.':
			return dpFunc(i+1, j)
		case '#':
			if j == len(partsInts) {
				return 0
			}
			if i+partsInts[j] > len(records) {
				return 0
			}
			for k := 1; k < partsInts[j]; k++ {
				if records[i+k] == '.' {
					return 0
				}
			}
			if i+partsInts[j] < len(records) && records[i+partsInts[j]] == '#' {
				return 0
			}
			return dpFunc(i+partsInts[j]+1, j+1)
		case '?':
			sum := dpFunc(i+1, j)
			if j < len(partsInts) {
				if i+partsInts[j] <= len(records) {
					canBeExtended := true
					for k := 1; k < partsInts[j]; k++ {
						if records[i+k] == '.' {
							canBeExtended = false
							break
						}
					}
					if i+partsInts[j] < len(records) && records[i+partsInts[j]] == '#' {
						canBeExtended = false
					}
					if canBeExtended {
						sum += dpFunc(i+partsInts[j]+1, j+1)
					}
				}
			}
			return sum
		}
		fmt.Println("Error: Invalid input, invalid record", records[i])
		return -1
	}

	return dpFunc(0, 0)
}

func solve(input []string) {
	result := 0
	for _, line := range input {
		sol := solveLine(line)
		fmt.Println("Line:", line, "solutions", sol)
		result += sol
	}
	fmt.Println("Result: ", result)
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

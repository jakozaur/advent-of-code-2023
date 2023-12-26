package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateHash(input string) int {
	hash := 0
	for _, c := range input {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func solve(input []string) {
	result := 0
	strs := strings.Split(input[0], ",")
	for _, s := range strs {
		partialResult := calculateHash(s)
		fmt.Println("Partial ", s, "result:", partialResult)
		result += partialResult
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

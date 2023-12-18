package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func isDigit(a byte) bool {
	return a >= '0' && a <= '9'
}

func isSymbol(a byte) bool {
	r := rune(a)
	return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '.'
}

func solve(input []string) {
	height := len(input)
	width := len(input[0])

	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}

	dfs := func(x, y int) {

	}

	dfs = func(x, y int) {
		if x < 0 || x >= width || y < 0 || y >= height {
			return
		}
		if matrix[x][y] == 1 {
			return
		}
		c := input[x][y]
		if isDigit(c) {
			matrix[x][y] = 1
			dfs(x, y-1)
			dfs(x, y+1)
		}
	}

	for x := range matrix {
		for y := range matrix[x] {
			if isSymbol(input[x][y]) {
				matrix[x][y] = 1
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						dfs(x+i, y+j)
					}
				}
			}
		}
	}

	fmt.Println("After numbers")
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 1 {
				fmt.Printf("%c", input[i][j])
			} else {
				print(" ")
			}
		}
		println()
	}

	results := []int{}

	for i := range matrix {
		currentNum := -1
		for j := range matrix[i] {
			c := input[i][j]
			if matrix[i][j] == 1 && isDigit(c) {
				if currentNum == -1 {
					currentNum = 0
				}
				currentNum = currentNum*10 + int(c-'0')
			} else if currentNum != -1 {
				results = append(results, currentNum)
				currentNum = -1
			}
		}

		if currentNum != -1 {
			results = append(results, currentNum)
			currentNum = -1
		}
	}

	fmt.Println("Results: ", results)
	sum := 0
	for _, r := range results {
		sum += r
	}
	fmt.Println("Sum: ", sum)
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

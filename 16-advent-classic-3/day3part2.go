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

	dfs := func(x, y, marker int) bool {
		return false
	}

	dfs2 := func(x, y, rest int) int {
		return 0
	}

	dfs = func(x, y, marker int) bool {
		if x < 0 || x >= width || y < 0 || y >= height {
			return false
		}
		if matrix[x][y] == marker {
			return false
		}
		c := input[x][y]
		if isDigit(c) {
			matrix[x][y] = marker
			dfs(x, y-1, marker)
			dfs(x, y+1, marker)
			return true
		}
		return false
	}

	dfs2 = func(x, y, rest int) int {
		if x < 0 || x >= width || y < 0 || y >= height {
			return 0
		}
		if matrix[x][y] == 2 {
			return 0
		}
		c := input[x][y]
		if isDigit(c) {
			matrix[x][y] = 2
			result := int(c-'0') + rest*10
			if r := dfs2(x, y-1, 0); r > 0 {
				result += r * 10
			}
			if r := dfs2(x, y+1, result); r > 0 {
				result = r
			}
			return result
		}
		return 0
	}

	newResult := []int{}

	for x := range matrix {
		for y := range matrix[x] {
			if input[x][y] == '*' {
				matrix[x][y] = 1
				count := 0
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						if dfs(x+i, y+j, 1) {
							count += 1
						}
					}
				}

				if count == 2 {
					matrix[x][y] = 2
					for i := -1; i <= 1; i++ {
						for j := -1; j <= 1; j++ {
							r := dfs2(x+i, y+j, 0)
							if r > 0 {
								newResult = append(newResult, r)
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("After numbers")
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 2 {
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
			if matrix[i][j] == 2 && isDigit(c) {
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

	fmt.Println("Results: ", newResult)
	sum := 0
	for i := 0; i < len(newResult); i += 2 {
		sum += newResult[i] * newResult[i+1]
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

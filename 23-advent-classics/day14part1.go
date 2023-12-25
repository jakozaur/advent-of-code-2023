package main

import (
	"bufio"
	"fmt"
	"os"
)

func printBoard(board [][]int) {
	for _, row := range board {
		for _, v := range row {
			switch v {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func titlNorth(board [][]int) {
	for i, row := range board {
		for j, v := range row {
			if v == 2 {
				for x := i - 1; x >= 0 && board[x][j] == 0; x-- {
					board[x][j] = 2
					board[x+1][j] = 0
				}
			}
		}
	}
}

func countWeights(board [][]int) int {
	result := 0

	for y, row := range board {
		for _, v := range row {
			if v == 2 {
				result += len(board) - y
			}
		}
	}
	return result
}

func solve(input []string) {
	board := make([][]int, len(input))
	for i := range input {
		board[i] = make([]int, len(input[0]))
		for j := range input[i] {
			switch input[i][j] {
			case '#':
				board[i][j] = 1
			case '.':
				board[i][j] = 0
			case 'O':
				board[i][j] = 2
			}
		}
	}

	titlNorth(board)
	printBoard(board)

	result := countWeights(board)

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

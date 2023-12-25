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

func titlWest(board [][]int) {
	for i, row := range board {
		for j, v := range row {
			if v == 2 {
				for x := j - 1; x >= 0 && board[i][x] == 0; x-- {
					board[i][x] = 2
					board[i][x+1] = 0
				}
			}
		}
	}
}

func titlSouth(board [][]int) {
	for i := len(board) - 1; i >= 0; i-- {
		row := board[i]
		for j, v := range row {
			if v == 2 {
				for x := i + 1; x < len(board) && board[x][j] == 0; x++ {
					board[x][j] = 2
					//fmt.Println("Setting", x-1, j, "to 0")
					board[x-1][j] = 0
				}
			}
		}
	}
}

func titlEast(board [][]int) {
	for i, row := range board {
		for j := len(row) - 1; j >= 0; j-- {
			v := row[j]
			if v == 2 {
				for x := j + 1; x < len(row) && board[i][x] == 0; x++ {
					board[i][x] = 2
					board[i][x-1] = 0
				}
			}
		}
	}
}

func titlCycle(board [][]int) {
	titlNorth(board)
	titlWest(board)
	titlSouth(board)
	titlEast(board)
}

func copy2DSlice(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func printDiff(boardA, boardB [][]int) {
	for i := range boardA {
		for j := range boardA[i] {
			if boardA[i][j] != boardB[i][j] {
				switch boardA[i][j] {
				case 0:
					fmt.Print("-")
				case 1:
					fmt.Print("E")
				case 2:
					fmt.Print("+")
				}
			} else {
				switch boardA[i][j] {
				case 0:
					fmt.Print(".")
				case 1:
					fmt.Print("#")
				case 2:
					fmt.Print("O")
				}
			}
		}
		fmt.Println()
	}
}

func boardEqual(boardA, boardB [][]int) bool {
	for i := range boardA {
		for j := range boardA[i] {
			if boardA[i][j] != boardB[i][j] {
				return false
			}
		}
	}
	return true
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

	cycles := 1000000000
	cycleCache := [][][]int{}
mainLoop:
	for i := 0; i < cycles; i++ {
		origBoard := copy2DSlice(board)
		cycleCache = append(cycleCache, origBoard)
		titlCycle(board)
		fmt.Println("Cycle", i+1)
		for j, prevBoard := range cycleCache {
			if boardEqual(prevBoard, board) {
				fmt.Println("Found cycle", i, "with", j)
				fmt.Println("Debug len(cycleCache)", len(cycleCache))
				cycleLength := (i + 1) - j
				remainingCycles := cycles - (i + 1)
				remainingCycles %= cycleLength
				board = cycleCache[j+remainingCycles]
				break mainLoop
			}
		}
		//printDiff(board, origBoard)
	}

	//fmt.Println("Final board:")
	//printBoard(board)

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

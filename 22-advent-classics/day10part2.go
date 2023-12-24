package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func (p *Point) isValid(input []string) bool {
	return p.x >= 0 && p.y >= 0 && p.y < len(input) && p.x < len(input[0])
}

func (p *Point) isValid2(board [][]int) bool {
	return p.x >= 0 && p.y >= 0 && p.y < len(board) && p.x < len(board[0])
}

func (p *Point) connected(input []string) []Point {
	c := input[p.y][p.x]
	potential := []Point{}
	switch c {
	case '|':
		potential = []Point{{p.x, p.y - 1}, {p.x, p.y + 1}}
	case '-':
		potential = []Point{{p.x - 1, p.y}, {p.x + 1, p.y}}
	case 'L':
		potential = []Point{{p.x + 1, p.y}, {p.x, p.y - 1}}
	case 'J':
		potential = []Point{{p.x - 1, p.y}, {p.x, p.y - 1}}
	case '7':
		potential = []Point{{p.x - 1, p.y}, {p.x, p.y + 1}}
	case 'F':
		potential = []Point{{p.x + 1, p.y}, {p.x, p.y + 1}}
	}
	result := []Point{}
	for _, p := range potential {
		if p.isValid(input) {
			result = append(result, p)
		}
	}
	return result
}

func (p *Point) around(input []string) []Point {
	potential := []Point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}}
	result := []Point{}
	for _, p := range potential {
		if p.isValid(input) {
			result = append(result, p)
		}
	}
	return result
}

func (p *Point) around2(board [][]int) []Point {
	potential := []Point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}}
	result := []Point{}
	for _, p := range potential {
		if p.isValid2(board) {
			result = append(result, p)
		}
	}
	return result
}

func (p *Point) aroundConnected(input []string) []Point {
	potential := p.around(input)
	result := []Point{}
	//fmt.Println("Potential:", potential, "for", p)
	for _, connectTo := range potential {
		back := connectTo.connected(input)
		//fmt.Println("Back:", back, "for", connectTo)
		for _, connectBack := range back {
			if connectBack.x == p.x && connectBack.y == p.y {
				result = append(result, connectTo)
				break
			}
		}
	}
	return result
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, v := range row {
			if v == -1 {
				fmt.Print(".")
			} else if v >= 0 {
				fmt.Print("X")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
}

func flood(board [][]int) {
	queue := []Point{}
	for y := range board {
		queue = append(queue, Point{0, y})
		queue = append(queue, Point{len(board[0]) - 1, y})
	}
	for x := range board[0] {
		queue = append(queue, Point{x, 0})
		queue = append(queue, Point{x, len(board) - 1})
	}

	queueNext := []Point{}
	for len(queue) > 0 {
		for _, p := range queue {
			if board[p.y][p.x] == -1 {
				board[p.y][p.x] = -10
				queueNext = append(queueNext, p.around2(board)...)
			}
		}
		queue = queueNext
		queueNext = []Point{}
	}
}

func upscaled(board [][]int, input []string) [][]int {
	result := make([][]int, len(board)*3)
	for i := range result {
		result[i] = make([]int, len(board[0])*3)
		for j := range result[i] {
			result[i][j] = -1
		}
	}

	for y := range board {
		for x := range board[y] {
			if board[y][x] >= 0 {
				result[y*3+1][x*3+1] = 1
				switch input[y][x] {
				case '|':
					result[y*3][x*3+1] = 1
					result[y*3+2][x*3+1] = 1
				case '-':
					result[y*3+1][x*3] = 1
					result[y*3+1][x*3+2] = 1
				case 'L':
					result[y*3+1][x*3+2] = 1
					result[y*3][x*3+1] = 1
				case 'J':
					result[y*3+1][x*3] = 1
					result[y*3][x*3+1] = 1
				case '7':
					result[y*3+1][x*3] = 1
					result[y*3+2][x*3+1] = 1
				case 'F':
					result[y*3+1][x*3+2] = 1
					result[y*3+2][x*3+1] = 1
				case 'S':
					sPoint := Point{x, y}
					queue := sPoint.aroundConnected(input)
					if len(queue) != 2 {
						fmt.Println("Error: Expected two connected to S", queue)
					}
					for _, p := range queue {
						//fmt.Println("S connected to", sPoint, "p", p)
						result[y*3+1+(p.y-sPoint.y)][x*3+1+(p.x-sPoint.x)] = 1
					}
				}
			}
		}
	}

	return result
}

func moveFlood(upscaled [][]int, board [][]int) {
	for y := range board {
		for x := range board[y] {
			if upscaled[y*3+1][x*3+1] == -10 {
				board[y][x] = -10
			}
		}
	}
}

func countNotFlooded(board [][]int) int {
	count := 0
	for y := range board {
		for x := range board[y] {
			if board[y][x] == -1 {
				count += 1
			}
		}
	}
	return count
}

func solve(input []string) {
	board := make([][]int, len(input))
	for i := range board {
		board[i] = make([]int, len(input[0]))
		for j := range board[i] {
			board[i][j] = -1
		}
	}

	var sPos Point
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'S' {
				sPos = Point{j, i}
			}
		}
	}

	board[sPos.y][sPos.x] = 0
	queue := sPos.aroundConnected(input)
	queueNext := []Point{}
	nextDist := 0
	for len(queue) > 0 {
		nextDist += 1
		for _, p := range queue {
			if board[p.y][p.x] == -1 {
				board[p.y][p.x] = nextDist
				queueNext = append(queueNext, p.connected(input)...)
			}
		}
		queue = queueNext
		queueNext = []Point{}
	}

	fmt.Println("Board:")
	printBoard(board)

	upscaledBoard := upscaled(board, input)
	fmt.Println("Upscaled board:")
	flood(upscaledBoard)
	printBoard(upscaledBoard)

	moveFlood(upscaledBoard, board)
	fmt.Println("Board with flood:")
	printBoard(board)

	result := countNotFlooded(board)

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

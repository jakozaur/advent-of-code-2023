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
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
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

	result := 0

	board[sPos.y][sPos.x] = 0
	queue := sPos.aroundConnected(input)
	queueNext := []Point{}
	nextDist := 0
	for len(queue) > 0 {
		nextDist += 1
		for _, p := range queue {
			if board[p.y][p.x] == -1 {
				board[p.y][p.x] = nextDist
				result = nextDist
				queueNext = append(queueNext, p.connected(input)...)
			}
		}
		queue = queueNext
		queueNext = []Point{}
	}

	fmt.Println("Board:")
	printBoard(board)

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

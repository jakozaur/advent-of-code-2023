package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	x, y int
	name int
}

func newRightDirection() Direction {
	return Direction{1, 0, 1}
}

func newLeftDirection() Direction {
	return Direction{-1, 0, 2}
}

func newDownDirection() Direction {
	return Direction{0, 1, 3}
}

func newUpDirection() Direction {
	return Direction{0, -1, 4}
}

type Beam struct {
	x, y      int
	direction Direction
}

func (b *Beam) isValid(input []string) bool {
	return b.x >= 0 && b.y >= 0 && b.y < len(input) && b.x < len(input[0])
}

func solveBeam(input []string, initialBeam Beam) int {
	visited := make([][]int, len(input))
	for i := range visited {
		visited[i] = make([]int, len(input[0]))
	}

	beams := []Beam{initialBeam}
	newBeams := []Beam{}
	for len(beams) > 0 {
		//fmt.Println("Beams", len(beams), beams)

		for _, beam := range beams {
			if 0 != (visited[beam.y][beam.x] & (1 << beam.direction.name)) {
				continue
			}
			visited[beam.y][beam.x] |= 1 << beam.direction.name
			switch input[beam.y][beam.x] {
			case '|':
				up := newUpDirection()
				down := newDownDirection()
				newBeams = append(newBeams, Beam{beam.x + up.x, beam.y + up.y, up}, Beam{beam.x + down.x, beam.y + down.y, down})
			case '-':
				left := newLeftDirection()
				right := newRightDirection()
				newBeams = append(newBeams, Beam{beam.x + left.x, beam.y + left.y, left}, Beam{beam.x + right.x, beam.y + right.y, right})
			case '\\':
				var newDirection Direction
				if beam.direction.name == newRightDirection().name {
					newDirection = newDownDirection()
				} else if beam.direction.name == newLeftDirection().name {
					newDirection = newUpDirection()
				} else if beam.direction.name == newUpDirection().name {
					newDirection = newLeftDirection()
				} else if beam.direction.name == newDownDirection().name {
					newDirection = newRightDirection()
				} else {
					fmt.Println("Error: Unknown direction", beam.direction)
				}
				newBeams = append(newBeams, Beam{beam.x + newDirection.x, beam.y + newDirection.y, newDirection})
			case '/':
				var newDirection Direction
				if beam.direction.name == newRightDirection().name {
					newDirection = newUpDirection()
				} else if beam.direction.name == newLeftDirection().name {
					newDirection = newDownDirection()
				} else if beam.direction.name == newUpDirection().name {
					newDirection = newRightDirection()
				} else if beam.direction.name == newDownDirection().name {
					newDirection = newLeftDirection()
				} else {
					fmt.Println("Error: Unknown direction", beam.direction)
				}
				newBeams = append(newBeams, Beam{beam.x + newDirection.x, beam.y + newDirection.y, newDirection})
			case '.':
				newBeams = append(newBeams, Beam{beam.x + beam.direction.x, beam.y + beam.direction.y, beam.direction})
			}
		}

		beams = []Beam{}
		for _, beam := range newBeams {
			if beam.isValid(input) {
				beams = append(beams, beam)
			}
		}
		newBeams = []Beam{}
	}

	illuminated := 0

	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] > 0 {
				illuminated += 1
				//fmt.Print("X")
			} else {
				//fmt.Print(".")
			}
		}
		//fmt.Println()
	}

	return illuminated
}

func solve(input []string) {
	illuminated := 0
	for y := range input {
		candidate := solveBeam(input, Beam{0, y, newRightDirection()})
		if candidate > illuminated {
			illuminated = candidate
		}

		candidate = solveBeam(input, Beam{len(input[0]) - 1, y, newLeftDirection()})
		if candidate > illuminated {
			illuminated = candidate
		}
	}

	for x := range input[0] {
		candidate := solveBeam(input, Beam{x, 0, newDownDirection()})
		if candidate > illuminated {
			illuminated = candidate
		}

		candidate = solveBeam(input, Beam{x, len(input) - 1, newUpDirection()})
		if candidate > illuminated {
			illuminated = candidate
		}
	}
	fmt.Println("Result: ", illuminated)
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

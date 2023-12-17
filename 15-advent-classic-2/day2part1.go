package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(input []string) {
	sum := 0

	for _, line := range input {
		mainParts := strings.Split(line, ":")
		if len(mainParts) != 2 {
			fmt.Println("Error: Invalid input, no ':'", line)
		}
		gameParts := strings.Split(mainParts[0], "Game ")
		if len(gameParts) != 2 {
			fmt.Println("Error: Invalid input, no 'Game '", line)
		}
		gameId, _ := strconv.Atoi(gameParts[1])
		validGame := true

		for _, move := range strings.Split(mainParts[1], ";") {
			dices := make(map[string]int)
			for _, dice := range strings.Split(move, ",") {
				numberColor := strings.Split(strings.TrimSpace(dice), " ")
				if len(gameParts) != 2 {
					fmt.Println("Error: Invalid input no (num color)", dice, line)
				}
				num, _ := strconv.Atoi(numberColor[0])
				color := numberColor[1]
				if color != "red" && color != "green" && color != "blue" {
					fmt.Println("Error: Invalid input, invalid color name", color, dice, line)
				}
				dices[color] = num
			}
			valueGreen, existsGreen := dices["green"]
			if !existsGreen {
				valueGreen = 0
			}
			valueRed, existsRed := dices["red"]
			if !existsRed {
				valueRed = 0
			}
			valueBlue, existsBlue := dices["blue"]
			if !existsBlue {
				valueBlue = 0
			}

			if valueGreen <= 13 && valueRed <= 12 && valueBlue <= 14 {
				//fmt.Println("Game is valid", gameId)
			} else {
				validGame = false
			}

		}
		if validGame {
			fmt.Println("Game is valid", gameId)
			sum += gameId
		} else {
			fmt.Println("Game is invalid", gameId)
		}
	}

	fmt.Println(sum)
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseNumbers(numbers string) []int {
	result := []int{}
	for _, number := range strings.Split(numbers, " ") {
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

func solve(input []string) {
	sum := 0

	for _, line := range input {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Error: Invalid input, no ':'", line)
		}
		// cardIds := strings.Split(parts[0], " ")
		// if len(cardIds) != 2 {
		// 	fmt.Println("Error: Invalid input, no 'card X'", line)
		// }
		// cardId, err := strconv.Atoi(cardIds[1])
		// if err != nil {
		// 	fmt.Println("Error: Invalid input, invalid card id", line, err)
		// }

		numbers := strings.Split(parts[1], "|")
		if len(numbers) != 2 {
			fmt.Println("Error: Invalid input, no '|' in numbers", numbers, "line", line)
		}

		selectedNumbers := parseNumbers(numbers[0])
		draftedNumbers := parseNumbers(numbers[1])

		draftedMap := make(map[int]bool)
		for _, v := range draftedNumbers {
			draftedMap[v] = true
		}

		winningNumbers := []int{}
		for _, v := range selectedNumbers {
			if draftedMap[v] {
				winningNumbers = append(winningNumbers, v)

			}
		}
		fmt.Println("Winning number: ", winningNumbers)
		pointValue := 1
		for i := 0; i < len(winningNumbers); i++ {
			sum += pointValue
			if i > 0 {
				pointValue *= 2
			}
		}
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

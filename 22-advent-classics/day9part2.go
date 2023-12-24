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

func allZero(numbers []int) bool {
	for _, v := range numbers {
		if v != 0 {
			return false
		}
	}
	return true
}

func predict(numbers []int) int {
	pyramid := [][]int{}
	newNumbers := make([]int, len(numbers)+1)
	copy(newNumbers[1:], numbers)
	pyramid = append(pyramid, newNumbers)

	for !allZero(pyramid[len(pyramid)-1]) {
		prev := pyramid[len(pyramid)-1]
		newNumbers := make([]int, len(prev)-1)
		for i := 1; i < len(prev)-1; i++ {
			newNumbers[i] = prev[i+1] - prev[i]
		}
		pyramid = append(pyramid, newNumbers)
	}

	for i := len(pyramid) - 2; i >= 0; i-- {
		pyramid[i][0] = pyramid[i][1] - pyramid[i+1][0]
	}

	fmt.Println("Pyramid:")
	for _, numbers := range pyramid {
		fmt.Println(numbers)
	}

	return pyramid[0][0]
}

func solve(input []string) {
	result := 0
	for _, line := range input {
		numbers := parseNumbers(line)
		result += predict(numbers)
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

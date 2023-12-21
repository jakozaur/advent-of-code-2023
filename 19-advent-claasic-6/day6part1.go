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

func expectPrefix(line, expectedPrefix string) {
	if line[:len(expectedPrefix)] != expectedPrefix {
		fmt.Println("Error: Invalid input, expected prefix", expectedPrefix, "got", line)
	}
}

func solve(input []string) {
	expectPrefix(input[0], "Time:")
	expectPrefix(input[1], "Distance:")
	times := parseNumbers(input[0][len("Time:"):])
	distances := parseNumbers(input[1][len("Distance:"):])

	if len(times) != len(distances) {
		fmt.Println("Error: Invalid input, len(times) != len(distance)", input[0], input[1])
	}

	result := 1

	for i := range times {
		winningCombination := 0
		time := times[i]
		dist := distances[i]
		for hold := 1; hold < time; hold++ {
			traveled := hold * (time - hold)
			if traveled > dist {
				winningCombination += 1
			}
		}
		fmt.Println("Time", time, "Dist", dist, "Winning combinations", winningCombination)
		result *= winningCombination
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

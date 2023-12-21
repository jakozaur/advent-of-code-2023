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

func findFirstGoodHold(time, dist int64) int64 {
	for hold := int64(1); hold < time; hold++ {
		traveled := hold * (time - hold)
		if traveled > dist {
			return hold
		}
	}
	return -1
}

func findLastGoodHold(time, dist int64) int64 {
	for hold := time - 1; hold > 0; hold-- {
		traveled := hold * (time - hold)
		if traveled > dist {
			return hold
		}
	}
	return -1
}

func solve(input []string) {
	expectPrefix(input[0], "Time:")
	expectPrefix(input[1], "Distance:")
	timeStr := strings.ReplaceAll(input[0][len("Time:"):], " ", "")
	time, err := strconv.ParseInt(timeStr, 10, 64)
	distStr := strings.ReplaceAll(input[1][len("Distance:"):], " ", "")
	dist, err2 := strconv.ParseInt(distStr, 10, 64)
	if err != nil || err2 != nil {
		fmt.Println("Error: Invalid input, invalid number", err, err2, timeStr, distStr)
	}

	// define hold as int64
	winningCombination := findLastGoodHold(time, dist) - findFirstGoodHold(time, dist) + 1

	fmt.Println("Time", time, "Dist", dist, "Winning combinations", winningCombination)
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

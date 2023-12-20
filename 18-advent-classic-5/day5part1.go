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

type Translation struct {
	destination, source, length int
}

func solve(input []string) {
	expectPrefix(input[0], "seeds:")
	seeds := parseNumbers(input[0][len("seeds:"):])
	values := seeds

	fmt.Println("Initial seeds", values)

	expectPrefix(input[2], "seed-to-soil map:")
	idx := 2
	for idx < len(input) {
		idx += 1
		// read translations
		translations := []Translation{}
		for ; idx < len(input) && len(input[idx]) > 0; idx++ {
			nums := parseNumbers(input[idx])
			if len(nums) != 3 {
				fmt.Println("Error: Invalid input, expected 3 numbers", input[idx])
			}
			translations = append(translations, Translation{nums[0], nums[1], nums[2]})
		}

		// do translations
		newValues := make([]int, len(values))
		for j := range values {
			newValues[j] = -1
		}

		for j, v := range values {
			for _, t := range translations {
				if t.source <= v && t.source+t.length > v {
					newValues[j] = t.destination + (v - t.source)
				}
			}
		}

		for j := range values {
			if newValues[j] == -1 {
				newValues[j] = values[j]
			}
		}

		fmt.Println("New values after translations", newValues)
		values = newValues

		for idx < len(input) && len(input[idx]) == 0 {
			idx += 1
		}
	}

	minn := 1000000000
	for _, v := range values {
		if v < minn {
			minn = v
		}
	}
	fmt.Println("Lowest location number", minn)
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

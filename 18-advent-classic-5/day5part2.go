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

type ValueRange struct {
	start, length int
}

func (v ValueRange) match(t *Translation) bool {
	return !(v.start+v.length <= t.source || v.start >= t.source+t.length)
}

func (v ValueRange) translate(t *Translation) (ValueRange, []ValueRange) {
	start := max(v.start, t.source)
	end := min(v.start+v.length, t.source+t.length)
	startOffset := start - t.source
	length := end - start

	crossedRange := ValueRange{t.destination + startOffset, length}
	otherRanges := []ValueRange{}
	if v.start < start {
		otherRanges = append(otherRanges, ValueRange{v.start, start - v.start})
	}
	if v.start+v.length > end {
		otherRanges = append(otherRanges, ValueRange{end, v.start + v.length - end})
	}

	return crossedRange, otherRanges
}

func solve(input []string) {
	expectPrefix(input[0], "seeds:")
	seeds := parseNumbers(input[0][len("seeds:"):])
	values := []ValueRange{}
	for i := 0; i < len(seeds); i += 2 {
		values = append(values, ValueRange{seeds[i], seeds[i+1]})
	}

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
		newValues := []ValueRange{}

	main:
		for j := 0; j < len(values); j++ {
			for _, t := range translations {
				if values[j].match(&t) {
					crossedRange, otherRanges := values[j].translate(&t)
					newValues = append(newValues, crossedRange)
					for _, r := range otherRanges {
						values = append(values, r)
					}
					continue main
				}
			}
			newValues = append(newValues, values[j])
		}

		fmt.Println("New values after translations", newValues)
		values = newValues

		for idx < len(input) && len(input[idx]) == 0 {
			idx += 1
		}
	}

	minn := 1000000000
	for _, v := range values {
		if v.start < minn {
			minn = v.start
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

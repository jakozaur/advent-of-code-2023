package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lense struct {
	label string
	focus int
}

func calculateHash(input string) int {
	hash := 0
	for _, c := range input {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func findLabel(label string, box []Lense) int {
	for i, l := range box {
		if l.label == label {
			return i
		}
	}
	return -1
}

func solve(input []string) {
	strs := strings.Split(input[0], ",")
	boxes := make([][]Lense, 256)

	for _, s := range strs {
		newLense := strings.Split(s, "=")
		if len(newLense) == 2 {
			label := newLense[0]
			focus, err := strconv.Atoi(newLense[1])
			if err != nil {
				fmt.Println("Error: Invalid input, invalid number", err, newLense[1])
			}
			boxId := calculateHash(label)
			foundLabel := findLabel(label, boxes[boxId])
			if foundLabel == -1 {
				boxes[boxId] = append(boxes[boxId], Lense{label, focus})
			} else {
				boxes[boxId][foundLabel].focus = focus
			}
		} else {
			parts := strings.Split(s, "-")
			if len(parts) != 2 {
				fmt.Println("Error: Invalid input, no '-'", s)
			}
			label := parts[0]
			boxId := calculateHash(label)
			foundLabel := findLabel(label, boxes[boxId])
			if foundLabel != -1 {
				box := boxes[boxId]
				boxes[boxId] = append(box[:foundLabel:foundLabel], box[foundLabel+1:]...)
			}
		}
		//partialResult := calculateHash(s)
		//fmt.Println("Partial ", s, "result:", partialResult)
		//result += partialResult
	}

	result := 0
	for i, box := range boxes {
		for _, lense := range box {
			fmt.Println("Box", i, "label", lense.label, "focus", lense.focus)
			result += (i + 1)
		}
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

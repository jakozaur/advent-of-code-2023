package main

import (
	"bufio"
	"fmt"
	"os"
)

func findDigits(line string) int {
	firstDigit := -1
	lastDigit := -1
	for _, c := range line {
		if c >= '0' && c <= '9' {
			lastDigit = int(c) - int('0')
			if firstDigit == -1 {
				firstDigit = lastDigit
			}
		}
	}
	if firstDigit == -1 || lastDigit == -1 {
		fmt.Println("Error: no two digits found in ", line)
	}
	return firstDigit*10 + lastDigit
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += findDigits(line)
	}
	fmt.Println(sum)

}

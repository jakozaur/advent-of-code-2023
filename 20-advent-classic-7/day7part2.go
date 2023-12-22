package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	deck     string
	strength int
	bid      int
}

func newHand(line string) Hand {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		fmt.Println("Error: Invalid input, no two parts", len(parts), line)
	}
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Error: Invalid input, invalid bid", err, line)
	}

	hash := make(map[rune]int)
	for _, c := range parts[0] {
		hash[c] += 1
	}
	counts := []int{}
	jokers := 0
	for k, v := range hash {
		if k == 'J' {
			jokers = v
		} else {
			counts = append(counts, v)
		}
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	if jokers == 5 {
		counts = append(counts, 5)
	} else {
		counts[0] += jokers
	}

	strength := 0
	switch {
	case counts[0] == 5:
		strength = 0
	case counts[0] == 4:
		strength = 1
	case counts[0] == 3 && counts[1] == 2:
		strength = 2
	case counts[0] == 3:
		strength = 3
	case counts[0] == 2 && counts[1] == 2:
		strength = 4
	case counts[0] == 2:
		strength = 5
	default:
		strength = 6
	}

	fmt.Println("deck", parts[0], "counts", counts, "bid", bid)

	return Hand{parts[0], strength, bid}
}

func runeToStrength(r byte) int {
	switch r {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 0
	case 'T':
		return 10
	default:
		return int(r - '0')
	}
}

func compareHands(a, b Hand) bool {
	if a.strength == b.strength {
		for i := 0; i < len(a.deck); i++ {
			c1 := 100 - runeToStrength(a.deck[i])
			c2 := 100 - runeToStrength(b.deck[i])
			if c1 != c2 {
				return c1 > c2
			}
		}
		return false
	} else {
		return a.strength > b.strength
	}
}

func solve(input []string) {
	hands := []Hand{}
	for _, line := range input {
		hands = append(hands, newHand(line))
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})

	fmt.Println("hands", hands)

	sum := 0
	for i, v := range hands {
		sum += v.bid * (i + 1)
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

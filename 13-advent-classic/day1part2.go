package main

import (
	"bufio"
	"fmt"
	"os"
)

type State struct {
	line       string
	position   int
	firstDigit int
	lastDigit  int
}

func newState(line string) State {
	return State{line: line, position: 0, firstDigit: -1, lastDigit: -1}
}

func (s *State) updateDigit(digit int) {
	if s.firstDigit == -1 {
		s.firstDigit = digit
	}
	s.lastDigit = digit
}

func (s *State) parseDigitNew() {
	for s.position < len(s.line) {
		c := s.line[s.position]
		s.position += 1
		switch c {
		case '1':
			s.updateDigit(1)
		case '2':
			s.updateDigit(2)
		case '3':
			s.updateDigit(3)
		case '4':
			s.updateDigit(4)
		case '5':
			s.updateDigit(5)
		case '6':
			s.updateDigit(6)
		case '7':
			s.updateDigit(7)
		case '8':
			s.updateDigit(8)
		case '9':
			s.updateDigit(9)
		case 'o':
			if s.position+1 < len(s.line) && s.line[s.position] == 'n' {
				if s.line[s.position+1] == 'e' {
					s.updateDigit(1)
					s.position += 1 // it can be eight later
				}
			}
		case 't':
			if s.position+1 < len(s.line) && s.line[s.position] == 'w' {
				if s.line[s.position+1] == 'o' {
					s.updateDigit(2)
					s.position += 1 // it can be one later
				}
			} else if s.position+3 < len(s.line) && s.line[s.position] == 'h' {
				if s.line[s.position+1] == 'r' {
					if s.line[s.position+2] == 'e' {
						if s.line[s.position+3] == 'e' {
							s.updateDigit(3)
							s.position += 2 // it can be eight later
						}
					}
				}
			}

		case 'f':
			if s.position+2 < len(s.line) && s.line[s.position] == 'o' {
				if s.line[s.position+1] == 'u' {
					if s.line[s.position+2] == 'r' {
						s.updateDigit(4)
						s.position += 3
					}
				}
			} else if s.position+2 < len(s.line) && s.line[s.position] == 'i' {
				if s.line[s.position+1] == 'v' {
					if s.line[s.position+2] == 'e' {
						s.updateDigit(5)
						s.position += 2 // it can be eight later
					}
				}
			}
		case 's':
			if s.position+1 < len(s.line) && s.line[s.position] == 'i' {
				if s.line[s.position+1] == 'x' {
					s.updateDigit(6)
					s.position += 2
				}
			} else if s.position+3 < len(s.line) && s.line[s.position] == 'e' {
				if s.line[s.position+1] == 'v' {
					if s.line[s.position+2] == 'e' {
						if s.line[s.position+3] == 'n' {
							s.updateDigit(7)
							s.position += 3 // it can be nine later
						}
					}
				}
			}
		case 'e':
			if s.position+3 < len(s.line) && s.line[s.position] == 'i' {
				if s.line[s.position+1] == 'g' {
					if s.line[s.position+2] == 'h' {
						if s.line[s.position+3] == 't' {
							s.updateDigit(8)
							s.position += 3 // it can be three later
						}
					}
				}
			}
		case 'n':
			if s.position+2 < len(s.line) && s.line[s.position] == 'i' {
				if s.line[s.position+1] == 'n' {
					if s.line[s.position+2] == 'e' {
						s.updateDigit(9)
						s.position += 2 // it can be eight later
					}
				}
			}
		default:
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		state := newState(line)
		state.parseDigitNew()
		//fmt.Println(state.firstDigit, state.lastDigit)
		sum += state.firstDigit*10 + state.lastDigit
	}
	fmt.Println(sum)

}

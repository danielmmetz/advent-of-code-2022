package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (string, error) {
	var total int
	for i, line := range lines {
		ls, rs, ok := strings.Cut(line, " ")
		if !ok {
			return "", fmt.Errorf("unexped values in line %d: %s", i, line)
		}
		l, err := parse1(ls)
		if err != nil {
			return "", fmt.Errorf("parsing lhs for line %d: %w", i, err)
		}
		r, err := parse1(rs)
		if err != nil {
			return "", fmt.Errorf("parsing rhs for line %d: %w", i, err)
		}
		total += score(r, l)
	}
	return strconv.Itoa(total), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	var total int
	for i, line := range lines {
		ls, rs, ok := strings.Cut(line, " ")
		if !ok {
			return "", fmt.Errorf("unexped values in line %d: %s", i, line)
		}
		l, r, err := parse2(ls, rs)
		if err != nil {
			return "", fmt.Errorf("parsing line %d: %w", i, err)
		}
		total += score(r, l)
	}
	return strconv.Itoa(total), nil
}

func parse1(s string) (rps, error) {
	switch s {
	case "A", "X":
		return rock, nil
	case "B", "Y":
		return paper, nil
	case "C", "Z":
		return scissors, nil
	default:
		return rock, fmt.Errorf("unexpected rps value %q", s)
	}
}

func parse2(l, r string) (rps, rps, error) {
	var left rps
	switch l {
	case "A":
		left = rock
	case "B":
		left = paper
	case "C":
		left = scissors
	default:
		return rock, rock, fmt.Errorf("unexpected rps lhs value %q", l)
	}
	switch r {
	case "X":
		return left, lose(left), nil
	case "Y":
		return left, left, nil
	case "Z":
		return left, beat(left), nil
	default:
		return rock, rock, fmt.Errorf("unexpected rps rhs value %q", r)
	}
}

type rps string

const (
	rock     rps = "rock"
	paper    rps = "paper"
	scissors rps = "scissors"
)

func beats(l, r rps) bool {
	switch l {
	case rock:
		return r == scissors
	case paper:
		return r == rock
	case scissors:
		return r == paper
	default:
		panic(fmt.Errorf("exhuasted rps evaluation, got %v", l))
	}
}

func ties(l, r rps) bool {
	return l == r
}

func beat(x rps) rps {
	switch x {
	case rock:
		return paper
	case paper:
		return scissors
	case scissors:
		return rock
	default:
		panic(fmt.Errorf("unexpected rps value: %s", x))
	}
}

func lose(x rps) rps {
	switch x {
	case rock:
		return scissors
	case paper:
		return rock
	case scissors:
		return paper
	default:
		panic(fmt.Errorf("unexpected rps value: %s", x))
	}
}

func score(l, r rps) int {
	var shape int
	switch l {
	case rock:
		shape = 1
	case paper:
		shape = 2
	case scissors:
		shape = 3
	}
	switch {

	case beats(l, r):
		return 6 + shape
	case ties(l, r):
		return 3 + shape
	default:
		return shape
	}
}

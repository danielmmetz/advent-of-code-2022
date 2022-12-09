package day03

import (
	"fmt"
	"strconv"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (string, error) {
	var total int
	for i, line := range lines {
		left, right, err := split(line)
		if err != nil {
			return "", fmt.Errorf("split line %d: %w", i, err)
		}
		r, ok := intersection2(left, right)
		if !ok {
			return "", fmt.Errorf("no intersection in halfs in line %d", i)
		}
		score, err := score(r)
		if err != nil {
			return "", fmt.Errorf("unable to score for line %d: %w", i, err)
		}
		total += score
	}
	return strconv.Itoa(total), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	if len(lines)%3 != 0 {
		return "", fmt.Errorf("got %d lines which is not divisible by 3", len(lines))
	}
	var total int
	for i := 0; i < len(lines); i += 3 {
		a, b, c := lines[i], lines[i+1], lines[i+2]
		r, ok := intersection3(a, b, c)
		if !ok {
			return "", fmt.Errorf("no intersection in halfs in line %d", i)
		}
		score, err := score(r)
		if err != nil {
			return "", fmt.Errorf("unable to score for line %d: %w", i, err)
		}
		total += score
	}
	return strconv.Itoa(total), nil
}

func split(s string) (left, right string, err error) {
	if len(s)%2 != 0 {
		return "", "", fmt.Errorf("non-even string length: %d", len(s))
	}
	half := len(s) / 2
	return s[:half], s[half:], nil
}

func score(r rune) (int, error) {
	if int('a') <= int(r) && int(r) <= int('z') {
		return int(r) - int('a') + 1, nil
	}
	if int('A') <= int(r) && int(r) <= int('Z') {
		return int(r) - int('A') + 27, nil
	}
	return 0, fmt.Errorf("unexpected input: %v", string(r))
}

// intersection2 returns a rune that's present in both left and right.
// Results may not be stable.
func intersection2(left, right string) (r rune, ok bool) {
	m := make(map[rune]bool)
	for _, r := range left {
		m[r] = true
	}
	for _, r := range right {
		if m[r] {
			return r, true
		}
	}
	return ' ', false
}

// intersection3 returns a rune that's present in a, b, and c
// Results may not be stable.
func intersection3(a, b, c string) (r rune, ok bool) {
	am, bm := make(map[rune]bool), make(map[rune]bool)
	for _, r := range a {
		am[r] = true
	}
	for _, r := range b {
		bm[r] = true
	}
	for _, r := range c {
		if am[r] && bm[r] {
			return r, true
		}
	}
	return ' ', false
}

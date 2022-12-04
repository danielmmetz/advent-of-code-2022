package day01

import (
	"strings"
	"testing"
)

var input = strings.Split(strings.TrimSpace(`
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`), "\n")

func TestPart1(t *testing.T) {
	t.Parallel()

	result, err := Solution{}.Part1(input)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if result != "24000" {
		t.Errorf("got %s, wanted 24000", result)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result, err := Solution{}.Part2(input)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if result != "45000" {
		t.Errorf("got %s, wanted 45000", result)
	}
}

package day01

import (
	"os"
	"strings"
	"testing"
)

func input(t *testing.T) []string {
	t.Helper()

	b, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("read input.txt: %v", err)
	}
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}

func TestPart1(t *testing.T) {
	t.Parallel()

	result, err := Solution{}.Part1(input(t))
	if err != nil {
		t.Fatalf(err.Error())
	}
	if result != "24000" {
		t.Errorf("got %s, wanted 24000", result)
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	result, err := Solution{}.Part2(input(t))
	if err != nil {
		t.Fatalf(err.Error())
	}
	if result != "45000" {
		t.Errorf("got %s, wanted 45000", result)
	}
}

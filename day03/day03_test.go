package day03

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
	var s Solution
	got, err := s.Part1(input(t))
	if err != nil {
		t.Fatal(err)
	}
	if got != "157" {
		t.Errorf("got %s, want 157", got)
	}
}

func TestPart2(t *testing.T) {
	var s Solution
	got, err := s.Part2(input(t))
	if err != nil {
		t.Fatal(err)
	}
	if got != "70" {
		t.Errorf("got %s, want 70", got)
	}
}

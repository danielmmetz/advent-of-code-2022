package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com.danielmmetz/advent-of-code-2022/day01"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	if err := mainE(ctx); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		// Only exit non-zero if our initial context has yet to be canceled.
		// Otherwise it's very likely that the error we're seeing is a result of our attempt at graceful shutdown.
		if ctx.Err() == nil {
			os.Exit(1)
		}
	}
}

func mainE(_ context.Context) error {
	day := flag.Int("day", 0, "which day's puzzle to run")
	part := flag.Int("part", 1, "which puzzle part to run")
	flag.Parse()

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("reading stdin: %w", err)
	}
	lines := strings.Split(string(b), "\n")

	var s any
	switch *day {
	case 1:
		s = day01.Solution{}
	}

	var result string
	switch *part {
	case 1:
		result, err = runPart1(s, lines)
	case 2:
		result, err = runPart2(s, lines)
	}
	if err != nil {
		return fmt.Errorf("day %d part %d: %w", *day, *part, err)
	}
	fmt.Println(result)
	return nil
}

func runPart1(s any, lines []string) (string, error) {
	p, ok := s.(Part1er)
	if !ok {
		return "", fmt.Errorf("not implemented")
	}
	return p.Part1(lines)
}

func runPart2(s any, lines []string) (string, error) {
	p, ok := s.(Part2er)
	if !ok {
		return "", fmt.Errorf("not implemented")
	}
	return p.Part2(lines)
}

type Part1er interface {
	Part1(lines []string) (string, error)
}

type Part2er interface {
	Part2(lines []string) (string, error)
}

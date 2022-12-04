package day01

import (
	"container/heap"
	"fmt"
	"strconv"
)

type Solution struct {
}

func (s Solution) Part1(lines []string) (string, error) {
	return solve(lines, 1)
}

func (s Solution) Part2(lines []string) (string, error) {
	return solve(lines, 3)
}

func solve(lines []string, top int) (string, error) {
	var ih maxHeap
	var current int
	for i, line := range lines {
		if line == "" {
			if current != 0 {
				heap.Push(&ih, current)
			}
			current = 0
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			return "", fmt.Errorf("line %d: converting to int: %w", i, err)
		}
		current += val
	}
	heap.Push(&ih, current)

	var total int
	for i := 0; i < top; i++ {
		total += heap.Pop(&ih).(int)
	}
	return strconv.Itoa(total), nil
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

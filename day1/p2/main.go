package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/mycreepy/aoc2022/common"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x int) {
	*h = append(*h, x)
	sort.Sort(h)
	for i, j := 0, h.Len()-1; i < j; i, j = i+1, j-1 {
		h.Swap(i, j)
	}
}

func (h *IntHeap) Pop() int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	lines, err := common.InputFileChan("day1/input.txt")
	if err != nil {
		panic(err)
	}

	bestElfs := make(IntHeap, 0)

	var current int

	for line := range lines {
		if line == "" {
			bestElfs.Push(current)
			current = 0

			if bestElfs.Len() > 3 {
				bestElfs.Pop()
			}

			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		current += val
	}

	fmt.Println(bestElfs.Pop() + bestElfs.Pop() + bestElfs.Pop())
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	elfs := make([]int, 0)

	var current int

	for scanner.Scan() {
		if scanner.Text() == "" {
			elfs = append(elfs, current)

			current = 0

			continue
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		current += val
	}

	sort.Ints(elfs)

	l := len(elfs) - 1

	fmt.Println(elfs[l] + elfs[l-1] + elfs[l-2])
}

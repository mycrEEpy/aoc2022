package main

import (
	"fmt"
	"strconv"

	"github.com/mycreepy/aoc2022/common"
)

func main() {
	lines, err := common.InputFileChan("day1/input.txt")
	if err != nil {
		panic(err)
	}

	var current, max int

	for line := range lines {
		if line == "" {
			if current > max {
				max = current
			}

			current = 0

			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		current += val
	}

	fmt.Println(max)
}

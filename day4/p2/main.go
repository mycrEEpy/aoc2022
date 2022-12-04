package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mycreepy/aoc2022/common"
)

func main() {
	lines, err := common.InputFileChan("day1/input.txt")
	if err != nil {
		panic(err)
	}

	var sum int

	for line := range lines {
		pairs := strings.Split(line, ",")

		contained, err := isContained(pairs[0], pairs[1])
		if err != nil {
			panic(err)
		}

		if contained {
			sum++
		}
	}

	fmt.Println(sum)
}

func isContained(a, b string) (bool, error) {
	aMin, aMax, err := splitAndConvertToInt(a)
	if err != nil {
		return false, err
	}

	bMin, bMax, err := splitAndConvertToInt(b)
	if err != nil {
		return false, err
	}

	// partial overlap
	if ((bMin >= aMin && bMin <= aMax) ||
		(bMax <= aMax && bMax >= aMin)) ||
		((aMin >= bMin && aMin <= bMax) ||
			(aMax <= bMax && aMax >= bMax)) {
		return true, nil
	}

	return false, nil
}

func splitAndConvertToInt(s string) (int, int, error) {
	split := strings.Split(s, "-")

	a, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}

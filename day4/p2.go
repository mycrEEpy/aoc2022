package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")

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

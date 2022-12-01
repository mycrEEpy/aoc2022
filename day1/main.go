package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var current, max int

	for scanner.Scan() {
		if scanner.Text() == "" {
			if current > max {
				max = current
			}

			current = 0

			continue
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		current += val
	}

	fmt.Println(max)
}

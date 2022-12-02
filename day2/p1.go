package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	win     = 6
	loss    = 0
	draw    = 3
	rock    = 1
	paper   = 2
	scissor = 3
)

func main() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var score int

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		enemy := split[0]
		me := split[1]

		switch enemy {
		case "A":
			switch me {
			case "X":
				score += draw + rock
			case "Y":
				score += win + paper
			case "Z":
				score += loss + scissor
			}
		case "B":
			switch me {
			case "X":
				score += loss + rock
			case "Y":
				score += draw + paper
			case "Z":
				score += win + scissor
			}
		case "C":
			switch me {
			case "X":
				score += win + rock
			case "Y":
				score += loss + paper
			case "Z":
				score += draw + scissor
			}
		}
	}

	fmt.Println(score)
}

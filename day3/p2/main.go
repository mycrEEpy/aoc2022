package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var sum int
	group := make([]string, 0, 3)

	for scanner.Scan() {
		group = append(group, scanner.Text())

		if len(group) < 3 {
			continue
		}

		for _, c := range group[0] {
			if strings.Contains(group[1], string(c)) && strings.Contains(group[2], string(c)) {
				switch {
				case c >= 65 && c <= 90:
					sum += int(c) - 38
				case c >= 97 && c <= 122:
					sum += int(c) - 96
				default:
					panic(fmt.Errorf("unexpected character: %d (%s)", c, string(c)))
				}

				break
			}
		}

		group = make([]string, 0, 3)
	}

	fmt.Println(sum)
}

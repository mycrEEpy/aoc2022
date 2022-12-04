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

	for scanner.Scan() {
		l := len(scanner.Text())
		first := scanner.Text()[:l/2]
		second := scanner.Text()[l/2:]

		for _, c := range first {
			if strings.Contains(second, string(c)) {
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
	}

	fmt.Println(sum)
}

package main

import (
	"fmt"
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
		l := len(line)
		first := line[:l/2]
		second := line[l/2:]

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

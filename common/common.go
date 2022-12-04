package common

import (
	"bufio"
	"os"
)

func InputFileChan(filepath string) (<-chan string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	resultChan := make(chan string)

	go func() {
		defer close(resultChan)

		for scanner.Scan() {
			resultChan <- scanner.Text()
		}
	}()

	return resultChan, nil
}

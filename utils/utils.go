package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInts(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	ints := []int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}

	return ints, nil
}

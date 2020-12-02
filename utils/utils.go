package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	s := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	return s, nil
}

func ReadInts(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	i := []int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		i = append(i, n)
	}

	return i, nil
}

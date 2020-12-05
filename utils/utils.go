package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

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
	defer f.Close()

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

func ReadContents(filename string) (string, error) {
	contents, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func ReplaceAtIndex(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

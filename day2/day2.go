package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var parser = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

type Policy struct {
	Check1 int
	Check2 int
	Letter string
}

type Password string

func (p Password) ToString() string {
	return string(p)
}

func main() {
	entries, err := utils.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Solution := part1(entries)
	part2Solution := part2(entries)

	fmt.Printf("Day 2: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 2: Part 2: = %+v\n", part2Solution)
}

func parseParts(line string) (Policy, Password) {
	parts := parser.FindStringSubmatch(line)

	check1, _ := strconv.Atoi(parts[1])
	check2, _ := strconv.Atoi(parts[2])
	letter := parts[3]
	password := parts[4]

	return Policy{Check1: check1, Check2: check2, Letter: letter}, Password(password)
}

func validatePasswordSledRental(policy Policy, password Password) bool {
	count := strings.Count(password.ToString(), policy.Letter)
	return count >= policy.Check1 && count <= policy.Check2
}

func validatePasswordOfficialTobogan(policy Policy, password Password) bool {
	return (string(password.ToString()[policy.Check1-1]) == policy.Letter) != (string(password.ToString()[policy.Check2-1]) == policy.Letter)
}

func part1(entries []string) int {
	validCount := 0
	for _, line := range entries {
		policy, password := parseParts(line)
		if validatePasswordSledRental(policy, password) {
			validCount++
		}
	}
	return validCount
}

func part2(entries []string) int {
	validCount := 0
	for _, line := range entries {
		policy, password := parseParts(line)
		if validatePasswordOfficialTobogan(policy, password) {
			validCount++
		}
	}
	return validCount
}

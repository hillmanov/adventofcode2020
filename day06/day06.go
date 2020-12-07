package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	contents, err := utils.ReadContents("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Solution := part1(contents)
	part2Solution := part2(contents)

	fmt.Printf("Day 06: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 06: Part 2: = %+v\n", part2Solution)
}

func part1(contents string) int {
	groups := flattenGroups(contents)

	uniqGroupSum := 0
	for _, group := range groups {
		uniqGroupSum += getUniqueAnswers(group)
	}

	return uniqGroupSum
}

func part2(contents string) int {
	groupedGroups := groupedGroups(contents)

	totalSum := 0
	for _, group := range groupedGroups {
		answerCounts := make(map[rune]int)
		for _, answer := range group {
			for _, letter := range answer {
				answerCounts[letter] += 1
			}
		}
		for _, count := range answerCounts {
			if count == len(group) {
				totalSum++
			}
		}
	}

	return totalSum
}

func flattenGroups(contents string) []string {
	block := regexp.MustCompile(`(?m)^\n`)
	groups := block.Split(contents, -1)

	flatten := regexp.MustCompile(`\s`)
	for i, group := range groups {
		groups[i] = flatten.ReplaceAllString(strings.TrimSpace(group), "")
	}
	return groups
}

func groupedGroups(contents string) [][]string {
	block := regexp.MustCompile(`(?m)^\n`)
	groups := block.Split(contents, -1)

	flatten := regexp.MustCompile(`\s`)
	groupedGroups := make([][]string, len(groups))
	for i, group := range groups {
		for _, line := range flatten.Split(group, -1) {
			if len(line) > 0 {
				groupedGroups[i] = append(groupedGroups[i], strings.TrimSpace(line))
			}
		}
	}
	return groupedGroups
}

func getUniqueAnswers(s string) int {
	u := make(map[string]bool)
	for _, letter := range s {
		u[string(letter)] = true
	}

	return len(u)
}

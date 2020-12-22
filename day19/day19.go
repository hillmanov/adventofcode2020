package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

type node [][]string

func main() {
	input, _ := utils.ReadLines("./input.txt")
	a, b, nodeMap, messages := parseInput(input)

	part1Solution := part1(a, b, nodeMap, messages)
	part2Solution := part2(a, b, nodeMap, messages)

	fmt.Printf("Day 19: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 19: Part 2: = %+v\n", part2Solution)
}

func part1(a, b string, nodeMap map[string]node, messages [][]string) int {
	count := 0
	for _, message := range messages {
		match := match(message, nodeMap["0"][0], nodeMap)
		if match {
			count++
		}
	}

	return count
}

func part2(a, b string, nodeMap map[string]node, messages [][]string) int {
	nodeMap["8"] = [][]string{{"42"}, {"42", "8"}}
	nodeMap["11"] = [][]string{{"42", "31"}, {"42", "11", "31"}}

	count := 0
	for _, message := range messages {
		match := match(message, nodeMap["0"][0], nodeMap)
		if match {
			count++
		}
	}

	return count
}

func match(message []string, pattern []string, nodeMap map[string]node) bool {
	if len(message) == 0 && len(pattern) == 0 {
		return true
	} else if (len(message) == 0 && len(pattern) > 0) || (len(message) > 0 && len(pattern) == 0) {
		return false
	} else if n, ok := nodeMap[string(pattern[0])]; ok {
		for i := 0; i < len(n); i++ {
			checkPattern := append(n[i], pattern[1:]...)
			if match(message, checkPattern, nodeMap) {
				return true
			}
		}
		return false
	} else if string(message[0]) == pattern[0] {
		return match(message[1:], pattern[1:], nodeMap)
	}
	return false
}

func parseInput(input []string) (a, b string, nodeMap map[string]node, messages [][]string) {
	nodeMap = map[string]node{}
	lineIndex := 0
	for i, line := range input {
		lineIndex = i
		if len(line) == 0 { // Rest is messages
			break
		}

		parts := strings.Split(line, ":")
		nodeID := strings.TrimSpace(parts[0])
		pattern := strings.TrimSpace(parts[1])

		if strings.Contains(line, "a") {
			a = nodeID
			continue
		}
		if strings.Contains(line, "b") {
			b = nodeID
			continue
		}

		node := [][]string{}
		for _, p := range strings.Split(pattern, "|") {
			sp := []string{}
			for _, pp := range strings.Split(p, " ") {
				if len(strings.TrimSpace(pp)) > 0 {
					sp = append(sp, strings.TrimSpace(pp))
				}
			}
			node = append(node, sp)
		}
		nodeMap[nodeID] = node
	}

	for _, line := range input[lineIndex+1:] {
		message := []string{}
		for _, r := range line {
			if string(r) == "a" {
				message = append(message, a)
			}
			if string(r) == "b" {
				message = append(message, b)
			}
		}
		messages = append(messages, message)
	}

	return a, b, nodeMap, messages
}

package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Solution := part1(lines)
	part2Solution := part2(lines)

	fmt.Printf("Day 17: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 17: Part 2: = %+v\n", part2Solution)
}

func part1(lines []string) int {
	pd := parseInput3D(lines)
	for i := 0; i < 6; i++ {
		pd.Cycle()
	}
	return pd.TotalActiveCount()
}

func part2(lines []string) int {
	pd := parseInput4D(lines)
	for i := 0; i < 6; i++ {
		pd.Cycle()
	}
	return pd.TotalActiveCount()
}

package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	mapPattern, err := utils.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Solution := part1(mapPattern)
	part2Solution := part2(mapPattern)

	fmt.Printf("Day 03: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 03: Part 2: = %+v\n", part2Solution)
}

func part1(mapPattern []string) int {
	return runMap(mapPattern, 3, 1)
}

func part2(mapPattern []string) int {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treeProduct := 1
	for _, slope := range slopes {
		treeProduct *= runMap(mapPattern, slope[0], slope[1])
	}

	return treeProduct
}

func runMap(mapPattern []string, slopeCol, slopeRow int) int {
	row := 0
	col := 0
	treeCount := 0

	for row < len(mapPattern) {
		rowPattern := mapPattern[row]
		if string(rowPattern[col]) == "#" {
			treeCount++
		}

		col = (col + slopeCol) % len(rowPattern)
		row = row + slopeRow
	}

	return treeCount
}

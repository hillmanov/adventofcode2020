package main

import (
	"adventofcode/utils"
	"fmt"
	"sort"
)

func main() {
	jolts, err := utils.ReadInts("./input.txt")
	if err != nil {
		panic(err)
	}

	// Add the charging port and final adapter

	part1Solution := part1(jolts)
	part2Solution := part2(jolts)

	fmt.Printf("Day 10: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 10: Part 2: = %+v\n", part2Solution)
}

func part1(jolts []int) int {
	jolts = append(jolts, 0)
	jolts = append(jolts, utils.MaxOf(jolts)+3)
	sort.Slice(jolts, func(i, j int) bool {
		return jolts[i] < jolts[j]
	})

	oneJoltDiffs := 0
	threeJoltDiffs := 0
	for i := 1; i < len(jolts); i++ {
		if jolts[i]-jolts[i-1] == 1 {
			oneJoltDiffs++
		}
		if jolts[i]-jolts[i-1] == 3 {
			threeJoltDiffs++
		}
	}

	return oneJoltDiffs * threeJoltDiffs
}

func part2(jolts []int) int {
	jolts = append(jolts, 0)
	jolts = append(jolts, utils.MaxOf(jolts)+3)

	sort.Slice(jolts, func(i, j int) bool {
		return jolts[i] < jolts[j]
	})

	// No idea how or why this works. Taken from https://github.com/jhillman/advent-of-code-2020/blob/master/day10/part2.c
	arrangements := make([]int, utils.MaxOf(jolts)+1)
	arrangements[0] = 1

	for i := 0; i < len(jolts)-1; i++ {
		joltage := jolts[i]

		for diff := 1; diff <= 3; diff++ {
			if joltage-diff >= 0 {
				arrangements[joltage] += arrangements[joltage-diff]
			}
		}
	}

	return utils.MaxOf(arrangements)
}

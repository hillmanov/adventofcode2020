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

	// part1Solution := part1(jolts)
	part2Solution := part2(jolts)

	// fmt.Printf("Day 10: Part 1: = %+v\n", part1Solution)
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
	// Get all permutation ranges

	sort.Slice(jolts, func(i, j int) bool {
		return jolts[i] < jolts[j]
	})

	fmt.Printf("jolts = %+v\n", jolts)

	// jolts = [1 2 3 4 7 8 9 10 11 14 17 18 19 20 23 24 25 28 31 32 33 34 35 38 39 42 45 46 47 48 49]
	options := make([]int, len(jolts))

	for i := 1; i < len(jolts); i++ {

	}
	fmt.Printf("options = %+v\n", options)

	return -1
}

/*

0 1 4 5 6 7
0 1 1 1 2 4

*/

// [0 1 4 5 6 7 10 11 12 15 16

/*
(0), 1, (4, 5, 6, 7), (10, 11, 12), 15, 16, 19, (22)
(0), 1, (4, 5, 6, 7), (10,     12), 15, 16, 19, (22)

(0), 1, (4, 5,    7), (10, 11, 12), 15, 16, 19, (22)
(0), 1, (4,    6, 7), (10, 11, 12), 15, 16, 19, (22)
(0), 1, (4,       7), (10, 11, 12), 15, 16, 19, (22)

(0), 1, (4, 5,    7), (10,     12), 15, 16, 19, (22)
(0), 1, (4,    6, 7), (10,     12), 15, 16, 19, (22)
(0), 1, (4,       7), (10,     12), 15, 16, 19, (22)
*/

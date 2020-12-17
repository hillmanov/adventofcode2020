package main

import (
	"fmt"
)

func main() {
	part1Solution := play([]int{2, 0, 1, 9, 5, 19}, 2020)
	part2Solution := play([]int{2, 0, 1, 9, 5, 19}, 30000000)

	fmt.Printf("Day 15: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 15: Part 2: = %+v\n", part2Solution)
}

func play(seed []int, target int) int {
	turnsSpoken := make(map[int][]int)
	var lastNumberSpoken int

	for turn, v := range seed {
		turnsSpoken[v] = append(turnsSpoken[v], turn+1)
		lastNumberSpoken = v
	}

	for turns := len(seed); turns < target; turns++ {
		turnsLastNumberHasBeenSpoken := turnsSpoken[lastNumberSpoken]

		if len(turnsLastNumberHasBeenSpoken) <= 1 {
			turnsSpoken[0] = []int{turnsSpoken[0][len(turnsSpoken[0])-1], turns + 1}
			lastNumberSpoken = 0
		} else {
			numberToSpeak := turnsLastNumberHasBeenSpoken[len(turnsLastNumberHasBeenSpoken)-1] - turnsLastNumberHasBeenSpoken[len(turnsLastNumberHasBeenSpoken)-2]
			turnsSpoken[numberToSpeak] = append(turnsSpoken[numberToSpeak], turns+1)
			lastNumberSpoken = numberToSpeak
		}
	}

	return lastNumberSpoken
}

package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	entries, _ := utils.IntsFromFile("./input.txt")

	fmt.Printf("Part 1: = %+v\n", part1(entries))
	fmt.Printf("Part 2: = %+v\n", part2(entries))
}

func part1(entries []int) int {
	var result int
	for i, entryI := range entries {
		for _, entryJ := range entries[i+1:] {
			if entryI+entryJ == 2020 {
				result = entryI * entryJ
				goto Return
			}
		}
	}
Return:
	return result
}

func part2(entries []int) int {
	var result int
	for i := 0; i < len(entries); i++ {
		for j := i + 1; j < len(entries); j++ {
			for k := j + 1; k < len(entries); k++ {
				if entries[i]+entries[j]+entries[k] == 2020 {
					result = entries[i] * entries[j] * entries[k]
					goto Return
				}
			}
		}
	}
Return:
	return result
}

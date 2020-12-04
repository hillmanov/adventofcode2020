package main

import (
	"adventofcode/utils"
	"errors"
	"fmt"
)

func main() {
	targetSum := 2020

	entries, err := utils.ReadInts("./input.txt")
	if err != nil {
		panic(err)
	}

	part1, err := findDouble(entries, targetSum)
	if err != nil {
		panic(err)
	}

	part2, err := findTriple(entries, targetSum)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 01: Part 1: = %+v\n", part1)
	fmt.Printf("Day 01: Part 2: = %+v\n", part2)
}

// Part 1
func findDouble(entries []int, target int) (int, error) {
	for i, entryI := range entries {
		for _, entryJ := range entries[i+1:] {
			if entryI+entryJ == target {
				return entryI * entryJ, nil
			}
		}
	}
	return -1, errors.New("No solution found")
}

// Part 2
func findTriple(entries []int, target int) (int, error) {
	for i := 0; i < len(entries); i++ {
		for j := i + 1; j < len(entries); j++ {
			for k := j + 1; k < len(entries); k++ {
				if entries[i]+entries[j]+entries[k] == target {
					return entries[i] * entries[j] * entries[k], nil
				}
			}
		}
	}
	return -1, errors.New("No solution found")
}

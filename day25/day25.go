package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
)

func main() {
	entries, _ := utils.ReadLines("./input.txt")
	cardPublicKeyStr, doorPublicKeyStr := entries[1], entries[0]
	cardPublicKey, _ := strconv.Atoi(cardPublicKeyStr)
	doorPublicKey, _ := strconv.Atoi(doorPublicKeyStr)

	part1Solution := part1(cardPublicKey, doorPublicKey)
	part2Solution := part2(cardPublicKey, doorPublicKey)

	fmt.Printf("Day 25: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 25: Part 2: = %+v\n", part2Solution)
}

func part1(cardPublicKey, doorPublicKey int) int {
	cardLoopSize := getLoopSize(7, cardPublicKey)
	return transform(doorPublicKey, cardLoopSize)
}

func part2(cardPublicKey, doorPublicKey int) int {
	return -1
}

func getLoopSize(subjectNumber int, target int) int {
	value := 1

	loopSize := 0
	for value != target {
		value *= subjectNumber
		value = value % 20201227
		loopSize++
	}

	return loopSize
}

func transform(subjectNumber int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subjectNumber
		value = value % 20201227
	}

	return value
}

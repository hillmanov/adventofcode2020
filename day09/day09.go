package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	numbers, err := utils.ReadInts("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Solution := part1(numbers, 25)
	part2Solution := part2(part1Solution, numbers)

	fmt.Printf("Day 09: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 09: Part 2: = %+v\n", part2Solution)
}

func part1(numbers []int, preambleLength int) int {
	for currentIndex := preambleLength; currentIndex < len(numbers); currentIndex++ {
		number := numbers[currentIndex]
		previousNumbers := numbers[currentIndex-preambleLength : currentIndex]

		if !isSumOfAnyTwoUniquePreviousNumbers(number, previousNumbers) {
			return number
		}
	}

	return -1
}

func part2(target int, numbers []int) int {
	// Get the index of the target
	var targetIndex int
	for i, number := range numbers {
		if number == target {
			targetIndex = i
		}
	}

	candidateNumbers := numbers[:targetIndex]

	nums := findContinuousNumbersThatAddTo(target, candidateNumbers)
	min, max := utils.MinMax(nums)

	return min + max
}

func isSumOfAnyTwoUniquePreviousNumbers(number int, previousNumbers []int) bool {
	for i := 0; i < len(previousNumbers); i++ {
		for j := i; j < len(previousNumbers); j++ {
			if previousNumbers[i] != previousNumbers[j] && previousNumbers[i]+previousNumbers[j] == number {
				return true
			}
		}
	}
	return false
}

func findContinuousNumbersThatAddTo(target int, numbers []int) []int {
	for startIndex := 0; startIndex < len(numbers); startIndex++ {
		for endIndex := startIndex + 1; endIndex < len(numbers); endIndex++ {
			candidateNumbers := numbers[startIndex:endIndex]
			sum := utils.SumOf(candidateNumbers)
			if sum == target {
				return candidateNumbers
			}
			if sum > target {
				break
			}
		}
	}
	return nil
}

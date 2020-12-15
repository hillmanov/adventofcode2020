package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Solution := part1(lines)
	part2Solution := part2(lines)

	fmt.Printf("Day 14: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 14: Part 2: = %+v\n", part2Solution)
}

func part1(lines []string) int {
	mem := make(map[int]int)

	var mask []rune
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			var maskValue string
			fmt.Sscanf(line, "mask = %v", &maskValue)
			mask = []rune(maskValue)
			continue
		}

		var address int
		var value int
		fmt.Sscanf(line, "mem[%d] = %d", &address, &value)
		binaryValue := fmt.Sprintf("%036s", strconv.FormatInt(int64(value), 2))

		maskedBinary := ""
		for i := 0; i < len(mask); i++ {
			if mask[i] == '0' || mask[i] == '1' {
				maskedBinary = maskedBinary + string(mask[i])
			} else {
				maskedBinary = maskedBinary + string(binaryValue[i])
			}
		}
		maskedValue, _ := strconv.ParseInt(maskedBinary, 2, 64)
		mem[address] = int(maskedValue)
	}

	sum := 0
	for _, value := range mem {
		sum += value
	}

	return sum
}

func part2(lines []string) int {
	mem := make(map[int]int)

	var mask []rune
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			var maskValue string
			fmt.Sscanf(line, "mask = %v", &maskValue)
			mask = []rune(maskValue)
			continue
		}

		var address int
		var value int

		fmt.Sscanf(line, "mem[%d] = %d", &address, &value)
		binaryValue := fmt.Sprintf("%036s", strconv.FormatInt(int64(address), 2))

		maskedBinary := ""
		for i := 0; i < len(mask); i++ {
			switch mask[i] {
			case '0':
				maskedBinary = maskedBinary + string(binaryValue[i])
			case '1':
				maskedBinary = maskedBinary + "1"
			case 'X':
				maskedBinary = maskedBinary + "X"
			}
		}

		for _, p := range permutations(maskedBinary) {
			permValue, _ := strconv.ParseInt(p, 2, 64)
			mem[int(permValue)] = value
		}
	}

	sum := 0
	for _, value := range mem {
		sum += value
	}

	return sum
}

func permutations(s string) []string {
	perms := []string{}

	indexesOfX := []int{}
	for index, char := range s {
		if char == 'X' {
			indexesOfX = append(indexesOfX, index)
		}
	}

	amountOfX := len(indexesOfX)
	amountOfPerms := int(math.Pow(2, float64(amountOfX)))

	for i := 0; i < amountOfPerms; i++ {
		perm := ""
		xIndex := amountOfX - 1
		for _, char := range s {
			if char == 'X' {
				permChar := strconv.Itoa(int(math.Floor(float64(i)/math.Pow(2, float64(xIndex)))) % 2)
				xIndex--
				perm = perm + permChar
			} else {
				perm = perm + string(char)
			}

		}
		perms = append(perms, perm)
	}
	return perms
}

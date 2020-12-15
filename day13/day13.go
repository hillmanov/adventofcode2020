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

	earliestDepartureTime, busses := parseInput(lines)

	part1Solution := part1(earliestDepartureTime, busses)
	part2Solution := part2()

	fmt.Printf("Day 13: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 13: Part 2: = %+v\n", part2Solution)
}

func part1(earliestDepartureTime int, busses []int) int {
	var bestBus int
	bestDepartureTimeDiff := math.MaxInt64

	for _, bus := range busses {
		nextAvailableDepartureTime := int(math.Ceil(float64(earliestDepartureTime)/float64(bus)) * float64(bus))
		departureTimeDiff := nextAvailableDepartureTime - earliestDepartureTime
		if departureTimeDiff < bestDepartureTimeDiff {
			bestDepartureTimeDiff = departureTimeDiff
			bestBus = bus
		}
	}

	return bestDepartureTimeDiff * bestBus
}

func part2() int {
	return -1
}

func parseInput(lines []string) (int, []int) {
	earliestDepartureTime, _ := strconv.Atoi(lines[0])
	busses := []int{}

	busEntries := strings.Split(lines[1], ",")
	for _, busEntry := range busEntries {
		if bus, err := strconv.Atoi(busEntry); err == nil {
			busses = append(busses, bus)
		}
	}

	return earliestDepartureTime, busses
}

package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"sort"
)

const (
	planeRows  = 128
	planeCols  = 8
	seatIDSalt = 8
)

func main() {
	boardingPasses, err := utils.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Solution := part1(boardingPasses)
	part2Solution := part2(boardingPasses)

	fmt.Printf("Day 05: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 05: Part 2: = %+v\n", part2Solution)
}

func part1(boardingPasses []string) int {
	maxSeatID := math.MinInt64
	for _, bp := range boardingPasses {
		rowSpecification := bp[:7]
		colSpecification := bp[7:]
		row, col := getColumnAndRow(rowSpecification, colSpecification)
		seatID := getSeatID(row, col, seatIDSalt)
		maxSeatID = utils.MaxInt(maxSeatID, seatID)
	}

	return maxSeatID
}

func part2(boardingPasses []string) int {
	seatIDs := []int{}

	for _, bp := range boardingPasses {
		rowSpecification := bp[:7]
		colSpecification := bp[7:]
		row, col := getColumnAndRow(rowSpecification, colSpecification)
		seatIDs = append(seatIDs, getSeatID(row, col, seatIDSalt))
	}

	sort.Slice(seatIDs, func(i, j int) bool {
		return seatIDs[i] < seatIDs[j]
	})

	offset := seatIDs[0]
	for i, seatID := range seatIDs {
		if seatID-offset != i {
			return seatID - 1
		}
	}

	return -1
}

func getColumnAndRow(rowSpecification string, colSpecification string) (int, int) {
	rows := makeList(planeRows)
	cols := makeList(planeCols)

	for _, r := range rowSpecification {
		switch r {
		case 'F':
			rows = rows[:len(rows)/2]
		case 'B':
			rows = rows[len(rows)/2:]
		}
	}

	for _, c := range colSpecification {
		switch c {
		case 'R':
			cols = cols[len(cols)/2:]
		case 'L':
			cols = cols[:len(cols)/2]
		}
	}

	return rows[0], cols[0]
}

func getSeatID(row, col, salt int) int {
	return (row * 8) + col
}

func makeList(size int) []int {
	l := make([]int, size)
	for i := range l {
		l[i] = i
	}
	return l
}

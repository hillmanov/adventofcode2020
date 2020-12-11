package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	seats, err := utils.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Solution := part1(seats)
	part2Solution := part2(seats)

	fmt.Printf("Day 11: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 11: Part 2: = %+v\n", part2Solution)
}

func part1(seats []string) int {
	currentSeats := seats
	for {
		newSeats, changed := adjacentTick(currentSeats)
		if changed == 0 {
			break
		}
		currentSeats = newSeats
	}

	totalOccupied := 0
	for _, row := range currentSeats {
		totalOccupied += countOf([]rune(row), '#')
	}

	return totalOccupied
}

func part2(seats []string) int {
	currentSeats := seats
	for {
		newSeats, changed := visibleTick(currentSeats)
		if changed == 0 {
			break
		}
		currentSeats = newSeats
	}

	totalOccupied := 0
	for _, row := range currentSeats {
		totalOccupied += countOf([]rune(row), '#')
	}

	return totalOccupied
}

func adjacentTick(seats []string) (newSeats []string, changed int) {
	newSeats = copyOf(seats)

	for rowIndex, row := range seats {
		for colIndex, seat := range row {
			adjacentSeats := getAdjacentSeats(seats, rowIndex, colIndex)
			switch seat {
			case 'L': // Vacant
				if countOf(adjacentSeats, '#') == 0 {
					newSeats[rowIndex] = replaceAtIndex(newSeats[rowIndex], colIndex, '#')
					changed++
				}
			case '#': // Occupied
				if countOf(adjacentSeats, '#') >= 4 {
					newSeats[rowIndex] = replaceAtIndex(newSeats[rowIndex], colIndex, 'L')
					changed++
				}
			}
		}
	}

	return newSeats, changed
}

func visibleTick(seats []string) (newSeats []string, changed int) {
	newSeats = copyOf(seats)

	for rowIndex, row := range seats {
		for colIndex, seat := range row {
			visibleSeats := getVisibleSeats(seats, rowIndex, colIndex)
			switch seat {
			case 'L': // Vacant
				if countOf(visibleSeats, '#') == 0 {
					newSeats[rowIndex] = replaceAtIndex(newSeats[rowIndex], colIndex, '#')
					changed++
				}
			case '#': // Occupied
				if countOf(visibleSeats, '#') >= 5 {
					newSeats[rowIndex] = replaceAtIndex(newSeats[rowIndex], colIndex, 'L')
					changed++
				}
			}
		}
	}

	return newSeats, changed
}

func getAdjacentSeats(seats []string, row, col int) []rune {
	adjacentSeats := []rune{}

	// U
	if row > 0 {
		adjacentSeats = append(adjacentSeats, rune(seats[row-1][col]))
	}
	// UR
	if row > 0 && col < len(seats[row])-1 {
		adjacentSeats = append(adjacentSeats, rune(seats[row-1][col+1]))
	}
	// R
	if col < len(seats[row])-1 {
		adjacentSeats = append(adjacentSeats, rune(seats[row][col+1]))
	}
	// DR
	if row < len(seats)-1 && col < len(seats[row])-1 {
		adjacentSeats = append(adjacentSeats, rune(seats[row+1][col+1]))
	}
	// D
	if row < len(seats)-1 {
		adjacentSeats = append(adjacentSeats, rune(seats[row+1][col]))
	}
	// DL
	if row < len(seats)-1 && col > 0 {
		adjacentSeats = append(adjacentSeats, rune(seats[row+1][col-1]))
	}
	// L
	if col > 0 {
		adjacentSeats = append(adjacentSeats, rune(seats[row][col-1]))
	}
	// UL
	if row > 0 && col > 0 {
		adjacentSeats = append(adjacentSeats, rune(seats[row-1][col-1]))
	}
	return adjacentSeats
}

func getVisibleSeats(seats []string, row, col int) []rune {
	adjacentSeats := []rune{}

	// U
	for r, c := row-1, col; r >= 0; {
		adjacentSeats = append(adjacentSeats, rune(seats[r][c]))
		if seats[r][c] != '.' {
			break
		}
		r--
	}
	// UR
	for r, c := row-1, col+1; r >= 0 && c < len(seats[r]); {
		adjacentSeats = append(adjacentSeats, rune(seats[r][c]))
		if seats[r][c] != '.' {
			break
		}
		r--
		c++
	}
	// R
	for r, c := row, col+1; c < len(seats[r]); {
		adjacentSeats = append(adjacentSeats, rune(seats[r][c]))
		if seats[r][c] != '.' {
			break
		}
		c++
	}
	// DR
	for r, c := row+1, col+1; r < len(seats) && c < len(seats[r]); {
		adjacentSeats = append(adjacentSeats, rune(seats[r][c]))
		if seats[r][c] != '.' {
			break
		}
		r++
		c++
	}
	// D
	for r, c := row+1, col; r < len(seats); {
		adjacentSeats = append(adjacentSeats, rune(seats[r][c]))
		if seats[r][c] != '.' {
			break
		}
		r++
	}
	// DL
	for r, c := row+1, col-1; r < len(seats) && c >= 0; {
		adjacentSeats = append(adjacentSeats, rune(seats[r][c]))
		if seats[r][c] != '.' {
			break
		}
		r++
		c--
	}
	// L
	for r, c := row, col-1; c >= 0; {
		adjacentSeats = append(adjacentSeats, rune(seats[r][c]))
		if seats[r][c] != '.' {
			break
		}
		c--
	}
	// UL
	for r, c := row-1, col-1; r >= 0 && c >= 0; {
		adjacentSeats = append(adjacentSeats, rune(seats[r][c]))
		if seats[r][c] != '.' {
			break
		}
		r--
		c--
	}
	return adjacentSeats
}

func countOf(seats []rune, match rune) int {
	count := 0
	for _, seat := range seats {
		if seat == match {
			count++
		}
	}
	return count
}

func replaceAtIndex(in string, i int, r rune) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func copyOf(seats []string) []string {
	newSeats := make([]string, len(seats))
	for i, line := range seats {
		newSeats[i] = line[:]
	}
	return newSeats
}

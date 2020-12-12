package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
)

type Instruction struct {
	Action rune
	Value  int
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	instructions := parseInstructions(lines)

	part1Solution := part1(instructions)
	part2Solution := part2(instructions)

	fmt.Printf("Day 12: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 12: Part 2: = %+v\n", part2Solution)
}

func part1(instructions []Instruction) int {
	x, y := 0, 0
	facing := 'E'
	degreesToHeading := map[int]rune{0: 'N', 90: 'E', 180: 'S', 270: 'W', -90: 'W', -180: 'S', -270: 'E'}
	headingToDegrees := map[rune]int{'N': 0, 'E': 90, 'S': 180, 'W': 270}

	for _, instruction := range instructions {
		switch instruction.Action {
		case 'N':
			y += instruction.Value
		case 'E':
			x += instruction.Value
		case 'S':
			y -= instruction.Value
		case 'W':
			x -= instruction.Value
		case 'R':
			facing = degreesToHeading[(headingToDegrees[facing]+instruction.Value)%360]
		case 'L':
			facing = degreesToHeading[(headingToDegrees[facing]-instruction.Value)%360]
		case 'F':
			switch facing {
			case 'N':
				y += instruction.Value
			case 'E':
				x += instruction.Value
			case 'S':
				y -= instruction.Value
			case 'W':
				x -= instruction.Value
			}
		}
	}
	return utils.Abs(x) + utils.Abs(y)
}

func part2(instructions []Instruction) int {
	shipX, shipY, waypointX, waypointY := 0, 0, 10, 1

	for _, instruction := range instructions {
		switch instruction.Action {
		case 'N':
			waypointY += instruction.Value
		case 'E':
			waypointX += instruction.Value
		case 'S':
			waypointY -= instruction.Value
		case 'W':
			waypointX -= instruction.Value
		case 'R':
			switch instruction.Value {
			case 90:
				waypointX, waypointY = waypointY, waypointX*-1
			case 180:
				waypointX, waypointY = waypointX*-1, waypointY*-1
			case 270:
				waypointX, waypointY = waypointY*-1, waypointX
			}
		case 'L':
			switch instruction.Value {
			case 90:
				waypointX, waypointY = waypointY*-1, waypointX
			case 180:
				waypointX, waypointY = waypointX*-1, waypointY*-1
			case 270:
				waypointX, waypointY = waypointY, waypointX*-1
			}
		case 'F':
			shipX, shipY = (shipX + (waypointX * instruction.Value)), (shipY + (waypointY * instruction.Value))
		}
	}
	return utils.Abs(shipX) + utils.Abs(shipY)
}

func parseInstructions(lines []string) []Instruction {
	instructions := []Instruction{}
	for _, line := range lines {
		value, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, Instruction{
			Action: rune(line[0]),
			Value:  value,
		})
	}
	return instructions
}

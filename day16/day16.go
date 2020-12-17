package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	Field  string
	Ranges [2][2]int
}

type Ticket []int

func (r Rule) isValid(v int) bool {
	return (v >= r.Ranges[0][0] && v <= r.Ranges[0][1]) || (v >= r.Ranges[1][0] && v <= r.Ranges[1][1])
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	rules, yourTicket, nearbyTickets := parseInput(lines)

	part1Solution := part1(rules, nearbyTickets)
	part2Solution := part2(rules, yourTicket, nearbyTickets)

	fmt.Printf("Day 16: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 16: Part 2: = %+v\n", part2Solution)
}

func part1(rules []Rule, nearbyTickets []Ticket) int {
	invalidValues := []int{}

	for _, nearbyTicket := range nearbyTickets {
		for _, value := range nearbyTicket {
			valueIsValid := false
			for _, rule := range rules {
				if rule.isValid(value) {
					valueIsValid = true
					break
				}
			}
			if !valueIsValid {
				invalidValues = append(invalidValues, value)
			}
		}
	}

	return utils.SumOf(invalidValues)
}

func part2(rules []Rule, yourTicket Ticket, nearbyTickets []Ticket) int {
	// Remove invalid nearby tickets
	validNearbyTickes := []Ticket{}
	for _, nearbyTicket := range nearbyTickets {
		allValuesValid := true
		for _, value := range nearbyTicket {
			valueValidForAnyRule := false
			for _, rule := range rules {
				valueValidForAnyRule = valueValidForAnyRule || rule.isValid(value)
			}
			allValuesValid = allValuesValid && valueValidForAnyRule
		}
		if allValuesValid {
			validNearbyTickes = append(validNearbyTickes, nearbyTicket)
		}
	}

	// Build rule matrix
	fieldCount := len(rules)
	ruleMatrix := make([][]bool, fieldCount, fieldCount)
	for ruleIndex := 0; ruleIndex < fieldCount; ruleIndex++ {
		ruleMatrix[ruleIndex] = make([]bool, fieldCount, fieldCount)
		for i := range ruleMatrix[ruleIndex] {
			ruleMatrix[ruleIndex][i] = true
		}
		for ticketValueIndex := 0; ticketValueIndex < fieldCount; ticketValueIndex++ {
			for _, ticket := range validNearbyTickes {
				ruleMatrix[ruleIndex][ticketValueIndex] = ruleMatrix[ruleIndex][ticketValueIndex] && rules[ruleIndex].isValid(ticket[ticketValueIndex])
			}
		}
	}

	// Find the "column" that has only 1 true in it - that row belongs to the column (position). Remove that rule index from consideration
	ruleToFieldPositionMap := map[int]int{}
	for len(ruleToFieldPositionMap) != fieldCount {
		for colIndex := 0; colIndex < fieldCount; colIndex++ {
			count, indexes := indexesAndCountsOfTrueForColumn(colIndex, ruleMatrix, ruleToFieldPositionMap)
			if count == 1 {
				ruleToFieldPositionMap[indexes[0]] = colIndex
			}
		}
	}

	product := 1
	for ruleIndex, valueIndex := range ruleToFieldPositionMap {
		if strings.HasPrefix(rules[ruleIndex].Field, "departure") {
			product *= yourTicket[valueIndex]
		}
	}

	return product
}

func indexesAndCountsOfTrueForColumn(col int, matrix [][]bool, ruleToFieldPositionMap map[int]int) (count int, indexes []int) {
	for rowIndex := range matrix {
		if _, ok := ruleToFieldPositionMap[rowIndex]; ok {
			continue
		}
		if matrix[rowIndex][col] {
			count++
			indexes = append(indexes, rowIndex)
		}
	}

	return count, indexes
}

func parseInput(lines []string) (rules []Rule, yourTicket Ticket, nearbyTickets []Ticket) {
	i := 0
	// Parse rules
	for ; len(lines[i]) > 0; i++ {
		rule := Rule{}
		parts := strings.Split(lines[i], ":")

		rule.Field = strings.TrimSpace(parts[0])

		var range1Low, range1High, range2Low, range2High int
		fmt.Sscanf(parts[1], "%d-%d or %d-%d", &range1Low, &range1High, &range2Low, &range2High)
		rule.Ranges = [2][2]int{{range1Low, range1High}, {range2Low, range2High}}

		rules = append(rules, rule)
	}

	// Parse your ticket
	i += 2
	yourTicket = parseIntsFromLine(lines[i])

	// Parse nearby tickets
	i += 3
	for ; i < len(lines); i++ {
		nearbyTickets = append(nearbyTickets, parseIntsFromLine(lines[i]))
	}

	return rules, yourTicket, nearbyTickets
}

func parseIntsFromLine(line string) []int {
	ints := strings.Split(line, ",")
	values := make([]int, 0, len(ints))
	for _, num := range ints {
		v, _ := strconv.Atoi(num)
		values = append(values, v)
	}
	return values
}

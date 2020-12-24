package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	entries, _ := utils.ReadLines("./input.txt")

	deck1, deck2 := parseDecks(entries)

	part1Solution := part1(deck1, deck2)
	part2Solution := part2(deck1, deck2)

	fmt.Printf("Day 22: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 22: Part 2: = %+v\n", part2Solution)
}

func part1(deck1, deck2 []int) int {
	var winningDeck []int
	winner, deck1, deck2 := playGame(deck1, deck2, false)

	switch winner {
	case 1:
		winningDeck = deck1
	case 2:
		winningDeck = deck2
	}

	return calculatedScore(winningDeck)
}

func part2(deck1, deck2 []int) int {
	var winningDeck []int
	winner, deck1, deck2 := playGame(deck1, deck2, true)

	switch winner {
	case 1:
		winningDeck = deck1
	case 2:
		winningDeck = deck2
	}

	return calculatedScore(winningDeck)
}

func playGame(deck1, deck2 []int, recursiveCombat bool) (int, []int, []int) {
	pastRounds := map[string]bool{}
	for len(deck1) != 0 && len(deck2) != 0 {
		if _, ok := pastRounds[key(deck1, deck2)]; ok && recursiveCombat {
			return 1, deck1, deck2
		}

		pastRounds[key(deck1, deck2)] = true
		switch {
		case recursiveCombat && len(deck1) > deck1[0] && len(deck2) > deck2[0]:
			switch winner, _, _ := playGame(utils.CopyOf(deck1[1:deck1[0]+1]), utils.CopyOf(deck2[1:deck2[0]+1]), true); winner {
			case 1:
				deck1 = append(deck1[1:], deck1[0])
				deck1 = append(deck1, deck2[0])
				deck2 = deck2[1:]
			case 2:
				deck2 = append(deck2[1:], deck2[0])
				deck2 = append(deck2, deck1[0])
				deck1 = deck1[1:]
			}
		case deck1[0] > deck2[0]:
			deck1 = append(deck1[1:], deck1[0])
			deck1 = append(deck1, deck2[0])
			deck2 = deck2[1:]
		case deck1[0] < deck2[0]:
			deck2 = append(deck2[1:], deck2[0])
			deck2 = append(deck2, deck1[0])
			deck1 = deck1[1:]
		}
	}

	switch {
	case len(deck1) > 0:
		return 1, deck1, deck2
	case len(deck2) > 1:
		return 2, deck1, deck2
	}

	return 0, nil, nil
}

func calculatedScore(deck []int) int {
	score := 0
	for i := 0; i < len(deck); i++ {
		score += deck[len(deck)-1-i] * (i + 1)
	}
	return score
}

func key(deck1, deck2 []int) string {
	return fmt.Sprintf("%v:%v", deck1, deck2)
}

func parseDecks(input []string) ([]int, []int) {
	deck1 := []int{}
	deck2 := []int{}

	lineIndex := 0
	for _, line := range input {
		lineIndex++
		if len(line) == 0 {
			break
		}
		if strings.HasPrefix(line, "Player") {
			continue
		}
		card, _ := strconv.Atoi(line)
		deck1 = append(deck1, card)
	}

	for _, line := range input[lineIndex+1:] {
		lineIndex++
		if len(line) == 0 {
			break
		}
		if strings.HasPrefix(line, "Player") {
			continue
		}
		card, _ := strconv.Atoi(line)
		deck2 = append(deck2, card)
	}

	return deck1, deck2
}

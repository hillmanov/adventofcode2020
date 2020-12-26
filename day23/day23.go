package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	part1Solution := part1([]int{2, 5, 3, 1, 4, 9, 8, 6, 7})
	fmt.Printf("Day 23: Part 1: = %+v\n", part1Solution)

	part2Solution := part2([]int{2, 5, 3, 1, 4, 9, 8, 6, 7})
	fmt.Printf("Day 23: Part 2: = %+v\n", part2Solution)
}

func part1(cups []int) string {
	currentCup := cups[0]
	for i := 0; i < 100; i++ {
		indexOfCurrentCup := utils.IndexOf(cups, currentCup)
		inPlayCups, pickedUpCups := split(cups, indexOfCurrentCup)
		destinationCup := getDestinationCup(inPlayCups, pickedUpCups, currentCup)
		cups = insertAt(inPlayCups, pickedUpCups, utils.IndexOf(inPlayCups, destinationCup))
		currentCup = cups[(utils.IndexOf(cups, currentCup)+1)%len(cups)]
	}

	labels := ""
	indexOfOne := utils.IndexOf(cups, 1)
	for i := 0; i < len(cups)-1; i++ {
		labels += strconv.Itoa(cups[(indexOfOne+1+i)%len(cups)])
	}
	return labels
}

func part2(startCups []int) int {
	cups := make([]int, 1000)
	copy(cups, startCups)

	for i := utils.MaxOf(cups) + 1; i < 100; i++ {
		cups[i] = i
	}
	fmt.Println("Done with that...")

	currentCup := cups[0]
	for i := 0; i < 100; i++ {
		fmt.Printf("i = %+v\n", i)
		indexOfCurrentCup := utils.IndexOf(cups, currentCup)
		inPlayCups, pickedUpCups := split(cups, indexOfCurrentCup)
		destinationCup := getDestinationCup(inPlayCups, pickedUpCups, currentCup)
		cups = insertAt(inPlayCups, pickedUpCups, utils.IndexOf(inPlayCups, destinationCup))
		currentCup = cups[(utils.IndexOf(cups, currentCup)+1)%len(cups)]
	}

	return utils.IndexOf(cups, 1)

}

func getDestinationCup(cups []int, pickedUpCups []int, currentCup int) int {
	destinationCup := currentCup - 1

	for {
		indexOfDestinationCup := utils.IndexOf(cups, destinationCup)
		if indexOfDestinationCup != -1 {
			return cups[indexOfDestinationCup]
		} else {
			destinationCup--
			if destinationCup <= 0 {
				destinationCup = len(cups) + len(pickedUpCups)
			}
		}
	}
}

func split(cups []int, afterIndex int) (inPlayCups []int, pickedUpCups []int) {
	for i := 0; i < 3; i++ {
		pickedUpCups = append(pickedUpCups, cups[(afterIndex+1+i)%len(cups)])
	}
	for _, cup := range cups {
		if utils.IndexOf(pickedUpCups, cup) == -1 {
			inPlayCups = append(inPlayCups, cup)
		}
	}

	return inPlayCups, pickedUpCups
}

func insertAt(cups []int, insert []int, at int) []int {
	newCups := make([]int, 0)
	newCups = append(newCups, cups[:at+1]...)
	newCups = append(newCups, insert...)
	newCups = append(newCups, cups[at+1:]...)
	return newCups
}

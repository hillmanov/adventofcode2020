package main

import (
	"fmt"
)

type Cup struct {
	Label int
	Next  *Cup
	Prev  *Cup
}

func (c *Cup) SnipChain() []*Cup {
	chain := []*Cup{c.Next, c.Next.Next, c.Next.Next.Next}

	chain[len(chain)-1].Prev = c
	c.Next = chain[len(chain)-1].Next

	chain[0].Prev = nil
	chain[len(chain)-1].Next = nil

	return chain
}

func (c *Cup) InsertChain(chain []*Cup) {
	chain[len(chain)-1].Next = c.Next
	chain[len(chain)-1].Next.Prev = chain[len(chain)-1]
	chain[0].Prev = c
	c.Next = chain[0]
}

func chainContains(label int, chain []*Cup) bool {
	return chain[0].Label == label || chain[1].Label == label || chain[2].Label == label
}

func resultLabel(cupOne *Cup) string {
	label := ""
	currentCup := cupOne.Next
	for currentCup != nil && currentCup != cupOne {
		label += fmt.Sprintf("%d", currentCup.Label)
		currentCup = currentCup.Next
	}
	return label
}

func main() {
	part1Solution := part1([]int{2, 5, 3, 1, 4, 9, 8, 6, 7})
	fmt.Printf("Day 23: Part 1: = %+v\n", part1Solution)

	part2Solution := part2([]int{2, 5, 3, 1, 4, 9, 8, 6, 7})
	fmt.Printf("Day 23: Part 2: = %+v\n", part2Solution)
}

func part1(labels []int) string {
	cupOne := play(labels, 0, 100)
	return resultLabel(cupOne)
}

func part2(labels []int) int {
	cupOne := play(labels, 1_000_000, 10_000_000)
	return cupOne.Next.Label * cupOne.Next.Next.Label
}

func play(labels []int, extra int, rounds int) *Cup {

	var cupRefs []*Cup
	if extra == 0 {
		cupRefs = make([]*Cup, len(labels)+1)
	} else {
		cupRefs = make([]*Cup, extra+1)
	}

	root := &Cup{Label: labels[0]}
	cupRefs[labels[0]] = root

	totalCups := 1
	currentCup := root
	for _, v := range labels[1:] {
		newCup := &Cup{Label: v}
		cupRefs[v] = newCup

		currentCup.Next = newCup
		newCup.Prev = currentCup
		newCup.Next = root
		currentCup = newCup
		totalCups++
	}

	for v := len(labels) + 1; v <= extra; v++ {
		newCup := &Cup{Label: v}
		cupRefs[v] = newCup

		currentCup.Next = newCup
		newCup.Prev = currentCup
		newCup.Next = root
		currentCup = newCup
		totalCups++
	}

	currentCup = root
	for i := 0; i < rounds; i++ {
		chain := currentCup.SnipChain()

		destinationCupLabel := currentCup.Label
		for {
			destinationCupLabel--
			if destinationCupLabel <= 0 {
				destinationCupLabel = totalCups
			}

			if chainContains(destinationCupLabel, chain) == false {
				break
			}
		}

		destinationCup := cupRefs[destinationCupLabel]
		destinationCup.InsertChain(chain)

		currentCup = currentCup.Next
	}

	return cupRefs[1]
}

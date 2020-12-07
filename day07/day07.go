package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

type bag struct {
	Name        string
	Contains    map[*bag]int // What it contains, and how many
	ContainedBy []*bag
}

func newBag(name string) *bag {
	return &bag{
		Name:        name,
		Contains:    map[*bag]int{},
		ContainedBy: []*bag{},
	}
}

func (b *bag) allContainerBags() []*bag {
	allBags := []*bag{}
	allBags = append(allBags, b.ContainedBy...)

	for _, containerBag := range b.ContainedBy {
		allBags = append(allBags, containerBag.allContainerBags()...)
	}

	uniquer := map[*bag]bool{}
	for _, bag := range allBags {
		uniquer[bag] = true
	}

	uniqueBags := make([]*bag, 0)
	for bag, _ := range uniquer {
		uniqueBags = append(uniqueBags, bag)
	}

	return uniqueBags
}

func (b *bag) allContainedBagsCount() int {
	totalCount := 0
	for containedBag, count := range b.Contains {
		totalCount += count + (count * containedBag.allContainedBagsCount())
	}
	return totalCount
}

func main() {
	input, err := utils.ReadLines("./input.txt")
	if err != nil {
		panic(err)
	}

	bags := parseRules(input)
	part1Solution := len(bags["shiny gold"].allContainerBags())
	part2Solution := bags["shiny gold"].allContainedBagsCount()

	fmt.Printf("Day 07: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 07: Part 2: = %+v\n", part2Solution)
}

func parseRules(rules []string) map[string]*bag {
	bags := make(map[string]*bag)

	for _, rule := range rules {
		parts := strings.Split(rule, "bags contain")
		containerBagName := strings.TrimSpace(parts[0])
		containedBags := parts[1]

		if _, ok := bags[containerBagName]; !ok {
			bags[containerBagName] = newBag(containerBagName)
		}

		for _, containedBag := range strings.Split(containedBags, ",") {
			var amount int
			var namePart1 string
			var namePart2 string
			fmt.Sscanf(containedBag, "%d %s %s", &amount, &namePart1, &namePart2)
			containedBagName := strings.TrimSpace(namePart1 + " " + namePart2)

			if _, ok := bags[containedBagName]; !ok {
				bags[containedBagName] = newBag(containedBagName)
			}

			if amount > 0 {
				bags[containerBagName].Contains[bags[containedBagName]] = amount
			}
			bags[containedBagName].ContainedBy = append(bags[containedBagName].ContainedBy, bags[containerBagName])
		}
	}

	return bags
}

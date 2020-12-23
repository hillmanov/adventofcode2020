package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	input, _ := utils.ReadLines("./input.txt")
	tiles := parseTiles(input)

	part1Solution := part1(tiles)
	part2Solution := part2()

	fmt.Printf("Day 20: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 20: Part 2: = %+v\n", part2Solution)
}

func part1(tiles []tile) int {
	// First, find the corners - these are those that have two edges that match NO others

	// corners := []tile{}
	for _, outerT := range tiles {
		edgeMatchCount := 0
		for _, innerT := range tiles {
			if outerT.ID == innerT.ID {
				continue
			}
			for _, outerTEdge := range outerT.Edges() {
				for _, innerTEdge := range innerT.Edges() {
					if outerTEdge == innerTEdge {
						edgeMatchCount++
					}
				}
			}
		}
		fmt.Printf("Tile %d has %d edges that match others\n", outerT.ID, edgeMatchCount)
	}

	return -1
}

func part2() int {
	return -1
}

type tile struct {
	ID   int
	Data [][]rune
}

func (t tile) Edges() []string {
	edges := []string{}
	// Top edge
	edges = append(edges, string(t.Data[0]))

	// Left and right edges
	var left, right string
	for i := 0; i < len(t.Data); i++ {
		left += string(t.Data[i][0])
		right += string(t.Data[i][len(t.Data[i])-1])
	}
	edges = append(edges, left)
	edges = append(edges, right)

	// Bottom edge
	edges = append(edges, string(t.Data[len(t.Data)-1]))

	return edges
}

func (t *tile) FlipHorizontal() {
	flippedData := [][]rune{}
	t.Data = flippedData
}

func (t *tile) FlipVertical() {
	flippedData := [][]rune{}
	t.Data = flippedData
}

func (t *tile) Rotate() {
	rotatedData := [][]rune{}
	t.Data = rotatedData
}

func parseTiles(input []string) []tile {
	tiles := []tile{}
	var t tile
	for _, line := range input {
		if strings.HasPrefix(line, "Tile") {
			t = tile{}
			fmt.Sscanf(line, "Tile %d", &t.ID)
		} else if len(line) != 0 {
			t.Data = append(t.Data, []rune(line))
		} else {
			tiles = append(tiles, t)
		}
	}
	return tiles
}

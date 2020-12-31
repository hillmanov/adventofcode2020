package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
)

const (
	white = 0
	black = 1
)

type color int

type Tile struct {
	X int
	Y int
	Z int
}

func neighborsAndCounts(t Tile, tiles map[Tile]color) (whiteCount, blackCount int) {
	neighborE := Tile{X: t.X - 1, Y: t.Y + 1, Z: t.Z}
	neighborSE := Tile{X: t.X, Y: t.Y + 1, Z: t.Z - 1}
	neighborSW := Tile{X: t.X + 1, Y: t.Y, Z: t.Z - 1}
	neighborW := Tile{X: t.X + 1, Y: t.Y - 1, Z: t.Z}
	neighborNW := Tile{X: t.X, Y: t.Y - 1, Z: t.Z + 1}
	neighborNE := Tile{X: t.X - 1, Y: t.Y, Z: t.Z + 1}

	neighbors := []Tile{neighborE, neighborSE, neighborSW, neighborW, neighborNW, neighborNE}

	for _, neighbor := range neighbors {
		tColor := tiles[neighbor]
		if tColor == white {
			whiteCount++
		} else if tColor == black {
			blackCount++
		}
	}

	return whiteCount, blackCount
}

func main() {
	// defer profile.Start().Stop()
	lines, _ := utils.ReadLines("./input.txt")
	tileDirections := parseTileDirections(lines)
	tiles := map[Tile]color{}

	part1Solution := part1(tileDirections, tiles)
	fmt.Printf("Day 24: Part 1: = %+v\n", part1Solution)

	part2Solution := part2(tiles)
	fmt.Printf("Day 24: Part 2: = %+v\n", part2Solution)
}

func part1(tileDirections [][]string, tiles map[Tile]color) int {

	for _, directions := range tileDirections {
		location := Tile{X: 0, Y: 0}
		for _, direction := range directions {
			switch direction {
			case "e":
				location.X--
				location.Y++
			case "se":
				location.Y++
				location.Z--
			case "sw":
				location.X++
				location.Z--
			case "w":
				location.X++
				location.Y--
			case "nw":
				location.Y--
				location.Z++
			case "ne":
				location.X--
				location.Z++
			}
		}
		tiles[location] = (tiles[location] + 1) % 2
	}

	blackTiles := 0
	for _, tColor := range tiles {
		if tColor == black {
			blackTiles++
		}
	}

	return blackTiles
}

func part2(tiles map[Tile]color) int {
	blackTiles := 0

	for i := -200; i < 200; i++ {
		for j := -200; j < 200; j++ {
			x := -i
			z := j
			y := -x - z
			t := Tile{X: x, Y: y, Z: z}
			if _, ok := tiles[t]; !ok {
				tiles[t] = white
			}
		}
	}

	for i := 0; i < 100; i++ {
		// Make a copy of all the tiles
		newTiles := map[Tile]color{}
		for t, tColor := range tiles {
			newTiles[Tile{X: t.X, Y: t.Y, Z: t.Z}] = tColor
		}

		for tile, tColor := range tiles {
			_, blackCount := neighborsAndCounts(tile, tiles)
			if tColor == black && (blackCount == 0 || blackCount > 2) {
				newTiles[tile] = white
			} else if tColor == white && blackCount == 2 {
				newTiles[tile] = black
			}
		}
		tiles = newTiles
	}

	for _, tColor := range tiles {
		if tColor == black {
			blackTiles++
		}
	}

	return blackTiles
}

func parseTileDirections(lines []string) [][]string {
	re := regexp.MustCompile("(e|se|sw|w|nw|ne)")
	tileDirections := [][]string{}
	for _, line := range lines {
		tileDirections = append(tileDirections, re.FindAllString(line, -1))
	}

	return tileDirections
}

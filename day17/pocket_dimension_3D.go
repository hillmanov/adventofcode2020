package main

import (
	"math"
)

type PocketDimension3D map[int]map[int]map[int]rune

func parseInput3D(lines []string) PocketDimension3D {
	pocketDimension := PocketDimension3D{}

	zIndex := 0
	pocketDimension[zIndex] = map[int]map[int]rune{}

	for yIndex, line := range lines {
		for xIndex, char := range line {
			if _, ok := pocketDimension[zIndex][yIndex]; !ok {
				pocketDimension[zIndex][yIndex] = map[int]rune{}
			}
			pocketDimension[zIndex][yIndex][xIndex] = char
		}
	}

	return pocketDimension
}

func (pd PocketDimension3D) Grow(zIndex, yIndex, xIndex int) {
	if _, ok := pd[zIndex]; !ok {
		pd[zIndex] = map[int]map[int]rune{}
	}
	if _, ok := pd[zIndex][yIndex]; !ok {
		pd[zIndex][yIndex] = map[int]rune{}
	}
}

func (pd PocketDimension3D) Copy() PocketDimension3D {
	zMin, zMax, xMin, xMax, yMin, yMax := pd.MinMax()

	newPd := PocketDimension3D{}

	for zIndex := zMin; zIndex <= zMax; zIndex++ {
		for yIndex := yMin; yIndex <= yMax; yIndex++ {
			for xIndex := xMin; xIndex <= xMax; xIndex++ {
				newPd.Set(zIndex, yIndex, xIndex, pd[zIndex][yIndex][xIndex])
			}
		}
	}
	return newPd
}

func (pd PocketDimension3D) TotalActiveCount() int {
	zMin, zMax, xMin, xMax, yMin, yMax := pd.MinMax()

	activeCount := 0
	for zIndex := zMin; zIndex <= zMax; zIndex++ {
		for yIndex := yMin; yIndex <= yMax; yIndex++ {
			for xIndex := xMin; xIndex <= xMax; xIndex++ {
				if pd[zIndex][yIndex][xIndex] == '#' {
					activeCount++
				}
			}
		}
	}
	return activeCount
}

func (pd PocketDimension3D) Set(zIndex, yIndex, xIndex int, value rune) {
	pd.Grow(zIndex, yIndex, xIndex)
	pd[zIndex][yIndex][xIndex] = value
}

func (pd PocketDimension3D) Get(zIndex, yIndex, xIndex int) rune {
	pd.Grow(zIndex, yIndex, xIndex)
	if value, ok := pd[zIndex][yIndex][xIndex]; !ok {
		pd.Set(zIndex, yIndex, xIndex, '.')
		return '.'
	} else {
		return value
	}
}

func (pd PocketDimension3D) GetActiveNeighborCount(zIndexCheck, yIndexCheck, xIndexCheck int) (active int) {
	for zIndex := zIndexCheck - 1; zIndex <= zIndexCheck+1; zIndex++ {
		for yIndex := yIndexCheck - 1; yIndex <= yIndexCheck+1; yIndex++ {
			for xIndex := xIndexCheck - 1; xIndex <= xIndexCheck+1; xIndex++ {
				if !(zIndexCheck == zIndex && yIndexCheck == yIndex && xIndexCheck == xIndex) && pd.Get(zIndex, yIndex, xIndex) == '#' {
					active++
				}
			}
		}
	}

	return active
}

func (pd PocketDimension3D) Cycle() {
	oldPd := pd.Copy()
	zMin, zMax, xMin, xMax, yMin, yMax := oldPd.MinMax()

	for zIndex := zMin - 1; zIndex <= zMax+1; zIndex++ {
		for yIndex := yMin - 1; yIndex <= yMax+1; yIndex++ {
			for xIndex := xMin - 1; xIndex <= xMax+1; xIndex++ {
				currentValue := oldPd.Get(zIndex, yIndex, xIndex)
				// Make sure to grow the original as well
				pd.Get(zIndex, yIndex, xIndex)
				activeCount := oldPd.GetActiveNeighborCount(zIndex, yIndex, xIndex)

				switch currentValue {
				case '#':
					if activeCount != 2 && activeCount != 3 {
						pd.Set(zIndex, yIndex, xIndex, '.')
					}
				case '.':
					if activeCount == 3 {
						pd.Set(zIndex, yIndex, xIndex, '#')
					}
				}
			}
		}
	}
}

func (pd PocketDimension3D) MinMax() (zMin, zMax, xMin, xMax, yMin, yMax int) {
	zMin = math.MaxInt64
	zMax = math.MinInt64
	xMin = math.MaxInt64
	xMax = math.MinInt64
	yMin = math.MaxInt64
	yMax = math.MinInt64

	for zKey, zDimension := range pd {
		if zKey < zMin {
			zMin = zKey
		}
		if zKey > zMax {
			zMax = zKey
		}
		for yKey, yDimension := range zDimension {
			if yKey < yMin {
				yMin = yKey
			}
			if yKey > yMax {
				yMax = yKey
			}

			for xKey := range yDimension {
				if xKey < xMin {
					xMin = xKey
				}
				if xKey > xMax {
					xMax = xKey
				}
			}
		}
	}

	return zMin, zMax, xMin, xMax, yMin, yMax
}

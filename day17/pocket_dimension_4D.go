package main

import (
	"math"
)

type PocketDimension4D map[int]map[int]map[int]map[int]rune

func parseInput4D(lines []string) PocketDimension4D {
	pocketDimension := PocketDimension4D{}

	wIndex := 0
	zIndex := 0
	pocketDimension[wIndex] = map[int]map[int]map[int]rune{}
	pocketDimension[wIndex][zIndex] = map[int]map[int]rune{}

	for yIndex, line := range lines {
		for xIndex, char := range line {
			if _, ok := pocketDimension[wIndex][zIndex][yIndex]; !ok {
				pocketDimension[wIndex][zIndex][yIndex] = map[int]rune{}
			}
			pocketDimension[wIndex][zIndex][yIndex][xIndex] = char
		}
	}

	return pocketDimension
}

func (pd PocketDimension4D) Grow(wIndex, zIndex, yIndex, xIndex int) {
	if _, ok := pd[wIndex]; !ok {
		pd[wIndex] = map[int]map[int]map[int]rune{}
	}
	if _, ok := pd[wIndex][zIndex]; !ok {
		pd[wIndex][zIndex] = map[int]map[int]rune{}
	}
	if _, ok := pd[wIndex][zIndex][yIndex]; !ok {
		pd[wIndex][zIndex][yIndex] = map[int]rune{}
	}
}

func (pd PocketDimension4D) Copy() PocketDimension4D {
	wMin, wMax, zMin, zMax, xMin, xMax, yMin, yMax := pd.MinMax()

	newPd := PocketDimension4D{}
	for wIndex := wMin; wIndex <= wMax; wIndex++ {
		for zIndex := zMin; zIndex <= zMax; zIndex++ {
			for yIndex := yMin; yIndex <= yMax; yIndex++ {
				for xIndex := xMin; xIndex <= xMax; xIndex++ {
					newPd.Set(wIndex, zIndex, yIndex, xIndex, pd[wIndex][zIndex][yIndex][xIndex])
				}
			}
		}
	}
	return newPd
}

func (pd PocketDimension4D) TotalActiveCount() int {
	wMin, wMax, zMin, zMax, xMin, xMax, yMin, yMax := pd.MinMax()

	activeCount := 0
	for wIndex := wMin; wIndex <= wMax; wIndex++ {
		for zIndex := zMin; zIndex <= zMax; zIndex++ {
			for yIndex := yMin; yIndex <= yMax; yIndex++ {
				for xIndex := xMin; xIndex <= xMax; xIndex++ {
					if pd[wIndex][zIndex][yIndex][xIndex] == '#' {
						activeCount++
					}
				}
			}
		}
	}
	return activeCount
}

func (pd PocketDimension4D) Set(wIndex, zIndex, yIndex, xIndex int, value rune) {
	pd.Grow(wIndex, zIndex, yIndex, xIndex)
	pd[wIndex][zIndex][yIndex][xIndex] = value
}

func (pd PocketDimension4D) Get(wIndex, zIndex, yIndex, xIndex int) rune {
	pd.Grow(wIndex, zIndex, yIndex, xIndex)
	if value, ok := pd[wIndex][zIndex][yIndex][xIndex]; !ok {
		pd.Set(wIndex, zIndex, yIndex, xIndex, '.')
		return '.'
	} else {
		return value
	}
}

func (pd PocketDimension4D) GetActiveNeighborCount(wIndexCheck, zIndexCheck, yIndexCheck, xIndexCheck int) (active int) {
	for wIndex := wIndexCheck - 1; wIndex <= wIndexCheck+1; wIndex++ {
		for zIndex := zIndexCheck - 1; zIndex <= zIndexCheck+1; zIndex++ {
			for yIndex := yIndexCheck - 1; yIndex <= yIndexCheck+1; yIndex++ {
				for xIndex := xIndexCheck - 1; xIndex <= xIndexCheck+1; xIndex++ {
					if !(wIndexCheck == wIndex && zIndexCheck == zIndex && yIndexCheck == yIndex && xIndexCheck == xIndex) && pd.Get(wIndex, zIndex, yIndex, xIndex) == '#' {
						active++
					}
				}
			}
		}
	}

	return active
}

func (pd PocketDimension4D) Cycle() {
	oldPd := pd.Copy()
	wMin, wMax, zMin, zMax, xMin, xMax, yMin, yMax := pd.MinMax()

	for wIndex := wMin - 1; wIndex <= wMax+1; wIndex++ {
		for zIndex := zMin - 1; zIndex <= zMax+1; zIndex++ {
			for yIndex := yMin - 1; yIndex <= yMax+1; yIndex++ {
				for xIndex := xMin - 1; xIndex <= xMax+1; xIndex++ {
					currentValue := oldPd.Get(wIndex, zIndex, yIndex, xIndex)
					// Make sure to grow the original as well
					pd.Get(wIndex, zIndex, yIndex, xIndex)
					activeCount := oldPd.GetActiveNeighborCount(wIndex, zIndex, yIndex, xIndex)

					switch currentValue {
					case '#':
						if activeCount != 2 && activeCount != 3 {
							pd.Set(wIndex, zIndex, yIndex, xIndex, '.')
						}
					case '.':
						if activeCount == 3 {
							pd.Set(wIndex, zIndex, yIndex, xIndex, '#')
						}
					}
				}
			}
		}
	}
}

func (pd PocketDimension4D) MinMax() (wMin, wMax, zMin, zMax, xMin, xMax, yMin, yMax int) {
	wMin = math.MaxInt64
	wMin = math.MaxInt64
	zMax = math.MinInt64
	zMax = math.MinInt64
	xMin = math.MaxInt64
	xMax = math.MinInt64
	yMin = math.MaxInt64
	yMax = math.MinInt64

	for wKey, wDimension := range pd {
		if wKey < wMin {
			wMin = wKey
		}
		if wKey > wMax {
			wMax = wKey
		}

		for zKey, zDimension := range wDimension {
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
	}

	return wMin, wMax, zMin, zMax, xMin, xMax, yMin, yMax
}

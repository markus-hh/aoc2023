package main

import (
	"fmt"

	"example.com/aoc/util"
)

func parseLine(line string) (galaxies []bool) {
	for _, positionRune := range line {
		if positionRune == '#' {
			galaxies = append(galaxies, true)
		} else {
			galaxies = append(galaxies, false)
		}
	}
	return
}

func expandVertically(galaxiesInput [][]bool) (galaxies [][]bool) {
	for _, galaxyLine := range galaxiesInput {
		galaxies = append(galaxies, util.Copy(galaxyLine))
		if !util.Any(galaxyLine, util.Identity) {
			galaxies = append(galaxies, util.Copy(galaxyLine))
		}
	}

	return
}

func expandHorizontally(galaxiesInput [][]bool) (galaxies [][]bool) {
	expandIndices := []bool{}

	for index := 0; index < len(galaxiesInput[0]); index++ {
		expandIndex := true

		for row := 0; row < len(galaxiesInput); row++ {
			if galaxiesInput[row][index] {
				expandIndex = false
				break
			}
		}

		expandIndices = append(expandIndices, expandIndex)
	}

	for rowIndex := 0; rowIndex < len(galaxiesInput); rowIndex++ {
		galaxyRowInput := galaxiesInput[rowIndex]
		galaxyRow := []bool{}

		for columnIndex := 0; columnIndex < len(galaxiesInput[0]); columnIndex++ {
			galaxyRow = append(galaxyRow, galaxyRowInput[columnIndex])
			if expandIndices[columnIndex] {
				galaxyRow = append(galaxyRow, galaxyRowInput[columnIndex])
			}
		}

		galaxies = append(galaxies, galaxyRow)
	}

	return
}

func findGalaxyPositions(galaxies [][]bool) (positions []util.Position2D) {
	for rowIndex := 0; rowIndex < len(galaxies); rowIndex++ {
		galaxyRow := galaxies[rowIndex]

		for columnIndex := 0; columnIndex < len(galaxyRow); columnIndex++ {
			if galaxyRow[columnIndex] {
				positions = append(positions, util.Position2D{X: columnIndex, Y: rowIndex})
			}
		}
	}
	return
}

func solvePart1(input string) {
	lines := util.SplitLines(input)
	galaxies := util.MapFunc(lines, parseLine)
	galaxies = expandHorizontally(galaxies)
	galaxies = expandVertically(galaxies)

	positions := findGalaxyPositions(galaxies)
	util.SortPositions(positions)

	sum := 0
	for firstPositionIndex := 0; firstPositionIndex < len(positions); firstPositionIndex++ {
		firstPosition := positions[firstPositionIndex]

		for secondPositionIndex := firstPositionIndex + 1; secondPositionIndex < len(positions); secondPositionIndex++ {
			secondPosition := positions[secondPositionIndex]
			sum += util.HammingDistance(firstPosition, secondPosition)
		}
	}

	fmt.Println(sum)
}

func main() {
	input := util.FetchInput(11)
	solvePart1(input)
}

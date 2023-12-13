package main

import (
	"fmt"

	"example.com/aoc/util"
)

func parseInput(lines []string) (patterns [][]string) {
	pattern := []string{}

	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, pattern)
			pattern = []string{}
		} else {
			pattern = append(pattern, line)
		}
	}

	patterns = append(patterns, pattern)
	return
}

func verticalDifferences(pattern []string, index1 int, index2 int) int {
	differences := 0
	for rowIndex := range pattern {
		if pattern[rowIndex][index1] != pattern[rowIndex][index2] {
			differences++
		}
	}
	return differences
}

func verticalEquals(pattern []string, index1 int, index2 int) bool {
	for rowIndex := range pattern {
		if pattern[rowIndex][index1] != pattern[rowIndex][index2] {
			return false
		}
	}

	return true
}

func isVerticalMirror(pattern []string, startingIndex int) bool {
	for offset := 0; offset < min(startingIndex, len(pattern[0])-startingIndex); offset++ {
		if !verticalEquals(pattern, startingIndex-offset-1, startingIndex+offset) {
			return false
		}
	}
	return true
}

func isVerticalMirrorWithSmudge(pattern []string, startingIndex int) bool {
	differences := 0

	for offset := 0; offset < min(startingIndex, len(pattern[0])-startingIndex); offset++ {
		differences += verticalDifferences(pattern, startingIndex-offset-1, startingIndex+offset)
	}
	return differences == 1
}

func isHorizontalMirror(pattern []string, startingIndex int) bool {
	for offset := 0; offset < min(startingIndex, len(pattern)-startingIndex); offset++ {
		if pattern[startingIndex-offset-1] != pattern[startingIndex+offset] {
			return false
		}
	}
	return true
}

func isHorizontalMirrorWithSmudge(pattern []string, startingIndex int) bool {
	differences := 0

	for offset := 0; offset < min(startingIndex, len(pattern)-startingIndex); offset++ {
		differences += horizontalIndices(pattern, startingIndex-offset-1, startingIndex+offset)
	}
	return differences == 1
}

func horizontalIndices(pattern []string, index1 int, index2 int) int {
	differences := 0
	for columnIndex := 0; columnIndex < len(pattern[0]); columnIndex++ {
		if pattern[index1][columnIndex] != pattern[index2][columnIndex] {
			differences++
		}
	}
	return differences
}

func findHorizontalMirrorIndices(pattern []string) (indices []int) {
	for index := 1; index < len(pattern); index++ {
		if isHorizontalMirror(pattern, index) {
			indices = append(indices, index)
		}
	}
	return
}

func findVerticalMirrorIndices(pattern []string) (indices []int) {
	for index := 1; index < len(pattern[0]); index++ {
		if isVerticalMirror(pattern, index) {
			indices = append(indices, index)
		}
	}
	return
}

func findHorizontalMirrorIndicesWithSmudge(pattern []string) (indices []int) {
	for index := 1; index < len(pattern); index++ {
		if isHorizontalMirrorWithSmudge(pattern, index) {
			indices = append(indices, index)
		}
	}
	return
}

func findVerticalMirrorIndicesWithSmudge(pattern []string) (indices []int) {
	for index := 1; index < len(pattern[0]); index++ {
		if isVerticalMirrorWithSmudge(pattern, index) {
			indices = append(indices, index)
		}
	}
	return
}

func summarizePattern(horizontalIndices []int, verticalIndices []int) int {
	sum := 0
	for _, index := range verticalIndices {
		sum += index
	}
	for _, index := range horizontalIndices {
		sum += 100 * index
	}
	return sum
}

func solvePart1(input string) {
	patterns := parseInput(util.SplitLines(input))

	sum := 0
	for index := range patterns {
		horizontalIndices := findHorizontalMirrorIndices(patterns[index])
		verticalIndices := findVerticalMirrorIndices(patterns[index])
		sum += summarizePattern(horizontalIndices, verticalIndices)
	}

	fmt.Println(sum)
}

func solvePart2(input string) {
	patterns := parseInput(util.SplitLines(input))

	sum := 0
	for index := range patterns {
		horizontalIndices := findHorizontalMirrorIndicesWithSmudge(patterns[index])
		verticalIndices := findVerticalMirrorIndicesWithSmudge(patterns[index])
		sum += summarizePattern(horizontalIndices, verticalIndices)
	}

	fmt.Println(sum)
}

func main() {
	input := util.FetchInput(13)
	solvePart2(input)
}

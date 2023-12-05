package main

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/aoc/util"
)

type MapRange struct {
	sourceStart int
	destinationStart int
	length int
}

func determineMappedIndex(category []MapRange, index int) int {
	for _, mapRange := range category {
		if index >= mapRange.sourceStart && index < mapRange.sourceStart + mapRange.length {
			offset := mapRange.destinationStart - mapRange.sourceStart
			return index + offset
		}
	}

	return index
}

func determineMappedIndexAllCategories(categories [][]MapRange, index int) int {
	for categoryIndex := 0; categoryIndex < len(categories) - 1; categoryIndex++ {
		index = determineMappedIndex(categories[categoryIndex], index)
	}
	return index
}

func extractSeedNumbers(line string) (seedNumbers []int) {
	seedNumbersRaw, _ := strings.CutPrefix(line, "seeds: ")
	for _, seedNumberRaw := range strings.Split(seedNumbersRaw, " ") {
		seedNumber, _ := strconv.Atoi(seedNumberRaw)
		seedNumbers = append(seedNumbers, seedNumber)
	}
	return
}

func extractMapRange(line string) (mapRange MapRange) {
	rangeParts := strings.Split(line, " ")
	mapRange.destinationStart, _ = strconv.Atoi(rangeParts[0])
	mapRange.sourceStart, _ = strconv.Atoi(rangeParts[1])
	mapRange.length, _ = strconv.Atoi(rangeParts[2])
	return
}

func extractCategory(lines []string, startIndex int) (mapRanges []MapRange) {
	for offset := 1; startIndex + offset < len(lines); offset++ {
		line := lines[startIndex + offset]
		if line == "" {
			break
		}

		mapRange := extractMapRange(line)
		mapRanges = append(mapRanges, mapRange)
	}

	return
}

func extractCategories(lines []string) (categories [][]MapRange) {
	for index := 0; index < len(lines); index++ {
		line := lines[index]
		if(line == "") {
			category := extractCategory(lines, index + 1)
			categories = append(categories, category)
		}
	}

	return
}

func solvePart1(input string) {
	lines := strings.Split(input, "\n")
	seedNumbers := extractSeedNumbers(lines[0])

	categoryLines := lines[1:]
	categories := extractCategories(categoryLines)

	smallestLocation := determineMappedIndexAllCategories(categories, seedNumbers[0])
	for seedIndex := 1; seedIndex < len(seedNumbers); seedIndex++ {
		smallestLocation = min(smallestLocation, determineMappedIndexAllCategories(categories, seedNumbers[seedIndex]))
	}

	fmt.Println(smallestLocation)
}

func main() {
	input := util.FetchInput(5)
	solvePart1(input)
}
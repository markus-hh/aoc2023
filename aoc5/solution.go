package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"example.com/aoc/util"
)

type MapRange struct {
	sourceStart      int
	destinationStart int
	length           int
}

func determineMappedIndex(category []MapRange, index int) int {
	for _, mapRange := range category {
		if index >= mapRange.sourceStart && index < mapRange.sourceStart+mapRange.length {
			return mapIndex(mapRange, index)
		}
	}

	return index
}

func mapIndex(mapRange MapRange, index int) int {
	offset := mapRange.destinationStart - mapRange.sourceStart
	return index + offset
}

func determineMappedIndexAllCategories(categories [][]MapRange, index int) int {
	for categoryIndex := 0; categoryIndex < len(categories)-1; categoryIndex++ {
		index = determineMappedIndex(categories[categoryIndex], index)
	}
	return index
}

func findMappedMapRangesAllCategories(sortedCategories [][]MapRange, mapRange MapRange) []MapRange {
	mapRanges := []MapRange{mapRange}
	for _, sortedCategory := range sortedCategories {
		mapRanges = findMappedRanges(sortedCategory, mapRanges)
	}

	return mapRanges
}

func findMappedRanges(sortedCategory []MapRange, mapRanges []MapRange) (totalMapRanges []MapRange) {
	for _, mapRange := range mapRanges {
		currentMappedMapRanges := findMappedRange(sortedCategory, mapRange)
		for _, mappedMapRange := range currentMappedMapRanges {
			totalMapRanges = append(totalMapRanges, mappedMapRange)
		}
	}

	return
}

func findMappedRange(sortedCategory []MapRange, mapRange MapRange) (mappedMapRanges []MapRange) {
	firstIndex := mapRange.destinationStart
	lastIndex := mapRange.destinationStart + mapRange.length - 1

	firstCategoryMapRange := sortedCategory[0]

	for currentIndex := firstIndex; currentIndex <= lastIndex; currentIndex++ {
		var endIndex int
		var mappedIndex int
		var length int

		if currentIndex < firstCategoryMapRange.sourceStart {
			endIndex = min(firstCategoryMapRange.sourceStart-1, lastIndex)
			length = endIndex - currentIndex + 1
			mappedIndex = currentIndex
		} else {
			rangeBeforeIndex := findRangeIndexBefore(sortedCategory, currentIndex)
			rangeBefore := sortedCategory[rangeBeforeIndex]

			rangeHit := currentIndex < rangeBefore.sourceStart+rangeBefore.length
			if rangeHit {
				// range hit
				endIndex = min(rangeBefore.sourceStart+rangeBefore.length-1, lastIndex)
				mappedIndex = mapIndex(rangeBefore, currentIndex)
				length = endIndex - currentIndex + 1
			} else {
				// range miss
				if rangeBeforeIndex == len(sortedCategory)-1 {
					endIndex = lastIndex
				} else {
					endIndex = min(lastIndex, sortedCategory[rangeBeforeIndex + 1].sourceStart-1)
				}

				length = endIndex - currentIndex + 1
				mappedIndex = currentIndex
			}
		}

		mappedMapRange := MapRange{
			sourceStart:      currentIndex,
			destinationStart: mappedIndex,
			length:           length,
		}

		mappedMapRanges = append(mappedMapRanges, mappedMapRange)
		currentIndex = endIndex
	}

	return
}

func findRangeIndexBefore(sortedCategory []MapRange, targetIndex int) int {
	for rangeIndex := 0; rangeIndex < len(sortedCategory); rangeIndex++ {
		if sortedCategory[rangeIndex].sourceStart > targetIndex {
			return rangeIndex - 1
		}
	}

	return len(sortedCategory) - 1
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

func extractCategory(lines []string, startIndex int) (category []MapRange) {
	for offset := 1; startIndex+offset < len(lines); offset++ {
		line := lines[startIndex+offset]
		if line == "" {
			break
		}

		mapRange := extractMapRange(line)
		category = append(category, mapRange)
	}

	return
}

func sortCategory(category []MapRange) []MapRange {
	slices.SortFunc(category,
		func(a, b MapRange) int {
			return cmp.Compare(a.sourceStart, b.sourceStart)
		})
	return category
}

func extractCategories(lines []string) (categories [][]MapRange) {
	for index := 0; index < len(lines); index++ {
		line := lines[index]
		if line == "" && index != len(lines) - 1 {
			category := extractCategory(lines, index+1)
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

func solvePart2(input string) {
	lines := strings.Split(input, "\n")
	seedNumbers := extractSeedNumbers(lines[0])

	categoryLines := lines[1:]
	categories := extractCategories(categoryLines)

	for index := 0; index < len(categories); index++ {
		categories[index] = sortCategory(categories[index])
	}

	sortedCategories := categories
	var totalMappedRanges []MapRange

	for seedIndex := 0; 2*seedIndex < len(seedNumbers); seedIndex++ {
		seedNumberStart := seedNumbers[2*seedIndex]
		seedNumberLength := seedNumbers[2*seedIndex+1]

		mapRange := MapRange {
			sourceStart: seedNumberStart,
			destinationStart: seedNumberStart,
			length: seedNumberLength,
		}
		mappedMapRanges := findMappedMapRangesAllCategories(sortedCategories, mapRange)
		
		for _, mappedMapRange := range mappedMapRanges {
			totalMappedRanges = append(totalMappedRanges, mappedMapRange)
		}
	}

	smallestLocation := totalMappedRanges[0].destinationStart
	for index := 1; index < len(totalMappedRanges); index++ {
		mappedMapRange := totalMappedRanges[index]
		smallestLocation = min(smallestLocation, mappedMapRange.destinationStart)
	}

	fmt.Println(smallestLocation)
}

func main() {
	input := util.FetchInput(5)
	solvePart2(input)
}

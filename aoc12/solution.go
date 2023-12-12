package main

import (
	"fmt"
	"os"
	"strings"

	"example.com/aoc/util"
)

const NO_SPRING = 0
const SPRING = 1
const UNKNOWN = 2

type Springs struct {
	occurences []int
	groupSizes []int
}

func parseSpringOccurence(occurenceRaw rune) int {
	switch occurenceRaw {
	case '.':
		return NO_SPRING
	case '#':
		return SPRING
	case '?':
		return UNKNOWN
	default:
		fmt.Println("error during parsing")
		os.Exit(1)
		return -1
	}
}

func isSpring(occurence int) bool {
	return occurence == SPRING
}

func getSpringsCombinations(springs Springs) int {
	if len(springs.groupSizes) == 0 {
		if util.Any(springs.occurences, isSpring) {
			return 0
		} else {
			return 1
		}
	}

	actualGroupSize := springs.groupSizes[0]
	currentGroupSize := 0
	sum := 0
	for index := 0; index < len(springs.occurences); index++ {
		occurence := springs.occurences[index]

		if occurence == NO_SPRING {
			if currentGroupSize == actualGroupSize {
				return sum + getSpringsCombinations(Springs{occurences: springs.occurences[index+1:], groupSizes: springs.groupSizes[1:]})
			} else if currentGroupSize > 0 {
				return sum
			}
		} else if occurence == SPRING {
			currentGroupSize++

			if currentGroupSize > actualGroupSize {
				return sum
			}
		} else if occurence == UNKNOWN {
			if currentGroupSize == 0 {
				sum += getSpringsCombinations(Springs{occurences: springs.occurences[index+1:], groupSizes: springs.groupSizes})
			} else if currentGroupSize == actualGroupSize {
				return sum + getSpringsCombinations(Springs{occurences: springs.occurences[index+1:], groupSizes: springs.groupSizes[1:]})
			}
			currentGroupSize++
		}
	}

	if currentGroupSize == actualGroupSize && len(springs.groupSizes) == 1 {
		return sum + 1
	}

	return 0
}

func parseLine(line string) Springs {
	lineParts := strings.Split(line, " ")
	occurencesRaw := util.Runes(lineParts[0])
	occucences := util.MapFunc(occurencesRaw, parseSpringOccurence)

	groupSizesRaw := strings.Split(lineParts[1], ",")
	groupSizes := util.MapFunc(groupSizesRaw, util.AtoiUnsafe)

	return Springs{occurences: occucences, groupSizes: groupSizes}
}

func solvePart1(input string) {
	lines := util.SplitLines(input)
	springs := util.MapFunc(lines, parseLine)

	sum := 0
	for _, spring := range springs {
		combinations := getSpringsCombinations(spring)
		sum += combinations
	}

	fmt.Println(sum)
}

func main() {
	input := util.FetchInput(12)
	solvePart1(input)
}

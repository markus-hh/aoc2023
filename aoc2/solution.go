package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"example.com/aoc/util"
)

type bagSubset struct {
	redCubes int
	blueCubes int
	greenCubes int
}

func isGamePossible(bagCapacity bagSubset, bagSubsets []bagSubset) bool {
	for _, subset := range bagSubsets {
		if(subset.blueCubes > bagCapacity.blueCubes ||
			subset.greenCubes > bagCapacity.greenCubes ||
			subset.redCubes > bagCapacity.redCubes) {
			return false
		}
	}

	return true
}

func main() {
	input := util.FetchInput(2)
	bagCapacity := bagSubset{
		redCubes: 12,
		greenCubes: 13,
		blueCubes: 14,
	}

	lines := util.SplitLines(input)
	
	sum := 0
	for _, line := range lines {
		id, subsets := parseLine(line)
		if isGamePossible(bagCapacity, subsets) {
			sum += id
		}
	}

	fmt.Println(sum)
}

func parseLine(line string) (id int, subsets []bagSubset) {
	regex := regexp.MustCompile(`Game (\d+): (.*)`)
	regexResult := regex.FindAllStringSubmatch(line, -1)[0]

	id, _ = strconv.Atoi(regexResult[1])
	subsets = parseSubsets(regexResult[2])

	return
}

func parseSubsets(lineSegment string) (subsets []bagSubset) {
	for _, splitLineSegment := range strings.Split(lineSegment, "; ") {
		subsets = append(subsets, parseSubset(splitLineSegment))
	}

	return subsets
}

func parseSubset(lineSegment string) bagSubset {
	subset := bagSubset{}
	cubeSets := strings.Split(lineSegment, ", ")
	
	for _, cubes := range cubeSets {
		cubeParts := strings.Split(cubes, " ")
		amount, _ := strconv.Atoi(cubeParts[0])

		switch cubeParts[1] {
		case "red": subset.redCubes = amount
		case "green": subset.greenCubes = amount
		case "blue": subset.blueCubes = amount
		}
	}

	return subset
}
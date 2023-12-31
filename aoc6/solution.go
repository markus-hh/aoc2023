package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"example.com/aoc/util"
)

func parseLines(lines []string) (timeLine string, distanceLine string) {
	timeLine = util.RemoveDuplicatesFromString(lines[0], " ")
	timeLine, _ = strings.CutPrefix(timeLine, "Time: ")
	distanceLine = util.RemoveDuplicatesFromString(lines[1], " ")
	distanceLine, _ = strings.CutPrefix(distanceLine, "Distance: ")
	return
}

func parseInput(lines []string) (times []int, distances []int) {
	timeLine, distanceLine := parseLines(lines)

	for _, rawTime := range strings.Split(timeLine, " ") {
		time, _ := strconv.Atoi(rawTime)
		times = append(times, time)
	}

	for _, rawDistance := range strings.Split(distanceLine, " ") {
		distance, _ := strconv.Atoi(rawDistance)
		distances = append(distances, distance)
	}

	return
}

func parseInputPart2(lines[] string) (time int, distance int) {
	timeLine, distanceLine := parseLines(lines)
	timeLine = strings.Replace(timeLine, " ", "", -1)
	distanceLine = strings.Replace(distanceLine, " ", "", -1)

	time, _ = strconv.Atoi(timeLine)
	distance, _ = strconv.Atoi(distanceLine)
	return
}

func determineWinCombinationAmount(time int, distance int) int {
	b := float64(-time)
	c := float64(distance)

	firstWinningHoldTime, lastWinningHoldTime := util.SolveQuadraticEquation(1, b, c)

	var firstWinningHoldTimeInt int
	if util.FloatEquals(firstWinningHoldTime, math.Floor(firstWinningHoldTime) + 1) {
		firstWinningHoldTimeInt = int(math.Floor(firstWinningHoldTime) + 2)
	} else {
		firstWinningHoldTimeInt = int(math.Floor(firstWinningHoldTime) + 1)
	}

	var lastWinningHoldTimeInt int
	if util.FloatEquals(lastWinningHoldTime, math.Ceil(lastWinningHoldTime) - 1) {
		lastWinningHoldTimeInt = int(math.Ceil(lastWinningHoldTime) - 2)
	} else {
		lastWinningHoldTimeInt = int(math.Ceil(lastWinningHoldTime) - 1)
	}

	return lastWinningHoldTimeInt - firstWinningHoldTimeInt + 1
}

func solvePart1(input string) {
	lines := util.SplitLines(input)
	times, distances := parseInput(lines)

	solution := 1
	for index := 0; index < len(times); index++ {
		time := times[index]
		distance := distances[index]
		combinationAmount := determineWinCombinationAmount(time, distance)
		solution *= combinationAmount
	}

	fmt.Println(solution)
}

func solvePart2(input string) {
	lines := util.SplitLines(input)
	time, distance := parseInputPart2(lines)

	combinationAmount := determineWinCombinationAmount(time, distance)
	fmt.Println(combinationAmount)
}

func main() {
	input := util.FetchInput(6)
	solvePart2(input)
}
package main

import (
	"fmt"
	"strings"

	"example.com/aoc/util"
)

func getSequenceValue(sequence string) int {
	value := 0
	for _, sequenceRune := range sequence {
		value += int(sequenceRune)
		value *= 17
		value %= 256
	}
	return value
}

func solvePart1(input string) {
	input = strings.TrimSpace(input)
	sequences := strings.Split(input, ",")

	sum := 0
	for _, sequence := range sequences {
		sum += getSequenceValue(sequence)
	}
	fmt.Println(sum)
}

func main() {
	input := util.FetchInput(15)
	solvePart1(input)
}
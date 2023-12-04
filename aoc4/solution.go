package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"example.com/aoc/util"
)

type ScratchCard struct {
	winningNumbers []int
	chosenNumbers  []int
}

func fillNumbersFromString(numbers []int, numbersRaw string) []int {
	for _, numberRaw := range util.SplitWithoutDuplicates(numbersRaw, " ") {
		number, _ := strconv.Atoi(numberRaw)
		numbers = append(numbers, number)
	}

	return numbers
}

func determineCardValue(scratchCard ScratchCard) (cardValue int) {
	for _, chosenNumber := range util.RemoveDuplicates[int](scratchCard.chosenNumbers) {
		isWinningNumber := slices.Contains(scratchCard.winningNumbers, chosenNumber)

		if isWinningNumber {
			if cardValue == 0 {
				cardValue = 1
			} else {
				cardValue *= 2
			}
		}
	}

	return
}

func parseScratchCard(line string) (card ScratchCard) {
	_, numbersRaw, _ := strings.Cut(line, ": ")
	splitNumbersRaw := strings.Split(numbersRaw, " | ")
	winningNumbersRaw := splitNumbersRaw[0]
	chosenNumbersRaw := splitNumbersRaw[1]

	card.winningNumbers = fillNumbersFromString(card.winningNumbers, winningNumbersRaw)
	card.chosenNumbers = fillNumbersFromString(card.chosenNumbers, chosenNumbersRaw)

	return
}

func parseScratchCards(lines []string) (cards []ScratchCard) {
	for _, line := range lines {
		cards = append(cards, parseScratchCard(line))
	}
	return
}

func solvePart1(input string) {
	lines := util.SplitLines(input)
	scratchCards := parseScratchCards(lines)

	sum := 0
	for _, scratchCard := range scratchCards {
		sum += determineCardValue(scratchCard)
	}

	fmt.Println(sum)
}

func main() {
	input := util.FetchInput(4)

	solvePart1(input)
	// fmt.Println(input)
}

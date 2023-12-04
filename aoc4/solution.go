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

func determineMatchNumber(scratchCard ScratchCard) (numberOfMatches int) {
	for _, chosenNumber := range util.RemoveDuplicates[int](scratchCard.chosenNumbers) {
		isWinningNumber := slices.Contains(scratchCard.winningNumbers, chosenNumber)

		if isWinningNumber {
			numberOfMatches++
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

func solvePart2(input string) {
	lines := util.SplitLines(input)
	scratchCards := parseScratchCards(lines)

	var scratchCardAmounts = make([]int, len(scratchCards))
	for index := 0; index < len(scratchCards); index++ {
		scratchCardAmounts[index] = 1
	}

	sum := 0
	for index := 0; index < len(scratchCards); index++ {
		scratchCard := scratchCards[index]
		scratchCardAmount := scratchCardAmounts[index]
		sum += scratchCardAmount

		matchNumber := determineMatchNumber(scratchCard)

		for lookahead := 1; lookahead <= matchNumber; lookahead++ {
			scratchCardAmounts[index + lookahead] += scratchCardAmount
		}
	}

	fmt.Println(sum)
}

func main() {
	input := util.FetchInput(4)

	solvePart2(input)
	// fmt.Println(input)
}

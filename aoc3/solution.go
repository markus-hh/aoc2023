package main

import (
	"fmt"
	"strconv"
	"unicode"

	"example.com/aoc/util"
)

func isSymbolInString(input string) bool {
	for _, rune := range input {
		if(!unicode.IsDigit(rune) && rune != '.') {
			return true
		} 
	}

	return false
}

func isPartNumber(lines []string, lineIndex int, numberStartIndex int, numberEndIndex int) bool {
	currentLine := lines[lineIndex]
	currntLineFirstIndex := 0
	currentLineLastIndex := len(currentLine) - 1

	leftCheckIndex := max(numberStartIndex - 1, currntLineFirstIndex)
	rightCheckIndex := min(numberEndIndex + 1, currentLineLastIndex)

	topCheck    := lineIndex != 0 && isSymbolInString(lines[lineIndex - 1][leftCheckIndex:rightCheckIndex + 1])
	bottomCheck := lineIndex != len(lines) - 1 && isSymbolInString(lines[lineIndex + 1][leftCheckIndex:rightCheckIndex + 1])
	leftCheck  := numberStartIndex != currntLineFirstIndex && isSymbolInString(currentLine[numberStartIndex - 1:numberStartIndex])
	rightCheck := numberEndIndex   != currentLineLastIndex && isSymbolInString(currentLine[numberEndIndex   + 1:numberEndIndex + 2])

	return topCheck || bottomCheck || leftCheck || rightCheck
}

func extractPartNumbersAsStrings(lines []string, lineIndex int) (numbers []string) {
	line := lines[lineIndex]

	currentNumber := ""
	currentNumberStartIndex := 0

	for currentIndex, rune := range line {
		if !unicode.IsDigit(rune) {
			if(currentNumber != "") {
				if(isPartNumber(lines, lineIndex, currentNumberStartIndex, currentIndex - 1)) {
					numbers = append(numbers, currentNumber)	
				}
				
				currentNumber = ""
			}

			currentNumberStartIndex = currentIndex + 1
			continue
		}

		currentNumber += string(rune)
	}

	if(currentNumber != "" && isPartNumber(lines, lineIndex, currentNumberStartIndex, len(line) - 1)) {
		numbers = append(numbers, currentNumber)
	}
	return
}

func findPartNumberSum(lines []string, index int) (sum int) {
	numbers := extractPartNumbersAsStrings(lines, index)
	for _, number := range numbers {
		numberAsInt, _ := strconv.Atoi(number)
		sum += numberAsInt
	}

	return
}

func solvePart1(input string) {
	lines := util.SplitLines(input)

	sum := 0

	for lineIndex := 0; lineIndex < len(lines); lineIndex++ {
		lineSum := findPartNumberSum(lines, lineIndex)
		sum += lineSum
	}

	fmt.Println(sum)
}

func main()  {
	input := util.FetchInput(3)
	solvePart1(input)
}
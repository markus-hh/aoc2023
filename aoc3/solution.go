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

func extractPartNumbersIndexed(lines []string, lineIndex int) ([]int) {
	line := lines[lineIndex]
	var numberAtPosition []int
	for i := 0; i < len(line); i++ {
		numberAtPosition = append(numberAtPosition, -1)
	}

	currentNumber := ""
	currentNumberStartIndex := 0

	for currentIndex, rune := range line {
		if !unicode.IsDigit(rune) {
			if(currentNumber != "") {
				if(isPartNumber(lines, lineIndex, currentNumberStartIndex, currentIndex - 1)) {
					currentNumberInt, _ := strconv.Atoi(currentNumber)
					for i := currentNumberStartIndex; i < currentIndex; i++ {
						numberAtPosition[i] = currentNumberInt
					}	
				}
				
				currentNumber = ""
			}

			currentNumberStartIndex = currentIndex + 1
			continue
		}

		currentNumber += string(rune)
	}

	if(currentNumber != "" && isPartNumber(lines, lineIndex, currentNumberStartIndex, len(line) - 1)) {
		currentNumberInt, _ := strconv.Atoi(currentNumber)
		for i := currentNumberStartIndex; i < len(line); i++ {
			numberAtPosition[i] = currentNumberInt
		}
	}

	return numberAtPosition
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

func buildPartNumberMap(lines []string) (partNumberMap [][]int) {
	for i := 0; i < len(lines); i++ {
		partNumberMap = append(partNumberMap, extractPartNumbersIndexed(lines, i))
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

func findDistinctPartNumbersInRange(numberAtPositionLine []int, startIndex int, endIndex int) (partNumbers[]int) {
	isNumberFocused := false
	for index := startIndex; index <= endIndex; index++ {
		currentNumber := numberAtPositionLine[index]
		if currentNumber == -1 {
			isNumberFocused = false
		} else if(!isNumberFocused) {
			isNumberFocused = true
			partNumbers = append(partNumbers, currentNumber)
		}
	}

	return
}

func findSurroundingPartNumbers(numberAtPosition [][]int, lineIndex int, currentIndex int) (partNumbers []int) {
	lineLastIndex := len(numberAtPosition[lineIndex]) - 1
	currntLineFirstIndex := 0
	currentLineLastIndex := len(numberAtPosition[lineIndex]) - 1

	leftCheckIndex := max(currentIndex - 1, currntLineFirstIndex)
	rightCheckIndex := min(currentIndex + 1, currentLineLastIndex)
	

	if(currentIndex != 0 && numberAtPosition[lineIndex][currentIndex - 1] != -1) {
		partNumbers = append(partNumbers, numberAtPosition[lineIndex][currentIndex - 1])
	}

	if(currentIndex != lineLastIndex && numberAtPosition[lineIndex][currentIndex + 1] != -1) {
		partNumbers = append(partNumbers, numberAtPosition[lineIndex][currentIndex + 1])
	}

	if(lineIndex != 0) {
		partNumbers = append(partNumbers, findDistinctPartNumbersInRange(numberAtPosition[lineIndex - 1], leftCheckIndex, rightCheckIndex)...)
	}

	if(lineIndex != len(numberAtPosition) - 1) {
		partNumbers = append(partNumbers, findDistinctPartNumbersInRange(numberAtPosition[lineIndex + 1], leftCheckIndex, rightCheckIndex)...)
	}

	return
}

func findGearRatioSumInLine(numbersAtPosition [][]int, lines []string, lineIndex int) (sum int) {
	line := lines[lineIndex]

	for currentIndex, currentRune := range line {
		if(currentRune != '*') {
			continue
		}

		surroundingPartNumbers := findSurroundingPartNumbers(numbersAtPosition, lineIndex, currentIndex)
		if(len(surroundingPartNumbers) == 2) {
			sum += surroundingPartNumbers[0] * surroundingPartNumbers[1]
		}
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

func solvePart2(input string) {
	lines := util.SplitLines(input)
	sum := 0
	partNumberMap := buildPartNumberMap(lines)

	for index := 0; index < len(lines); index++ {
		sum += findGearRatioSumInLine(partNumberMap, lines, index)
	}

	fmt.Println(sum)
}

func main()  {
	input := util.FetchInput(3)
	solvePart2(input)
}
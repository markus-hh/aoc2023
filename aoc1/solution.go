package main

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/aoc/util"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "0": 0,
}

func main() {

	input := util.FetchInput(1)

	sum := 0
	for _, line := range strings.Split(input, "\n") {
        firstNumber := extractNumberForwards(line)
        lastNumber := extractNumberBackwards(reverseString(line))

        combinedNumber, _ := strconv.Atoi(strconv.Itoa(firstNumber) + strconv.Itoa(lastNumber))
		sum += combinedNumber
	}

	fmt.Println(sum)
}

func reverseString(input string) string {
    reversedString := ""
    for _, character := range input {
        reversedString = string(character) + reversedString
    }
    return reversedString
}

func extractNumberBackwards(reversedLine string) int {
	if reversedLine == "" {
		return 0
	}

	lineSegment := ""

	for _, character := range reversedLine {
		lineSegment = string(character) + lineSegment

        for writtenDigit, numericDigit := range digitMap {
            if strings.Contains(lineSegment, writtenDigit) {
                return numericDigit
            }
        }
	}

    return 0
}


func extractNumberForwards(line string) int {
	if line == "" {
		return 0
	}

	lineSegment := ""

	for _, character := range line {
		lineSegment += string(character)

        for writtenDigit, numericDigit := range digitMap {
            if strings.Contains(lineSegment, writtenDigit) {
                return numericDigit
            }
        }
	}

	return 0
}

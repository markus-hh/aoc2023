package main

import (
	"fmt"
	"strings"

	"example.com/aoc/util"
)

func predictNextNumber(numbers []int) int {
	nextNumberAggregate := util.LastFrom(numbers)

	for {
		if len(numbers) == 1 {
			nextNumberAggregate += numbers[0]
			break
		}

		numbersNextLine := make([]int, len(numbers)-1)
		for index := 0; index < len(numbersNextLine); index++ {
			delta := numbers[index+1] - numbers[index]
			numbersNextLine[index] = delta
		}

		numbers = numbersNextLine
		finished := !util.Any(numbers, isNotZero)

		if finished {
			break
		}

		nextNumberAggregate += util.LastFrom(numbers)
	}

	return nextNumberAggregate
}

func predictPreviousNumber(numbers []int) int {
	previousNumberAggregate := numbers[0]

	subtractNextNumber := true
	for {
		if len(numbers) == 1 {
			previousNumberAggregate += numbers[0]
			break
		}

		numbersNextLine := make([]int, len(numbers)-1)
		for index := 0; index < len(numbersNextLine); index++ {
			delta := numbers[index+1] - numbers[index]
			numbersNextLine[index] = delta
		}

		numbers = numbersNextLine
		finished := !util.Any(numbers, isNotZero)

		if finished {
			break
		}

		if subtractNextNumber {
			previousNumberAggregate -= numbers[0]
		} else {
			previousNumberAggregate += numbers[0]
		}
		subtractNextNumber = !subtractNextNumber
	}

	return previousNumberAggregate
}

func isNotZero(number int) bool {
	return number != 0
}

func extractNumbersFromLine(line string) []int {
	numberStrings := strings.Split(line, " ")
	return util.MapFunc(numberStrings, util.AtoiUnsafe)
}

func solvePart2(input string) {
	lines := util.SplitLines(input)

	sum := 0
	for _, line := range lines {
		numbers := extractNumbersFromLine(line)
		predictedNumber := predictPreviousNumber(numbers)
		sum += predictedNumber
	}

	fmt.Println(sum)
}

func solvePart1(input string) {
	lines := util.SplitLines(input)

	sum := 0
	for _, line := range lines {
		numbers := extractNumbersFromLine(line)
		predictedNumber := predictNextNumber(numbers)
		sum += predictedNumber
	}

	fmt.Println(sum)
}

func main() {
	input := util.FetchInput(9)
	solvePart2(input)
}

package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"example.com/aoc/util"
)

const JOKER_KEY = 0

type Hand struct {
	orders []int
	bid    int
}

type GroupedHands struct {
	fiveOfAKinds  []Hand
	fourOfAKinds  []Hand
	fullHouses    []Hand
	threeOfAKinds []Hand
	twoPairs      []Hand
	onePairs      []Hand
	highCards     []Hand
}

func sortHands(hands []Hand) []Hand {
	slices.SortFunc(hands,
		func(a, b Hand) int {
			return slices.Compare(a.orders, b.orders)
		})
	return hands
}

func sortGroupedHands(groupedHands GroupedHands) (sortedHands []Hand) {
	sortedHands = util.AddAll(sortedHands, sortHands(groupedHands.highCards))
	sortedHands = util.AddAll(sortedHands, sortHands(groupedHands.onePairs))
	sortedHands = util.AddAll(sortedHands, sortHands(groupedHands.twoPairs))
	sortedHands = util.AddAll(sortedHands, sortHands(groupedHands.threeOfAKinds))
	sortedHands = util.AddAll(sortedHands, sortHands(groupedHands.fullHouses))
	sortedHands = util.AddAll(sortedHands, sortHands(groupedHands.fourOfAKinds))
	sortedHands = util.AddAll(sortedHands, sortHands(groupedHands.fiveOfAKinds))
	return
}

func parseHand(line string) Hand {
	handParts := strings.Split(line, " ")
	orders := []int{}

	for _, rawCard := range handParts[0] {
		orders = append(orders, mapCardToOrder(rawCard))
	}
	bid, _ := strconv.Atoi(handParts[1])

	return Hand{orders: orders, bid: bid}
}

func groupHands(hands []Hand) (groupdHands GroupedHands) {
	for _, hand := range hands {
		histogram := util.MakeHistogram[int](hand.orders)
		jokerCards := histogram[JOKER_KEY]
		delete(histogram, JOKER_KEY)
		// guard value to ensure there is always at least one element in sortedHandCards

		sortedHandCards := util.SortMapByValue[int, int](histogram)
		slices.Reverse(sortedHandCards)

		var highestCardOccurence int
		if jokerCards == 5 {
			highestCardOccurence = 5
		} else {
			highestCardOccurence = sortedHandCards[0].Value + jokerCards
		}
		
		if highestCardOccurence == 5 {
			groupdHands.fiveOfAKinds = append(groupdHands.fiveOfAKinds, hand)
		} else if highestCardOccurence == 4 {
			groupdHands.fourOfAKinds = append(groupdHands.fourOfAKinds, hand)
		} else if highestCardOccurence == 3 && sortedHandCards[1].Value == 2 {
			groupdHands.fullHouses = append(groupdHands.fullHouses, hand)
		} else if highestCardOccurence == 3 {
			groupdHands.threeOfAKinds = append(groupdHands.threeOfAKinds, hand)
		} else if highestCardOccurence == 2 && sortedHandCards[1].Value == 2 {
			groupdHands.twoPairs = append(groupdHands.twoPairs, hand)
		} else if highestCardOccurence == 2 {
			groupdHands.onePairs = append(groupdHands.onePairs, hand)
		} else {
			groupdHands.highCards = append(groupdHands.highCards, hand)
		}
	}
	return
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
func mapCardToOrder(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		// part 1 solution: return 11
		return JOKER_KEY
	case 'T':
		return 10
	default:
		order, _ := strconv.Atoi(string(card))
		return order
	}
}

func solvePart1(input string) {
	lines := util.SplitLines(input)
	hands := util.MapFunc[string, Hand](lines, parseHand)
	groupedHands := groupHands(hands)
	sortedHands := sortGroupedHands(groupedHands)

	sum := 0
	for index, hand := range sortedHands {
		rank := index + 1
		value := rank * hand.bid
		sum += value
	}

	fmt.Println(sum)
}

func solvePart2(input string) {
	solvePart1(input)
}

func main() {
	input := util.FetchInput(7)
	solvePart2(input)
}

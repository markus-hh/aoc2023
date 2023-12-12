package util

import (
	"cmp"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Position2D struct {
	X int
	Y int
}

type MapElement[T, U any] struct {
	Key   T
	Value U
}

func FetchInput(day int) string {
	if len(os.Args) == 1 {
		fmt.Println("No session id specified, cannot fetch input!")
		os.Exit(1)
	}

	sessionId := os.Args[1]
	url := "https://adventofcode.com/2023/day/" + strconv.Itoa(day) + "/input"
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Cookie", "session="+sessionId)
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	response, error := http.DefaultClient.Do(request)

	if error != nil {
		return "error"
	}

	defer response.Body.Close()
	body, error := io.ReadAll(response.Body)

	return string(body)
}

func SplitLines(input string) (lines []string) {
	trimmedInput := strings.Trim(input, "\n")
	for _, line := range strings.Split(trimmedInput, "\n") {
		lines = append(lines, line)
	}
	return
}

func RemoveDuplicatesFromString(input string, seperator string) string {
	for {
		lengthBefore := len(input)
		input = strings.ReplaceAll(input, seperator+seperator, seperator)
		lengthAfter := len(input)

		if lengthBefore == lengthAfter {
			return input
		}
	}
}

func SplitWithoutDuplicates(input string, seperator string) []string {
	processedInput := strings.Trim(RemoveDuplicatesFromString(input, seperator), " ")
	return strings.Split(processedInput, seperator)
}

func RemoveDuplicates[T comparable](inputSlice []T) (sliceWithoutDuplicates []T) {
	elementMap := make(map[T]bool)

	for _, element := range inputSlice {
		elementMap[element] = true
	}

	keys := make([]T, 0, len(elementMap))
	for k := range elementMap {
		keys = append(keys, k)
	}

	return keys
}

func SolveQuadraticEquation(a float64, b float64, c float64) (leftX float64, rightX float64) {
	rootValue := math.Sqrt(b*b - 4*a*c)
	leftX = (-b - rootValue) / (2 * a)
	rightX = (-b + rootValue) / (2 * a)
	return
}

func FloatEquals(a float64, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

// https://stackoverflow.com/a/71624929
func MapFunc[T, U any](originalSlice []T, mappingFunc func(T) U) []U {
	mappedSlice := make([]U, len(originalSlice))
	for index := range originalSlice {
		mappedSlice[index] = mappingFunc(originalSlice[index])
	}
	return mappedSlice
}

func FilterFunc[T any](originalSlice []T, filterFunc func(T) bool) []T {
	filteredSlice := []T{}
	for _, element := range originalSlice {
		if filterFunc(element) {
			filteredSlice = append(filteredSlice, element)
		}
	}
	return filteredSlice
}

func InitSlice[T any](slice []T, value T) []T {
	for index := range slice {
		slice[index] = value
	}
	return slice
}

func FillSlice[T any](slice []T, value T, amount int) []T {
	for index := 0; index < amount; index++ {
		slice = append(slice, value)
	}
	return slice
}

func MakeHistogram[T comparable](slice []T) (histogram map[T]int) {
	histogram = map[T]int{}
	for _, element := range slice {
		histogram[element]++
	}
	return
}

func SortMapByKey[K cmp.Ordered, V any](inputMap map[K]V) (sortedElements []MapElement[K, V]) {
	for key, value := range inputMap {
		element := MapElement[K, V]{
			Key:   key,
			Value: value,
		}
		sortedElements = append(sortedElements, element)
	}

	slices.SortFunc(sortedElements,
		func(a, b MapElement[K, V]) int {
			return cmp.Compare(a.Key, b.Key)
		})
	return
}

func SortMapByValue[K comparable, V cmp.Ordered](inputMap map[K]V) (sortedElements []MapElement[K, V]) {
	for key, value := range inputMap {
		element := MapElement[K, V]{
			Key:   key,
			Value: value,
		}
		sortedElements = append(sortedElements, element)
	}

	slices.SortFunc(sortedElements,
		func(a, b MapElement[K, V]) int {
			return cmp.Compare(a.Value, b.Value)
		})
	return
}

func AddAll[T any](container []T, newElements []T) []T {
	for _, element := range newElements {
		container = append(container, element)
	}
	return container
}

func AtoiUnsafe(input string) int {
	number, _ := strconv.Atoi(input)
	return number
}

func LastFrom[T any](slice []T) T {
	return slice[len(slice)-1]
}

func Any[T any](slice []T, predicateFunc func(T)bool) bool {
	for _, element := range slice {
		if predicateFunc(element) {
			return true
		}
	}

	return false
}

func NotAny[T any](slice []T, predicateFunc func(T)bool) bool {
	for _, element := range slice {
		if predicateFunc(element) {
			return false
		}
	}

	return true
}

func Min[T cmp.Ordered](slice []T) T {
	minValue := slice[0]
	for index := 1; index < len(slice); index++ {
		minValue = min(minValue, slice[index])
	}

	return minValue
}

func Identity[T any](value T) T {
	return value
}

func Copy[T any](slice []T) []T {
	copiedSlice := make([]T, len(slice))
	copy(copiedSlice, slice)
	return copiedSlice
}

func SortPositions(positions []Position2D) {
	sort.Slice(positions, func(i, j int) bool {
		if positions[i].Y != positions[j].Y {
			return cmp.Less(positions[i].Y, positions[j].Y)
		}
		return cmp.Less(positions[i].X, positions[j].X)
	})
}

func HammingDistance(firstPosition Position2D, secondPosition Position2D) int {
	return int(math.Abs(float64(firstPosition.X) - float64(secondPosition.X))) + int(math.Abs(float64(firstPosition.Y) - float64(secondPosition.Y)))
}

func Runes(line string) (runes []rune) {
	for _, lineRune := range line {
		runes = append(runes, lineRune)
	}
	return
} 
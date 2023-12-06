package util

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			lines = append(lines, line)
		}
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
	leftX = (-b - rootValue) / (2*a)
	rightX = (-b + rootValue) / (2*a)
	return
}

func FloatEquals(a float64, b float64) bool {
	return math.Abs(a - b) < 1e-9
}
package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func FetchInput(number int) string {
	if len(os.Args) == 1 {
		fmt.Println("No session id specified, cannot fetch input!")
		os.Exit(1)
	}

	sessionId := os.Args[1]
	url := "https://adventofcode.com/2023/day/" + strconv.Itoa(number) + "/input"
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

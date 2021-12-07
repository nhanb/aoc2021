package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func downloadText(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "session=foo")
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	check(err)
	return string(body)
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

func readLines(path string) []string {
	text := strings.Trim(readFile(path), "\n")
	return strings.Split(text, "\n")
}

func stringsToInts(strings []string) []int {
	// Not having a generic map() is pretty goofy ngl
	ints := make([]int, len(strings))
	for index, str := range strings {
		ints[index], _ = strconv.Atoi(str)
	}
	return ints
}

func binaryToDecimal(input string) int64 {
	result, _ := strconv.ParseInt(string(input), 2, 64)
	return result
}

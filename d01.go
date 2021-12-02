package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func downloadText(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "session="+AOC_SESSION)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func main() {
	text := downloadText("https://adventofcode.com/2021/day/1/input")
	fmt.Println(">>", text)
}

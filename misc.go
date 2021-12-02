package main

import (
	"io"
	"log"
	"net/http"
	"os"
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

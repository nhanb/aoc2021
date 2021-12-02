package main

import (
	"strconv"
	"strings"
)

func main() {
	text := strings.Trim(readFile("d01.input"), "\n")
	numIncreases := 0
	lines := strings.Split(text, "\n")
	previousNum, err := strconv.Atoi(lines[0])
	check(err)
	for _, line := range lines[1:] {
		currentNum, err := strconv.Atoi(line)
		check(err)
		if currentNum > previousNum {
			numIncreases++
			println("Increased")
		} else {
			println("---")
		}
		previousNum = currentNum
	}
	println("Result:", numIncreases)
}

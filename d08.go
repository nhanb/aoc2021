package main

import (
	"fmt"
	"strings"
)

type Mapping struct {
	one   rune
	four  rune
	seven rune
	eight rune
}

func day08() {
	lines := readLines("d08.input")
	fmt.Println(lines)

	count := 0
	for _, line := range lines {
		digitsPart := strings.Split(line, "|")[1][1:]
		digits := strings.Split(digitsPart, " ")
		for _, digit := range digits {
			if find(len(digit), []int{2, 3, 4, 7}) != -1 {
				count++
			}
		}
		fmt.Println(digits)
	}
	fmt.Println("Count:", count)
}

func find(needle int, haystack []int) int {
	for i, el := range haystack {
		if el == needle {
			return i
		}
	}
	return -1
}

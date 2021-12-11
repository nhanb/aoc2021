package main

import (
	"fmt"
	"regexp"
)

type VentLine struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func day05() {
	linePattern := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
	lines := readLines("d05.input")
	var ventLines []VentLine
	for _, line := range lines {
		matches := linePattern.FindAllStringSubmatch(line, -1)[0]
		coords := stringsToInts([]string{
			matches[1],
			matches[2],
			matches[3],
			matches[4],
		})
		ventLines = append(
			ventLines,
			VentLine{coords[0], coords[1], coords[2], coords[3]},
		)
	}
	fmt.Println(ventLines)
}

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
	var ventsMap [1000][1000]int

	// Parse input
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

	// Populate the vents map
	for _, vline := range ventLines {
		if vline.x1 == vline.x2 {
			if vline.y1 == vline.y2 {
				fmt.Println("Single dot:", vline.x1, vline.x2)
				ventsMap[vline.y1][vline.x1]++
			} else {
				ys := orderAscending(vline.y1, vline.y2)
				fmt.Println("Ys:", ys, "- X:", vline.x1)
				for y := ys[0]; y <= ys[1]; y++ {
					ventsMap[y][vline.x1]++
				}
			}
		} else if vline.y1 == vline.y2 {
			xs := orderAscending(vline.x1, vline.x2)
			fmt.Println("Xs:", xs)
			for x := xs[0]; x <= xs[1]; x++ {
				ventsMap[vline.y1][x]++
			}
		}
	}

	twoOrMoreOverlaps := 0
	for _, row := range ventsMap {
		for _, value := range row {
			if value >= 2 {
				twoOrMoreOverlaps++
			}
		}
	}
	println("Part 1 result:", twoOrMoreOverlaps)
}

func orderAscending(a int, b int) [2]int {
	if a < b {
		return [2]int{a, b}
	} else {
		return [2]int{b, a}
	}
}

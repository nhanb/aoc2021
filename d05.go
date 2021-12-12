package main

import (
	"fmt"
	"regexp"
)

type Point struct {
	x int
	y int
}

type VentLine struct {
	start Point
	end   Point
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
			VentLine{
				Point{coords[0], coords[1]},
				Point{coords[2], coords[3]},
			},
		)
	}

	// Populate the vents map
	for _, vline := range ventLines {
		start := vline.start
		end := vline.end
		points := pointsInLine(start, end)
		for _, point := range points {
			ventsMap[point.y][point.x]++
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
	fmt.Println("Part 1 result:", twoOrMoreOverlaps)
}

func orderAscending(a int, b int) [2]int {
	if a < b {
		return [2]int{a, b}
	} else {
		return [2]int{b, a}
	}
}

func pointsInLine(start Point, end Point) []Point {
	if start == end {
		return []Point{start}
	}

	var points []Point

	if start.x == end.x {
		ys := orderAscending(start.y, end.y)
		for y := ys[0]; y <= ys[1]; y++ {
			points = append(points, Point{start.x, y})
		}
	} else if start.y == end.y {
		xs := orderAscending(start.x, end.x)
		for x := xs[0]; x <= xs[1]; x++ {
			points = append(points, Point{x, start.y})
		}
	} else { // diagonal

	}

	return points
}

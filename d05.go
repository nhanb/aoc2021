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
	fmt.Println("Result:", twoOrMoreOverlaps)
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

	if start.x == end.x { // vertical
		ys := inclusiveIntRange(start.y, end.y)
		for _, y := range ys {
			points = append(points, Point{start.x, y})
		}
	} else if start.y == end.y { // horizontal
		xs := inclusiveIntRange(start.x, end.x)
		for _, x := range xs {
			points = append(points, Point{x, start.y})
		}
	} else { // diagonal
		xs := inclusiveIntRange(start.x, end.x)
		ys := inclusiveIntRange(start.y, end.y)
		for i := 0; i < len(xs); i++ {
			points = append(points, Point{xs[i], ys[i]})
		}
	}

	return points
}

func inclusiveIntRange(start int, end int) []int {
	var result []int
	if start < end {
		for i := start; i <= end; i++ {
			result = append(result, i)
		}
	} else {
		for i := start; i >= end; i-- {
			result = append(result, i)
		}
	}
	return result
}

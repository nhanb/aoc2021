package main

import (
	"log"
	"strconv"
	"strings"
)

func day02Part1() {
	horizontalPosition := 0
	depth := 0

	lines := readLines("d02.input")
	for _, line := range lines {
		tuple := strings.Split(line, " ")
		command := tuple[0]
		value, _ := strconv.Atoi(tuple[1])
		switch command {
		case "forward":
			horizontalPosition += value
		case "up":
			depth -= value
		case "down":
			depth += value
		default:
			log.Fatal("Unexpected input:", line)
		}
	}
	println("Horizontal position:", horizontalPosition)
	println("Dept:", depth)
	println("Result:", horizontalPosition*depth)
}

func day02Part2() {
	horizontalPosition := 0
	depth := 0
	aim := 0

	lines := readLines("d02.input")
	for _, line := range lines {
		tuple := strings.Split(line, " ")
		command := tuple[0]
		value, _ := strconv.Atoi(tuple[1])
		switch command {
		case "forward":
			horizontalPosition += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		default:
			log.Fatal("Unexpected input:", line)
		}
	}
	println("Horizontal position:", horizontalPosition)
	println("Dept:", depth)
	println("Result:", horizontalPosition*depth)
}

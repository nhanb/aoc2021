package main

import (
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {
	text := strings.Trim(readFile("d01.input"), "\n")
	lines := strings.Split(text, "\n")
	numIncreases := 0
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

func part2() {
	text := strings.Trim(readFile("d01.input"), "\n")
	lines := strings.Split(text, "\n")
	numIncreases := 0
	n1, _ := strconv.Atoi(lines[0])
	n2, _ := strconv.Atoi(lines[1])
	n3, _ := strconv.Atoi(lines[2])
	previousSum := n1 + n2 + n3

	for _, line := range lines[3:] {
		n1 = n2
		n2 = n3
		n3, _ = strconv.Atoi(line)
		currentSum := n1 + n2 + n3
		println(n1, n2, n3, previousSum, currentSum)
		if currentSum > previousSum {
			numIncreases++
			println("+++")
		} else {
			println("---")
		}
		previousSum = currentSum

	}
	println("Result:", numIncreases)
}

package main

import (
	"log"
	"strconv"
)

func day03Part1() {
	lines := readLines("d03.input")
	lineLength := len(lines[0])
	significantBits := make([]int, lineLength)
	for _, line := range lines {
		for i, char := range line {
			switch char {
			case '1':
				significantBits[i]++
			case '0':
				significantBits[i]--
			}
		}
	}

	gammaInBinary := make([]rune, lineLength)
	epsilonInBinary := make([]rune, lineLength)
	for i, bit := range significantBits {
		if bit > 0 {
			gammaInBinary[i] = '1'
			epsilonInBinary[i] = '0'
		} else if bit < 0 {
			gammaInBinary[i] = '0'
			epsilonInBinary[i] = '1'
		} else {
			log.Fatalln("Unable to determine significant bit at index:", i)
		}
	}

	gamma, _ := strconv.ParseInt(string(gammaInBinary), 2, 64)
	epsilon, _ := strconv.ParseInt(string(epsilonInBinary), 2, 64)
	println("Gamma:", gamma)
	println("Epsilon:", epsilon)
	println("Result:", gamma*epsilon)
}

func day03Part2() {
	lines := readLines("d03.input")
	ogRating := binaryToDecimal(findMatchingLine(lines, 0, true))
	csRating := binaryToDecimal(findMatchingLine(lines, 0, false))
	println("Oxygen generator rating:", ogRating)
	println("CO2 scrubber rating:", csRating)
	println("Result:", ogRating*csRating)
}

func findMatchingLine(lines []string, currentIndex int, mostCommon bool) string {
	if len(lines) == 1 {
		return lines[0]
	}

	var linesWith0 []string
	var linesWith1 []string

	for _, line := range lines {
		switch line[currentIndex] {
		case '0':
			linesWith0 = append(linesWith0, line)
		case '1':
			linesWith1 = append(linesWith1, line)
		}
	}

	var linesMatchingMostCommon []string
	var linesMatchingLeastCommon []string
	if len(linesWith0) > len(linesWith1) {
		linesMatchingMostCommon = linesWith0
		linesMatchingLeastCommon = linesWith1
	} else {
		linesMatchingMostCommon = linesWith1
		linesMatchingLeastCommon = linesWith0
	}

	if mostCommon {
		return findMatchingLine(linesMatchingMostCommon, currentIndex+1, mostCommon)
	} else {
		return findMatchingLine(linesMatchingLeastCommon, currentIndex+1, mostCommon)
	}
}

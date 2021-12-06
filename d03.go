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

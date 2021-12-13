package main

import (
	"fmt"
	"strings"
)

func day06() {
	//input := "3,4,3,1,2"
	//fishes := stringsToInts(strings.Split(input, ","))
	input := readFile("d06.input")
	input = input[:len(input)-1]
	fishes := stringsToInts(strings.Split(input, ","))

	for day := 1; day <= 80; day++ {
		for i, fish := range fishes {
			if fish == 0 {
				fishes[i] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[i]--
			}
		}
		//fmt.Print("Day ", day, ": ", fishes, "\n")
	}

	fmt.Println("Result:", len(fishes))
}

func sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

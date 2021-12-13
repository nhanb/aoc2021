package main

import (
	"fmt"
	"strings"
)

type FishChunk struct {
	count int
	timer int
}

func day06() {
	//input := "3,4,3,1,2"
	//fishes := stringsToInts(strings.Split(input, ","))
	input := readFile("d06.input")
	input = input[:len(input)-1]
	fishes := stringsToInts(strings.Split(input, ","))

	var chunks []FishChunk
	currentChunk := FishChunk{1, fishes[0]}
	for _, fish := range fishes[1:] {
		if fish == currentChunk.timer {
			currentChunk.count++
		} else {
			chunks = append(chunks, currentChunk)
			currentChunk = FishChunk{1, fish}
		}
	}
	chunks = append(chunks, currentChunk)

	for day := 1; day <= 256; day++ {
		newChunk := FishChunk{0, 8}
		for i, _ := range chunks {
			if chunks[i].timer == 0 {
				chunks[i].timer = 6
				newChunk.count += chunks[i].count
			} else {
				chunks[i].timer--
			}
		}
		if newChunk.count > 0 {
			chunks = append(chunks, newChunk)
		}
		//fmt.Print("Day ", day, ": ", chunks, "\n")
	}

	fishCount := 0
	for _, chunk := range chunks {
		fishCount += chunk.count
	}

	fmt.Println(chunks)
	fmt.Println("Result:", fishCount)
}

func sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

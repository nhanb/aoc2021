package main

import (
	"fmt"
	"strings"
)

func day07() {
	input := readFile("d07.input")
	input = input[:len(input)-1]
	positions := stringsToInts(strings.Split(input, ","))
	fmt.Println(positions)
	fmt.Println("Num positions:", len(positions))

	//mymap := make(map[int]bool)
	//for _, val := range positions {
	//mymap[val] = true
	//}
	//fmt.Println("Uniques:", len(mymap))

	maxX := max(positions)
	lowestFuel := 9223372036854775807
	for x := 0; x <= maxX; x++ {
		fuel := calculateFuelForX(x, positions)
		if fuel < lowestFuel {
			lowestFuel = fuel
		}
	}
	fmt.Println("lowestFuel:", lowestFuel)
}

func calculateFuelForX(x int, positions []int) int {
	fuel := 0
	for _, position := range positions {
		fuel += abs(position - x)
	}
	return fuel
}

func abs(value int) int {
	if value > 0 {
		return value
	} else {
		return -value
	}
}

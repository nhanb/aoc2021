package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day04Part1() {
	lines := readLines("d04.input")
	draws := stringsToInts(strings.Split(lines[0], ","))
	boards := parseBoards(lines[1:])
	var drawn []int

	for _, draw := range draws {
		drawn = append(drawn, draw)
	}

	fmt.Println(draws)
	fmt.Println(boards[0])
	fmt.Println(boards[len(boards)-1])
	fmt.Println("Number of boards:", len(boards))
}

type Board [5][5]int
type BoardPosition struct {
	x int
	y int
}

func parseBoard(lines []string) Board {
	numbersRegex := regexp.MustCompile(`\d+`)
	var board Board
	for y, line := range lines {
		numbers := numbersRegex.FindAllString(line, -1)
		for x, num := range numbers {
			board[y][x], _ = strconv.Atoi(num)
		}
	}
	return board
}

func parseBoards(lines []string) []Board {
	var boards []Board
	for i := 0; i < len(lines); i += 6 {
		boards = append(boards, parseBoard(lines[i+1:i+6]))
	}
	return boards
}

func (board Board) String() string {
	result := ""
	for _, line := range board {
		for _, num := range line {
			result += fmt.Sprintf("%v ", num)
		}
		result += "\n"
	}
	return result
}

func bingo(board Board, drawn []int) bool {
	if len(drawn) < len(board) {
		return false
	}

	latestDraw := drawn[len(drawn)-1]
	positionInBoard := BoardPosition{-1, -1}
	for y := range board {
		for x := range board[0] {
			if board[y][x] == latestDraw {
				positionInBoard.x = x
				positionInBoard.y = y
				break
			}
		}
	}
	if positionInBoard.x == -1 {
		return false
	}

	//println("Position in board:", positionInBoard.x, positionInBoard.y)

	isHorizontalBingo := true
	isVerticalBingo := true
	for i := 0; i < 5; i++ {
		if !intInSlice(board[positionInBoard.y][i], drawn) {
			isHorizontalBingo = false
		}
		if !intInSlice(board[i][positionInBoard.x], drawn) {
			isVerticalBingo = false
		}
	}

	//println("isHorizontalBingo:", isHorizontalBingo)
	//println("isVerticalBingo:", isVerticalBingo)

	return isHorizontalBingo || isVerticalBingo
}

func intInSlice(needle int, haystack []int) bool {
	for _, el := range haystack {
		if needle == el {
			return true
		}
	}
	return false
}

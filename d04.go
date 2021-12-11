package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day04() {
	lines := readLines("d04.input")
	draws := stringsToInts(strings.Split(lines[0], ","))
	boards := parseBoards(lines[1:])
	wonBoardIndexes := make(map[int]bool)
	var drawn []int

	for turn, draw := range draws {
		drawn = append(drawn, draw)
		for boardIndex, board := range boards {
			if bingo(board, drawn) {
				println(
					"Turn", turn,
					"- Bingo on board number", boardIndex+1,
					"- Score:", bingoScore(board, drawn),
				)
				wonBoardIndexes[boardIndex] = true
			}
			if len(wonBoardIndexes) == len(boards) {
				return
			}
		}
	}
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

	return isHorizontalBingo || isVerticalBingo
}

func bingoScore(board Board, drawn []int) int {
	latestDraw := drawn[len(drawn)-1]
	score := 0
	for _, row := range board {
		for _, num := range row {
			if !intInSlice(num, drawn) {
				score += num
			}
		}
	}
	score *= latestDraw
	return score
}

func intInSlice(needle int, haystack []int) bool {
	for _, el := range haystack {
		if needle == el {
			return true
		}
	}
	return false
}

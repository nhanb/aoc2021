package main

import (
	"testing"
)

func TestBingo(t *testing.T) {
	board := Board{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	assert(t, !bingo(board, []int{1, 2}), "Should short circuit")
	assert(t, !bingo(board, []int{1, 2, 3, 5, 6}), "Should fail because missing 4")
	assert(t, bingo(board, []int{1, 2, 3, 5, 6, 4}), "Should pass horizontal")
	assert(t, bingo(board, []int{0, 7, 2, 8, 12, 17, 22}), "Should pass vertical")
}

func assert(t *testing.T, condition bool, message string) {
	if !condition {
		t.Errorf(message)
	} else {
		println(".")
	}
}

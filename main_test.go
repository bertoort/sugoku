package main

import (
	"fmt"
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
	"testing"
)

func Test_boardBasic(t *testing.T) {
	var input [9][9]int
	input[0] = [9]int{1, 0, 3, 0, 0, 6, 0, 8, 0}
	input[1] = [9]int{0, 5, 0, 0, 8, 0, 1, 2, 0}
	input[2] = [9]int{7, 0, 9, 1, 0, 3, 0, 5, 6}
	input[3] = [9]int{0, 3, 0, 0, 6, 7, 0, 9, 0}
	input[4] = [9]int{5, 0, 7, 8, 0, 0, 0, 3, 0}
	input[5] = [9]int{8, 0, 1, 0, 3, 0, 5, 0, 7}
	input[6] = [9]int{0, 4, 0, 0, 7, 8, 0, 1, 0}
	input[7] = [9]int{6, 0, 8, 0, 0, 2, 0, 4, 0}
	input[8] = [9]int{0, 1, 2, 0, 4, 5, 0, 7, 8}
	if board.Basic() != input {
		fmt.Println(board.Basic())
		t.Error("should output a 2D matrix input")
	}
}

func Test_puzzleSetUp(t *testing.T) {
	newBoard := puzzle.New()
	input := board.Basic()
	newBoard.FillPuzzle(input)
	b, _ := newBoard.Display()
	if b != board.Basic() {
		fmt.Println(b)
		t.Error("should equal: ", board.Basic())
	}
}

func Test_solve(t *testing.T) {
	newBoard := puzzle.New()
	input := board.Basic()
	newBoard.FillPuzzle(input)
	newBoard.Solve()
	b, _ := newBoard.Display()
	if b != board.Solved() {
		fmt.Println(b)
		t.Error("should equal solved board: ", board.Solved())
	}
}

func Test_validate(t *testing.T) {
	newBoard := puzzle.New()
	// import standard input
	input := board.Basic()
	newBoard.FillPuzzle(input)
	newBoard.Solve()
	v := newBoard.Validate()
	if !v {
		fmt.Println(v)
		t.Error("should equal true")
	}
	// import unsolvable input
	input = board.Broken()
	newBoard.FillPuzzle(input)
	newBoard.Solve()
	v = newBoard.Validate()
	if v {
		fmt.Println(v)
		t.Error("should equal false")
	}
}

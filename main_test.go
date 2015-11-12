package main

import (
	"fmt"
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
	"testing"
)

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

package main

import (
	"fmt"
	// "github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
)

// Command: go run main.go

func main() {
	v := puzzle.Generate("medium")
	// basicInput := board.Basic()
	sudoku := puzzle.New(v)
	sudoku.Grade()
	sudoku.Solve()
	board, status, dif := sudoku.Display()
	fmt.Println(board, status, dif)
}

package main

import (
	"fmt"
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
)

// Command: go run cli.go

func main() {
	i := board.Basic()
	sudoku := puzzle.New(i)
	sudoku.Grade()
	err := sudoku.SlowSolve()
	board, status, dif := sudoku.Display()
	if !err {
		fmt.Println(board, status, dif)
	}
}

package main

import (
	"fmt"
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
)

// Command: go run main.go

func main() {
	basicInput := board.Medium()
	newBoard := puzzle.New(basicInput)
	newBoard.Solve()
	board, status := newBoard.Display()
	fmt.Println(board, status)
}

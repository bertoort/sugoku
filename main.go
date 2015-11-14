package main

import (
	"fmt"
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
)

func main() {
	newBoard := puzzle.New()
	basicInput := board.Medium()
	newBoard.FillPuzzle(basicInput)
	newBoard.Solve()
	board, status := newBoard.Display()
	fmt.Println(board, status)
}

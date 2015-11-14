package main

import (
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
	"testing"
)

func Benchmark_solve(b *testing.B) {
	newBoard := puzzle.New()
	input := board.Basic()
	newBoard.FillPuzzle(input)
	for n := 0; n < b.N; n++ {
		newBoard.Solve()
	}
}

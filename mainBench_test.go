package main

import (
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
	"testing"
)

// Command: go test -bench=.

func Benchmark_solveBasic(b *testing.B) {
	input := board.Basic()
	newBoard := puzzle.New(input)
	for n := 0; n < b.N; n++ {
		newBoard.Solve()
	}
}

func Benchmark_solveMedium(b *testing.B) {
	input := board.Medium()
	newBoard := puzzle.New(input)
	for n := 0; n < b.N; n++ {
		newBoard.Solve()
	}
}

func Benchmark_solveHard(b *testing.B) {
	input := board.Hard()
	newBoard := puzzle.New(input)
	for n := 0; n < b.N; n++ {
		newBoard.Solve()
	}
}

// ************
// Slow Solving
// ************

func Benchmark_slowSolveBasic(b *testing.B) {
	input := board.Basic()
	newBoard := puzzle.New(input)
	for n := 0; n < b.N; n++ {
		newBoard.SlowSolve()
	}
}

func Benchmark_slowSolveMedium(b *testing.B) {
	input := board.Medium()
	newBoard := puzzle.New(input)
	for n := 0; n < b.N; n++ {
		newBoard.SlowSolve()
	}
}

func Benchmark_slowSolveHard(b *testing.B) {
	input := board.Hard()
	newBoard := puzzle.New(input)
	for n := 0; n < b.N; n++ {
		newBoard.SlowSolve()
	}
}

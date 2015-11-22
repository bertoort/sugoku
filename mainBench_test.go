package main

import (
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/puzzle"
	"testing"
)

// Command: go test -bench=.

func Benchmark_solveBasic(b *testing.B) {
	e := puzzle.New(board.Basic())
	m := puzzle.New(board.Medium())
	h := puzzle.New(board.Hard())
	for n := 0; n < b.N; n++ {
		e.Solve()
		m.Solve()
		h.Solve()
	}
}

// ************
// Slow Solving
// ************

func Benchmark_slowSolveBasic(b *testing.B) {
	e := puzzle.New(board.Basic())
	m := puzzle.New(board.Medium())
	h := puzzle.New(board.Hard())
	for n := 0; n < b.N; n++ {
		e.SlowSolve()
		m.SlowSolve()
		h.SlowSolve()
	}
}

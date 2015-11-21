package puzzle

import (
	"github.com/bertoort/sugoku/board"
	"github.com/bertoort/sugoku/logic"
	"math/rand"
	"time"
)

// Generate will generate a random board by input: easy, medium, hard
func Generate(dif string) [9][9]int {
	board := RandomBoard()
	n := logic.LevelNumber(dif)
	newBoard := RemoveVals(board, n)
	return newBoard
}

// GenRandom will generate a random board
func GenRandom() [9][9]int {
	dif := logic.RandomDif()
	board := RandomBoard()
	n := logic.LevelNumber(dif)
	newBoard := RemoveVals(board, n)
	return newBoard
}

// RemoveVals takes a solved puzzle and sets it to a difficulty
func RemoveVals(board [9][9]int, vals int) [9][9]int {
	for i, row := range board {
		for j := range row {
			t := time.Now()
			rand.Seed(int64(t.Nanosecond()))
			r := rand.Intn(2)
			if board[i][j] != 0 && vals > 0 {
				if r == 0 {
					board[i][j] = 0
					vals--
				}
			}
		}
	}
	if vals > 0 {
		board = RemoveVals(board, vals)
	}
	return board
}

// RandomBoard will generate a random solved board
func RandomBoard() [9][9]int {
	input := board.Empty()
	sudoku := New(input)
	sudoku.RandomFill()
	sudoku.Solve()
	board, _, _ := sudoku.Display()
	return board
}

// RandomFill will fill the the first row randomly
func (p *Puzzle) RandomFill() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	logic.Shuffle(list)
	for i := range p.Board[0] {
		p.Board[0][i].val = list[i]
	}
}

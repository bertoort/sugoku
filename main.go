package main

import (
	"fmt"
	// "github.com/bertoort/sugoku/puzzle"
)

func main() {
	newBoard := Puzzle{}
	newBoard.createPuzzle()
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
	newBoard.fillPuzzle(input)
	board := newBoard.display()
	fmt.Println(board)
	newBoard.solve()
	board = newBoard.display()
	fmt.Println(board)
}

// Puzzle class
type Puzzle struct {
	board [9][9]square
}

type square struct {
	val int
	x   int
	y   int
	b   int
	not []int
	row []int
	col []int
	box []int
}

func (p *Puzzle) display() [9][9]int {
	var board [9][9]int
	for i, row := range p.board {
		var newRow [9]int
		for j := range row {
			newRow[j] = p.board[i][j].val
		}
		board[i] = newRow
	}
	return board
}

func (p *Puzzle) solve() {
	var delta bool
	for i, row := range p.board {
		for j := range row {
			if p.board[i][j].val == 0 {
				p.board[i][j].check(p)
				delta = p.board[i][j].evaluate()
			}
		}
	}
	solved := p.solved()
	if !solved && delta {
		p.solve()
	}
}

func (p *Puzzle) solved() bool {
	s := true
	for i, row := range p.board {
		for j := range row {
			if p.board[i][j].val == 0 {
				s = false
			}
		}
	}
	return s
}

func (s *square) check(p *Puzzle) {
	for i, row := range p.board {
		for j := range row {
			if p.board[i][j].val != 0 {
				if s.x != p.board[i][j].x || s.y != p.board[i][j].y {
					s.compare(p.board[i][j])
				}
			}
		}
	}
}

func (s *square) evaluate() bool {
	if len(s.not) == 8 {
		s.val = addValue(s.not)
		return true
	} else if len(s.row) == 8 {
		s.val = addValue(s.row)
		return true
	} else if len(s.col) == 8 {
		s.val = addValue(s.col)
		return true
	} else if len(s.box) == 8 {
		s.val = addValue(s.box)
		return true
	}
	return false
}

func addValue(a []int) int {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var n int
	for _, v := range list {
		if !numInSlice(v, a) {
			n = v
		}
	}
	return n
}

func numInSlice(num int, list []int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

func (s *square) compare(c square) {
	if s.x == c.x {
		appendVal(s, c.val, "x")
	}
	if s.y == c.y {
		appendVal(s, c.val, "y")
	}
	if s.b == c.b {
		appendVal(s, c.val, "b")
	}
}

func appendVal(s *square, n int, t string) {
	// is there a way to index structs? (select with bracket notation)
	switch t {
	case "x":
		if !duplicate(s.row, n) {
			s.row = append(s.row, n)
		}
	case "y":
		if !duplicate(s.col, n) {
			s.col = append(s.col, n)
		}
	case "b":
		if !duplicate(s.box, n) {
			s.box = append(s.box, n)
		}
	}
	if !duplicate(s.not, n) {
		s.not = append(s.not, n)
	}
}

func duplicate(not []int, n int) bool {
	var d bool
	for _, num := range not {
		if num == n {
			d = true
		}
	}
	return d
}

func (p *Puzzle) fillPuzzle(input [9][9]int) {
	for i, row := range p.board {
		for j := range row {
			p.board[i][j].val = input[i][j]
		}
	}
}

func (p *Puzzle) createPuzzle() [9][9]square {
	var puzzle [9][9]square
	for i, row := range puzzle {
		for j := range row {
			puzzle[i][j].x = j
			puzzle[i][j].y = i
			if puzzle[i][j].x < 3 {
				if puzzle[i][j].y < 3 {
					puzzle[i][j].b = 0
				} else if puzzle[i][j].y > 5 {
					puzzle[i][j].b = 6
				} else {
					puzzle[i][j].b = 3
				}
			} else if puzzle[i][j].x > 5 {
				if puzzle[i][j].y < 3 {
					puzzle[i][j].b = 2
				} else if puzzle[i][j].y > 5 {
					puzzle[i][j].b = 8
				} else {
					puzzle[i][j].b = 5
				}
			} else {
				if puzzle[i][j].y < 3 {
					puzzle[i][j].b = 1
				} else if puzzle[i][j].y > 5 {
					puzzle[i][j].b = 7
				} else {
					puzzle[i][j].b = 4
				}
			}
		}
	}
	p.board = puzzle
	return puzzle
}

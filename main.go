package main

import (
	"fmt"
)

// Puzzle class
type Puzzle struct {
	board [9][9]square
}

type square struct {
	val int
	x   int
	y   int
	box int
	not []int
}

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
	newBoard.solve()
	board := newBoard.display()
	fmt.Println(board)
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
	for i, row := range p.board {
		for j := range row {
			p.board[i][j].check(p)
		}
	}
}

func (s *square) check(p *Puzzle) {
	for i, row := range p.board {
		for j := range row {
			if s.val == 0 && p.board[i][j].val != 0 {
				if s.x != p.board[i][j].x || s.y != p.board[i][j].y {
					s.compare(p.board[i][j])
				}
			}
			if len(s.not) == 8 {
				list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
				var n int
				for _, v := range list {
					if !numInSlice(v, s.not) {
						n = v
					}
				}
				s.val = n
			}
		}
	}
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
	if s.x == c.x || s.y == c.y {
		s.not = append(s.not, c.val)
	} else if s.box == c.box {
		s.not = append(s.not, c.val)
	}
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
					puzzle[i][j].box = 0
				} else if puzzle[i][j].y > 5 {
					puzzle[i][j].box = 6
				} else {
					puzzle[i][j].box = 3
				}
			} else if puzzle[i][j].x > 5 {
				if puzzle[i][j].y < 3 {
					puzzle[i][j].box = 2
				} else if puzzle[i][j].y > 5 {
					puzzle[i][j].box = 8
				} else {
					puzzle[i][j].box = 5
				}
			} else {
				if puzzle[i][j].y < 3 {
					puzzle[i][j].box = 1
				} else if puzzle[i][j].y > 5 {
					puzzle[i][j].box = 7
				} else {
					puzzle[i][j].box = 4
				}
			}
		}
	}
	p.board = puzzle
	return puzzle
}

package main

import (
	"fmt"
	// "github.com/bertoort/sugoku/puzzle"
)

func main() {
	newBoard := Puzzle{}
	newBoard.status = "unsolved"
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
	board, status := newBoard.display()
	fmt.Println(board, status)
	newBoard.solve()
	board, status = newBoard.display()
	fmt.Println(board, status, v)
}

// Puzzle class
type Puzzle struct {
	status string
	board  [9][9]square
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

func (p *Puzzle) display() ([9][9]int, string) {
	var board [9][9]int
	for i, row := range p.board {
		var newRow [9]int
		for j := range row {
			newRow[j] = p.board[i][j].val
		}
		board[i] = newRow
	}
	return board, p.status
}

func (p *Puzzle) solve() {
	deltas := 0
	for i, row := range p.board {
		for j := range row {
			if p.board[i][j].val == 0 {
				p.board[i][j].check(p)
				deltas += p.board[i][j].evaluate()
			}
		}
	}
	solved := p.solved()
	if !solved && deltas > 0 {
		p.solve()
	} else if !solved {
		fmt.Println("guessing")
		result, err := p.guess(0)
		if err {
			p.status = "unsolvable"
		} else {
			p.board = result.board
			p.status = "solved"
		}
	} else {
		p.status = "solved"
	}
}

func (p Puzzle) validate() bool {
	for i, row := range p.board {
		for j := range row {
			if p.board[i][j].val != 0 {
				r := p.board[i][j].checkUniqueness(p)
				if !r {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

func (s *square) checkUniqueness(p Puzzle) bool {
	for i, row := range p.board {
		for j := range row {
			if s.x != p.board[i][j].x || s.y != p.board[i][j].y {
				if s.x == p.board[i][j].x && s.val == p.board[i][j].val {
					return false
				} else if s.y == p.board[i][j].y && s.val == p.board[i][j].val {
					return false
				} else if s.b == p.board[i][j].b && s.val == p.board[i][j].val {
					return false
				}
			}
		}
	}
	return true
}

func (p *Puzzle) guess(n int) (Puzzle, bool) {
	values, _ := p.display()
	mirror := Puzzle{}
	mirror.status = "unsolved"
	mirror.createPuzzle()
	mirror.fillPuzzle(values)
	deltas := 0
	first := true
	for i, row := range mirror.board {
		for j := range row {
			if mirror.board[i][j].val == 0 {
				mirror.board[i][j].check(&mirror)
				deltas += mirror.board[i][j].evaluate()
				if first {
					mirror.board[i][j].assignVal(n)
					first = false
				}
			}
		}
	}
	solved := mirror.solved()
	if !solved && deltas > 0 {
		fmt.Println("inside guessing solving")
		mirror.solve()
	} else if !solved {
		mirror.status = "unsolvable"
		return mirror, true
	}
	return mirror, false
}

func (s *square) assignVal(n int) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	possibilities := []int{}
	for i := 0; i < 9; i++ {
		var u bool
		for j := 0; j < len(s.not); j++ {
			if s.not[j] == list[i] {
				u = true
			}
		}
		if !u {
			possibilities = append(possibilities, list[i])
		}
	}
	fmt.Println(possibilities)
	s.val = possibilities[n]
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

func (s *square) evaluate() int {
	if len(s.not) == 8 {
		s.val = addValue(s.not)
		return 1
	} else if len(s.row) == 8 {
		s.val = addValue(s.row)
		return 1
	} else if len(s.col) == 8 {
		s.val = addValue(s.col)
		return 1
	} else if len(s.box) == 8 {
		s.val = addValue(s.box)
		return 1
	}
	return 0
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

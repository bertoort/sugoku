package puzzle

import (
	"github.com/bertoort/sugoku/logic"
)

// square are the individual blocks in a sudoku board
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

// **************
// square methods
// **************

// CheckUniqueness is used while validating to make sure the value
// isn't repeated in the same row, col or box
func (s *square) CheckUniqueness(p Puzzle) bool {
	for i, row := range p.Board {
		for j := range row {
			if s.x != p.Board[i][j].x || s.y != p.Board[i][j].y {
				if s.x == p.Board[i][j].x && s.val == p.Board[i][j].val {
					return false
				} else if s.y == p.Board[i][j].y && s.val == p.Board[i][j].val {
					return false
				} else if s.b == p.Board[i][j].b && s.val == p.Board[i][j].val {
					return false
				}
			}
		}
	}
	return true
}

// AssignVal finds and assigns a possible value to an empty square
func (s *square) AssignVal(n int) bool {
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
	if len(possibilities) > n {
		s.val = possibilities[n]
		return false
	}
	return true
}

// Check compares the row, col and box of a square and fills the 'not' values
func (s *square) Check(p *Puzzle) {
	for i, row := range p.Board {
		for j := range row {
			if p.Board[i][j].val != 0 {
				if s.x != p.Board[i][j].x || s.y != p.Board[i][j].y {
					s.Compare(p.Board[i][j])
				}
			}
		}
	}
}

// Evaluate checks if a square is the remaining value of a row, col or box
func (s *square) Evaluate() int {
	if len(s.not) == 8 {
		s.val = logic.AddValue(s.not)
		return 1
	} else if len(s.row) == 8 {
		s.val = logic.AddValue(s.row)
		return 1
	} else if len(s.col) == 8 {
		s.val = logic.AddValue(s.col)
		return 1
	} else if len(s.box) == 8 {
		s.val = logic.AddValue(s.box)
		return 1
	}
	return 0
}

// Compare checks the position of the square to other values
// and adds them to the 'not' lists
func (s *square) Compare(c square) {
	if s.x == c.x {
		AppendVal(s, c.val, "x")
	}
	if s.y == c.y {
		AppendVal(s, c.val, "y")
	}
	if s.b == c.b {
		AppendVal(s, c.val, "b")
	}
}

// AppendVal will include the value to the 'not' lists if it's not a duplicate
func AppendVal(s *square, n int, t string) {
	switch t {
	case "x":
		if !logic.Duplicate(s.row, n) {
			s.row = append(s.row, n)
		}
	case "y":
		if !logic.Duplicate(s.col, n) {
			s.col = append(s.col, n)
		}
	case "b":
		if !logic.Duplicate(s.box, n) {
			s.box = append(s.box, n)
		}
	}
	if !logic.Duplicate(s.not, n) {
		s.not = append(s.not, n)
	}
}

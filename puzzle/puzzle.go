package puzzle

import (
	"github.com/bertoort/sugoku/logic"
)

// Puzzle class
type Puzzle struct {
	Status string
	Board  [9][9]square
}

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

//New puzzle with methods for exporting
func New() Puzzle {
	new := Puzzle{
		Status: "unsolved",
	}
	new.CreatePuzzle()
	return new
}

// **************
// puzzle methods
// **************

// Display returns the board values in a 2D matrix and the sudoku status
func (p *Puzzle) Display() ([9][9]int, string) {
	var board [9][9]int
	for i, row := range p.Board {
		var newRow [9]int
		for j := range row {
			newRow[j] = p.Board[i][j].val
		}
		board[i] = newRow
	}
	return board, p.Status
}

// Solve is the main method that solved the sudoku puzzle
func (p *Puzzle) Solve() {
	deltas := 0
	for i, row := range p.Board {
		for j := range row {
			if p.Board[i][j].val == 0 {
				p.Board[i][j].Check(p)
				deltas += p.Board[i][j].Evaluate()
			}
		}
	}
	solved := p.Solved()
	if !solved && deltas > 0 {
		p.Solve()
	} else if !solved {
		result, err := p.Guess(0)
		if err {
			p.Status = "unsolvable"
		} else {
			p.Board = result.Board
			if p.Validate() {
				p.Status = "solved"
			} else {
				p.Status = "broken"
			}
		}
	} else {
		if p.Validate() {
			p.Status = "solved"
		} else {
			p.Status = "broken"
		}
	}
}

// Validate checks if the sudoku has a valid solution
func (p Puzzle) Validate() bool {
	for i, row := range p.Board {
		for j := range row {
			if p.Board[i][j].val != 0 {
				r := p.Board[i][j].CheckUniqueness(p)
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

// Guess finds the first empty value, adds a possible value,
// and recursively tries to solve the puzzle
func (p *Puzzle) Guess(n int) (Puzzle, bool) {
	values, _ := p.Display()
	mirror := Puzzle{}
	mirror.Status = "unsolved"
	mirror.CreatePuzzle()
	mirror.FillPuzzle(values)
	deltas := 0
	first := true
	for i, row := range mirror.Board {
		for j := range row {
			if mirror.Board[i][j].val == 0 {
				mirror.Board[i][j].Check(&mirror)
				deltas += mirror.Board[i][j].Evaluate()
				if first {
					mirror.Board[i][j].AssignVal(n)
					first = false
				}
			}
		}
	}
	solved := mirror.Solved()
	if !solved && deltas > 0 {
		mirror.Solve()
	} else if !solved {
		mirror.Status = "unsolvable"
		return mirror, true
	}
	return mirror, false
}

// Solved quickly checks that there aren't any empty values
func (p *Puzzle) Solved() bool {
	s := true
	for i, row := range p.Board {
		for j := range row {
			if p.Board[i][j].val == 0 {
				s = false
			}
		}
	}
	return s
}

// FillPuzzle takes a 2D matrix and fills the puzzle board with it
func (p *Puzzle) FillPuzzle(input [9][9]int) {
	for i, row := range p.Board {
		for j := range row {
			p.Board[i][j].val = input[i][j]
		}
	}
}

// CreatePuzzle adds an empty sudoku board to a puzzle class
func (p *Puzzle) CreatePuzzle() {
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
	p.Board = puzzle
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
func (s *square) AssignVal(n int) {
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
	s.val = possibilities[n]
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

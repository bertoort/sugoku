package puzzle

// Puzzle class
type Puzzle struct {
	Status string
	Board  [9][9]square
}

// **************
// puzzle methods
// **************

//New puzzle with methods for exporting
func New(input [9][9]int) Puzzle {
	new := Puzzle{
		Status: "unsolved",
	}
	new.CreatePuzzle(input)
	return new
}

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

// FindValues automatically fills a square with the correct possible values
func (p *Puzzle) FindValues() (bool, bool) {
	c, broken := 0, false
	for i, row := range p.Board {
		for j := range row {
			if p.Board[i][j].val == 0 {
				p.Board[i][j].Check(p)
				if len(p.Board[i][j].not) >= 9 {
					return false, true
				}
				c += p.Board[i][j].Evaluate()
			}
		}
	}
	solved := p.Solved()
	if !solved && c > 0 {
		solved, broken = p.FindValues()
	}
	return solved, broken
}

// Solve is the main method that solved the sudoku puzzle
func (p *Puzzle) Solve() {
	solved, _ := p.FindValues()
	if !solved {
		result, err := p.Guess(0)
		if err {
			p.Status = "unsolvable"
		} else {
			p.Board = result.Board
			p.Status = "solved"
		}
	} else if p.Validate() {
		p.Status = "solved"
	} else {
		p.Status = "broken"
	}
}

// Guess finds the first empty value, adds a possible value,
// and recursively tries to solve the puzzle
func (p *Puzzle) Guess(t int) (Puzzle, bool) {
	max, first, err := true, true, false
	values, _ := p.Display()
	mirror := New(values)
	for i, row := range mirror.Board {
		for j := range row {
			if mirror.Board[i][j].val == 0 && first {
				mirror.Board[i][j].Check(&mirror)
				max = mirror.Board[i][j].AssignVal(t)
				first = false
			}
		}
	}
	solved, err := mirror.FindValues()
	if max {
		return mirror, true
	} else if err {
		mirror, err = p.Guess(t + 1)
		if err {
			return mirror, true
		}
	} else if !solved {
		mirror, err = mirror.Guess(0)
		if err {
			mirror, err = p.Guess(t + 1)
			if err {
				return mirror, true
			}
		}
	} else if !mirror.Validate() {
		mirror, err = p.Guess(t + 1)
		if err {
			return mirror, true
		}
	}
	return mirror, false
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
	p.Status = "unsolved"
	for i, row := range p.Board {
		for j := range row {
			p.Board[i][j].val = input[i][j]
		}
	}
}

// CreatePuzzle adds a sudoku board to a puzzle class
func (p *Puzzle) CreatePuzzle(input [9][9]int) {
	var puzzle [9][9]square
	for i, row := range puzzle {
		for j := range row {
			puzzle[i][j].val = input[i][j]
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

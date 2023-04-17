package puzzle

// Puzzle class
type Puzzle struct {
	Status     string
	Difficulty string
	Board      [9][9]square
}

// **************
// puzzle methods
// **************

//New puzzle with methods for exporting
func New(input [9][9]int) Puzzle {
	new := Puzzle{
		Status:     "unsolved",
		Difficulty: "unknown",
	}
	new.CreatePuzzle(input)
	return new
}

// Display returns the board values in a 2D matrix and the sudoku status
func (p *Puzzle) Display() ([9][9]int, string, string) {
	var board [9][9]int
	for i, row := range p.Board {
		var newRow [9]int
		for j := range row {
			newRow[j] = p.Board[i][j].val
		}
		board[i] = newRow
	}
	return board, p.Status, p.Difficulty
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

// Grade assigns a difficulty to a puzzle
// hard: <27, medium: 28-32, easy: 33>
func (p *Puzzle) Grade() {
	z := 0
	for i, row := range p.Board {
		for j := range row {
			if p.Board[i][j].val != 0 {
				z++
			}
		}
	}
	if z <= 27 {
		p.Difficulty = "hard"
	} else if z >= 33 {
		p.Difficulty = "easy"
	} else {
		p.Difficulty = "medium"
	}
}

// Validate checks if the sudoku has a valid solution
func (p Puzzle) Validate() bool {
	for i, row := range p.Board {
		for j := range row {
			if p.Board[i][j].val != 0 && p.Board[i][j].val <= 9 {
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

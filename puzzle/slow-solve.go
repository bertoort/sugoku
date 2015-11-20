package puzzle

// *****************
// puzzle slow solve
// *****************

// SlowSolve initiates guessing sudoku puzzle
func (p *Puzzle) SlowSolve() {
	result, err := p.SlowGuess(0)
	if err {
		p.Status = "unsolvable"
	} else {
		p.Board = result.Board
		p.Status = "solved"
	}
}

// SlowGuess finds the first empty value, adds a possible value,
// and recursively tries to solve the puzzle
func (p *Puzzle) SlowGuess(t int) (Puzzle, bool) {
	max, first, err := true, true, false
	values, _, _ := p.Display()
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
	solved := mirror.Solved()
	if max {
		return mirror, true
	} else if !solved {
		mirror, err = mirror.SlowGuess(0)
		if err {
			mirror, err = p.Guess(t + 1)
			if err {
				return mirror, true
			}
		}
	} else if !mirror.Validate() {
		mirror, err = p.SlowGuess(t + 1)
		if err {
			return mirror, true
		}
	}
	return mirror, false
}

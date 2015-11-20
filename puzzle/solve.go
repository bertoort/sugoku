package puzzle

// **************
// puzzle solving
// **************

// Solve is the main method that solved the sudoku puzzle
func (p *Puzzle) Solve() {
	solved, _ := p.quickFill()
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

// quickFill automatically fills a square with the correct possible values
func (p *Puzzle) quickFill() (bool, bool) {
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
		solved, broken = p.quickFill()
	}
	return solved, broken
}

// Guess finds the first empty value, adds a possible value,
// and recursively tries to solve the puzzle
func (p *Puzzle) Guess(t int) (Puzzle, bool) {
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
	solved, err := mirror.quickFill()
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

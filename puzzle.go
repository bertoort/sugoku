package puzzle

type square struct {
	val int
	x   int
	y   int
	box int
}

func puzzle() [9][9]square {
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
	return puzzle
}

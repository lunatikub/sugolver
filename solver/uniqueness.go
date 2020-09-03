package solver

func (c *cell) findCandidate(val int) bool {
	for v := range c.candidates {
		if v == val {
			return true
		}
	}
	return false
}

func (s *Solver) uniquenessCol(col int, yCell int, val int) bool {
	for y := 0; y < NrLine; y++ {
		if y != yCell && s.grid[y][col].findCandidate(val) {
			return false
		}
	}
	return true
}

func (s *Solver) uniquenessLine(line int, xCell int, val int) bool {
	for x := 0; x < NrCol; x++ {
		if x != xCell && s.grid[line][x].findCandidate(val) {
			return false
		}
	}
	return true
}

func (s *Solver) uniquenessBlock(line int, col int, val int) bool {
	yBlock := line - line%nrBlockLine
	xBlock := col - col%nrBlockCol

	for y := yBlock; y < yBlock+nrBlockCol; y++ {
		for x := xBlock; x < xBlock+nrBlockLine; x++ {
			if (y != line || x != col) && s.grid[y][x].findCandidate(val) {
				return false
			}
		}
	}
	return true
}

// if a cell C can contain several values, but one of these
// values ​​V is not possible in any other cell of its line, columns or
// block, then cell C contains the value V
func (s *Solver) uniqueness() {
	for y, line := range s.grid {
		for x, cell := range line {
			if cell.typeCell == emptyCell {
				for v := range cell.candidates {
					if s.uniquenessLine(y, x, v) ||
						s.uniquenessCol(x, y, v) ||
						s.uniquenessBlock(y, x, v) {
						s.set(y, x, v, uniquenessCell)
						s.updateCandidates(y, x, v)
						s.nrUniqueness++
					}
				}
			}
		}
	}
}

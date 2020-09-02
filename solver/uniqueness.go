package solver

func (c *cell) findCandidate(val int) bool {
	for v := range c.candidates {
		if v == val {
			return true
		}
	}
	return false
}

func (b *Board) uniquenessCol(col int, yCell int, val int) bool {
	for y := 0; y < nrCol; y++ {
		if y != yCell && b.cells[y][col].findCandidate(val) {
			return false
		}
	}
	return true
}

func (b *Board) uniquenessLine(line int, xCell int, val int) bool {
	for x := 0; x < nrCol; x++ {
		if x != xCell && b.cells[line][x].findCandidate(val) {
			return false
		}
	}
	return true
}

func (b *Board) uniquenessBlock(line int, col int, val int) bool {
	yBlock := line - line%nrBlockLine
	xBlock := col - col%nrBlockCol

	for y := yBlock; y < yBlock+nrBlockCol; y++ {
		for x := xBlock; x < xBlock+nrBlockLine; x++ {
			if (y != line || x != col) && b.cells[y][x].findCandidate(val) {
				return false
			}
		}
	}
	return true
}

// Uniqueness if a cell C can contain several values, but one of these
// values ​​V is not possible in any other cell of its line, columns or
// block, then cell C contains the value V
func (b *Board) Uniqueness() {
	for y, line := range b.cells {
		for x, cell := range line {
			if cell.typeCell == emptyCell {
				for v := range cell.candidates {
					if b.uniquenessLine(y, x, v) ||
						b.uniquenessCol(x, y, v) ||
						b.uniquenessBlock(y, x, v) {
						b.set(y, x, v, uniquenessCell)
						b.updateCandidates(y, x, v)
						b.nrUniqueness++
					}
				}
			}
		}
	}
}

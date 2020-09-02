package solver

// Backtracking Explore all valid possibilites for each
// candidate of each cell to find a solution
func (b *Board) Backtracking() {
	for y, line := range b.cells {
		for x, cell := range line {
			if cell.val == 0 {
				for v := range cell.candidates {
					if b.isValidSet(y, x, v) {
						b.set(y, x, v, candidateCell)
						b.Backtracking()
						b.reset(y, x, v)
					}
				}
				return
			}
		}
	}
	b.DumpBoard()
}

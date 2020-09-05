package solver

// Explore all valid possibilites for each
// candidate of each cell to find a solution
func (s *Solver) backtracking() bool {
	for y, line := range s.grid {
		for x, cell := range line {
			if cell.val == 0 {
				for v := range cell.candidates {
					if s.isValidSet(y, x, v) {
						s.set(y, x, v, candidateCell)
						s.nrBacktracking++
						if s.backtracking() {
							return true // a solution has been found
						}
						s.reset(y, x, v)
						s.nrBacktracking--
					}
				}
				return false // backtrack
			}
		}
	}
	return true
}

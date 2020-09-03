package solver

// Backtracking Explore all valid possibilites for each
// candidate of each cell to find a solution
func (s *Solver) Backtracking() {
	for y, line := range s.grid {
		for x, cell := range line {
			if cell.val == 0 {
				for v := range cell.candidates {
					if s.isValidSet(y, x, v) {
						s.set(y, x, v, candidateCell)
						s.Backtracking()
						s.reset(y, x, v)
					}
				}
				return
			}
		}
	}
	s.DumpGrid()
}

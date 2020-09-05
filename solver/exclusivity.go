package solver

func (s *Solver) pushExclusivity(c *cell, y int, x int) {
	if len(c.candidates) == 1 {
		s.excluStack = append(s.excluStack, coord{y, x})
	}
}

func (s *Solver) popExclusivity() coord {
	n := len(s.excluStack) - 1
	c := s.excluStack[n]
	s.excluStack = s.excluStack[:n]
	return c
}

// If we have found the value V of a cell C then this
// value is removed from all cases in the same block as C,
// in other words we update the lines, the cols and the block boolean vectors.
func (s *Solver) exclusivity() {
	if len(s.excluStack) == 0 {
		return
	}
	for {
		coord := s.popExclusivity()
		cell := s.grid[coord.y][coord.x]
		for v := range cell.candidates {
			s.set(coord.y, coord.x, v, exclusivityCell)
			s.updateCandidates(coord.y, coord.x, v)
			s.nrExclusivity++
		}
		if len(s.excluStack) == 0 {
			break
		}
	}
}

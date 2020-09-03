package solver

func (s *Solver) pushExclusivity(c *cell, y int, x int) {
	if len(c.candidates) == 1 {
		s.exclusivity = append(s.exclusivity, coord{y, x})
	}
}

func (s *Solver) popExclusivity() coord {
	n := len(s.exclusivity) - 1
	c := s.exclusivity[n]
	s.exclusivity = s.exclusivity[:n]
	return c
}

// Exclusivity if we have found the value V of a cell C then this
// value is removed from all cases in the same block as C,
// in other words we update the lines, the cols and the block boolean vectors.
func (s *Solver) Exclusivity() {
	for {
		coord := s.popExclusivity()
		cell := s.grid[coord.y][coord.x]
		for v := range cell.candidates {
			s.set(coord.y, coord.x, v, exclusivityCell)
			s.updateCandidates(coord.y, coord.x, v)
			s.nrExclusivity++
		}
		if len(s.exclusivity) == 0 {
			break
		}
	}
}

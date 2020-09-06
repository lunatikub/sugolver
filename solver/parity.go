package solver

import (
	"reflect"
)

func (s *Solver) updateCol(c *cell, xCell int, y1 int, y2 int) {
	for y := 0; y < NrLine; y++ {
		if y != y1 && y != y2 {
			for v := range c.candidates {
				s.updateCandidate(y, xCell, v)
			}
		}
	}
}

func (s *Solver) parityCol(yCell int, xCell int) {
	cell := s.grid[yCell][xCell]
	for y := 0; y < NrLine; y++ {
		c := s.grid[y][xCell]
		if y != yCell && len(c.candidates) == 2 {
			if reflect.DeepEqual(cell.candidates, c.candidates) {
				s.updateCol(&c, xCell, yCell, y)
			}
		}
	}
}

func (s *Solver) updateLine(c *cell, yCell int, x1 int, x2 int) {
	for x := 0; x < NrCol; x++ {
		if x != x1 && x != x2 {
			for v := range c.candidates {
				s.updateCandidate(yCell, x, v)
			}
		}
	}
}

func (s *Solver) parityLine(yCell int, xCell int) {
	cell := s.grid[yCell][xCell]
	for x := 0; x < NrCol; x++ {
		c := s.grid[yCell][x]
		if x != xCell && len(c.candidates) == 2 {
			if reflect.DeepEqual(cell.candidates, c.candidates) {
				s.updateLine(&c, yCell, xCell, x)
			}
		}
	}
}

func (s *Solver) updateBlock(c *cell, yBlock int, xBlock int, y1 int, x1 int, y2 int, x2 int) {
	for y := yBlock; y < yBlock+nrBlockLine; y++ {
		for x := xBlock; x < xBlock+nrBlockCol; x++ {
			if (x != x1 || y != y1) && (x != y2 || y != y2) {
				for v := range c.candidates {
					s.updateCandidate(y, x, v)
				}
			}
		}
	}
}

func (s *Solver) parityBlock(yCell int, xCell int) {
	yBlock := yCell - yCell%nrBlockLine
	xBlock := xCell - xCell%nrBlockCol
	cell := s.grid[yCell][xCell]

	for y := yBlock; y < yBlock+nrBlockCol; y++ {
		for x := xBlock; x < xBlock+nrBlockLine; x++ {
			c := s.grid[y][x]
			if (x != xCell || y != yCell) && len(c.candidates) == 2 {
				if reflect.DeepEqual(cell.candidates, c.candidates) {
					s.updateBlock(&c, yBlock, xBlock, yCell, xCell, y, x)
				}
			}
		}
	}
}

// If a couple of cells (C,C') on the same line or on the
// same column or in the same block can only contain the same pair
// of values {U,V} then we remove these two values of the candidates
// from lines, cols and blocks of C and C' cells candidates.
func (s *Solver) parity() {
	for y, line := range s.grid {
		for x, cell := range line {
			if len(cell.candidates) == 2 {
				s.parityCol(y, x)
				s.parityLine(y, x)
				s.parityBlock(y, x)
			}
		}
	}
}

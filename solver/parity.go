package solver

import (
	"fmt"
	"reflect"
)

func (s *Solver) parityCol(yCell int, xCell int) {
	cell := s.grid[yCell][xCell]
	for y := 0; y < NrLine; y++ {
		c := s.grid[y][xCell]
		if y != yCell && len(c.candidates) == 2 {
			if reflect.DeepEqual(cell.candidates, c.candidates) {
				fmt.Println("COL (Y: ", yCell, ", X: ", xCell, ")<->(Y:", y, ", X: ", xCell, ")")
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
				fmt.Println("LINE (Y: ", yCell, ", X: ", xCell, ")<->(Y:", yCell, ", X: ", x, ")")
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
					fmt.Println("BLOCK (Y: ", yCell, ", X: ", xCell, ")<->(Y:", y, ", X: ", x, ")")
				}
			}
		}
	}
}

// Parity if a couple of cells (C,C') on the same line or on the
// same column or in the same block can only contain the same pair
// of values {U,V} then we remove these two values of the candidates
// from lines, cols and blocks of C and C' cells candidates.
func (s *Solver) Parity() {
	for y, line := range s.grid {
		for x, cell := range line {
			if len(cell.candidates) == 2 {
				fmt.Println("found couple: (Y:", y, ", X:", x, ")")
				s.parityCol(y, x)
				s.parityLine(y, x)
				s.parityBlock(y, x)
			}
		}
	}
}

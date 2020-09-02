package solver

func (b *Board) pushExclusivity(c *cell, y int, x int) {
	if len(c.candidates) == 1 {
		b.exclusivity = append(b.exclusivity, coord{y, x})
	}
}

func (b *Board) popExclusivity() coord {
	n := len(b.exclusivity) - 1
	c := b.exclusivity[n]
	b.exclusivity = b.exclusivity[:n]
	return c
}

// Exclusivity if we have found the value V of a cell C then this
// value is removed from all cases in the same block as C,
// in other words we update the lines, the cols and the block boolean vectors.
func (b *Board) Exclusivity() {
	for {
		coord := b.popExclusivity()
		cell := b.cells[coord.y][coord.x]
		for v := range cell.candidates {
			b.set(coord.y, coord.x, v, exclusivityCell)
			b.updateCandidates(coord.y, coord.x, v)
			b.nrExclusivity++
		}
		if len(b.exclusivity) == 0 {
			break
		}
	}
}

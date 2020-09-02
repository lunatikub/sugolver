package solver

// Sudoku constants
const (
	nrLine      = 9
	nrCol       = 9
	nrBlock     = 9
	nrVal       = 9
	nrBlockLine = 3
	nrBlockCol  = 3
)

const (
	emptyCell = iota
	initialCell
	candidateCell
	exclusivityCell
	uniquenessCell
)

type coord struct {
	y int
	x int
}

type cell struct {
	val        int
	typeCell   int
	candidates map[int]struct{}
}

// Board playing 9x9 sudoku
type Board struct {
	cells       [nrCol][nrLine]cell
	lines       [nrLine][nrVal]bool
	cols        [nrCol][nrVal]bool
	blocks      [nrBlock][nrVal]bool
	exclusivity []coord
	nrCandidate uint
	nrEmpty     uint
	// statistics
	nrInitVal       uint
	nrInitCandidate uint
	nrExclusivity   uint
	nrUniqueness    uint
}

func getBlockID(y int, x int) int {
	return x/nrBlockCol + (y/nrBlockLine)*nrBlockCol
}

func (b *Board) set(y int, x int, v int, t int) {
	b.cells[y][x].val = v
	b.cells[y][x].typeCell = t
	b.lines[y][v-1] = true
	b.cols[x][v-1] = true
	b.blocks[getBlockID(y, x)][v-1] = true
	b.nrEmpty--
}

func (b *Board) reset(y int, x int, v int) {
	b.cells[y][x].val = 0
	b.cells[y][x].typeCell = emptyCell
	b.lines[y][v-1] = false
	b.cols[x][v-1] = false
	b.blocks[getBlockID(y, x)][v-1] = false
	b.nrEmpty++
}

func (b *Board) isValidSet(y int, x int, v int) bool {
	return !b.lines[y][v-1] && !b.cols[x][v-1] && !b.blocks[getBlockID(y, x)][v-1]
}

func (c *cell) setCandidate(b *Board, y int, x int) {
	for v := 0; v < nrVal; v++ {
		if !b.lines[y][v] && !b.cols[x][v] && !b.blocks[getBlockID(y, x)][v] {
			c.candidates[v+1] = struct{}{}
			b.nrCandidate++
		}
	}
}

// SetCandidates find the potential
// candidates for each cell of the board
func (b *Board) setCandidates() {
	for y, line := range b.cells {
		for x, cell := range line {
			if b.cells[y][x].typeCell == emptyCell {
				cell.setCandidate(b, y, x)
				b.pushExclusivity(&cell, y, x)
			}
		}
	}
	b.nrInitCandidate = b.nrCandidate
}

// Init the board
func (b *Board) init(initialValues *[9][9]int) {
	b.nrEmpty = nrCol * nrLine
	for y, line := range initialValues {
		for x, v := range line {
			b.cells[y][x].candidates = make(map[int]struct{})
			if v != 0 {
				b.set(y, x, v, initialCell)
				b.nrInitVal++
			}
		}
	}

}

// New Create a new board with the initial values
func New(initValues *[9][9]int) *Board {
	b := new(Board)
	b.init(initValues)
	b.setCandidates()
	return b
}

func (b *Board) updateCandidate(y int, x int, v int) {
	cell := b.cells[y][x]
	if _, ok := cell.candidates[v]; ok {
		delete(cell.candidates, v)
		b.nrCandidate--
		b.pushExclusivity(&cell, y, x)
	}
}

func (b *Board) updateCandidatesLine(line int, v int) {
	for x := 0; x < nrCol; x++ {
		b.updateCandidate(line, x, v)
	}
}

func (b *Board) updateCandidatesCol(col int, v int) {
	for y := 0; y < nrLine; y++ {
		b.updateCandidate(y, col, v)
	}
}

func (b *Board) updateCandidatesBlock(line int, col int, v int) {
	yBlock := line - line%nrBlockLine
	xBlock := col - col%nrBlockCol

	for y := yBlock; y < yBlock+nrBlockCol; y++ {
		for x := xBlock; x < xBlock+nrBlockLine; x++ {
			b.updateCandidate(y, x, v)
		}
	}
}

func (b *Board) updateCandidates(line int, col int, v int) {
	b.updateCandidatesLine(line, v)
	b.updateCandidatesCol(col, v)
	b.updateCandidatesBlock(line, col, v)
}

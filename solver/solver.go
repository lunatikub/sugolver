package solver

// Sudoku constants
const (
	NrLine      = 9
	NrCol       = 9
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

// Grid Sudoku 9x9 playing grid
type Grid [NrLine][NrCol]int

type coord struct {
	y int
	x int
}

type cell struct {
	val        int
	typeCell   int
	candidates map[int]struct{}
}

// Solver Sudoku solver
type Solver struct {
	grid        [NrCol][NrLine]cell
	lines       [NrLine][nrVal]bool
	cols        [NrCol][nrVal]bool
	blocks      [nrBlock][nrVal]bool
	excluStack  []coord
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

func (s *Solver) set(y int, x int, v int, t int) {
	s.grid[y][x].val = v
	s.grid[y][x].typeCell = t
	s.lines[y][v-1] = true
	s.cols[x][v-1] = true
	s.blocks[getBlockID(y, x)][v-1] = true
	s.nrEmpty--
}

func (s *Solver) reset(y int, x int, v int) {
	s.grid[y][x].val = 0
	s.grid[y][x].typeCell = emptyCell
	s.lines[y][v-1] = false
	s.cols[x][v-1] = false
	s.blocks[getBlockID(y, x)][v-1] = false
	s.nrEmpty++
}

func (s *Solver) isValidSet(y int, x int, v int) bool {
	return !s.lines[y][v-1] && !s.cols[x][v-1] && !s.blocks[getBlockID(y, x)][v-1]
}

func (c *cell) setCandidate(s *Solver, y int, x int) {
	for v := 0; v < nrVal; v++ {
		if !s.lines[y][v] && !s.cols[x][v] && !s.blocks[getBlockID(y, x)][v] {
			c.candidates[v+1] = struct{}{}
			s.nrCandidate++
		}
	}
}

// SetCandidates find the potential
// candidates for each cell of the grid
func (s *Solver) setCandidates() {
	for y, line := range s.grid {
		for x, cell := range line {
			if cell.typeCell == emptyCell {
				cell.setCandidate(s, y, x)
				s.pushExclusivity(&cell, y, x)
			}
		}
	}
	s.nrInitCandidate = s.nrCandidate
}

// Init the grid
func (s *Solver) init(initialValues *Grid) {
	s.nrEmpty = NrCol * NrLine
	for y, line := range initialValues {
		for x, v := range line {
			s.grid[y][x].candidates = make(map[int]struct{})
			if v != 0 {
				s.set(y, x, v, initialCell)
				s.nrInitVal++
			}
		}
	}

}

// New Create a new grid with the initial values
func New(initValues *Grid) *Solver {
	s := new(Solver)
	s.init(initValues)
	s.setCandidates()
	return s
}

func (s *Solver) updateCandidate(y int, x int, v int) {
	cell := s.grid[y][x]
	if _, ok := cell.candidates[v]; ok {
		delete(cell.candidates, v)
		s.nrCandidate--
		s.pushExclusivity(&cell, y, x)
	}
}

func (s *Solver) updateCandidatesLine(line int, v int) {
	for x := 0; x < NrCol; x++ {
		s.updateCandidate(line, x, v)
	}
}

func (s *Solver) updateCandidatesCol(col int, v int) {
	for y := 0; y < NrLine; y++ {
		s.updateCandidate(y, col, v)
	}
}

func (s *Solver) updateCandidatesBlock(line int, col int, v int) {
	yBlock := line - line%nrBlockLine
	xBlock := col - col%nrBlockCol

	for y := yBlock; y < yBlock+nrBlockCol; y++ {
		for x := xBlock; x < xBlock+nrBlockLine; x++ {
			s.updateCandidate(y, x, v)
		}
	}
}

func (s *Solver) updateCandidates(line int, col int, v int) {
	s.updateCandidatesLine(line, v)
	s.updateCandidatesCol(col, v)
	s.updateCandidatesBlock(line, col, v)
}

// Solve Resolver a Sudoku
func (s *Solver) Solve(doExclu bool, doUniq bool) {
	if doExclu {
		s.exclusivity()
	}
	if s.nrEmpty == 0 {
		return
	}
	if doUniq {
		s.uniqueness()
	}
	if s.nrEmpty == 0 {
		return
	}
	s.backtracking()
}

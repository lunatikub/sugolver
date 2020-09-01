package solver

import (
	"fmt"

	"github.com/fatih/color"
)

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
	excluCell
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
	// stats
	nrInitVal       uint
	nrInitCandidate uint
	nrCandidate     uint
	nrExclusivity   uint
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
}

func (b *Board) reset(y int, x int, v int) {
	b.cells[y][x].val = 0
	b.cells[y][x].typeCell = emptyCell
	b.lines[y][v-1] = false
	b.cols[x][v-1] = false
	b.blocks[getBlockID(y, x)][v-1] = false
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

// Exclusivity if we have found the value V of a cell C then this
// value is removed from all cases in the same block as C,
// in other words we update the lines, the cols and the block boolean vectors.
func (b *Board) Exclusivity() {
	for {
		coord := b.popExclusivity()
		cell := b.cells[coord.y][coord.x]
		for v := range cell.candidates {
			b.set(coord.y, coord.x, v, excluCell)
			b.updateCandidates(coord.y, coord.x, v)
			b.nrExclusivity++
		}
		if len(b.exclusivity) == 0 {
			break
		}
	}
}

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

// DumpBoard Debug function to dump a sudoku board
func (b *Board) DumpBoard() {
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	yellow := color.New(color.FgYellow)

	for y, line := range b.cells {
		if y%3 == 0 {
			fmt.Println("-------------------------")
		}
		for x, cell := range line {
			if x%3 == 0 {
				fmt.Print("| ")
			}
			if cell.val != 0 {
				if cell.typeCell == initialCell {
					red.Print(cell.val, " ")
				} else if cell.typeCell == candidateCell {
					blue.Print(cell.val, " ")
				} else if cell.typeCell == excluCell {
					yellow.Print(cell.val, " ")
				}
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("-------------------------")
	fmt.Println("")
}

// DumpStats Debug function to dump the stats
func (b *Board) DumpStats() {
	fmt.Println("number of initial values : ", b.nrInitVal)
	fmt.Println("number of candidates     : ", b.nrInitCandidate)
	fmt.Println("number of exclusivity    : ", b.nrExclusivity)
}

// DumpCandidates Debug function to dump the candidates
func (b *Board) DumpCandidates() {
	blue := color.New(color.FgBlue)
	yellow := color.New(color.FgYellow)

	for y, line := range b.cells {
		yellow.Print("Y:", y, " ")
		for x, cell := range line {
			if len(cell.candidates) != 0 {
				yellow.Print("X:", x)
				fmt.Print("(")
				for v := range cell.candidates {
					blue.Print(v)
					fmt.Print(",")
				}
				fmt.Print(") ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

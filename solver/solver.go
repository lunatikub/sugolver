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
)

type cell struct {
	val        int // value of the cell
	typeCell   int // type of the cell
	candidates map[int]struct{}
}

// Board playing 9x9 sudoku
type Board struct {
	cells  [nrCol][nrLine]cell
	lines  [nrLine][nrVal]bool
	cols   [nrCol][nrVal]bool
	blocks [nrBlock][nrVal]bool
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
			}
		}
	}
}

// Init the board
func (b *Board) init(initialValues *[9][9]int) {
	for y, line := range initialValues {
		for x, v := range line {
			b.cells[y][x].candidates = make(map[int]struct{})
			if v != 0 {
				b.set(y, x, v, initialCell)
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

// Solve Explore all valid possibilites for each
// candidate of each cell to find a solution by backtracking
func (b *Board) Solve() {
	for y, line := range b.cells {
		for x, cell := range line {
			if cell.val == 0 {
				for v := range cell.candidates {
					if b.isValidSet(y, x, v) {
						b.set(y, x, v, candidateCell)
						b.Solve()
						b.reset(y, x, v)
					}
				}
				return
			}
		}
	}
	b.Dump()
}

// Dump Debug function to dunmp a sudoku board
func (b *Board) Dump() {
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)

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
				}
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("-------------------------")
}

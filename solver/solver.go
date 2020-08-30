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

// Enumeration of cell type
const (
	empty = iota
	initial
)

type cell struct {
	v uint // value of the cell
	t uint // type of the cell
}

// Board playing 9x9 sudoku
type Board struct {
	cells  [nrCol][nrLine]cell
	lines  [nrLine][nrVal]bool  // lines value flags
	cols   [nrCol][nrVal]bool   // cols value flags
	blocks [nrBlock][nrVal]bool // block value flags
}

// Solver resolve a suduko playing 9x9 board
type Solver struct {
	B Board
}

func getBlockID(y uint, x uint) uint {
	return x/nrBlockCol + (y/nrBlockLine)*nrBlockCol
}

func (b *Board) set(y uint, x uint, v uint, t uint) {
	b.cells[y][x].v = v
	b.cells[y][x].t = t
	b.lines[y][v-1] = true
	b.cols[x][v-1] = true
	b.blocks[getBlockID(y, x)][v-1] = true
}

// Init the board with the initial values
func (b *Board) Init(initialValues *[9][9]uint) {
	for y, line := range initialValues {
		for x, v := range line {
			if v != 0 {
				b.set(uint(y), uint(x), v, initial)
			}
		}
	}
}

// Dump the playing sudoku board
func (b *Board) Dump() {
	red := color.New(color.FgRed)

	for y, line := range b.cells {
		if y%3 == 0 {
			fmt.Println("-------------------------")
		}
		for x, cell := range line {
			if x%3 == 0 {
				fmt.Print("| ")
			}
			if cell.v != 0 {
				red.Print(cell.v, " ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("-------------------------")
}

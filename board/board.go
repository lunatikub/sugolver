package board

import (
	"fmt"

	"github.com/fatih/color"
)

// Board constants
const (
	NrLine      = 9
	NrCol       = 9
	NrBlock     = 9
	NrVal       = 9
	NrBlockLine = 3
	NrBlockCol  = 3
)

// Enumeration of cell type
const (
	Empty = iota
	Initial
)

// Cell of a board
type Cell struct {
	v uint // value of the cell
	t uint // type of the cell
}

// Board Playing sudoku
type Board struct {
	cells  [NrCol][NrLine]Cell
	lines  [NrLine][NrVal]bool  // lines values flag
	cols   [NrCol][NrVal]bool   // cols values flag
	blocks [NrBlock][NrVal]bool // blocks values flag
}

func getBlockID(y uint, x uint) uint {
	return x/NrBlockCol + (y/NrBlockLine)*NrBlockCol
}

// Set the value and the type of a cell
func (b *Board) Set(y uint, x uint, v uint, t uint) {
	b.cells[y][x].v = v
	b.cells[y][x].t = t
	b.lines[y][v-1] = true
	b.cols[x][v-1] = true
	b.blocks[getBlockID(y, x)][v-1] = true
}

// Get the value of a cell
func (b *Board) Get(y uint, x uint) uint {
	return b.cells[y][x].v
}

// Init the board with initial values
func (b *Board) Init(initialValues *[9][9]uint) {
	for y, line := range initialValues {
		for x, v := range line {
			if v != 0 {
				b.Set(uint(y), uint(x), v, Initial)
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

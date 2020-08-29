package board

// Board constants
const (
	NrLine      = 9
	NrCol       = 9
	NrBlock     = 9
	NrVal       = 9
	NrBlockLine = 3
	NrBlockCol  = 3
)

func getBlockID(y uint, x uint) uint {
	return x/NrBlockCol + (y/NrBlockLine)*NrBlockCol
}

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

// Set the value and the type of a cell
func (b *Board) Set(y uint, x uint, v uint, t uint) {
	b.cells[y][x].v = v
	b.cells[y][x].t = t
	b.lines[y][v] = true
	b.cols[x][v] = true
	b.blocks[getBlockID(y, x)][v] = true
}

// Get the value of a cell
func (b *Board) Get(y uint, x uint) uint {
	return b.cells[y][x].v
}

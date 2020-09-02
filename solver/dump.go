package solver

import (
	"fmt"

	"github.com/fatih/color"
)

// DumpBoard Debug function to dump a sudoku board
func (b *Board) DumpBoard() {
	red := color.New(color.FgRed)
	blue := color.New(color.FgBlue)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)

	fmt.Println("number of empty cell: ", b.nrEmpty)
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
				} else if cell.typeCell == exclusivityCell {
					yellow.Print(cell.val, " ")
				} else if cell.typeCell == uniquenessCell {
					green.Print(cell.val, " ")
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
	fmt.Println("number of uniqueness     : ", b.nrUniqueness)
}

// DumpCandidates Debug function to dump the candidates
func (b *Board) DumpCandidates() {
	blue := color.New(color.FgBlue)
	yellow := color.New(color.FgYellow)

	fmt.Println("number of candidates: ", b.nrCandidate)
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

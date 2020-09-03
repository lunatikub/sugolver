package solver

import (
	"fmt"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed)
var blue = color.New(color.FgBlue)
var yellow = color.New(color.FgYellow)
var green = color.New(color.FgGreen)

// DumpGrid Debug function to dump a Sudoku grid
func (s *Solver) DumpGrid() {
	fmt.Println("number of empty cell: ", s.nrEmpty)
	for y, line := range s.grid {
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
func (s *Solver) DumpStats() {
	fmt.Println("number of initial values : ", s.nrInitVal)
	fmt.Println("number of candidates     : ", s.nrInitCandidate)
	fmt.Println("number of exclusivity    : ", s.nrExclusivity)
	fmt.Println("number of uniqueness     : ", s.nrUniqueness)
}

// DumpCandidates Debug function to dump the candidates
func (s *Solver) DumpCandidates() {
	fmt.Println("number of candidates: ", s.nrCandidate)
	for y, line := range s.grid {
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

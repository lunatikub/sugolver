package solver

import (
	"fmt"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed)
var blue = color.New(color.FgBlue)
var yellow = color.New(color.FgYellow)
var green = color.New(color.FgGreen)

func (c *cell) dump() {
	if c.val != 0 {
		if c.typeCell == initialCell {
			red.Print(c.val, " ")
		} else if c.typeCell == candidateCell {
			blue.Print(c.val, " ")
		} else if c.typeCell == exclusivityCell {
			yellow.Print(c.val, " ")
		} else if c.typeCell == uniquenessCell {
			green.Print(c.val, " ")
		}
	} else {
		fmt.Print(". ")
	}
}

// DumpGrid Debug function to dump a Sudoku grid
func (s *Solver) DumpGrid() {
	fmt.Print("[")
	red.Print("initial ")
	blue.Print("backtracking ")
	yellow.Print("exclusivity ")
	green.Print("uniqueness ")
	fmt.Println("]")

	for y, line := range s.grid {
		if y%3 == 0 {
			fmt.Println("-------------------------")
		}
		for x, cell := range line {
			if x%3 == 0 {
				fmt.Print("| ")
			}
			cell.dump()
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
	fmt.Println("number of backtracking   : ", s.nrBacktracking)
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

// DumpSolution Debug function to dump the solution
func (s *Solver) DumpSolution() {
	for _, line := range s.grid {
		for _, cell := range line {
			fmt.Print(cell.val)
		}
	}
}

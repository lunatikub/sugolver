package main

import solver "github.com/lunatikub/sugolver/solver"

func main() {

	var simple = solver.Grid{
		{5, 0, 0, 0, 8, 3, 2, 0, 0},
		{0, 4, 7, 0, 0, 5, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 9, 4},
		{0, 0, 0, 6, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 9, 7, 4, 6, 0},
		{0, 0, 8, 2, 3, 0, 0, 0, 5},
		{8, 0, 5, 7, 0, 6, 3, 0, 0},
		{4, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 0, 5, 9, 0, 0, 0},
	}

	s := solver.New(&simple)
	s.DumpGrid()
	s.DumpCandidates()
	s.Uniqueness()
	s.DumpGrid()
	s.DumpStats()
}

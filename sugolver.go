package main

import solver "github.com/lunatikub/sugolver/solver"

func main() {

	var easySudoku = [9][9]uint{
		{8, 6, 9, 0, 0, 0, 0, 0, 4},
		{0, 4, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 1, 3, 9, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 3, 0, 0, 0, 8, 0, 0},
		{0, 0, 6, 0, 0, 0, 1, 4, 0},
		{0, 1, 0, 0, 8, 3, 0, 5, 0},
		{0, 3, 4, 0, 5, 0, 0, 0, 7},
	}

	var s solver.Solver
	s.B.Init(&easySudoku)
	s.B.Dump()
}

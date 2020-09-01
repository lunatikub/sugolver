package main

import board "github.com/lunatikub/sugolver/solver"

func main() {

	var simple1 = [9][9]int{
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

	b := board.New(&simple1)
	b.DumpBoard()
	b.DumpCandidates()

	for {
		if nr := b.Exclusivity(); nr == 0 {
			break
		}
	}

	b.DumpBoard()
	b.DumpCandidates()
	b.Backtracking()
}

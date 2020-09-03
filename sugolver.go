package main

import (
	"bufio"
	"flag"
	"strings"

	solver "github.com/lunatikub/sugolver/solver"
)

func parseGrid(intput string) solver.Grid {
	var grid solver.Grid
	var x, y int

	r := bufio.NewReader(strings.NewReader(intput))
	for {
		if c, _, err := r.ReadRune(); err == nil {
			if string(c) == "." {
				grid[y][x] = 0
			} else {
				grid[y][x] = int(c) - '0'
			}
			x++
			if x > solver.NrCol-1 {
				x = 0
				y++
			}
		} else {
			break
		}
	}
	return grid
}

// var simple = solver.Grid{
// 	{5, 0, 0, 0, 8, 3, 2, 0, 0},
// 	{0, 4, 7, 0, 0, 5, 0, 0, 0},
// 	{0, 0, 0, 0, 0, 0, 0, 9, 4},
// 	{0, 0, 0, 6, 0, 0, 0, 0, 0},
// 	{0, 0, 0, 0, 9, 7, 4, 6, 0},
// 	{0, 0, 8, 2, 3, 0, 0, 0, 5},
// 	{8, 0, 5, 7, 0, 6, 3, 0, 0},
// 	{4, 0, 0, 0, 0, 0, 0, 0, 0},
// 	{0, 0, 3, 0, 5, 9, 0, 0, 0},
// }

type options struct {
	grid     string
	solution bool
	stats    bool
}

func getOptions() *options {
	opts := new(options)

	flag.StringVar(&opts.grid, "grid", "", "grid to solve <1..34.5...6...7[....]>")
	flag.BoolVar(&opts.solution, "solution", false, "dump the solution")
	flag.BoolVar(&opts.stats, "stats", false, "dump the statistics")

	flag.Parse()

	if opts.grid == "" {
		panic("Option '--grid' is mandatory")
	}

	return opts
}

func main() {
	opts := getOptions()
	grid := parseGrid(opts.grid)

	s := solver.New(&grid)

	s.Solve(false, false)

	if opts.solution {
		s.DumpSolution()
	}
	if opts.stats {
		s.DumpStats()
	}
}

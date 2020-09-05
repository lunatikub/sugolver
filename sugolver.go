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

type options struct {
	grid        string
	dump        string
	stats       bool
	exclusivity bool
	uniqueness  bool
	parity      bool
	pretty      bool
}

func getOptions() *options {
	opts := new(options)

	flag.StringVar(&opts.grid, "grid", "", "grid to solve <1..34.5...6...7[....]>")
	flag.StringVar(&opts.dump, "dump", "", "dump the solution [solution, pretty]")
	flag.BoolVar(&opts.stats, "stats", false, "dump the statistics")
	flag.BoolVar(&opts.exclusivity, "exclusivity", false, "enable exclusivity")
	flag.BoolVar(&opts.uniqueness, "uniqueness", false, "enable uniqueness")
	flag.BoolVar(&opts.parity, "parity", false, "enable parity")

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

	s.Solve(opts.exclusivity, opts.uniqueness)

	if opts.dump != "" {
		if opts.dump == "solution" {
			s.DumpSolution()
		} else if opts.dump == "pretty" {
			s.DumpGrid()
		}
	}

	if opts.stats {
		s.DumpStats()
	}
}

package solver

import "testing"

func testBlockID(t *testing.T, y int, x int, blockIDExpected int) {
	var blockID int

	blockID = getBlockID(y, x)
	if blockID != blockIDExpected {
		t.Errorf("getBlockID s incorrect, got: %d, expected: %d", blockID, blockIDExpected)
	}
}

func TestBlockID(t *testing.T) {
	testBlockID(t, 8, 8, 8)
	testBlockID(t, 1, 2, 0)
	testBlockID(t, 2, 6, 2)
	testBlockID(t, 8, 4, 7)
}

func testSet(t *testing.T, s *Solver, y int, x int, v int) {
	s.set(y, x, v, initialCell)
	res := s.grid[y][x].val
	if res != v {
		t.Errorf("Set/Get is incorrect, got: %d, expected: %d", res, v)
	}
	if s.lines[y][v-1] != true {
		t.Errorf("Set of val %d is incorrect for the line:%d", v, y)
	}
	if s.cols[x][v-1] != true {
		t.Errorf("Set of val %d is incorrect for the col:%d", v, x)
	}
	blockID := getBlockID(y, x)
	if s.blocks[blockID][v-1] != true {
		t.Errorf("Set of val %d is incorrect for the block:%d", v, blockID)
	}
}

func TestSet(t *testing.T) {
	var s Solver
	testSet(t, &s, 3, 3, 1)
	testSet(t, &s, 6, 5, 2)
	testSet(t, &s, 3, 3, 2)
}

func TestIsValid(t *testing.T) {
	var s Solver
	if s.isValidSet(0, 0, 1) == false {
		t.Errorf("Valid set expected")
	}
	s.set(0, 0, 1, candidateCell)
	if s.isValidSet(0, 1, 1) == true {
		t.Errorf("Unvalid set expected")
	}
	if s.isValidSet(1, 0, 1) == true {
		t.Errorf("Unvalid set expected")
	}
	if s.isValidSet(1, 1, 1) == true {
		t.Errorf("Unvalid set expected")
	}
}

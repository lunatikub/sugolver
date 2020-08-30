package solver

import "testing"

func testBlockID(t *testing.T, y uint, x uint, blockIDExpected uint) {
	var blockID uint

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

func testSet(t *testing.T, b *Board, y uint, x uint, v uint) {
	b.set(y, x, v, initial)
	res := b.cells[y][x].v
	if res != v {
		t.Errorf("Set/Get is incorrect, got: %d, expected: %d", res, v)
	}
	if b.lines[y][v-1] != true {
		t.Errorf("Set of val %d is incorrect for the line:%d", v, y)
	}
	if b.cols[x][v-1] != true {
		t.Errorf("Set of val %d is incorrect for the col:%d", v, x)
	}
	blockID := getBlockID(y, x)
	if b.blocks[blockID][v-1] != true {
		t.Errorf("Set of val %d is incorrect for the block:%d", v, blockID)
	}
}

func TestSet(t *testing.T) {
	var b Board
	testSet(t, &b, 3, 3, 1)
	testSet(t, &b, 6, 5, 2)
	testSet(t, &b, 3, 3, 2)
}

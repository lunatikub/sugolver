package board

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

func testSetGet(t *testing.T, b *Board, y uint, x uint, v uint) {
	b.Set(y, x, v, Initial)
	res := b.Get(y, x)
	if res != v {
		t.Errorf("Set/Get is incorrect, got: %d, expected: %d", res, v)
	}
	if b.lines[y][v] != true {
		t.Errorf("Set is incorrect, lines[%d][%d] is not set", y, v)
	}
	if b.cols[x][v] != true {
		t.Errorf("Set is incorrect, cols[%d][%d] is not set", x, v)
	}
	blockID := getBlockID(y, x)
	if b.blocks[blockID][v] != true {
		t.Errorf("Set is incorrect, blocks[%d][%d] is not set", blockID, v)
	}
}

func TestSetGet(t *testing.T) {
	var b Board
	testSetGet(t, &b, 3, 3, 1)
	testSetGet(t, &b, 6, 5, 2)
	testSetGet(t, &b, 3, 3, 2)
}

package main

import (
	"testing"
)

func Test(t *testing.T) {
	expectedCells := [][]int{
		[]int{0, 1},
		[]int{0, 0},
		[]int{1, 1},
		[]int{1, 0},
	}

	distOrderCells := allCellsDistOrder(2, 2, 0, 1)

    t.Logf("\n%v\n", expectedCells)
    t.Logf("\n%v\n", distOrderCells)

	for i, cell := range distOrderCells {
		expectedCell := expectedCells[i]
		if expectedCell[0] != cell[0] || expectedCell[1] != cell[1] {
			t.Logf("Expected: %v , Returned: %v", expectedCell, cell)
		}
	}

	if len(expectedCells) != len(distOrderCells) {
		t.Error("length no mathc expected")
	}
}

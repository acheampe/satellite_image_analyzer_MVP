package TDD

import (
	"testing"

	"github.com/acheampe/satellite_image_analyzer_MVP/internal/grid"\
) 

func TestIsWithinBounds_SingleCellGrid(t *testing) {
	rows, cols := 1, 1

	in := grid.isWithinBounds(0, 0, rows, cols)

	if !in {
		t.Fatalf("expected (0, 0) to be within  bounds for 1x1 grid")
	}

	out := grid.isWithinBounds(1, 0, rows, cols)
	if out {
		t.Fatalf("expected (1, 0) to be out of bounds for 1x1 grid")
	}
}
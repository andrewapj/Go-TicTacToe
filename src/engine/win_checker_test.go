package engine

import (
	"github.com/andrewapj/go-tictactoe/src/grid"
	"testing"
)

func TestHorizontal(t *testing.T) {

	for y := uint8(0); y < 3; y++ {
		g := grid.New()
		_ = g.SelectSquare(0, y, grid.SquareO)
		_ = g.SelectSquare(1, y, grid.SquareO)
		_ = g.SelectSquare(2, y, grid.SquareO)

		if !CheckWin(&g) {
			t.Fatal("Expected a win")
		}
	}
}

func TestVertical(t *testing.T) {

	for x := uint8(0); x < 3; x++ {
		g := grid.New()
		_ = g.SelectSquare(x, 0, grid.SquareO)
		_ = g.SelectSquare(x, 1, grid.SquareO)
		_ = g.SelectSquare(x, 2, grid.SquareO)

		if !CheckWin(&g) {
			t.Fatal("Expected a win")
		}
	}
}

func TestDiagonal(t *testing.T) {

	g := grid.New()

	_ = g.SelectSquare(0, 0, grid.SquareO)
	_ = g.SelectSquare(1, 1, grid.SquareO)
	_ = g.SelectSquare(2, 2, grid.SquareO)

	if !CheckWin(&g) {
		t.Fatal("Expected a win")
	}

	g2 := grid.New()

	_ = g2.SelectSquare(2, 0, grid.SquareO)
	_ = g2.SelectSquare(1, 1, grid.SquareO)
	_ = g2.SelectSquare(0, 2, grid.SquareO)

	if !CheckWin(&g2) {
		t.Fatal("Expected a win")
	}
}

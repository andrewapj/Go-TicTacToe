package grid

import (
	"testing"
)

/**
Selecting a square
*/

func TestSetsCellIfEmpty(t *testing.T) {
	grid := New()

	_ = grid.SelectSquare(0, 0, SquareO)

	if grid.square[0][0] != SquareO {
		t.Fatalf("Failed to set the square state correctly")
	}
}

func TestErrorIfSquareOutsideRange(t *testing.T) {
	grid := New()

	if err := grid.SelectSquare(3, 0, SquareO); err == nil {
		t.Fatalf("Expected an error")
	}
	if err := grid.SelectSquare(0, 3, SquareO); err == nil {
		t.Fatalf("Expected an error")
	}
	if err := grid.SelectSquare(3, 3, SquareO); err == nil {
		t.Fatalf("Expected an error")
	}
}

func TestErrorIfSquareTaken(t *testing.T) {
	grid := New()

	grid.square[0][0] = "X"

	if err := grid.SelectSquare(0, 0, SquareX); err == nil {
		t.Fatalf("Expected an error when selecting a square that is taken")
	}
}

func TestErrorIfSquareStateInvalid(t *testing.T) {
	grid := New()

	if err := grid.SelectSquare(0, 0, "A"); err == nil {
		t.Fatalf("Expected an error when setting an invalid square state")
	}
}

/**
Getting a square
*/

func TestGetSquareCorrectly(t *testing.T) {
	grid := New()

	if s, _ := grid.GetSquare(0, 0); s != "" {
		t.Fatalf("Expected to get an empty square")
	}
}

func TestErrorIfGetSquareOutOfBounds(t *testing.T) {
	grid := New()

	if _, err := grid.GetSquare(3, 0); err == nil {
		t.Fatalf("Expected an error")
	}
}

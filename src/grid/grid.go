package grid

import (
	"errors"
)

const (
	//Size - The length and height of the grid
	Size uint8 = 3
	//SquareX - The symbol to use for player X
	SquareX = "X"
	//SquareO - The symbol to use for player O
	SquareO = "O"
	//SquareEmpty - The symbol to use for an empty square
	SquareEmpty = ""
)

//Grid - represents the game grid
type Grid struct {
	square [Size][Size]string
}

//New - Creates a new game grid
func New() Grid {
	return Grid{square: [Size][Size]string{
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}}
}

//SelectSquare - Selects a square on the grid and sets it's state
func (g *Grid) SelectSquare(x uint8, y uint8, state string) error {

	if x < 0 || x >= Size || y < 0 || y >= Size {
		return errors.New("invalid square selected")
	}

	if g.square[y][x] != SquareEmpty {
		return errors.New("square already taken")
	}

	if state != SquareO && state != SquareX {
		return errors.New("invalid square state chosen")
	}

	g.square[y][x] = state
	return nil
}

//GetSquare - Gets a square within the grid
func (g *Grid) GetSquare(x uint8, y uint8) (string, error) {

	if x < 0 || x >= Size || y < 0 || y >= Size {
		return "", errors.New("invalid square selected")
	}

	return g.square[y][x], nil
}

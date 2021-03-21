package grid

import (
	"errors"
)

const (
	Size        uint8 = 3
	SquareX           = "X"
	SquareO           = "O"
	SquareEmpty       = ""
)

type Grid struct {
	square [Size][Size]string
}

func New() Grid {
	return Grid{square: [Size][Size]string{
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
		{SquareEmpty, SquareEmpty, SquareEmpty},
	}}
}

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

func (g *Grid) GetSquare(x uint8, y uint8) (string, error) {

	if x < 0 || x >= Size || y < 0 || y >= Size {
		return "", errors.New("invalid square selected")
	}

	return g.square[y][x], nil
}

package engine

import "github.com/andrewapj/go-tictactoe/src/grid"

//CheckWin - Checks if a player has won the game
func CheckWin(g *grid.Grid) bool {
	return checkHorizontal(g) || checkVertical(g) || checkDiagonal(g)
}

func checkHorizontal(g *grid.Grid) bool {
	for y := uint8(0); y < 3; y++ {
		s1, _ := g.GetSquare(0, y)
		s2, _ := g.GetSquare(1, y)
		s3, _ := g.GetSquare(2, y)

		if s1 == s2 && s2 == s3 && s1 != grid.SquareEmpty {
			return true
		}
	}

	return false
}

func checkVertical(g *grid.Grid) bool {

	for x := uint8(0); x < 3; x++ {
		s1, _ := g.GetSquare(x, 0)
		s2, _ := g.GetSquare(x, 1)
		s3, _ := g.GetSquare(x, 2)

		if s1 == s2 && s2 == s3 && s1 != grid.SquareEmpty {
			return true
		}
	}

	return false
}

func checkDiagonal(g *grid.Grid) bool {

	s1, _ := g.GetSquare(0, 0)
	s2, _ := g.GetSquare(1, 1)
	s3, _ := g.GetSquare(2, 2)

	if s1 == s2 && s2 == s3 && s1 != grid.SquareEmpty {
		return true
	}

	s1, _ = g.GetSquare(2, 0)
	s2, _ = g.GetSquare(1, 1)
	s3, _ = g.GetSquare(0, 2)

	if s1 == s2 && s2 == s3 && s1 != grid.SquareEmpty {
		return true
	}

	return false
}

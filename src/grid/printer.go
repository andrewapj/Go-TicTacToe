package grid

import (
	"fmt"
	"github.com/andrewapj/go-tictactoe/src/input"
)

func PrintGrid(g *Grid) {

	for y := uint8(0); y < Size; y++ {
		for x := uint8(0); x < Size; x++ {
			s, _ := g.GetSquare(x, y)
			fmt.Printf("|")
			if s == SquareEmpty {
				fmt.Printf(" (%d) ", input.CalculateMarkerNumber(x, y))
			} else {
				fmt.Printf("  %s  ", s)
			}
		}
		fmt.Printf("|")
		fmt.Println("")
	}
}

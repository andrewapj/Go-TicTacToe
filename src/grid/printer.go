package grid

import "fmt"

func PrintGrid(g Grid) {

	for y := uint8(0); y < Size; y++ {
		for x := uint8(0); x < Size; x++ {
			s, _ := g.GetSquare(x, y)
			fmt.Printf("|")
			if s == SquareEmpty {
				fmt.Printf(" (%d) ", calculateMarkerNumber(x, y))
			} else {
				fmt.Printf("  %s  ", s)
			}
		}
		fmt.Printf("|")
		fmt.Println("")
	}
}

func calculateMarkerNumber(x uint8, y uint8) uint8 {
	return (y * 3) + (x + 1)
}

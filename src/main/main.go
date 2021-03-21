package main

import (
	"github.com/andrewapj/go-tictactoe/src/grid"
)

func main() {

	gameGrid := grid.New()
	grid.PrintGrid(gameGrid)
}

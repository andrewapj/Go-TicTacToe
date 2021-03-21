package main

import (
	"github.com/andrewapj/go-tictactoe/src/engine"
	"github.com/andrewapj/go-tictactoe/src/grid"
)

func main() {

	gameGrid := grid.New()
	engine.Play(&gameGrid)
}

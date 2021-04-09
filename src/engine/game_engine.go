package engine

import (
	"fmt"
	"github.com/andrewapj/go-tictactoe/src/grid"
	"github.com/andrewapj/go-tictactoe/src/input"
)

//Play - Main game loop that plays the game
func Play(g *grid.Grid) {

	currentTurn := swapPlayerTurn("")
	totalTurns := 0

	fmt.Println("Welcome to Tic Tac Toe")

	grid.PrintGrid(g)

	for {
		fmt.Println()
		fmt.Printf("Player %s's turn. Please select a square: ", currentTurn)

		choice, err := input.GetInput()
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		x, y := input.CalculateIndexNumber(choice)
		if err = g.SelectSquare(x, y, currentTurn); err != nil {
			fmt.Println("Invalid square selection")
		} else {
			totalTurns++
		}

		grid.PrintGrid(g)

		if CheckWin(g) {
			fmt.Println()
			fmt.Printf("Player %s wins!", currentTurn)
			fmt.Println()
			break
		}

		if totalTurns == 9 {
			fmt.Println()
			fmt.Print("It's a draw!")
			fmt.Println()
			break
		}

		currentTurn = swapPlayerTurn(currentTurn)
	}
}

func swapPlayerTurn(player string) string {
	if player == "" || player == grid.SquareO {
		return grid.SquareX
	}
	return grid.SquareO
}

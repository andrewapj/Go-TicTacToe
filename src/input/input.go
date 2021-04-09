package input

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

//GetInput - Gets input from the user, for square selection
func GetInput() (uint8, error) {
	scanner.Scan()
	input := scanner.Text()

	return checkInput(input)
}

func checkInput(s string) (uint8, error) {
	val, err := strconv.ParseInt(s, 10, 8)

	if err != nil {
		return 0, err
	} else if val < 1 || val > 9 {
		return 0, errors.New("invalid value")
	}

	return uint8(val), err
}

//CalculateMarkerNumber - Calculates the number on the grid, which allows the user to make their selection
func CalculateMarkerNumber(x uint8, y uint8) uint8 {
	return (y * 3) + (x + 1)
}

//CalculateIndexNumber - Calculates the index in the grid from the marker number selected by the user
func CalculateIndexNumber(i uint8) (x uint8, y uint8) {

	if i >= 1 && i <= 3 {
		y = 0
	} else if i >= 4 && i <= 6 {
		y = 1
	} else {
		y = 2
	}

	if i == 1 || i == 4 || i == 7 {
		x = 0
	} else if i == 2 || i == 5 || i == 8 {
		x = 1
	} else {
		x = 2
	}

	return x, y
}

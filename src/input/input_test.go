package input

import (
	"strconv"
	"testing"
)

func TestValidInput(t *testing.T) {

	for i := 1; i <= 9; i++ {
		val, err := checkInput(strconv.Itoa(i))
		if err != nil {
			t.Fatalf("Got an error when parsing %d : %s", i, err.Error())
		}
		if val != uint8(i) {
			t.Fatalf("Incorrect conversion, expected %d, got %d", i, val)
		}
	}
}

func TestInvalidInputWithNonNumber(t *testing.T) {

	invalidInputs := []string{"", " ", "A", "0", "10", "99999", "@", "5.1"}

	for _, input := range invalidInputs {
		if _, err := checkInput(input); err == nil {
			t.Fatal("Expected an error when parsing invalid input")
		}
	}
}

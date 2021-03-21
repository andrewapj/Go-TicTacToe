package custom_errors

import "fmt"

type SquareTakenError struct {
}

func (e *SquareTakenError) Error() string {
	return fmt.Sprint("Square already taken")
}

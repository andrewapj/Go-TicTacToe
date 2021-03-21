package custom_errors

import "fmt"

type OutOfBoundsError struct {
	Message string
}

func (e *OutOfBoundsError) Error() string {
	return fmt.Sprintf(e.Message)
}

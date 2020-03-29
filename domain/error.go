package domain

import "fmt"

type BackendError struct {
	Msg string
	Err error
}

func (e *BackendError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}

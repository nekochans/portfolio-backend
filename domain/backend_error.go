package domain

type BackendError struct {
	Msg string
	Err error
}

func (e *BackendError) Error() string {
	return e.Msg
}

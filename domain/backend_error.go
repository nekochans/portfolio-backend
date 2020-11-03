package domain

type BackendError struct {
	Message string
	Err     error
}

func (e *BackendError) Error() string {
	return e.Message
}

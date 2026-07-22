package model

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrNotFound      Error = "resource not found"
	ErrAlreadyExists Error = "resource already exists"
)

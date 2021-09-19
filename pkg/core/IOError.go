package core

import "fmt"

type IOError struct {
	Path string
	Err  error
}

func NewIOError(path string, err error) *IOError {
	return &IOError{path, err}
}

func (ioErr *IOError) Error() string {
	return fmt.Sprintf("path: %s\n error: %v", ioErr.Path, ioErr.Err)
}

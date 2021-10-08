package core

import "fmt"

type ExecError struct {
	Command string
	Err     error
}

func NewExecError(command string, err error) *ExecError {
	return &ExecError{command, err}
}

func (execErr ExecError) Error() string {
	return fmt.Sprintf("command: %s\n error: %v", execErr.Command, execErr.Err)
}


package core

import "fmt"

type DBError struct {
	ActionRequested string
	Err error
}

func NewDBError(actionRequested string, err error) *DBError {
	return &DBError{actionRequested, err}
}

func (dbError *DBError) Error() string {
	return  fmt.Sprintf("trying to do: %s\n error: %v", dbError.ActionRequested, dbError.Err);
}

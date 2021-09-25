package core

import (
	"fmt"
	"reflect"
)

type TypeError struct {
	From reflect.Type
	To   reflect.Type
}

func NewTypeError(from interface{}, to interface{}) *TypeError {
	return &TypeError{
		From: reflect.TypeOf(from),
		To:   reflect.TypeOf(to),
	}
}

func (typeErr TypeError) Error() string {
	return fmt.Sprintf("invalid type provided, failed to convert %s to %s", typeErr.From, typeErr.To)
}

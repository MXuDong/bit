package bit

import "fmt"

type Error struct {
	value string
}

func gen(value string) Error {
	return Error{value: value}
}

func (e *Error) Error(i ...interface{}) error {
	return fmt.Errorf(e.value, i...)
}

package congo

import (
	"errors"
	"fmt"
)

var (
	ErrNotImplemented = errors.New("Method not implemented")
	ErrNotSupported   = errors.New("Method not supported")
)

type ErrKeyNotFound struct {
	Key    string
	Config Config
}

func (e ErrKeyNotFound) Error() string {
	return fmt.Sprintf("Key %q no found in %s", e.Key, e.Config)
}

type ErrInvalidValueType struct {
	ValueLocation ValueLocation
	Value         interface{}
}

func (e ErrInvalidValueType) Error() string {
	return fmt.Sprintf("Cannot get value of type '%T' from %q", e.Value, e.ValueLocation)
}

func IsKeyNotFound(e error) bool {
	_, ok := e.(ErrKeyNotFound)
	return ok
}
func IsInvalidValueType(e error) bool {
	_, ok := e.(ErrInvalidValueType)
	return ok
}

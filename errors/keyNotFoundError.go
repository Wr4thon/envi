package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

type KeyNotFoundError struct {
	EnviError
	key engine.Key
}

func NewKeyNotFoundError(
	key engine.Key,
) *KeyNotFoundError {
	return &KeyNotFoundError{
		EnviError: NewEnviError(ErrKeyNotFound),
		key:       key,
	}
}

func (e KeyNotFoundError) Error() string {
	return "\n\trequested key: \"" + e.Key().Value() + "\": " + e.EnviError.Error()
}

func (e *KeyNotFoundError) Key() engine.Key {
	return e.key
}

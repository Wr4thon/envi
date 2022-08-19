package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

// KeyNotFoundError is used when you request a key that is unknown.
type KeyNotFoundError struct {
	EnviError
	key engine.Key
}

// NewKeyNotFoundError creates a new KeyNotFoundError.
func NewKeyNotFoundError(
	key engine.Key,
) *KeyNotFoundError {
	return &KeyNotFoundError{
		EnviError: NewEnviError(ErrKeyNotFound),
		key:       key,
	}
}

// Error implements the error interface.
func (e KeyNotFoundError) Error() string {
	return "\n\trequested key: \"" + e.Key().Value() + "\": " + e.EnviError.Error()
}

// Key can be used to retrieve the key associated with this error.
func (e *KeyNotFoundError) Key() engine.Key {
	return e.key
}

package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

// UnmarshalError is used when a problem occurs
type UnmarshalError struct {
	VariableError
}

// WrapUnmarshalError can be used to provide additional information about the
// error. When the cause is nil, nil will be returned.
func WrapUnmarshalError(
	cause error,
	variable engine.Var,
) error {
	if cause == nil {
		return nil
	}

	return NewUnmarshalError(
		cause,
		variable,
	)
}

// NewUnmarshalError creates a new UnmarshalError.
func NewUnmarshalError(
	cause error,
	variable engine.Var,
) *UnmarshalError {
	return &UnmarshalError{
		VariableError: NewVariableError(cause, variable),
	}
}

// Error implements the error interface.
func (e UnmarshalError) Error() string {
	return "\n\tcould not unmarshal data: " + e.VariableError.Error()
}

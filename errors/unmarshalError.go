package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

type UnmarshalError struct {
	VariableError
}

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

func NewUnmarshalError(
	cause error,
	variable engine.Var,
) *UnmarshalError {
	return &UnmarshalError{
		VariableError: NewVariableError(cause, variable),
	}
}

func (e UnmarshalError) Error() string {
	return "\n\tcould not unmarshal data: " + e.VariableError.Error()
}

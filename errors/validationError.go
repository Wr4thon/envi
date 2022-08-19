package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

type ValidationError struct {
	VariableError
}

func WrapValidationError(
	cause error,
	v engine.Var,
) error {
	if cause == nil {
		return nil
	}

	return NewValidationError(cause, v)
}

func NewValidationError(
	cause error,
	v engine.Var,
) ValidationError {
	return ValidationError{
		VariableError: NewVariableError(cause, v),
	}
}

// TODO: Validator Name?
func (e ValidationError) Error() string {
	return "\n\terror while validating variable: " + e.VariableError.Error()
}

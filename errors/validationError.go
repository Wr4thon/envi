package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

// ValidationError is used when an error occurs during the validation of values.
type ValidationError struct {
	VariableError
}

// WrapValidationError can be used to provide additional information about the
// error. When the cause is nil, nil will be returned.
func WrapValidationError(
	cause error,
	v engine.Var,
) error {
	if cause == nil {
		return nil
	}

	return NewValidationError(cause, v)
}

// NewValidationError creates a new ValidationError.
func NewValidationError(
	cause error,
	v engine.Var,
) ValidationError {
	return ValidationError{
		VariableError: NewVariableError(cause, v),
	}
}

// Error implements the error interface.
// TODO: Validator Name?
func (e ValidationError) Error() string {
	return "\n\terror while validating variable: " + e.VariableError.Error()
}

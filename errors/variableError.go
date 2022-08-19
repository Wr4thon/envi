package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

// VariableError is used when a problem is encountered while working with
// variables.
type VariableError struct {
	EnviError

	variable engine.Var
}

// WrapVariableError can be used to provide additional information about the
// error. When the cause is nil, nil will be returned.
func WrapVariableError(
	cause error,
	variable engine.Var,
) error {
	if cause == nil {
		return nil
	}

	return NewVariableError(cause, variable)
}

// NewVariableError creates a new VariableError.
func NewVariableError(
	cause error,
	variable engine.Var,
) VariableError {
	return VariableError{
		EnviError: NewEnviError(cause),
		variable:  variable,
	}
}

// Error implements the error interface.
func (e VariableError) Error() string {
	return "\n\tvariable with key: \"" +
		e.Key().Value() +
		"\": " +
		e.EnviError.Error()
}

// Key can be used to get the key that causes problems.
func (e VariableError) Key() engine.Key {
	return e.variable.Key()
}

// Variable can be used to access the variable that causes problems.
func (e VariableError) Variable() engine.Var {
	return e.variable
}

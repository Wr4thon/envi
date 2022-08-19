package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

type VariableError struct {
	EnviError

	variable engine.Var
}

func WrapVariableError(
	cause error,
	variable engine.Var,
) error {
	if cause == nil {
		return nil
	}

	return NewVariableError(cause, variable)
}

func NewVariableError(
	cause error,
	variable engine.Var,
) VariableError {
	return VariableError{
		EnviError: NewEnviError(cause),
		variable:  variable,
	}
}

func (e VariableError) Error() string {
	return "\n\tvariable with key: \"" +
		e.Key().Value() +
		"\": " +
		e.EnviError.Error()
}

func (e VariableError) Key() engine.Key {
	return e.variable.Key()
}

func (e VariableError) Variable() engine.Var {
	return e.variable
}

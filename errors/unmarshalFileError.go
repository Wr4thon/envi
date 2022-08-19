package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

type pathVar interface {
	engine.Var
	Path() string
}

type UnmarshalFileError struct {
	FileError
}

func WrapUnmarshalFileError(
	cause error,
	variable pathVar,
) error {
	if cause == nil {
		return nil
	}

	return NewUnmarshalFileError(cause, variable)
}

func NewUnmarshalFileError(
	cause error,
	variable pathVar) UnmarshalFileError {

	return UnmarshalFileError{
		FileError: NewFileError(
			NewUnmarshalError(
				cause,
				variable,
			),
			variable,
		),
	}
}

func (e UnmarshalFileError) Error() string {
	return "\n\tfailed to unmarshal file:" + e.FileError.Error()
}

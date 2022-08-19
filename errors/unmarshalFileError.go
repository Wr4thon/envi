package errors

import (
	"github.com/Clarilab/envi/v2/engine"
)

type pathVar interface {
	engine.Var
	Path() string
}

// UnmarshalFileError is used when an error occurs while unmarshaling a file.
type UnmarshalFileError struct {
	FileError
}

// WrapUnmarshalFileError can be used to provide additional information about
// the error. When the cause is nil, nil will be returned.
func WrapUnmarshalFileError(
	cause error,
	variable pathVar,
) error {
	if cause == nil {
		return nil
	}

	return NewUnmarshalFileError(cause, variable)
}

// NewUnmarshalFileError creates a new UnmarshalFileError.
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

// Error implements the error interface.
func (e UnmarshalFileError) Error() string {
	return "\n\tfailed to unmarshal file:" + e.FileError.Error()
}

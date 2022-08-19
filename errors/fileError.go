package errors

// FileError is used to provide information when something went wrong
// that involves a file.
type FileError struct {
	VariableError
	filePath string
}

// WrapFileError can be used to provide additional information about the
// file involved when an error occurred.
// When the cause is nil, nil will be returned.
func WrapFileError(
	cause error,
	v pathVar,
) error {
	if cause == nil {
		return nil
	}

	return NewFileError(
		cause,
		v,
	)
}

// NewFileError creates a new FileError.
func NewFileError(
	cause error,
	v pathVar,
) FileError {
	return FileError{
		VariableError: NewVariableError(cause, v),
		filePath:      v.Path(),
	}
}

// Error implements the error interface.
func (e FileError) Error() string {
	return "\n\terror with file \"" + e.filePath + "\": " + e.VariableError.Error()
}

// FilePath is used to retrieve the filePath associated with this error.
func (e FileError) FilePath() string {
	return e.filePath
}

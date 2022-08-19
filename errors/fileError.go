package errors

type FileError struct {
	VariableError
	filePath string
}

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

func NewFileError(
	cause error,
	v pathVar,
) FileError {
	return FileError{
		VariableError: NewVariableError(cause, v),
		filePath:      v.Path(),
	}
}

func (e FileError) Error() string {
	return "\n\terror with file \"" + e.filePath + "\": " + e.VariableError.Error()
}

func (e FileError) FilePath() string {
	return e.filePath
}

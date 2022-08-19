package envi

import (
	"fmt"
	"strings"
	"syscall"
)

var (
	ErrMissingFile error = syscall.Errno(syscall.ENOENT)
)

func fileError(path string, cause error) error {
	return &FileError{
		EnviError: &EnviError{
			cause: cause,
		},
		filePath: path,
	}
}

type EnviError struct {
	cause error
}

func (e *EnviError) Unwrap() error {
	return e.Cause()
}

func (e *EnviError) Cause() error {
	return e.cause
}

type FileError struct {
	*EnviError
	filePath string
}

func (e *FileError) FilePath() string {
	return e.filePath
}

func (e *FileError) Error() string {
	return fmt.Sprintf("missing file '%s': %v", e.filePath, e.EnviError.Cause())
}

func unmarshalError(cause error) error {
	return &UnmarshalError{
		EnviError: &EnviError{
			cause: cause,
		},
	}
}

type UnmarshalError struct {
	*EnviError
}

func (e *UnmarshalError) Error() string {
	return fmt.Sprintf("could not unmarshal data: %v", e.EnviError.Cause())
}

func unmarshalFileError(path string, cause error) error {
	return &UnmarshalFileError{
		FileError: FileError{
			EnviError: &EnviError{
				cause: cause,
			},
			filePath: path,
		},
	}
}

type UnmarshalFileError struct {
	FileError
}

func (e *UnmarshalFileError) Cause() error {
	return e.cause
}

func (e *UnmarshalFileError) FilePath() string {
	return e.filePath
}

func (e *UnmarshalFileError) Error() string {
	return fmt.Sprintf("failed to unmarshal file '%s': %v", e.filePath, e.EnviError.Cause())
}

// RequiredEnvVarsMissing says, that a required Environment Variable is not given.
type RequiredEnvVarsMissing struct {
	missingVars []string
}

func (e *RequiredEnvVarsMissing) Error() string {
	return fmt.Sprintf("One or more required environment variables are missing\nThe missing variables are: %s", e.printMissingVars())
}

func (e *RequiredEnvVarsMissing) printMissingVars() string {
	return strings.Join(e.missingVars, ", ")
}

func (e *RequiredEnvVarsMissing) MissingVars() []string {
	result := make([]string, len(e.missingVars))
	copy(result, e.missingVars)
	return result
}

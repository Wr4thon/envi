package envi

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrValidationFailed       error = errors.New("variable validation failed")
	ErrRequiredVariableNotSet error = errors.New("required variable not set")
	ErrUnknownSource          error = errors.New("source is unknown or not set")
	ErrUnknownEncoding        error = errors.New("encoding is unknown or not set")
)

// RequiredEnvVarsMissing says, that a required Environment Variable is not given.
type RequiredEnvVarsMissing struct {
	MissingVars []string
}

func (e *RequiredEnvVarsMissing) Error() string {
	return fmt.Sprintf("One or more required environment variables are missing\nThe missing variables are: %s", e.printMissingVars())
}

func (e *RequiredEnvVarsMissing) printMissingVars() string {
	return strings.Join(e.MissingVars, ", ")
}

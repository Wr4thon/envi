package errors

import (
	"fmt"
	"strings"
)

// RequiredEnvVarsMissing says, that a required Environment Variable is not given.
type RequiredEnvVarsMissing struct {
	missingVars []string
}

func (e RequiredEnvVarsMissing) Error() string {
	return fmt.Sprintf("One or more required environment variables are missing\nThe missing variables are: %s", e.printMissingVars())
}

func (e *RequiredEnvVarsMissing) printMissingVars() string {
	return strings.Join(e.missingVars, ", ")
}

// MissingVars is used to get a list of all missing variables.
func (e *RequiredEnvVarsMissing) MissingVars() []string {
	result := make([]string, len(e.missingVars))
	copy(result, e.missingVars)
	return result
}

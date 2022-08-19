package errors

// EnviError is the base error that can be used to get basic unwrap and cause
// functionality. Also provides a default `Error() string` method.
type EnviError struct {
	cause error
}

// WrapEnviError can be used to wrap an error inside an `EnviError`.
// When the cause is nil, nil will be returned.
func WrapEnviError(
	cause error,
) error {
	if cause == nil {
		return nil
	}

	return NewEnviError(cause)
}

// NewEnviError creates a new EnviError.
func NewEnviError(
	cause error,
) EnviError {
	return EnviError{
		cause: cause,
	}
}

// Error implements the error interface.
func (e EnviError) Error() string {
	return "\n\t" + e.cause.Error()
}

// Unwrap is used by the errors package
func (e EnviError) Unwrap() error {
	return e.Cause()
}

// Cause is used to get the cause of an error;
// This is used by the errors package.
func (e EnviError) Cause() error {
	return e.cause
}

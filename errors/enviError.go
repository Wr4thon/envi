package errors

type EnviError struct {
	cause error
}

func WrapEnviError(
	cause error,
) error {
	if cause == nil {
		return nil
	}

	return NewEnviError(cause)
}

func NewEnviError(
	cause error,
) EnviError {
	return EnviError{
		cause: cause,
	}
}

func (e EnviError) Error() string {
	return "\n\t" + e.cause.Error()
}

func (e EnviError) Unwrap() error {
	return e.Cause()
}

func (e EnviError) Cause() error {
	return e.cause
}

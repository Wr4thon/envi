package envi

import (
	"strconv"

	"github.com/pkg/errors"
)

func ValidatorBool(v Value) error {
	s := v.String()
	if s != "true" && s != "false" {
		return errors.Wrapf(ErrValidationFailed, "value '%s' is not a valid boolean value", s)
	}

	return nil
}

func ValidatorInt(v Value) error {
	s := v.String()
	if _, err := strconv.Atoi(s); err != nil {
		return errors.Wrapf(ErrValidationFailed, "value '%s' is not a valid integer value", s)
	}

	return nil
}

package envi

import "github.com/pkg/errors"

func ValidatorBool(v Value) error {
	s := v.String()
	if s != "true" && s != "false" {
		return errors.Wrapf(ErrValidationFailed, "value %s is not a valid boolean value", s)
	}

	return nil
}

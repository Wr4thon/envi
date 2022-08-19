package variables

import (
	"github.com/pkg/errors"

	"github.com/Clarilab/envi/v2/engine"
	enviErrors "github.com/Clarilab/envi/v2/errors"
)

type Variable struct {
	factory           func() interface{}
	validators        []func(value interface{}) error
	callback          func(interface{})
	autoValidateOnSet bool
	defaultValue      interface{}
	key               engine.Key

	value interface{}
}

func NewVariable[T any](
	key engine.Key,
	factory engine.Factory[*T],
	opts ...Opt,
) *Variable {
	v := &Variable{
		factory: func() interface{} {
			return factory()
		},
		key: key,
	}

	for i := range opts {
		opts[i](v)
	}

	return v
}

func (v *Variable) Load() error {
	v.value = v.defaultValue
	return nil
}

func (v *Variable) Instantiate() (interface{}, error) {
	val := v.factory()
	if err := v.Set(val); err != nil {
		return nil, errors.Wrapf(
			err,
			"error while setting variable value",
		)
	}

	return val, nil
}

func (v *Variable) Value() interface{} {
	return v.value
}

func (v *Variable) Key() engine.Key {
	return v.key
}

func (v *Variable) Validate(value interface{}) error {
	for i := range v.validators {
		if err := v.validators[i](value); err != nil {
			return errors.Wrap(
				enviErrors.WrapValidationError(
					err,
					v,
				),
				"error while validating value",
			)
		}
	}

	return nil
}

func (v *Variable) Set(value interface{}) error {
	if v.autoValidateOnSet {
		if err := v.Validate(value); err != nil {
			return errors.Wrapf(err, "validation of value failed")
		}
	}

	v.value = value

	return nil
}

package variables

import "github.com/Clarilab/envi/v2/engine"

// Callback is a generic function contract that can be used
// to receive a callback whenever a value changes.
type Callback[T any] func(newVal T)

// Opt can be used to set certain values in the Variable struct.
type Opt func(*Variable)

// WithValidator adds a validator to a variable
func WithValidator[T any](validator engine.Validator[T]) Opt {
	return func(v *Variable) {
		v.validators = append(v.validators, func(value interface{}) error {
			tVal, ok := value.(T)
			if !ok {
				return nil // wrong Type
			}

			return validator(tVal)
		})
	}
}

// WithValueChangedCallback registers a callback that is called when a
// value changes.
func WithValueChangedCallback[T any](callback Callback[T]) Opt {
	return func(v *Variable) {
		v.callback = func(newValue interface{}) {
			val, ok := newValue.(T)
			if !ok {
				// hmmm
				return
			}

			callback(val)
		}
	}
}

// WithAutoValidateOnSet sets the `autoValidateOnSet` flag to true.
// This activates the validation feature when a variable loads a value.
func WithAutoValidateOnSet() Opt {
	return func(v *Variable) {
		v.autoValidateOnSet = true
	}
}

// WithDefaultValue sets a default value for a variable.
func WithDefaultValue[T any](val T) Opt {
	return func(v *Variable) {
		v.defaultValue = val
	}
}

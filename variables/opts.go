package variables

import "github.com/Clarilab/envi/v2/engine"

type Callback[T any] func(newVal T)

type Opt func(*Variable)

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

func WithAutoValidateOnSet() Opt {
	return func(v *Variable) {
		v.autoValidateOnSet = true
	}
}

func WithDefaultValue[T any](val T) Opt {
	return func(v *Variable) {
		v.defaultValue = val
	}
}

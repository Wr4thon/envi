package envi

import (
	"github.com/Clarilab/envi/v2/engine"
	"github.com/Clarilab/envi/v2/variables"
)

// Opt is a function contract that is used to configure the Envi struct.
type Opt func(*Envi)

func providerFunc[T any](s T) func() T {
	return func() T {
		return s
	}
}

// WithContinueOnError let's you set a function that can be used to skip
// certain errors. For example when you try to load an optional file.
func WithContinueOnError(continueOnError func(error) bool) Opt {
	return func(e *Envi) {
		e.continueOnError = continueOnError
	}
}

// WithJSONFile creates an Opt that reads a file from the drive,
// unmarshalls it via json.Unmarshal and validates the result if needed.
func WithJSONFile[T any](
	filePath string,
	key engine.Key,
	factory engine.Factory[*T],
	opts ...variables.Opt,
) Opt {
	return WithVar(
		key,
		variables.NewJSONFileVariable(
			key,
			providerFunc(filePath),
			factory,
			opts...,
		),
	)
}

// WithYAMLFile creates an Opt that reads a file from the drive,
// unmarshalls it via yaml.Unmarshal and validates the result if needed.
func WithYAMLFile[T any](
	filePath string,
	key engine.Key,
	factory engine.Factory[*T],
	opts ...variables.Opt,
) Opt {
	return WithVar(
		key,
		variables.NewYAMLFileVariable(
			key,
			providerFunc(filePath),
			factory,
			opts...,
		),
	)
}

// WithVar can be used to register own variable types, sources, ... to specific
// keys.
func WithVar(key engine.Key, v engine.Var) Opt {
	return func(e *Envi) {
		if e.variables == nil {
			e.variables = make(map[engine.Key]engine.Var)
		}

		e.variables[key] = v
	}
}

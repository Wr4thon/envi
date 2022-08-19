package envi

import (
	"github.com/Clarilab/envi/v2/engine"
	"github.com/Clarilab/envi/v2/variables"
)

type Opt func(*Envi)

func WithContinueOnError(continueOnError func(error) bool) Opt {
	return func(e *Envi) {
		e.continueOnError = continueOnError
	}
}

func WithJSONFile[T any](filePath string, key engine.Key, factory engine.Factory[*T], opts ...variables.Opt) Opt {
	return WithVar(
		key,
		variables.NewJSONFileVariable(key, filePath, factory, opts...),
	)
}

func WithYAMLFile[T any](filePath string, key engine.Key, factory engine.Factory[*T], opts ...variables.Opt) Opt {
	return WithVar(
		key,
		variables.NewYAMLFileVariable(key, filePath, factory, opts...),
	)
}

func WithVar(key engine.Key, v engine.Var) Opt {
	return func(e *Envi) {
		if e.variables == nil {
			e.variables = make(map[engine.Key]engine.Var)
		}

		e.variables[key] = v
	}
}

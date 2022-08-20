package engine

// Factory is used to instantiate a variable.
type Factory[T any] func() T

// Validator is used to validate a configuration or value.
type Validator[T any] func(value T) error

// Var is the interface a custom variable needs to implement
// to be able to register it.
type Var interface {
	Load() error
	Key() Key
	Value() interface{}
}

type GenVar[T any] interface {
	Var
	GenValue() T
}

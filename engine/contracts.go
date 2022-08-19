package engine

type Factory[T any] func() T

type Validator[T any] func(value T) error

type Var interface {
	Load() error
	Value() interface{}
	Key() Key
}

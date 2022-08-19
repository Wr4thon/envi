package variables

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/Clarilab/envi/v2/engine"
	enviErrors "github.com/Clarilab/envi/v2/errors"
)

type unmarshal func([]byte, interface{}) error

// NewJSONFileVariable create a new instance of UnmarshalFileVariable
// which is configured to unmarshal a files content with the json unmarshaler.
func NewJSONFileVariable[T any](
	key engine.Key,
	filePath func() string,
	factory engine.Factory[*T],
	opts ...Opt,
) *UnmarshalFileVariable {
	return NewUnmarshalFileVariable(
		key,
		filePath,
		factory,
		json.Unmarshal,
		opts...,
	)
}

// NewYAMLFileVariable create a new instance of UnmarshalFileVariable
// which is configured to unmarshal a files content with the yaml unmarshaler.
func NewYAMLFileVariable[T any](
	key engine.Key,
	filePath func() string,
	factory engine.Factory[*T],
	opts ...Opt,
) *UnmarshalFileVariable {
	return NewUnmarshalFileVariable(
		key,
		filePath,
		factory,
		yaml.Unmarshal,
		opts...,
	)
}

// NewUnmarshalFileVariable can be used to add a UnmarshalFileVariable with a
// custom unmarshaler, that adheres to the function contract.
func NewUnmarshalFileVariable[T any](
	key engine.Key,
	filePath func() string,
	factory engine.Factory[*T],
	unmarshal unmarshal,
	opts ...Opt,
) *UnmarshalFileVariable {
	v := &UnmarshalFileVariable{
		Variable: NewVariable(
			key,
			factory,
			opts...,
		),
		unmarshal: unmarshal,
		filePath:  filePath,
	}

	return v
}

// UnmarshalFileVariable proved the functionality to load a file from the drive,
// unmarshal it with the provided marshaler and validates the value.
type UnmarshalFileVariable struct {
	*Variable

	filePath  func() string
	unmarshal unmarshal
}

// Load is used to load the value of the variable.
func (ufv *UnmarshalFileVariable) Load() error {
	data, err := loadFileContent(ufv)
	if err != nil {
		return errors.Wrapf(
			err,
			"error while loading file",
		)
	}

	value, err := ufv.Instantiate()
	if err != nil {
		return errors.Wrap(
			err,
			"error while instantiating variable",
		)
	}

	if err = ufv.unmarshal(data, value); err != nil {
		return errors.Wrap(
			enviErrors.WrapUnmarshalFileError(
				err,
				ufv,
			),
			"unable to unmarshal file content",
		)
	}

	return nil
}

// Path is used to get the path to the file this variable is operating on.
func (ufv *UnmarshalFileVariable) Path() string {
	return ufv.filePath()
}

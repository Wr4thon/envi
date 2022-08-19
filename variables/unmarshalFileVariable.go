package variables

import (
	"encoding/json"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/Clarilab/envi/v2/engine"
	enviErrors "github.com/Clarilab/envi/v2/errors"
)

type unmarshal func([]byte, interface{}) error

func NewJSONFileVariable[T any](key engine.Key, filePath string, factory engine.Factory[*T], opts ...Opt) *UnmarshalFileVariable {
	return NewUnmarshalFileVariable(key, filePath, factory, json.Unmarshal, opts...)
}

func NewYAMLFileVariable[T any](key engine.Key, filePath string, factory engine.Factory[*T], opts ...Opt) *UnmarshalFileVariable {
	return NewUnmarshalFileVariable(key, filePath, factory, yaml.Unmarshal, opts...)
}

func NewUnmarshalFileVariable[T any](key engine.Key, filePath string, factory engine.Factory[*T], unmarshal unmarshal, opts ...Opt) *UnmarshalFileVariable {
	v := &UnmarshalFileVariable{
		Variable:  NewVariable(key, factory, opts...),
		unmarshal: unmarshal,
		filePath:  filePath,
	}

	return v
}

type UnmarshalFileVariable struct {
	*Variable

	filePath  string
	unmarshal unmarshal
}

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
		return errors.Wrap(err, "error while instantiating variable")
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

func (ufv *UnmarshalFileVariable) Path() string {
	return ufv.filePath
}

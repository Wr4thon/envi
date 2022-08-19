package envi

import (
	"reflect"

	"github.com/Clarilab/envi/v2/engine"
	enviErrors "github.com/Clarilab/envi/v2/errors"

	"github.com/pkg/errors"
)

type Envi struct {
	variables       map[engine.Key]engine.Var
	continueOnError func(error) bool
}

func NewEnvi(opts ...Opt) Envi {
	envi := &Envi{}

	for i := range opts {
		opts[i](envi)
	}

	return *envi
}

func (e *Envi) LoadKey(k engine.Key) error {
	var v engine.Var
	var ok bool
	if v, ok = e.variables[k]; !ok {
		return enviErrors.NewKeyNotFoundError(k)
	}

	return e.load(k, v)
}

func (e *Envi) load(k engine.Key, v engine.Var) error {
	err := v.Load()
	if err != nil {
		err = errors.Wrapf(
			err,
			"error while loading variable \"%s\"",
			k,
		)
	}

	return err
}

func (e *Envi) Load() error {
	for k, v := range e.variables {
		if err := e.load(k, v); err != nil {
			if e.continueOnError != nil &&
				e.continueOnError(err) {
				continue
			}

			return err
		}
	}

	return nil
}

func (e *Envi) Get(key engine.Key, outPtr interface{}) error {
	variable, ok := e.variables[key]
	if !ok {
		return errors.Wrap(
			enviErrors.NewKeyNotFoundError(key),
			"requested key does not exist",
		)
	}

	return assignValueToPointer(variable.Value(), outPtr)
}

func (e *Envi) GetConfig(callback func(k engine.Key, val engine.Var) error) error {
	for k, v := range e.variables {
		if err := callback(k, v); err != nil {
			return errors.Wrap(err, "error while getting config")
		}
	}

	return nil
}

func assignValueToPointer(v interface{}, out interface{}) error {
	if out == nil {
		return errors.New("not a pointer")
	}

	optrT := reflect.ValueOf(out)
	if optrT.Kind() != reflect.Pointer ||
		optrT.IsNil() {
		return errors.New("not a pointer")
	}

	if v == nil {
		return errors.New("value not set yet")
	}

	var val reflect.Value = reflect.ValueOf(v)
	if !val.Type().ConvertibleTo(optrT.Type()) &&
		!val.Type().AssignableTo(optrT.Type()) {
		return errors.New("value not assignable")
	}

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	optrT.Elem().Set(val)

	return nil
}

package variables

import (
	"os"

	"github.com/pkg/errors"

	"github.com/Clarilab/envi/v2/engine"
	enviErrors "github.com/Clarilab/envi/v2/errors"
)

type pathVar interface {
	engine.Var
	Path() string
}

type FileVariable struct {
	*Variable
	path string
}

func NewFileVariable[T any](key engine.Key, path string, factory engine.Factory[*T], opts ...Opt) *FileVariable {
	return &FileVariable{
		path:     path,
		Variable: NewVariable(key, factory, opts...),
	}
}

func (fv *FileVariable) Load() error {
	return nil
}

func loadFileContent(fv pathVar) ([]byte, error) {
	data, err := os.ReadFile(fv.Path())
	if err != nil {
		return nil, errors.Wrapf(
			enviErrors.WrapFileError(
				err,
				fv,
			),
			"error while reading file content",
		)
	}

	return data, nil
}

func (fv *FileVariable) Path() string {
	return fv.path
}

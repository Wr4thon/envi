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

// FileVariable wraps a Variable. Can be used to get a specific file into RAM.
type FileVariable struct {
	*Variable
	path string
}

// NewFileVariable creates a new FileVariable.
func NewFileVariable[T any](
	key engine.Key,
	path string,
	factory engine.Factory[*T],
	opts ...Opt,
) *FileVariable {
	return &FileVariable{
		path:     path,
		Variable: NewVariable(key, factory, opts...),
	}
}

// Load loads the value
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

// Path is used to access the information which file is supposed to be accessed.
func (fv *FileVariable) Path() string {
	return fv.path
}

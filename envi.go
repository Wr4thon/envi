package envi

import (
	"os"

	"github.com/pkg/errors"
)

var (
	envVariableConfig VariableConfig = VariableConfig{Source: SourceEnvironment}
)

type (
	Envi interface {
		// ToEnv writes all key-value pairs to the environment.
		ToEnv()

		// ToMap returns a map, containing all key-value pairs.
		ToMap() Values

		// Load all configured variables
		Load() (Values, error)
	}

	envi struct {
		loadedVars map[string]Value
		variables  []variable
	}
)

func NewEnvi(opts ...Opt) Envi {
	e := &envi{
		loadedVars: make(map[string]Value),
	}

	for opt := range opts {
		opts[opt](e)
	}

	return e
}

func (envi *envi) Load() (Values, error) {
	var missingRequiredFields []string

	for _, variable := range envi.variables {
		var value []byte
		var err error
		switch variable.conf.Source & 0x0f {
		case SourceEnvironment:
			value = []byte(os.Getenv(variable.key))
		case SourceLiteral:
			value = []byte(variable.conf.Value)
		case SourceFile:
			if value, err = os.ReadFile(variable.conf.Path); err != nil {
				if !variable.required {
					break
				}
				return nil, errors.Wrapf(err, "error while reading file %s", variable.conf.Path)
			}

			if variable.conf.ValuesToMap {
				values := map[string]string{}
				switch variable.conf.Encoding {
				case EncodingJSON:
					err = Value(value).UnmarshalJson(&values)
				case EncodingYAML:
					err = Value(value).UnmarshalYaml(&values)
				default:
					err = errors.Wrapf(ErrUnknownEncoding, "encoding '%d' of variable %s is unknown", variable.conf.Encoding, variable.key)
				}

				if err != nil {
					if !variable.required {
						break
					}
					return nil, err
				}

				for key := range values {
					envi.loadedVars[key] = Value(values[key])
				}

				continue
			}
		default:
			return nil, errors.Wrapf(ErrUnknownSource, "variable '%s' has an unknown source: %d", variable.key, variable.conf.Source)
		}

		if len(value) == 0 {
			if variable.required {
				missingRequiredFields = append(missingRequiredFields, variable.key)
				continue
			}

			value = []byte(variable.conf.Value)
		}

		if variable.conf.Validator != nil {
			if err = variable.conf.Validator(value); err != nil {
				return nil, errors.Wrapf(err, "validation for variable '%s' failed", variable.key)
			}
		}

		envi.loadedVars[variable.key] = value
	}

	if len(missingRequiredFields) > 0 {
		return nil, &RequiredEnvVarsMissing{
			MissingVars: missingRequiredFields,
		}
	}

	return envi.loadedVars, nil
}

func (envi *envi) ToEnv() {
	for key := range envi.loadedVars {
		os.Setenv(key, envi.loadedVars[key].String())
	}
}

func (envi *envi) ToMap() Values {
	return envi.loadedVars
}

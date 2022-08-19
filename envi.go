package envi

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Envi interface {
	// FromMap loads the given key-value pairs and loads them into the local map.
	FromMap(map[string]string)

	// LoadEnv loads the given keys from environment.
	LoadEnv(vars ...string)

	// LoadFile loads a string value under given key from a file.
	LoadFile(key, filePath string) error

	// LoadJSON loads key-value pairs from one or many json blobs.
	LoadJSON(...[]byte) error

	// LoadJSONFiles loads key-value pairs from one or more json files.
	LoadJSONFiles(...string) error

	// LoadYAML loads key-value pairs from one or many yaml blobs.
	LoadYAML(...[]byte) error

	// LoadYAMLFiles loads key-value pairs from one or more yaml files.
	LoadYAMLFiles(...string) error

	// EnsureVars checks, if all given keys have a non-empty value.
	EnsureVars(...string) error

	// ToEnv writes all key-value pairs to the environment.
	ToEnv()

	// ToMap returns a map, containing all key-value pairs.
	ToMap() map[string]string
}

type envi struct {
	loadedVars map[string]string
}

func NewEnvi() Envi {
	return &envi{
		loadedVars: make(map[string]string),
	}
}

func (envi *envi) FromMap(vars map[string]string) {
	for key := range vars {
		envi.loadedVars[key] = vars[key]
	}
}

func (envi *envi) LoadEnv(vars ...string) {
	for _, key := range vars {
		envi.loadedVars[key] = os.Getenv(key)
	}
}

func (envi *envi) LoadFile(key, filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Wrap(
			fileError(
				filePath,
				err,
			),
			"failed to load file",
		)
	}

	envi.loadedVars[key] = string(data)

	return nil
}

func (envi *envi) LoadJSONFiles(paths ...string) error {
	return errors.Wrap(
		envi.loadFiles(
			envi.LoadJSON,
			paths...,
		),
		"error while loading JSON file",
	)
}

func (envi *envi) LoadJSON(data ...[]byte) error {
	return errors.Wrap(
		envi.load(
			func(data []byte, out interface{}) error {
				return errors.Wrap(
					json.Unmarshal(
						data,
						&out,
					),
					"error while unmarshaling data from json",
				)
			},
			data...,
		),
		"error while loading JSON file",
	)
}

func (envi *envi) LoadYAMLFiles(paths ...string) error {
	return errors.Wrap(
		envi.loadFiles(
			envi.LoadYAML,
			paths...,
		),
		"error while loading YAML file",
	)
}

func (envi *envi) LoadYAML(data ...[]byte) error {
	return errors.Wrap(
		envi.load(
			func(data []byte, out interface{}) error {
				const errMsg = "error while unmarshaling data from yaml"
				var err error
				if err = yaml.Unmarshal(data, out); err == nil {
					return nil
				}

				typeErr := &yaml.TypeError{}
				switch {
				case errors.As(err, &typeErr):
					return errors.Wrap(
						unmarshalError(
							err,
						),
						errMsg,
					)
				default:
					return errors.Wrap(
						err,
						"unexpected error while unmarshaling data from yaml",
					)
				}
			},
			data...,
		),
		"error while loading YAML file",
	)
}

func (envi *envi) loadFiles(
	unmarshalFile func(...[]byte) error,
	paths ...string,
) error {
	var errMsg = "error while loading file"

	for i := range paths {
		path := paths[i]

		fileContent, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrap(
				fileError(
					path,
					err,
				),
				errMsg,
			)
		}

		err = unmarshalFile(fileContent)
		if err != nil {
			return errors.Wrap(
				unmarshalFileError(
					path,
					err,
				),
				errMsg,
			)
		}
	}

	return nil
}

func (envi *envi) load(
	unmarshal func([]byte, interface{}) error,
	blobs ...[]byte,
) error {
	for i := range blobs {
		var decoded map[string]string

		err := unmarshal(blobs[i], &decoded)
		if err != nil {
			return errors.Wrap(
				err,
				"failed to unmarshal",
			)
		}

		for key := range decoded {
			envi.loadedVars[key] = decoded[key]
		}
	}

	return nil
}

func (envi *envi) EnsureVars(requiredVars ...string) error {
	var missingVars []string

	for _, key := range requiredVars {
		if envi.loadedVars[key] == "" {
			missingVars = append(missingVars, key)
		}
	}

	if len(missingVars) > 0 {
		return &RequiredEnvVarsMissing{missingVars: missingVars}
	}

	return nil
}

func (envi *envi) ToEnv() {
	for key := range envi.loadedVars {
		os.Setenv(key, envi.loadedVars[key])
	}
}

func (envi *envi) ToMap() map[string]string {
	return envi.loadedVars
}

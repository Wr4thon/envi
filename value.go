package envi

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type Value []byte

func (v Value) String() string {
	return string(v)
}

func (v Value) UnmarshalJson(val interface{}) error {
	return json.Unmarshal([]byte(v), val)
}

func (v Value) UnmarshalYaml(val interface{}) error {
	return yaml.Unmarshal([]byte(v), val)
}

func (v Value) Bool() bool {
	return v.String() == "true"
}

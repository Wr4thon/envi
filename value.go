package envi

import (
	"encoding/json"
	"strconv"

	"github.com/mitchellh/mapstructure"
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

func (value Value) Bool() bool {
	return value.String() == "true"
}

func (value Value) Int() (int, error) {
	return strconv.Atoi(value.String())
}

type Values map[string]Value

func (values Values) Decode(result interface{}) error {
	// var decodeMap map[string]interface{} = map[string]interface{}{}
	// for key := range values {
	// 	decodeMap[key] = values[key].String()
	// }

	return mapstructure.Decode(values, &result)
}

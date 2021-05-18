package envi

const (
	SourceEnvironment Source = 0x01 << iota
	SourceLiteral     Source = 0x01 << iota
	SourceFile        Source = 0x01 << iota
)

const (
	EncodingJSON Encoding = 0x01 << iota
	EncodingYAML Encoding = 0x01 << iota
)

type (
	Source   int
	Encoding uint8

	VariableValueValidator func(Value) error

	VariableConfig struct {
		Source      Source
		Encoding    Encoding
		Value       string
		Path        string
		Validator   VariableValueValidator
		ValuesToMap bool
	}

	variable struct {
		key      string
		required bool
		conf     VariableConfig
	}
)

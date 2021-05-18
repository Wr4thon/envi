package envi

type Opt func(*envi)

func WithDefaultValues(values map[string]string) Opt {
	return func(e *envi) {
		for key := range values {
			e.loadedVars[key] = Value(values[key])
		}
	}
}

func WithDefaultValueEnvVariable(key string, required bool, defaultValue string) Opt {
	cfg := envVariableConfig
	cfg.Value = defaultValue
	return WithVariableConfig(key, required, cfg)
}

func WithEnvVariable(key string) Opt {
	return WithVariableConfig(key, false, envVariableConfig)
}

func WithRequiredEnvVariable(key string) Opt {
	return WithVariableConfig(key, true, envVariableConfig)
}

func WithFile(key string, required bool, path string) Opt {
	return WithVariableConfig(
		key,
		required,
		VariableConfig{
			Source: SourceFile,
			Path:   path,
		},
	)
}

func WithFileToMap(path string, encoding Encoding, required bool) Opt {
	return WithVariableConfig(
		path,
		required,
		VariableConfig{
			ValuesToMap: true,
			Source:      SourceFile,
			Encoding:    encoding,
			Path:        path,
		},
	)
}

func WithVariableConfig(key string, required bool, conf VariableConfig) Opt {
	return func(e *envi) {
		variable := variable{
			key:      key,
			required: required,
			conf:     conf,
		}

		e.variables = append(e.variables, variable)
	}
}

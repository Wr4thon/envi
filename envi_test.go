package envi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	editor    = "EDITOR"
	pager     = "PAGER"
	home      = "HOME"
	mail      = "MAIL"
	url       = "URL"
	file      = "FILE"
	someValue = "SOMEVALUE"
)

type data struct {
	URL    string `json:"URL"`
	EDITOR string `json:"EDITOR"`
	HOME   string `json:"HOME"`
}

func Test_FromMap(t *testing.T) {
	payload := make(map[string]string)
	payload[editor] = "vim"
	payload[pager] = "less"

	e := NewEnvi(WithDefaultValues(payload))

	assert.Len(t, e.ToMap(), 2)
}

func Test_LoadEnv(t *testing.T) {
	e := NewEnvi(
		WithRequiredEnvVariable(editor),
		WithRequiredEnvVariable(pager),
		WithRequiredEnvVariable(home),
	)

	_, err := e.Load()

	assert.NoError(t, err)
	assert.Len(t, e.ToMap(), 3)
}

func Test_LoadJSONFromFile(t *testing.T) {
	t.Run("load nothing", func(t *testing.T) {
		e := NewEnvi()

		_, err := e.Load()

		assert.NoError(t, err)
		assert.Len(t, e.ToMap(), 0)
	})

	t.Run("a valid json file", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/valid1.json", EncodingJSON, true),
		)

		_, err := e.Load()

		assert.NoError(t, err)
		assert.Len(t, e.ToMap(), 3)
	})

	t.Run("2 valid json files", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/valid1.json", EncodingJSON, true),
			WithFileToMap("testdata/valid2.json", EncodingJSON, true),
		)

		_, err := e.Load()

		assert.NoError(t, err)
		assert.Len(t, e.ToMap(), 4)
	})

	t.Run("an invalid json file", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/invalid.json", EncodingJSON, true),
		)

		_, err := e.Load()

		assert.Error(t, err)
	})

	t.Run("a missing file", func(t *testing.T) {
		e := NewEnvi(
			WithFile("iDoNotExist", true, "testdata/idontexist.json"),
		)

		_, err := e.Load()

		assert.Error(t, err)
	})

	t.Run("an existing and a missing file", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/valid1.json", EncodingJSON, true),
			WithFileToMap("testdata/idontexist.json", EncodingJSON, false),
		)

		_, err := e.Load()

		assert.NoError(t, err)
	})

	t.Run("an existing and a missing required file", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/valid1.json", EncodingJSON, true),
			WithFileToMap("testdata/idontexist.json", EncodingJSON, true),
		)

		_, err := e.Load()

		assert.Error(t, err)
	})
}

func Test_LoadYAMLFomFile(t *testing.T) {
	t.Run("load nothing", func(t *testing.T) {
		e := NewEnvi()

		_, err := e.Load()

		assert.NoError(t, err)
		assert.Len(t, e.ToMap(), 0)
	})

	t.Run("a valid yaml file", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/valid1.yaml", EncodingYAML, true),
		)

		_, err := e.Load()

		assert.NoError(t, err)
		assert.Len(t, e.ToMap(), 3)
	})

	t.Run("2 valid yaml files", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/valid1.yaml", EncodingYAML, true),
			WithFileToMap("testdata/valid2.yaml", EncodingYAML, true),
		)

		_, err := e.Load()

		assert.NoError(t, err)
		assert.Len(t, e.ToMap(), 4)
	})

	t.Run("an invalid yaml file", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/invalid.yaml", EncodingYAML, true),
		)

		_, err := e.Load()

		assert.Error(t, err)
	})

	t.Run("a missing file", func(t *testing.T) {
		e := NewEnvi(
			WithFile("iDoNotExist", true, "testdata/idontexist.yaml"),
		)

		_, err := e.Load()

		assert.Error(t, err)
	})

	t.Run("an existing and a missing file", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/valid1.yaml", EncodingYAML, true),
			WithFileToMap("testdata/idontexist.yaml", EncodingYAML, false),
		)

		_, err := e.Load()

		assert.NoError(t, err)
	})

	t.Run("an existing and a missing required file", func(t *testing.T) {
		e := NewEnvi(
			WithFileToMap("testdata/valid1.yaml", EncodingYAML, true),
			WithFileToMap("testdata/idontexist.yaml", EncodingYAML, true),
		)

		_, err := e.Load()

		assert.Error(t, err)
	})
}

func Test_EnsureVars(t *testing.T) {
	t.Run("all ensured vars are present", func(t *testing.T) {
		payload := make(map[string]string)
		payload[editor] = "vim"
		payload[pager] = "less"

		e := NewEnvi(
			WithDefaultValues(payload),
			WithRequiredEnvVariable(editor),
			WithRequiredEnvVariable(pager),
		)

		_, err := e.Load()

		assert.NoError(t, err)
	})

	t.Run("one ensured var is missing", func(t *testing.T) {
		payload := make(map[string]string)
		payload[editor] = "vim"
		payload[pager] = "less"

		e := NewEnvi(
			WithDefaultValues(payload),
			WithRequiredEnvVariable(editor),
			WithRequiredEnvVariable(pager),
			WithRequiredEnvVariable(someValue),
		)

		_, err := e.Load()

		assert.Error(t, err)
	})

	t.Run("all ensured vars are missing", func(t *testing.T) {
		payload := make(map[string]string)
		payload[editor] = "vim"
		payload[pager] = "less"

		e := NewEnvi(
			WithDefaultValues(payload),
			WithRequiredEnvVariable(home),
			WithRequiredEnvVariable(mail),
			WithRequiredEnvVariable(url),
		)

		_, err := e.Load()

		assert.Error(t, err)
	})
}

func Test_ToEnv(t *testing.T) {
	payload := make(map[string]string)
	payload["SCHURZLPURZ"] = "yes, indeed"

	e := NewEnvi(
		WithDefaultValues(payload),
	)

	e.ToEnv()

	assert.Equal(t, "yes, indeed", os.Getenv("SCHURZLPURZ"))
}

func Test_ToMap(t *testing.T) {
	payload := make(map[string]string)
	payload[editor] = "vim"
	payload[pager] = "less"

	e := NewEnvi(
		WithDefaultValues(payload),
	)

	vars := e.ToMap()

	assert.Len(t, vars, 2)
}

func Test_LoadFile(t *testing.T) {
	t.Run("no file", func(t *testing.T) {
		e := NewEnvi(
			WithFile(file, true, ""),
		)

		_, err := e.Load()

		assert.Error(t, err)
		assert.Len(t, e.ToMap(), 0)
	})

	t.Run("file with string content", func(t *testing.T) {
		e := NewEnvi(
			WithVariableConfig(file, true, VariableConfig{
				Source: SourceFile,
				Path:   filepath.Join("testdata/valid.txt"),
			}),
		)

		_, err := e.Load()

		assert.NoError(t, err)
		assert.Len(t, e.ToMap(), 1)
		assert.Equal(t, "valid string", e.ToMap()["FILE"].String())
	})
}

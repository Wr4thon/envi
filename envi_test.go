package envi_test

import (
	"errors"
	"testing"

	"github.com/Clarilab/envi/v2"
	"github.com/Clarilab/envi/v2/engine"
	enviErrors "github.com/Clarilab/envi/v2/errors"
	"github.com/stretchr/testify/require"
)

func Test_Load(t *testing.T) {
	type (
		testCase struct {
			envi   envi.Envi
			act    func(t *testing.T, envi *testCase)
			assert func(t *testing.T, envi *testCase)

			err error
		}

		Valid1 struct {
			Editor string `json:"EDITOR"`
			Home   string `json:"HOME"`
			Url    string `json:"URL"`
		}
	)

	var (
		valid1JSON   engine.Key = "valid1.json"
		fileNotFound engine.Key = "fileNotFound"
	)

	sFunc := func(s string) func() string { return func() string { return s } }

	var (
		valid1JSONPath   = sFunc("testdata/" + valid1JSON.Value())
		fileNotFoundPath = sFunc(fileNotFound.Value())
	)

	var (
		nopFactory    engine.Factory[*Valid1] = nil
		valid1Factory engine.Factory[*Valid1] = func() *Valid1 { return &Valid1{} }
	)

	var (
		defaultAct = func(t *testing.T, tc *testCase) { tc.err = tc.envi.Load() }
	)

	tt := map[string]*testCase{
		valid1JSON.Value(): {
			envi: envi.NewEnvi(
				envi.WithJSONFile(valid1JSONPath, valid1JSON, valid1Factory),
			),
			assert: func(t *testing.T, tc *testCase) {
				var res Valid1
				require.NoError(t, tc.envi.Get(valid1JSON, &res))
				require.Equal(t, "emacs", res.Editor)
			},
		},
		fileNotFound.Value(): {
			envi: envi.NewEnvi(
				envi.WithJSONFile(fileNotFoundPath, fileNotFound, nopFactory),
			),
			act: func(t *testing.T, tc *testCase) {
				tc.err = tc.envi.Load()
			},
			assert: func(t *testing.T, tc *testCase) {
				var fileError enviErrors.FileError
				require.ErrorAs(t, tc.err, &fileError)
				require.ErrorIs(t, fileError, enviErrors.ErrMissingFile)
				require.Equal(t, fileNotFound, fileError.Key())
				require.Equal(t, fileNotFoundPath(), fileError.FilePath())
			},
		},
		fileNotFound.Value() + "_and_" + valid1JSON.Value() + "_skipOnError": {
			envi: envi.NewEnvi(
				envi.WithJSONFile(valid1JSONPath, valid1JSON, valid1Factory),
				envi.WithJSONFile(fileNotFoundPath, fileNotFound, nopFactory),
				envi.WithContinueOnError(func(err error) bool {
					var fileError enviErrors.FileError
					return errors.As(err, &fileError) &&
						fileError.Key() == fileNotFound &&
						errors.Is(err, enviErrors.ErrMissingFile)
				}),
			),
			assert: func(t *testing.T, tc *testCase) {
				var res Valid1
				require.NoError(t, tc.envi.Get(valid1JSON, &res))
				require.Equal(t, "emacs", res.Editor)
			},
		},
	}

	for i := range tt {
		tc := tt[i]
		t.Run(i, func(t *testing.T) {
			t.Parallel()

			act := tc.act
			if act == nil {
				act = defaultAct
			}

			act(t, tc)

			if tc.assert == nil {
				return
			}

			tc.assert(t, tc)
		})
	}
}

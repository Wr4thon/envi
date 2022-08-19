package envi_test

import (
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
		valid1       engine.Key = "valid1.json"
		fileNotFound engine.Key = "fileNotFound"
	)

	var (
		valid1Path       string = "testdata/valid1.json"
		fileNotFoundPath string = "fileNotFound"
	)

	var (
		nopFactory    engine.Factory[*Valid1] = nil
		valid1Factory engine.Factory[*Valid1] = func() *Valid1 { return &Valid1{} }
	)

	var (
		defaultAct = func(t *testing.T, tc *testCase) { tc.err = tc.envi.Load() }
	)

	tt := map[string]*testCase{
		valid1.Value(): {
			envi: envi.NewEnvi(
				envi.WithJSONFile(valid1Path, valid1, valid1Factory),
			),
			assert: func(t *testing.T, tc *testCase) {
				var res Valid1
				require.NoError(t, tc.envi.Get(valid1, &res))
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
				require.Equal(t, fileNotFoundPath, fileError.FilePath())
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

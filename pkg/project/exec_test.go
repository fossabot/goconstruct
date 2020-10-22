package project

import (
	"crypto/sha256"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadDynamicConfig(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     map[string]interface{}
		wantErr  bool
	}{
		{
			name:     "multi-line",
			filename: filepath.Join("test-fixtures", "nesteddir", "variables.toml"),
			want: map[string]interface{}{
				"foo": "foo",
				"bar": "bar",
				"baz": "baz",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := readDynamicConfig(tt.filename)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("readDynamicConfig() mismatch (-want +got):\n%s", diff)
			}

			assertErr := func(x error, b bool) bool {
				return (x == nil && b == false) || (x != nil && b == true)
			}

			if ok := assertErr(gotErr, tt.wantErr); !ok {
				t.Errorf("readDynamicConfig() \n\twantErr: '%+v'\n\tgotErr: '%+v'", tt.wantErr, gotErr)
			}
		})
	}
}

func TestCopyDir(t *testing.T) {
	for _, tt := range []struct {
		name       string
		inputFile  string
		outputFile string
		wantErr    bool
	}{
		{
			name:       "nested-dir",
			inputFile:  filepath.Join("test-fixtures", "nesteddir"),
			outputFile: t.TempDir(),
			wantErr:    false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := copyDir(tt.outputFile, tt.inputFile)

			assertErr := func(x error, b bool) bool {
				return (x == nil && b == false) || (x != nil && b == true)
			}

			if ok := assertErr(gotErr, tt.wantErr); !ok {
				t.Errorf("copyDir() \n\twantErr: '%+v'\n\tgotErr: '%+v'", tt.wantErr, gotErr)
			}

			if sha256.Sum256(tt.inputFile) != sha256.Sum256(tt.outputFile) {
				t.Fatalf("Input and output files do not match\n"+
					"Input:\n%s\nOutput:\n%s\n", tt.inputFile, tt.outputFile)
			}
		})
	}

}

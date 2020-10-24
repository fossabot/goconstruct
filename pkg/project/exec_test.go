package project

import (
	"crypto/sha256"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func makeTestConfig(t *testing.T, b []byte) (string, func()) {
	tmpdir, err := ioutil.TempDir("", "make-test-file")
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(tmpdir, "variables.toml"), b, 0644)
	if err != nil {
		t.Fatal(err)
	}
	return filepath.Join(tmpdir, "variables.toml"), func() { os.RemoveAll(tmpdir) }
}

func TestReadDynamicConfig(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "multi-line",
			input: []byte(`
foo = "foo"
bar = "bar"
`),
			want: map[string]interface{}{
				"foo": "foo",
				"bar": "bar",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testConfig, cleanupFn := makeTestConfig(t, tt.input)
			defer cleanupFn()
			got, gotErr := readDynamicConfig(testConfig)

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
	tmpdir, err := ioutil.TempDir("", "copy-dir-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	tmplDir := filepath.Join(tmpdir, "templates")
	err = os.Mkdir(tmplDir, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	tmplSubDir := filepath.Join(tmplDir, "sub-dir")
	err = os.Mkdir(tmplSubDir, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(tmplSubDir, "main.txt"), []byte("hello"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	targetDir := filepath.Join(tmpdir, "target")
	os.Mkdir(targetDir, os.ModePerm)

	err = copyDir(targetDir, tmplDir)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = os.Lstat(filepath.Join(targetDir, "sub-dir", "main.txt")); os.IsNotExist(err) {
		t.Fatal("target sub-dir/main.go was not created")
	}

	in, err := ioutil.ReadFile(filepath.Join(tmplDir, "sub-dir", "main.txt"))
	if err != nil {
		t.Fatal("could not read file template sub-dir/main.txt")
	}
	out, err := ioutil.ReadFile(filepath.Join(targetDir, "sub-dir", "main.txt"))
	if err != nil {
		t.Fatal("could not read file target sub-dir/main.txt")
	}

	if sha256.Sum256(in) != sha256.Sum256(out) {
		t.Fatalf("Input and output files do not match\n"+
			"Input:\n%s\nOutput:\n%s\n", string(in), string(out))
	}
}

func TestRenderDir(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "render-dir-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	tmplDir := filepath.Join(tmpdir, "templates")
	err = os.Mkdir(tmplDir, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	tmplSubDir := filepath.Join(tmplDir, "sub-dir")
	err = os.Mkdir(tmplSubDir, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(tmplSubDir, "main.txt"), []byte(`
{{{.Foo}}}
hello
{{{.Bar}}}`), 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = renderDir(tmplDir, map[string]interface{}{
		"Foo": "foo",
		"Bar": "bar",
	})
	if err != nil {
		t.Fatal(err)
	}

	actual, err := ioutil.ReadFile(filepath.Join(tmplSubDir, "main.txt"))
	if err != nil {
		t.Fatal("could not read file template sub-dir/main.txt")
	}

	expected := []byte(`
foo
hello
bar`)

	if sha256.Sum256(actual) != sha256.Sum256(expected) {
		t.Fatalf("Input and output files do not match\n"+
			"Actual:\n%s\nExpected:\n%s\n", string(actual), string(expected))
	}
}

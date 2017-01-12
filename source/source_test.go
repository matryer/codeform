package source_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/matryer/codeform/source"
	"github.com/matryer/is"
)

// lookup is a source lookup that specifies a
// timeout on http requests.
var lookup = source.Lookup{
	Client: http.Client{
		Timeout: 2 * time.Second,
	},
}

func TestReader(t *testing.T) {
	is := is.New(t)
	s := source.Reader("verbatim.go", strings.NewReader(`This is a verbatim source file.`))
	defer s.Close()
	is.Equal(s.IsDir, false)
	is.Equal(s.Path, "verbatim.go")
	b, err := ioutil.ReadAll(s)
	is.NoErr(err)
	is.True(strings.Contains(string(b), `This is a verbatim source file.`))
}

func TestLocal(t *testing.T) {
	is := is.New(t)
	s, err := lookup.Get("./testdata/source.go")
	is.NoErr(err)
	defer s.Close()
	is.Equal(s.IsDir, false)
	is.Equal(s.Path, "./testdata/source.go")
	b, err := ioutil.ReadAll(s)
	is.NoErr(err)
	is.True(strings.Contains(string(b), `This is a local source file.`))
}

func TestTemplate(t *testing.T) {
	is := is.New(t)
	s, err := lookup.Get("template:inspect/packages")
	is.NoErr(err)
	defer s.Close()
	is.Equal(s.IsDir, false)
	is.True(strings.HasSuffix(s.Path, "github.com/matryer/codeform-templates/inspect/packages.tpl"))
	b, err := ioutil.ReadAll(s)
	is.NoErr(err)
	is.True(strings.Contains(string(b), `{{range .Packages}}`))
}

func TestDefault(t *testing.T) {
	is := is.New(t)
	s, err := lookup.Get("default")
	is.NoErr(err)
	defer s.Close()
	is.Equal(s.IsDir, false)
	is.True(strings.HasSuffix(s.Path, "github.com/matryer/codeform/source/default/source.go"))
	b, err := ioutil.ReadAll(s)
	is.NoErr(err)
	is.True(strings.Contains(string(b), `package defaultsource`))
}

func TestURL(t *testing.T) {
	is := is.New(t)
	s, err := lookup.Get("https://raw.githubusercontent.com/matryer/drop-test/master/greet.go")
	is.NoErr(err)
	defer s.Close()
	is.Equal(s.IsDir, false)
	is.Equal(s.Path, "https://raw.githubusercontent.com/matryer/drop-test/master/greet.go")
	b, err := ioutil.ReadAll(s)
	is.NoErr(err)
	is.True(strings.Contains(string(b), `func Greet(name string) string`))
}

func TestGoGet(t *testing.T) {
	is := is.New(t)
	s, err := lookup.Get("github.com/matryer/drop-test/explicit")
	is.NoErr(err)
	defer s.Close()
	is.Equal(s.IsDir, true)
	is.True(strings.HasSuffix(s.Path, `/src/github.com/matryer/drop-test/explicit`))
}

// TestGopath tests that the local GOPATH is checked.
func TestGoPath(t *testing.T) {
	is := is.New(t)
	originalGopath := os.Getenv("GOPATH")
	defer func() {
		os.Setenv("GOPATH", originalGopath)
	}()
	tmpgopath := filepath.Join(os.TempDir(), "codeform-test-gopath")
	tmpfiledir := filepath.Join(tmpgopath, "src/github.com/matryer/codeform-test")
	tmpfile := filepath.Join(tmpfiledir, "file.go")
	err := os.MkdirAll(tmpfiledir, 0777)
	is.NoErr(err)
	defer os.RemoveAll(tmpgopath)
	err = ioutil.WriteFile(tmpfile, []byte(`test file from GOPATH`), 0777)
	is.NoErr(err)
	os.Setenv("GOPATH", tmpgopath)
	s, err := lookup.Get("github.com/matryer/codeform-test/file.go")
	is.NoErr(err)
	defer s.Close()
	is.Equal(s.IsDir, false)
	is.Equal(s.Path, tmpfile)
	contentb, err := ioutil.ReadAll(s)
	is.NoErr(err)
	is.Equal(string(contentb), `test file from GOPATH`)
}

// TestGopath tests that the local GOROOT is checked.
func TestGoRoot(t *testing.T) {
	is := is.New(t)
	originalGoRoot := os.Getenv("GOROOT")
	defer func() {
		os.Setenv("GOROOT", originalGoRoot)
	}()
	tmpgoroot := filepath.Join(os.TempDir(), "codeform-test-goroot")
	tmpfiledir := filepath.Join(tmpgoroot, "src/github.com/matryer/codeform-test")
	tmpfile := filepath.Join(tmpfiledir, "file.go")
	err := os.MkdirAll(tmpfiledir, 0777)
	is.NoErr(err)
	defer os.RemoveAll(tmpgoroot)
	err = ioutil.WriteFile(tmpfile, []byte(`test file from GOROOT`), 0777)
	is.NoErr(err)
	os.Setenv("GOROOT", tmpgoroot)
	s, err := lookup.Get("github.com/matryer/codeform-test/file.go")
	is.NoErr(err)
	defer s.Close()
	is.Equal(s.IsDir, false)
	is.Equal(s.Path, tmpfile)
	contentb, err := ioutil.ReadAll(s)
	is.NoErr(err)
	is.Equal(string(contentb), `test file from GOROOT`)
}

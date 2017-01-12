// Package source provides a way to refer to source files.
// Valid sources include:
//     /local/path/to/file.go
//     /local/path/to/package
//     github.com/matryer/package
//     github.com/matryer/package/specific-file.go
//     https://domain.com/path/to/file
package source

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Source represents the source of a code file or
// package.
// Sources must always be closed.
type Source struct {
	readOnce   sync.Once
	readCloser io.ReadCloser
	closer     func() error
	Path       string
	IsDir      bool
}

// ErrNotFound indicates the source file could not be found.
type ErrNotFound string

func (e ErrNotFound) Error() string {
	return "not found: " + string(e)
}

// Lookup provides source lookup capabilities.
type Lookup struct {
	Client http.Client
}

// DefaultLookup is a default Lookup instance.
var DefaultLookup = &Lookup{}

// Reader gets a Source for the specified io.Reader.
func Reader(path string, r io.Reader) *Source {
	s := &Source{
		Path:       path,
		IsDir:      false,
		readCloser: ioutil.NopCloser(r),
	}
	return s
}

// Get gets the file at the given source.
// Valid sources include local paths, URLs, or go gettable
// paths.
// Sources must be closed.
// Valid sources include:
//     /local/path/to/file.go
//     /local/path/to/package
//     github.com/matryer/package
//     github.com/matryer/package/specific-file.go
//     https://domain.com/path/to/file
//
// Special strings
//
// The string "default" will be taken to mean the default source file which can be
// found in default/source.go.
// A special template prefix will indicate the template
// is hosted in the official https://github.com/matryer/codeform-templates repository.
//     template:testing/mocking/mock
func (l *Lookup) Get(src string) (*Source, error) {

	// handle special cases
	if src == "default" {
		src = "https://raw.githubusercontent.com/matryer/codeform/master/source/default/source.go"
	}
	if strings.HasPrefix(src, "template:") {
		src = fmt.Sprintf("https://raw.githubusercontent.com/matryer/codeform-templates/master/%s.tpl", strings.TrimPrefix(src, "template:"))
	}

	if s, err := tryLocal(l, src); err == nil {
		return s, nil
	}
	if s, err := tryURL(l, src); err == nil {
		return s, nil
	}
	if s, err := tryGoPath(l, src); err == nil {
		return s, nil
	}
	if s, err := tryGoRoot(l, src); err == nil {
		return s, nil
	}
	if s, err := tryGoGet(l, src); err == nil {
		return s, nil
	}
	return nil, ErrNotFound(src)
}

// MustLocal gets a local source.
// Panics if the local source is invalid.
func MustLocal(src string) *Source {
	s, err := tryLocal(nil, src)
	if err != nil {
		panic("codeform/source: " + err.Error())
	}
	return s
}

func tryLocal(l *Lookup, src string) (*Source, error) {
	info, err := os.Stat(src)
	if err != nil {
		return nil, err
	}
	s := &Source{
		Path:  src,
		IsDir: info.IsDir(),
	}
	return s, nil
}

func tryURL(l *Lookup, src string) (*Source, error) {
	u, err := url.Parse(src)
	if err != nil {
		return nil, err
	}
	res, err := l.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	s := &Source{
		Path:       src,
		IsDir:      false,
		readCloser: res.Body,
	}
	return s, nil
}

func tryGoPath(l *Lookup, src string) (*Source, error) {
	src = filepath.Join(os.Getenv("GOPATH"), "src", src)
	return tryLocal(l, src)
}

func tryGoRoot(l *Lookup, src string) (*Source, error) {
	src = filepath.Join(os.Getenv("GOROOT"), "src", src)
	return tryLocal(l, src)
}

func tryGoGet(l *Lookup, src string) (*Source, error) {
	repo, done, err := goget(src)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(repo)
	if err != nil {
		done()
		return nil, err
	}
	s := &Source{
		Path:  repo,
		IsDir: info.IsDir(),
		closer: func() error {
			done()
			return nil
		},
	}
	return s, nil
}

// Read implements io.Reader and is a safe way of reading
// files.
func (s *Source) Read(b []byte) (int, error) {
	if s.IsDir {
		return 0, errors.New("source: cannot read directories")
	}
	var err error
	s.readOnce.Do(func() {
		if s.readCloser != nil {
			return
		}
		s.readCloser, err = os.Open(s.Path)
	})
	if err != nil {
		return 0, err
	}
	return s.readCloser.Read(b)
}

// Close closes the Source and cleans up any used resources.
func (s *Source) Close() error {
	var err1, err2 error
	if rc := s.readCloser; rc != nil {
		err1 = rc.Close()
	}
	if c := s.closer; c != nil {
		err2 = c()
	}
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

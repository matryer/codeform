package source

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestGoGet(t *testing.T) {
	is := is.New(t)
	repo, done, err := goget("github.com/matryer/drop-test/explicit")
	is.NoErr(err) // goget
	defer done()
	is.True(strings.HasSuffix(repo, "src/github.com/matryer/drop-test/explicit"))
}

func TestSplit(t *testing.T) {
	is := is.New(t)
	var repo, path string

	repo, path = split("github.com/matryer")
	is.Equal(repo, "github.com/matryer")
	is.Equal(path, "")

	repo, path = split("github.com/matryer/codeform")
	is.Equal(repo, "github.com/matryer/codeform")
	is.Equal(path, "")

	repo, path = split("github.com/matryer/codeform/source")
	is.Equal(repo, "github.com/matryer/codeform")
	is.Equal(path, "source")

	repo, path = split("github.com/matryer/codeform/source/testdata")
	is.Equal(repo, "github.com/matryer/codeform")
	is.Equal(path, "source/testdata")

}

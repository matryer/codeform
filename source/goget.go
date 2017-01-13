package source

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// split breaks the repo and path out of the source.
// So that github.com/matryer/codeform/source becomes
// "github.com/matryer/codeform" and "source"
func split(src string) (string, string) {
	segs := strings.Split(src, "/")
	if len(segs) > 3 {
		return strings.Join(segs[0:3], "/"), strings.Join(segs[3:], "/")
	}
	return src, ""
}

func goget(src string) (string, func(), error) {
	done := func() {}
	tmp, err := ioutil.TempDir(os.TempDir(), ".codeform-")
	if err != nil {
		return "", done, err
	}
	done = func() {
		os.RemoveAll(tmp)
	}
	gopath := filepath.Join(tmp, "codeform-gopath")
	gopath, err = filepath.Abs(gopath)
	if err != nil {
		return "", done, err
	}
	err = os.MkdirAll(gopath, 0777)
	if err != nil {
		return "", done, err
	}
	source, path := split(src)
	//info("go get -d", source)
	goget := exec.Command("go", "get", "-d", source)
	env := []string{"GOPATH=" + gopath} // control GOPATH for this command
	env = append(env, os.Environ()...)  // but use rest of normal environemnt
	goget.Env = env

	// Omit error, `go get -d ` exits with status 1 if the package is not buildable
	_, _ = goget.CombinedOutput()
	fullpath := filepath.Join(gopath, "src", source, path)
	return fullpath, done, nil
}

type errGoGet struct {
	err    error
	source string
	output []byte
}

func (e errGoGet) Error() string {
	return "go get " + e.source + ": " + e.err.Error() + ": " + string(e.output)
}

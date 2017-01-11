package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/matryer/codeform/internal/version"
	"github.com/matryer/codeform/source"
	"github.com/matryer/codeform/tool"
	"github.com/pkg/errors"
)

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			io.WriteString(os.Stderr, "codeform: ")
			io.WriteString(os.Stderr, fatalErr.Error()+"\n")
			os.Exit(1)
		}
	}()
	var (
		src           = flag.String("src", ".", "code source")
		srcin         = flag.Bool("srcin", false, "take source from standard in")
		templatesrc   = flag.String("templatesrc", "", "template source")
		template      = flag.String("template", "", "template verbatim")
		targetPackage = flag.String("pkg", "", "target package")
		names         = flag.String("name", "", "comma separated list of things to include")
		packages      = flag.String("package", "", "comma separated list of packages to include")
		interfaces    = flag.String("interface", "", "comma separated list of interfaces to include")
		structs       = flag.String("struct", "", "comma separated list of structs to include")
		funcs         = flag.String("func", "", "comma separated list of funcs to include")
		out           = flag.String("out", "", "output file (defaults to standard out)")
		timeout       = flag.Duration("timeout", 2*time.Second, "timeout for HTTP requests")
		verbose       = flag.Bool("v", false, "verbose logging")
		nolinefeed    = flag.Bool("n", false, "suppress final linefeed")
		printVersion  = flag.Bool("version", false, "print version and exit")
	)
	flag.Usage = func() {
		fmt.Println("codeform", "v"+version.Current)
		flag.PrintDefaults()
	}
	flag.Parse()
	if *printVersion {
		fmt.Println(version.Current)
		return
	}
	log := func(args ...interface{}) {}
	logf := func(format string, args ...interface{}) {
		log(fmt.Sprintf(format, args...))
	}
	if *verbose {
		log = func(args ...interface{}) {
			fmt.Println(args...)
		}
	}
	var err error
	sourceLookup := source.Lookup{Client: http.Client{Timeout: *timeout}}
	if len(*src) == 0 && !*srcin {
		fatalErr = errors.New("provide src")
		return
	}
	logf("lookup: %s", *src)
	var codeSource *source.Source
	if *srcin {
		if len(*src) > 0 {
			logf("srcin: ignoring src")
		}
		codeSource = source.Reader("stdin", os.Stdin)
	} else {
		codeSource, err = sourceLookup.Get(*src)
		if err != nil {
			fatalErr = errors.Wrap(err, "lookup code source")
			return
		}
	}
	defer codeSource.Close()
	var templateSource *source.Source
	if len(*template)+len(*templatesrc) == 0 {
		fatalErr = errors.New("provide templatesrc or template")
		return
	}
	if len(*template) > 0 {
		if len(*templatesrc) > 0 {
			fatalErr = errors.New("provide templatesrc or template (not both)")
			return
		}
		log("using verbatim template")
		templateSource = source.Reader("verbatim-template", strings.NewReader(*template))
	} else {
		logf("lookup: %s", *templatesrc)
		templateSource, err = sourceLookup.Get(*templatesrc)
		if err != nil {
			fatalErr = errors.Wrap(err, "lookup template source")
			return
		}
	}
	defer templateSource.Close()
	var buf bytes.Buffer
	var w io.Writer
	w = &buf
	if len(*out) == 0 {
		w = os.Stdout
	}
	job := tool.Job{
		Log:           log,
		Code:          codeSource,
		Template:      templateSource,
		Names:         splitArgList(*names),
		Packages:      splitArgList(*packages),
		Interfaces:    splitArgList(*interfaces),
		Structs:       splitArgList(*structs),
		Funcs:         splitArgList(*funcs),
		TargetPackage: *targetPackage,
	}
	err = job.Execute(w)
	if err != nil {
		fatalErr = err
		return
	}
	if !*nolinefeed {
		w.Write([]byte("\n"))
	}
	if len(*out) > 0 {
		// save file
		err = ioutil.WriteFile(*out, buf.Bytes(), 0644)
		if err != nil {
			fatalErr = err
			return
		}
	}
}

func splitArgList(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, ",")
}

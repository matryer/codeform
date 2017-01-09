// Package editor provides the editor web tool.
// To run with Google App Engine, install the Google App Engine Go SDK
// and go `goapp serve` in this directory.
package editor

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func init() {
	http.Handle("/", New())
}

// New makes a new http.Handler that provides the editor.
func New() http.Handler {
	router := NewRouter()
	router.contexter = ContexterFunc(func(r *http.Request) context.Context {
		return appengine.NewContext(r)
	})
	routes(router)
	return router
}

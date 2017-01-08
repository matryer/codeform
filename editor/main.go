package editor

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func init() {
	router := NewRouter()
	router.contexter = ContexterFunc(func(r *http.Request) context.Context {
		return appengine.NewContext(r)
	})
	routes(router)
	http.Handle("/", router)
}

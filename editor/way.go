package editor

// lovingly stolen and modified from github.com/matryer/way

import (
	"net/http"
	"strings"

	"golang.org/x/net/context"
)

// wayContextKey is the context key type for storing
// parameters in context.Context.
type wayContextKey string

// Router routes HTTP requests.
type Router struct {
	routes []*route
	// NotFound is the http.Handler to call when no routes
	// match. By default uses http.NotFoundHandler().
	NotFound http.Handler

	// OnSystemError is called when a handler returns an error.
	OnSystemError func(context.Context, http.ResponseWriter, *http.Request, error)

	// contexter is the Contexter that generates context objects.
	contexter Contexter
}

// NewRouter makes a new Router.
func NewRouter() *Router {
	return &Router{
		NotFound: http.NotFoundHandler(),
		OnSystemError: func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}
}

func (r *Router) pathSegments(p string) []string {
	return strings.Split(strings.Trim(p, "/"), "/")
}

// Handler is http.Handler with context.
type Handler interface {
	ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

// HandlerFunc is http.HandlerFunc with context.
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (h HandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return h(ctx, w, r)
}

// Handle adds a handler with the specified method and pattern.
// Method can be any HTTP method string or "*" to match all methods.
// Pattern can contain path segments such as: /item/:id which is
// accessible via Param(context, "id").
// If pattern ends with trailing /, it acts as a prefix.
// Errors returned from handlers are considered system errors and are
// handled via Router.OnSystemError.
func (r *Router) Handle(method, pattern string, handler Handler) {
	route := &route{
		method:  strings.ToLower(method),
		segs:    r.pathSegments(pattern),
		handler: handler,
		prefix:  strings.HasSuffix(pattern, "/"),
	}
	r.routes = append(r.routes, route)
}

// HandleFunc is the http.HandlerFunc alternative to http.Handle.
func (r *Router) HandleFunc(method, pattern string, fn HandlerFunc) {
	r.Handle(method, pattern, fn)
}

// ServeHTTP routes the incoming http.Request based on method and path
// extracting path parameters as it goes.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r == nil {
		panic("nil router")
	}
	method := strings.ToLower(req.Method)
	segs := r.pathSegments(req.URL.Path)
	for _, route := range r.routes {
		if route.method != method && route.method != "*" {
			continue
		}
		if ctx, ok := route.match(r.contexter.Context(req), r, segs); ok {
			systemErr := route.handler.ServeHTTP(ctx, w, req)
			if systemErr != nil {
				r.OnSystemError(ctx, w, req, systemErr)
			}
			return
		}
	}
	r.NotFound.ServeHTTP(w, req)
}

// Param gets the path parameter from the specified Context.
// Returns an empty string if the parameter was not found.
func Param(ctx context.Context, param string) string {
	v := ctx.Value(wayContextKey(param))
	if v == nil {
		return ""
	}
	vStr, ok := v.(string)
	if !ok {
		return ""
	}
	return vStr
}

type route struct {
	method  string
	segs    []string
	handler Handler
	prefix  bool
}

func (r *route) match(ctx context.Context, router *Router, segs []string) (context.Context, bool) {
	if len(segs) > len(r.segs) && !r.prefix {
		return nil, false
	}
	for i, seg := range r.segs {
		if i > len(segs)-1 {
			return nil, false
		}
		isParam := false
		if strings.HasPrefix(seg, ":") {
			isParam = true
			seg = strings.TrimPrefix(seg, ":")
		}
		if !isParam { // verbatim check
			if seg != segs[i] {
				return nil, false
			}
		}
		if isParam {
			ctx = context.WithValue(ctx, wayContextKey(seg), segs[i])
		}
	}
	return ctx, true
}

// Contexter gets a context for a specific request.
type Contexter interface {
	Context(*http.Request) context.Context
}

// ContexterFunc is a function wrapper for Contexter.
type ContexterFunc func(*http.Request) context.Context

// Context gets a context for the specified request.
func (fn ContexterFunc) Context(r *http.Request) context.Context {
	return fn(r)
}

// ContexterBackground gets a Contexter that provides context.Background().
func ContexterBackground() Contexter {
	return ContexterFunc(func(*http.Request) context.Context {
		return context.Background()
	})
}

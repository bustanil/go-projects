package router

import (
	"bustanil.com/file-api/middleware"
	"net/http"
)

type Router struct {
	mux         *http.ServeMux
	middlewares []middleware.Middleware
	routes      map[pathMethod]http.HandlerFunc
}

type pathMethod struct {
	path   string
	method string
}

func NewRouter() *Router {
	return &Router{
		mux:         http.NewServeMux(),
		middlewares: make([]middleware.Middleware, 0),
		routes:      make(map[pathMethod]http.HandlerFunc),
	}
}

func (r *Router) WithMiddlewares(ms ...middleware.Middleware) *Router {
	r.middlewares = ms
	return r
}

func (r *Router) Build() *http.ServeMux {
	for pm, handlerFunc := range r.routes {
		if pm.method == http.MethodGet {
			r.mux.HandleFunc(pm.path, r.applyMiddlewares(get(handlerFunc), r.middlewares))
		}
		if pm.method == http.MethodPost {
			r.mux.HandleFunc(pm.path, r.applyMiddlewares(post(handlerFunc), r.middlewares))
		}
	}
	return r.mux
}

func get(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		} else {
			handlerFunc(writer, r)
		}
	}
}

func post(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		} else {
			handlerFunc(writer, r)
		}
	}
}

func (r *Router) applyMiddlewares(next http.HandlerFunc, m []middleware.Middleware) http.HandlerFunc {
	if len(m) == 0 {
		return next
	} else {
		return func(writer http.ResponseWriter, request *http.Request) {
			m[0].Intercept(writer, request, r.applyMiddlewares(next, m[1:]))
		}
	}
}

func (r *Router) GET(path string, handler http.HandlerFunc) *Router {
	r.routes[pathMethod{path, http.MethodGet}] = handler
	return r
}

func (r *Router) POST(path string, handler http.HandlerFunc) *Router {
	r.routes[pathMethod{path, http.MethodPost}] = handler
	return r
}

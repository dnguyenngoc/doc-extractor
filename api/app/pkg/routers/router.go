//doc-extractor/pkg/routers/router.go

package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	router := mux.NewRouter()
	return &Router{Router: router}
}

func (r *Router) Use(middleware func(http.Handler) http.Handler) {
	r.Router.Use(middleware)
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.HandleFunc(path, handler).Methods(http.MethodGet)
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.HandleFunc(path, handler).Methods(http.MethodPost)
}

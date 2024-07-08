package router

import (
	"github.com/gorilla/mux"
	"github.com/ternaryinvalid/todo1/internal/app/domain/config"
	"net/http"
)

type Router struct {
	router *mux.Router
	config config.Router
}

func New() *Router {
	router := mux.NewRouter()

	r := Router{
		router: router,
	}

	return &r
}

const (
	apiV1Prefix = "/api/v1"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.Handler
}

func (r *Router) Router() http.Handler {
	return r.router
}

func (r *Router) appendRoutesToRouter(subrouter *mux.Router, routes []Route) {
	for _, route := range routes {
		subrouter.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
}

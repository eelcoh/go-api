package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

// NewRoute creates a Route
func NewRoute(name, method, pattern string, handler http.HandlerFunc) Route {
	return Route{name, method, pattern, handler}
}

// NewRouter ...
func NewRouter(routes Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

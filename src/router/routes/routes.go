package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represent all the API routes
type Route struct {
	uri                    string
	method                 string
	function               func(http.ResponseWriter, *http.Request)
	authenticationRequired bool
}

// Config sets up all routes inside the router
func Config(r *mux.Router) *mux.Router {

	//Routes
	routes := append(routesUser, routesDoc)

	for _, route := range routes {
		r.HandleFunc(route.uri, route.function).Methods(route.method)
	}

	return r
}

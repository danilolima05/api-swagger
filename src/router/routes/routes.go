package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represent all the API routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

// Config put all teh routes inside the router
func Config(r *mux.Router) *mux.Router {

	//Routes from Users
	routes := routesUser

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	routesDoc := routesDoc

	for _, route := range routesDoc {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}

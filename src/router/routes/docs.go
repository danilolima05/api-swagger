package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesDoc = Route{
	uri:                    "/doc",
	method:                 http.MethodGet,
	function:               controllers.GetDocs,
	authenticationRequired: false,
}

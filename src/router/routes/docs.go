package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesDoc = []Route{

	{
		URI:                    "/docs",
		Method:                 http.MethodGet,
		Function:               controllers.GetDocs,
		AuthenticationRequired: false,
	},
}

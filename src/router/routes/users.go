package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUser = []Route{
	{
		uri:                    "/users",
		method:                 http.MethodGet,
		function:               controllers.GetUsers,
		authenticationRequired: false,
	},
	{
		uri:                    "/users",
		method:                 http.MethodPost,
		function:               controllers.CreateUser,
		authenticationRequired: false,
	},
}

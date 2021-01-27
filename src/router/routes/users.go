package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUser = []Route{
	// swagger:route GET /users users users listUsers
	//
	//
	//
	//

	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.GetUsers,
		AuthenticationRequired: false,
	},
}

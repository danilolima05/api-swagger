// Package classification of user API
//
// Documentation for User API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
//  swagger:meta

package controllers

import (
	"io/ioutil"
	"net/http"
)

// swagger:route GET /users users listUsers
// Return a random user
// responses:
//  200: userResponse

//GetUsers all users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	w.Write([]byte(body))
}

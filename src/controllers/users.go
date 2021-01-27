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

// swagger:route GET /users users GetUsers
// Return a random user
// responses:
//  200: randomUser return a random user

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

// swagger:route POST /users users CreateUser
// Return the same posted body - this is just for test
// responses:
//  200: createUserResponse return the payload

//CreateUser handles the users POST endpoint
func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// handle error
	}

	w.Write([]byte(requestBody))
}

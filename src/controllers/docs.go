// Package classification of Docs API
//
// Documentation for Doc API
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
	"log"
	"net/http"
)

// swagger:route GET /docs docs listTheDoc
// Return the api documentation

//GetDocs get the documentation
func GetDocs(w http.ResponseWriter, r *http.Request) {

	docFile, err := ioutil.ReadFile("swagger.yaml")
	if err != nil {
		log.Printf("docFile.Get err   #%v ", err)
	}

	w.Write([]byte(docFile))
}

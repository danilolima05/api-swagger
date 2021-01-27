package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
)

// swagger:route GET /doc doc GetDocumentation
// Return the api documentation

//GetDocs get the documentation
func GetDocs(w http.ResponseWriter, r *http.Request) {

	docFile, err := ioutil.ReadFile("swagger.yaml")
	if err != nil {
		log.Printf("docFile.Get err   #%v ", err)
	}

	w.Write([]byte(docFile))
}

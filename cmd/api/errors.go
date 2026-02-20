package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s %s => %s\n", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusInternalServerError, "something went wrong on the server")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s %s => %s\n", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s %s => %s\n", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusNotFound, "resource not found")
}

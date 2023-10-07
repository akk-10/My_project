package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.Handler(http.MethodGet, "/v1/healthcheck", http.HandlerFunc(app.healthcheckHandler))
	router.Handler(http.MethodPost, "/v1/cameras", http.HandlerFunc(app.createCameraHandler))
	router.Handler(http.MethodGet, "/v1/cameras/:id", http.HandlerFunc(app.showCameraHandler))

	return router
}

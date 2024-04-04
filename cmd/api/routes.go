package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/health", app.healthcheckHandler)

	protected := alice.New(app.requireActivatedUser)
	router.Handler(http.MethodGet, "/v1/movies", protected.ThenFunc(app.listMovieHandler))
	router.Handler(http.MethodPost, "/v1/movies", protected.ThenFunc(app.createMovieHandler))
	router.Handler(http.MethodGet, "/v1/movies/:id", protected.ThenFunc(app.showMovieHandler))
	router.Handler(http.MethodPatch, "/v1/movies/:id", protected.ThenFunc(app.updateMovieHandler))
	router.Handler(http.MethodDelete, "/v1/movies/:id", protected.ThenFunc(app.deleteMovieHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	standard := alice.New(app.recoverPanic, app.rateLimit, app.authenticate)
	return standard.Then(router)
}

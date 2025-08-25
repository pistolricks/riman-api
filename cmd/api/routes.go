package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/orders", app.requirePermission("orders:read", app.listOrdersHandler))
	router.HandlerFunc(http.MethodPost, "/v1/orders", app.requirePermission("orders:write", app.createOrderHandler))
	router.HandlerFunc(http.MethodGet, "/v1/orders/:id", app.requirePermission("orders:read", app.showOrderHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/orders/:id", app.requirePermission("orders:write", app.updateOrderHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/orders/:id", app.requirePermission("orders:write", app.deleteOrderHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/password", app.updateUserPasswordHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/activation", app.createActivationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/password-reset", app.createPasswordResetTokenHandler)

	router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}

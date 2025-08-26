package main

import (
	"net/http"
)

func (app *application) getShippingGet(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}

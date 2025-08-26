package main

import (
	"net/http"
)

func (app *application) getShippingGet(w http.ResponseWriter, r *http.Request) {
	app.cartV2.ShippingApi
}

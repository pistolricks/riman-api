package main

import (
	"net/http"
	"strconv"

	"github.com/pistolricks/riman-api/internal/cart_v1"
)

func (app *application) getShippingGetShippingMethods(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainOrdersFkStr := q.Get("mainOrdersFK")
	countryCode := q.Get("countryCode")
	mainOrdersFk64, _ := strconv.ParseInt(mainOrdersFkStr, 10, 32)
	res, h, err := app.cartV1.ShippingApi.ShippingGetShippingMethods(r.Context(), int32(mainOrdersFk64), countryCode)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"shippingMethods": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getShippingGetShippingSettings(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainOrdersFkStr := q.Get("mainOrdersFK")
	shippingChartTypeFkStr := q.Get("shippingChartTypeFK")
	mainOrdersFk64, _ := strconv.ParseInt(mainOrdersFkStr, 10, 32)
	shippingChartTypeFk64, _ := strconv.ParseInt(shippingChartTypeFkStr, 10, 32)
	res, h, err := app.cartV1.ShippingApi.ShippingGetShippingSettings(r.Context(), int32(mainOrdersFk64), int32(shippingChartTypeFk64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"shippingSettings": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) postShippingIsAddressValid(w http.ResponseWriter, r *http.Request) {
	var address cart_v1.AddressModel
	if err := app.readJSON(w, r, &address); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShippingApi.ShippingIsAddressValid(r.Context(), address)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"addressValidation": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

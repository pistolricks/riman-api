package main

import (
	"net/http"

	"github.com/pistolricks/riman-api/internal/cart_v1"
)

func (app *application) postShoppingAddressAddBillingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartAddressFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressAddBillingAddress(r.Context(), cartKey, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"billingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postShoppingAddressAddMailingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartAddressFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressAddMailingAddress(r.Context(), cartKey, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"mailingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postShoppingAddressAddShippingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartAddressFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressAddShippingAddress(r.Context(), cartKey, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"shippingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) patchShoppingAddressPatchBillingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartAddressViewModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressPatchBillingAddress(r.Context(), cartKey, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"billingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) patchShoppingAddressPatchMailingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartAddressViewModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressPatchMailingAddress(r.Context(), cartKey, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"mailingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) patchShoppingAddressPatchShippingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartAddressViewModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressPatchShippingAddress(r.Context(), cartKey, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"shippingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteShoppingAddressRemoveBillingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressRemoveBillingAddress(r.Context(), cartKey)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"billingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteShoppingAddressRemoveMailingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressRemoveMailingAddress(r.Context(), cartKey)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"mailingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteShoppingAddressRemoveShippingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressRemoveShippingAddress(r.Context(), cartKey)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"shippingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) putShoppingAddressUpdateBillingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartAddressFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressUpdateBillingAddress(r.Context(), cartKey, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"billingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) putShoppingAddressUpdateMailingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartAddressFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressUpdateMailingAddress(r.Context(), cartKey, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"mailingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) putShoppingAddressUpdateShippingAddress(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartAddressFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingAddressApi.ShoppingAddressUpdateShippingAddress(r.Context(), cartKey, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"shippingAddress": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

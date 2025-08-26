package main

import (
	"net/http"
	"strconv"

	cart_v1 "github.com/pistolricks/riman-api/internal/cart_v1"
	"github.com/antihax/optional"
)

func (app *application) putShoppingApplyDiscount(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.CartDiscountFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	h, err := app.cartV1.ShoppingApi.ShoppingApplyDiscount(r.Context(), model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": "ok"}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) deleteShoppingClearReferrer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	res, h, err := app.cartV1.ShoppingApi.ShoppingClearReferrer(r.Context(), id)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"cart": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postShoppingCreate(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.CartFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV1.ShoppingApi.ShoppingCreate(r.Context(), model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"cart": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postShoppingCreateMainOrder(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	res, h, err := app.cartV1.ShoppingApi.ShoppingCreateMainOrder(r.Context(), cartKey)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"mainOrderId": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getShoppingGetCalculatedTaxes(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	res, h, err := app.cartV1.ShoppingApi.ShoppingGetCalculatedTaxes(r.Context(), cartKey)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"tax": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getShoppingGetCampaigns(w http.ResponseWriter, r *http.Request) {
	opts := &cart_v1.ShoppingApiShoppingGetCampaignsOpts{}
	// no explicit optional fields in struct here beyond defaults; kept for future extension
	_ = opts
	res, h, err := app.cartV1.ShoppingApi.ShoppingGetCampaigns(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"campaigns": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getShoppingGetCart(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	res, h, err := app.cartV1.ShoppingApi.ShoppingGetCart(r.Context(), id)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"cart": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getShoppingGetCheckoutMessage(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v1.ShoppingApiShoppingGetCheckoutMessageOpts{}
	if v := q.Get("countryCode"); v != "" { opts.CountryCode = optional.NewString(v) }
	if v := q.Get("languageFK"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.LanguageFK = optional.NewInt32(int32(n)) } }
	if v := q.Get("cartKey"); v != "" { opts.CartKey = optional.NewInterface(v) }
	res, h, err := app.cartV1.ShoppingApi.ShoppingGetCheckoutMessage(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"messages": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postShoppingGetQualifiedSalesCampaignInfo(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.GetQualifiedSalesCampaignRequestModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV1.ShoppingApi.ShoppingGetQualifiedSalesCampaignInfo(r.Context(), model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"qualifiedSalesCampaign": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getShoppingGetRank(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	res, h, err := app.cartV1.ShoppingApi.ShoppingGetRank(r.Context(), id)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"rank": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getShoppingGetShippingOptions(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	res, h, err := app.cartV1.ShoppingApi.ShoppingGetShippingOptions(r.Context(), cartKey)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"shippingOptions": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) patchShoppingPatch(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var model cart_v1.CartPatchModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV1.ShoppingApi.ShoppingPatch(r.Context(), id, model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"cart": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) deleteShoppingRemoveCart(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	res, h, err := app.cartV1.ShoppingApi.ShoppingRemoveCart(r.Context(), id)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putShoppingSetShipping(w http.ResponseWriter, r *http.Request) {
	cartKey := r.URL.Query().Get("cartKey")
	var model cart_v1.CartSetShippingFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV1.ShoppingApi.ShoppingSetShipping(r.Context(), cartKey, model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putShoppingUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var model cart_v1.CartFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV1.ShoppingApi.ShoppingUpdate(r.Context(), id, model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"cart": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getShoppingValidateCart(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	cartKey := q.Get("cartKey")
 opts := &cart_v1.ShoppingApiShoppingValidateCartOpts{}
	if v := q.Get("orderNumber"); v != "" { opts.OrderNumber = optional.NewString(v) }
	res, h, err := app.cartV1.ShoppingApi.ShoppingValidateCart(r.Context(), cartKey, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"validation": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}

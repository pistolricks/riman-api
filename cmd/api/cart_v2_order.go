package main

import (
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	cart_v2 "github.com/pistolricks/riman-api/internal/cart_v2"
)

func (app *application) postOrderAddCCPaymentToOrderV2(w http.ResponseWriter, r *http.Request) {
	var model cart_v2.AddCcPaymentAndBillingFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV2.OrderApi.OrderAddCCPaymentToOrder(r.Context(), model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"payment": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postOrderAddProductV2(w http.ResponseWriter, r *http.Request) {
	var model cart_v2.AddProductFormModelv2
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV2.OrderApi.OrderAddProduct(r.Context(), model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postOrderApplyInteractiveDiscount(w http.ResponseWriter, r *http.Request) {
	var model cart_v2.ApplyIterativeDiscountFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV2.OrderApi.OrderApplyInterativeDiscount(r.Context(), model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postOrderCreateOrderV2(w http.ResponseWriter, r *http.Request) {
	var model cart_v2.NewOrderFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV2.OrderApi.OrderCreateOrder(r.Context(), model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"order": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getOrderGetOrderDetails(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainOrdersPK := q.Get("mainOrdersPK")
	opts := &cart_v2.OrderApiOrderGetOrderDetailsOpts{}
	if v := q.Get("model.mainOrdersPK"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.ModelMainOrdersPK = optional.NewInt32(int32(n)) } }
	if v := q.Get("model.email"); v != "" { opts.ModelEmail = optional.NewString(v) }
	if v := q.Get("model.culture"); v != "" { opts.ModelCulture = optional.NewString(v) }
	res, h, err := app.cartV2.OrderApi.OrderGetOrderDetails(r.Context(), mainOrdersPK, opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"details": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getOrderGetOrderTaxAmountString(w http.ResponseWriter, r *http.Request) {
	mainOrdersFkStr := r.URL.Query().Get("mainOrdersFk")
	mainOrdersFk64, _ := strconv.ParseInt(mainOrdersFkStr, 10, 32)
	res, h, err := app.cartV2.OrderApi.OrderGetOrderTaxAmountString(r.Context(), int32(mainOrdersFk64))
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"tax": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getOrderGetOrderType(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.OrderApiOrderGetOrderTypeOpts{}
	if v := q.Get("model.mainOrdersPK"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.ModelMainOrdersPK = optional.NewInt32(int32(n)) } }
	if v := q.Get("model.email"); v != "" { opts.ModelEmail = optional.NewString(v) }
	if v := q.Get("model.culture"); v != "" { opts.ModelCulture = optional.NewString(v) }
	res, h, err := app.cartV2.OrderApi.OrderGetOrderType(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"orderType": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putOrderSetShippingOnAutoShipOrder(w http.ResponseWriter, r *http.Request) {
	mainOrdersPKStr := r.URL.Query().Get("mainOrdersPK")
	mainOrdersPK64, _ := strconv.ParseInt(mainOrdersPKStr, 10, 32)
	res, h, err := app.cartV2.OrderApi.OrderSetShippingOnAutoShipOrder(r.Context(), int32(mainOrdersPK64))
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putOrderSetShippingOnOrderV2(w http.ResponseWriter, r *http.Request) {
	var model cart_v2.SetShippingFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV2.OrderApi.OrderSetShippingOnOrder(r.Context(), model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putOrderUpdateOrderItemsFromCart(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainOrdersFkStr := q.Get("mainOrdersFk")
	cartKey := q.Get("cartKey")
	mainOrdersFk64, _ := strconv.ParseInt(mainOrdersFkStr, 10, 32)
	res, h, err := app.cartV2.OrderApi.OrderUpdateOrderItemsFromCart(r.Context(), int32(mainOrdersFk64), cartKey)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putOrderUpdateOrderShippingAddressV2(w http.ResponseWriter, r *http.Request) {
	mainOrdersFKStr := r.URL.Query().Get("mainOrdersFK")
	mainOrdersFK64, _ := strconv.ParseInt(mainOrdersFKStr, 10, 32)
	var model cart_v2.UpdateOrderShipAddressFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.cartV2.OrderApi.OrderUpdateOrderShippingAddress(r.Context(), int32(mainOrdersFK64), model)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getOrderValidateRetailSignup(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.OrderApiOrderValidateRetailSignupOpts{}
	if v := q.Get("mainFK"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.MainFK = optional.NewInt32(int32(n)) } }
	res, h, err := app.cartV2.OrderApi.OrderValidateRetailSignup(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"retailSignup": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}

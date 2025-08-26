package main

import (
	"net/http"
	"strconv"

	"github.com/pistolricks/riman-api/internal/cart_v1"
)

func (app *application) postShoppingItemsAddAffiliateItems(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var items cart_v1.CartItemAffiliateModel
	if err := app.readJSON(w, r, &items); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsAddAffiliateItems(r.Context(), id, items)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"affiliateCart": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postShoppingItemsAddItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var model cart_v1.CartItemFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsAddItem(r.Context(), id, model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"item": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postShoppingItemsAddItems(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var models []cart_v1.CartItemFormModel
	if err := app.readJSON(w, r, &models); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsAddItems(r.Context(), id, models)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"items": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postShoppingItemsAddQuickCartItems(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var items cart_v1.CartItemAffiliateModel
	if err := app.readJSON(w, r, &items); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsAddQuickCartItems(r.Context(), id, items)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"affiliateCart": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getShoppingItemsGetCartItems(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsGetCartItems(r.Context(), id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"items": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) patchShoppingItemsPatchItem(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id := q.Get("id")
	itemIdStr := q.Get("itemId")
	itemId64, _ := strconv.ParseInt(itemIdStr, 10, 32)
	var model cart_v1.CartItemPatchModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsPatchItem(r.Context(), id, int32(itemId64), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"item": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteShoppingItemsRemoveCartItems(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsRemoveCartItems(r.Context(), id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteShoppingItemsRemoveCartItemsList(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var items []int32
	if err := app.readJSON(w, r, &items); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsRemoveCartItemsList(r.Context(), id, items)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteShoppingItemsRemoveItem(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id := q.Get("id")
	itemIdStr := q.Get("itemId")
	itemId64, _ := strconv.ParseInt(itemIdStr, 10, 32)
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsRemoveItem(r.Context(), id, int32(itemId64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) putShoppingItemsUpdateItem(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id := q.Get("id")
	itemIdStr := q.Get("itemId")
	itemId64, _ := strconv.ParseInt(itemIdStr, 10, 32)
	var model cart_v1.CartItemFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsUpdateItem(r.Context(), id, int32(itemId64), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"item": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getShoppingItemsValidateCartItems(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	res, h, err := app.cartV1.ShoppingItemsApi.ShoppingItemsValidateCartItems(r.Context(), id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"validation": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

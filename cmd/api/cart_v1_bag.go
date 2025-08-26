package main

import (
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	"github.com/pistolricks/riman-api/internal/cart_v1"
)

func (app *application) postBagCreate(w http.ResponseWriter, r *http.Request) {
	var input cart_v1.BagFormModel
	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.BagApi.BagCreate(r.Context(), input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"bag": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteBag(w http.ResponseWriter, r *http.Request) {
	bagIDStr := r.URL.Query().Get("bagId")
	bagID64, _ := strconv.ParseInt(bagIDStr, 10, 32)
	res, h, err := app.cartV1.BagApi.BagDelete(r.Context(), int32(bagID64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getBagAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v1.BagApiBagGetAllOpts{}
	if v := q.Get("bagPK"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.BagQueryModelBagPK = optional.NewInt32(int32(n))
		}
	}
	if v := q.Get("customerPK"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.BagQueryModelCustomerPK = optional.NewInt32(int32(n))
		}
	}
	if v := q.Get("cartType"); v != "" {
		opts.BagQueryModelCartType = optional.NewString(v)
	}
	if v := q.Get("countryCode"); v != "" {
		opts.BagQueryModelCountryCode = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.BagQueryModelCulture = optional.NewString(v)
	}
	if v := q.Get("siteUrl"); v != "" {
		opts.BagQueryModelSiteUrl = optional.NewString(v)
	}

	res, h, err := app.cartV1.BagApi.BagGetAll(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"bags": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getBagById(w http.ResponseWriter, r *http.Request) {
	bagIDStr := r.URL.Query().Get("bagId")
	bagID64, _ := strconv.ParseInt(bagIDStr, 10, 32)
	res, h, err := app.cartV1.BagApi.BagGetById(r.Context(), int32(bagID64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"bag": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

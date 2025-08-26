package main

import (
	"net/http"

	"github.com/antihax/optional"
	"github.com/pistolricks/riman-api/internal/cart_v1"
)

func (app *application) postProductGetProductConfiguration(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.ProductConfigurationFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ProductApi.ProductGetProductConfiguration(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"configuration": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) postProductGetProductDetail(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.ProductDetailFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ProductApi.ProductGetProductDetail(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"detail": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getProductGetProductRegulatoryDetails(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v1.ProductApiProductGetProductRegulatoryDetailsOpts{}
	if v := q.Get("sku"); v != "" {
		opts.Sku = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.Culture = optional.NewString(v)
	}
	res, h, err := app.cartV1.ProductApi.ProductGetProductRegulatoryDetails(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"regulatoryDetails": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getProductGetRafikiProduct(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.cartV1.ProductApi.ProductGetRafikiProduct(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"rafiki": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) postProductGetRetailProducts(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.ProductFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ProductApi.ProductGetRetailProducts(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"retailProducts": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getProductGetSamples(w http.ResponseWriter, r *http.Request) {
	siteURL := r.URL.Query().Get("siteUrl")
	res, h, err := app.cartV1.ProductApi.ProductGetSamples(r.Context(), siteURL)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"samples": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) postProductGetWholeSaleProducts(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.ProductFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ProductApi.ProductGetWholeSaleProducts(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"wholesaleProducts": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) postProductSubmitNvShadeRequest(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.ShadeRequestModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.ProductApi.ProductSubmitNvShadeRequest(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"shadeRequest": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

package main

import (
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	"github.com/pistolricks/riman-api/internal/cart_v2"
)

func (app *application) getProductGetBestSellers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetBestSellersOpts{}
	if v := q.Get("cartKey"); v != "" {
		opts.CartKey = optional.NewInterface(v)
	}
	if v := q.Get("productsToShow"); v != "" {
		opts.ProductsToShow = optional.NewString(v)
	}
	if v := q.Get("limit"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.Limit = optional.NewInt32(int32(n))
		}
	}
	if v := q.Get("cartType"); v != "" {
		opts.CartType = optional.NewString(v)
	}
	if v := q.Get("countryCode"); v != "" {
		opts.CountryCode = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.Culture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetBestSellers(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"bestSellers": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetBestsellers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetBestsellersOpts{}
	if v := q.Get("cartKey"); v != "" {
		opts.ModelCartKey = optional.NewInterface(v)
	}
	if v := q.Get("productsToShow"); v != "" {
		opts.ModelProductsToShow = optional.NewString(v)
	}
	if v := q.Get("limit"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.ModelLimit = optional.NewInt32(int32(n))
		}
	}
	if v := q.Get("cartType"); v != "" {
		opts.ModelCartType = optional.NewString(v)
	}
	if v := q.Get("countryCode"); v != "" {
		opts.ModelCountryCode = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.ModelCulture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetBestsellers(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"bestsellers": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetBrandSpotlight(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetBrandSpotlightOpts{}
	if v := q.Get("cartType"); v != "" {
		opts.CartType = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.Culture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetBrandSpotlight(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"brandSpotlight": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetCAPHeadersWithCVRanges(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	country := q.Get("country")
	langIdStr := q.Get("langId")
	langId64, _ := strconv.ParseInt(langIdStr, 10, 32)
	res, h, err := app.cartV2.ProductApi.ProductGetCAPHeadersWithCVRanges(r.Context(), country, int32(langId64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"capHeaders": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetCategories(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetCategoriesOpts{}
	if v := q.Get("cartType"); v != "" {
		opts.CartType = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.Culture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetCategories(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"categories": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetCategoryById(w http.ResponseWriter, r *http.Request) {
	categoryIdStr := r.URL.Query().Get("categoryId")
	categoryId64, _ := strconv.ParseInt(categoryIdStr, 10, 32)
	res, h, err := app.cartV2.ProductApi.ProductGetCategoryById(r.Context(), int32(categoryId64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"category": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetMayLoveProducts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	productPkStr := q.Get("productPK")
	productPk64, _ := strconv.ParseInt(productPkStr, 10, 32)
	opts := &cart_v2.ProductApiProductGetMayLoveProductsOpts{}
	if v := q.Get("suggestedQueryProductPk"); v != "" {
		opts.SuggestProductQueryCartType = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetMayLoveProducts(r.Context(), int32(productPk64), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"mayLove": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetProductBrands(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.cartV2.ProductApi.ProductGetProductBrands(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"brands": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetProductById(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	productFkStr := q.Get("productFk")
	cartType := q.Get("cartType")
	culture := q.Get("culture")
	productFk64, _ := strconv.ParseInt(productFkStr, 10, 32)
	res, h, err := app.cartV2.ProductApi.ProductGetProductById(r.Context(), int32(productFk64), cartType, culture)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"product": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetProductById_1(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetProductById_1Opts{}
	if v := q.Get("productPK"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.ProductPK = optional.NewInt32(int32(n))
		}
	}
	if v := q.Get("cartType"); v != "" {
		opts.CartType = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.Culture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetProductById_1(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"product": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetProductFunctions(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetProductFunctionsOpts{}
	if v := q.Get("culture"); v != "" {
		opts.Culture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetProductFunctions(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"productFunctions": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetProducts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetProductsOpts{}
	if v := q.Get("brandId"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.BrandId = optional.NewInt32(int32(n))
		}
	}
	if v := q.Get("productPk"); v != "" {

	}
	if v := q.Get("categoryId"); v != "" {

	}
	if v := q.Get("market"); v != "" {

	}
	if v := q.Get("promo"); v != "" {

	}
	if v := q.Get("cartType"); v != "" {
		opts.CartType = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.Culture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetProducts(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"products": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetRemainingProductQuantities(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetRemainingProductQuantitiesOpts{}
	if v := q.Get("modelProductIds"); v != "" {
		opts.ModelProductIds = optional.NewString(v)
	}
	if v := q.Get("mdelMainPK"); v != "" {
		//	opts.ModelMainPK = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetRemainingProductQuantities(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"remainingQuantities": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetRemainingSignUpProductQuantities(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetRemainingSignUpProductQuantitiesOpts{}
	if v := q.Get("modelProductIds"); v != "" {
		opts.ModelProductIds = optional.NewString(v)
	}
	if v := q.Get("modelMainPk"); v != "" {
		// opts.Market = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetRemainingSignUpProductQuantities(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"remainingSignUpQuantities": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetRitualProducts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	productPkStr := q.Get("productPK")
	productPk64, _ := strconv.ParseInt(productPkStr, 10, 32)
	opts := &cart_v2.ProductApiProductGetRitualProductsOpts{}
	if v := q.Get("suggestProductQueryCartType"); v != "" {
		opts.SuggestProductQueryCartType = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetRitualProducts(r.Context(), int32(productPk64), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"ritualProducts": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetRituals(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetRitualsOpts{}
	if v := q.Get("modelCartType"); v != "" {
		opts.ModelCartType = optional.NewString(v)
	}
	if v := q.Get("modelCulture"); v != "" {
		opts.ModelCulture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetRituals(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"rituals": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductGetSubCategories(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductGetSubCategoriesOpts{}
	if v := q.Get("cartType"); v != "" {
		opts.CartType = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.Culture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductGetSubCategories(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"subCategories": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductInvalidateCache(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.cartV2.ProductApi.ProductInvalidateCache(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getProductSearchProduct(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &cart_v2.ProductApiProductSearchProductOpts{}
	if v := q.Get("cartType"); v != "" {
		// if n, err := strconv.ParseInt(v, 10, 32); err == nil {
		// 	//opts.CategoryPK = optional.NewInt32(int32(n))
		// }
	}
	if v := q.Get("brandPK"); v != "" {
		//	if n, err := strconv.ParseInt(v, 10, 32); err == nil {
		//			opts.BrandPK = optional.NewInt32(int32(n))
		//	}
	}
	// if v := q.Get("category"); v != "" {
	// 	opts.Category = optional.NewString(v)
	// }
	// if v := q.Get("gender"); v != "" {
	// 	opts.Gender = optional.NewString(v)
	// }
	// if v := q.Get("market"); v != "" {
	// 	opts.Market = optional.NewString(v)
	// }
	// if v := q.Get("promo"); v != "" {
	// 	opts.Promo = optional.NewString(v)
	// }
	if v := q.Get("cartType"); v != "" {
		opts.CartType = optional.NewString(v)
	}
	if v := q.Get("culture"); v != "" {
		opts.Culture = optional.NewString(v)
	}
	res, h, err := app.cartV2.ProductApi.ProductSearchProduct(r.Context(), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"search": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

package main

import (
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	"github.com/pistolricks/riman-api/internal/cart_v1"
)

func (app *application) getLoyaltyPointsRanks(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.cartV1.LoyaltyPointsApi.LoyaltyPointsGetRanks(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"ranks": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getLoyaltyPointsStatus(w http.ResponseWriter, r *http.Request) {
	mainPkStr := r.URL.Query().Get("mainPk")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	res, h, err := app.cartV1.LoyaltyPointsApi.LoyaltyPointsGetStatus(r.Context(), int32(mainPk64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"status": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getLoyaltyPointsTransactions(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainPkStr := q.Get("mainPk")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	opts := &cart_v1.LoyaltyPointsApiLoyaltyPointsGetTransactionsOpts{}
	if v := q.Get("pageNumber"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.PageNumber = optional.NewInt32(int32(n))
		}
	}
	if v := q.Get("pageSize"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.PageSize = optional.NewInt32(int32(n))
		}
	}
	res, h, err := app.cartV1.LoyaltyPointsApi.LoyaltyPointsGetTransactions(r.Context(), int32(mainPk64), opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"transactions": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

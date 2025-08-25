package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pistolricks/riman-api/internal/data"
	"github.com/pistolricks/riman-api/internal/validator"
)

func (app *application) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID                      int64       `json:"id"`
		CreatedAt               time.Time   `json:"created_at"`
		OrderDate               string      `json:"orderDate"`
		MainOrdersPK            int         `json:"mainOrdersPK"`
		OrderType               string      `json:"orderType"`
		FinalOrderType          interface{} `json:"finalOrderType"`
		SiteURL                 string      `json:"siteURL"`
		EncOrderNumber          string      `json:"encOrderNumber"`
		CurrencySymbol          string      `json:"currencySymbol"`
		CurrencyCode            string      `json:"currencyCode"`
		PaidStatus              bool        `json:"paidStatus"`
		HasTaxInvoice           bool        `json:"hasTaxInvoice"`
		HasCommercialInvoice    bool        `json:"hasCommercialInvoice"`
		HasCreditNote           bool        `json:"hasCreditNote"`
		IsShippingPending       bool        `json:"isShippingPending"`
		IsPB                    bool        `json:"isPB"`
		IsPA                    bool        `json:"isPA"`
		IsCC                    bool        `json:"isCC"`
		MainFK                  int         `json:"mainFK"`
		MainOrderTypeFK         int         `json:"mainOrderTypeFK"`
		VoucherURL              interface{} `json:"voucherURL"`
		ShipCountry             string      `json:"shipCountry"`
		ShippingStatus          string      `json:"shippingStatus"`
		OrderShippingStatus     string      `json:"orderShippingStatus"`
		OrderTypeValue          interface{} `json:"orderTypeValue"`
		PaidStatusValue         string      `json:"paidStatusValue"`
		Quantity                int         `json:"quantity"`
		Email                   interface{} `json:"email"`
		Phone                   interface{} `json:"phone"`
		ShipFirstName           interface{} `json:"shipFirstName"`
		ShipLastName            interface{} `json:"shipLastName"`
		MarkedPaidDate          string      `json:"markedPaidDate"`
		Total                   float64     `json:"total"`
		ConvTotal               float64     `json:"convTotal"`
		ConvTotalFormat         string      `json:"convTotalFormat"`
		SubTotal                int         `json:"subTotal"`
		ConvSubtotal            int         `json:"convSubtotal"`
		ShipTotal               float64     `json:"shipTotal"`
		ConvShipTotal           float64     `json:"convShipTotal"`
		Taxes                   float64     `json:"taxes"`
		TaxLabel                string      `json:"taxLabel"`
		ProductTax              float64     `json:"productTax"`
		ShippingTax             float64     `json:"shippingTax"`
		TotalProductTax         float64     `json:"totalProductTax"`
		AdditionalTaxLabel      string      `json:"additionalTaxLabel"`
		AdditionalTax           interface{} `json:"additionalTax"`
		ConvTaxes               float64     `json:"convTaxes"`
		OrderProcessingFees     interface{} `json:"orderProcessingFees"`
		ConvOrderProcessingFees interface{} `json:"convOrderProcessingFees"`
		Discount                int         `json:"discount"`
		ConvDiscount            int         `json:"convDiscount"`
		RefundAmount            int         `json:"refundAmount"`
		ConvRefund              int         `json:"convRefund"`
		SalesCampaignFK         interface{} `json:"salesCampaignFK"`
		Paidstatusfk            int         `json:"paidstatusfk"`
		DeliveryDate            interface{} `json:"deliveryDate"`
		ShippingDetails         interface{} `json:"shippingDetails"`
		OrderItems              interface{} `json:"orderItems"`
		Payments                interface{} `json:"payments"`
		IsPrepaidOrder          interface{} `json:"isPrepaidOrder"`
		ShowInvoice             bool        `json:"showInvoice"`
		ShowOrderInvoice        bool        `json:"showOrderInvoice"`
		KrGuaranteeNo           string      `json:"krGuaranteeNo"`
		WeChatOrderNumber       interface{} `json:"weChatOrderNumber"`
		MemberID                interface{} `json:"memberID"`
		Version                 int         `json:"-"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	order := &data.Order{
		ID:                      input.ID,
		CreatedAt:               input.CreatedAt,
		OrderDate:               input.OrderDate,
		MainOrdersPK:            input.MainOrdersPK,
		OrderType:               input.OrderType,
		FinalOrderType:          input.FinalOrderType,
		SiteURL:                 input.SiteURL,
		EncOrderNumber:          input.EncOrderNumber,
		CurrencySymbol:          input.CurrencySymbol,
		CurrencyCode:            input.CurrencyCode,
		PaidStatus:              input.PaidStatus,
		HasTaxInvoice:           input.HasTaxInvoice,
		HasCommercialInvoice:    input.HasCommercialInvoice,
		HasCreditNote:           input.HasCreditNote,
		IsShippingPending:       input.IsShippingPending,
		IsPB:                    input.IsPB,
		IsPA:                    input.IsPA,
		IsCC:                    input.IsCC,
		MainFK:                  input.MainFK,
		MainOrderTypeFK:         input.MainOrderTypeFK,
		VoucherURL:              input.VoucherURL,
		ShipCountry:             input.ShipCountry,
		ShippingStatus:          input.ShippingStatus,
		OrderShippingStatus:     input.OrderShippingStatus,
		OrderTypeValue:          input.OrderTypeValue,
		PaidStatusValue:         input.PaidStatusValue,
		Quantity:                input.Quantity,
		Email:                   input.Email,
		Phone:                   input.Phone,
		ShipFirstName:           input.ShipFirstName,
		ShipLastName:            input.ShipLastName,
		MarkedPaidDate:          input.MarkedPaidDate,
		Total:                   input.Total,
		ConvTotal:               input.ConvTotal,
		ConvTotalFormat:         input.ConvTotalFormat,
		SubTotal:                input.SubTotal,
		ConvSubtotal:            input.ConvSubtotal,
		ShipTotal:               input.ShipTotal,
		ConvShipTotal:           input.ConvShipTotal,
		Taxes:                   input.Taxes,
		TaxLabel:                input.TaxLabel,
		ProductTax:              input.ProductTax,
		ShippingTax:             input.ShippingTax,
		TotalProductTax:         input.TotalProductTax,
		AdditionalTaxLabel:      input.AdditionalTaxLabel,
		AdditionalTax:           input.AdditionalTax,
		ConvTaxes:               input.ConvTaxes,
		OrderProcessingFees:     input.OrderProcessingFees,
		ConvOrderProcessingFees: input.ConvOrderProcessingFees,
		Discount:                input.Discount,
		ConvDiscount:            input.ConvDiscount,
		RefundAmount:            input.RefundAmount,
		ConvRefund:              input.ConvRefund,
		SalesCampaignFK:         input.SalesCampaignFK,
		Paidstatusfk:            input.Paidstatusfk,
		DeliveryDate:            input.DeliveryDate,
		ShippingDetails:         input.ShippingDetails,
		OrderItems:              input.OrderItems,
		Payments:                input.Payments,
		IsPrepaidOrder:          input.IsPrepaidOrder,
		ShowInvoice:             input.ShowInvoice,
		ShowOrderInvoice:        input.ShowOrderInvoice,
		KrGuaranteeNo:           input.KrGuaranteeNo,
		WeChatOrderNumber:       input.WeChatOrderNumber,
		MemberID:                input.MemberID,
		Version:                 input.Version,
	}

	v := validator.New()

	if data.ValidateOrder(v, order); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Orders.Insert(order)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/orders/%d", order.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"order": order}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	order, err := app.models.Orders.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"order": order}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	order, err := app.models.Orders.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title   *string       `json:"title"`
		Year    *int32        `json:"year"`
		Runtime *data.Runtime `json:"runtime"`
		Genres  []string      `json:"genres"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Orders.Update(order)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"order": order}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Orders.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "order successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listOrdersHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string
		Genres []string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Genres = app.readCSV(qs, "genres", []string{})

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	orders, metadata, err := app.models.Orders.GetAll(input.Title, input.Genres, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"orders": orders, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

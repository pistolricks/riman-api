package main

import (
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	"github.com/pistolricks/riman-api/internal/cart_v1"
)

func (app *application) postOrderAddCCPaymentToOrder(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.AddCcPaymentFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderAddCCPaymentToOrder(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"payment": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderAddLogOrderStatus(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.AddLogOrderStatusFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderAddLogOrderStatus(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderAddMainRetailInfo(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.AddMainRetailInfoFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderAddMainRetailInfo(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderAddProcessingFees(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.AddProcessingFeesFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderAddProcessingFees(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"processingFees": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderAddProduct(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.AddProductFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderAddProduct(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderApplyDiscountCode(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainOrdersPKStr := q.Get("mainOrdersPK")
	mainOrdersPK64, _ := strconv.ParseInt(mainOrdersPKStr, 10, 32)
	var model cart_v1.ApplyOrderDiscountFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderApplyDiscountCode(r.Context(), int32(mainOrdersPK64), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderApplyRimanDollarsPayment(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainOrdersPKStr := q.Get("mainOrdersPK")
	mainOrdersPK64, _ := strconv.ParseInt(mainOrdersPKStr, 10, 32)
	var model cart_v1.RimanDollarsPaymentModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderApplyRimanDollarsPayment(r.Context(), int32(mainOrdersPK64), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"payment": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderCancelMerchantOrder(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.OrderResponseModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderCancelMerchantOrder(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getOrderCheckOrderPaidStatus(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("orderId")
	res, h, err := app.cartV1.OrderApi.OrderCheckOrderPaidStatus(r.Context(), orderId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"status": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getOrderCheckOrderPaymentStarted(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("orderId")
	res, h, err := app.cartV1.OrderApi.OrderCheckOrderPaymentStarted(r.Context(), orderId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"status": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getOrderClearMerchantPendingStatus(w http.ResponseWriter, r *http.Request) {
	orderIdStr := r.URL.Query().Get("orderId")
	orderId64, _ := strconv.ParseInt(orderIdStr, 10, 32)
	res, h, err := app.cartV1.OrderApi.OrderClearMerchantPendingStatus(r.Context(), int32(orderId64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getOrderConvertToLocalCurrency(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainOrdersFKStr := q.Get("mainOrdersFK")
	amountStr := q.Get("amount")
	mainOrdersFK64, _ := strconv.ParseInt(mainOrdersFKStr, 10, 32)
	amount, _ := strconv.ParseFloat(amountStr, 64)
	res, h, err := app.cartV1.OrderApi.OrderConvertToLocalCurrency(r.Context(), int32(mainOrdersFK64), amount)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderCreateAutoshipOrder(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.cartV1.OrderApi.OrderCreateAutoshipOrder(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"autoshipOrder": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderCreateOrder(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.NewOrderFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderCreateOrder(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"order": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderCreateOrderClimbersClub(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.cartV1.OrderApi.OrderCreateOrderClimbersClub(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderCreatePrepaidOrder(w http.ResponseWriter, r *http.Request) {
	var products []cart_v1.CartItemFormModel
	if err := app.readJSON(w, r, &products); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderCreatePrepaidOrder(r.Context(), products)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"autoshipOrder": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderCreateRMA(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.RmaFormModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderCreateRMA(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"rma": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getOrderDeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderIdStr := r.URL.Query().Get("orderId")
	orderId64, _ := strconv.ParseInt(orderIdStr, 10, 32)
	res, h, err := app.cartV1.OrderApi.OrderDeleteOrder(r.Context(), int32(orderId64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderGet3Ds2(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.OrderComplete3Ds2Model
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderGet3Ds2(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderGet3DsTokenUrl(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.OrderCompleteTokenModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderGet3DsTokenUrl(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) postOrderGet3DsUrl(w http.ResponseWriter, r *http.Request) {
	var model cart_v1.OrderCompleteModel
	if err := app.readJSON(w, r, &model); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	res, h, err := app.cartV1.OrderApi.OrderGet3DsUrl(r.Context(), model)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getOrderGetAcceptedCreditCardTypes(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	countryCode := q.Get("countryCode")
	currencyCode := q.Get("currencyCode")
	opts := &cart_v1.OrderApiOrderGetAcceptedCreditCardTypesOpts{}
	if v := q.Get("mainOrderTypeFK"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.MainOrderTypeFK = optional.NewInt32(int32(n))
		}
	}
	if v := q.Get("mainOrderPk"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 32); err == nil {
			opts.MainOrderPk = optional.NewInt32(int32(n))
		}
	}
	res, h, err := app.cartV1.OrderApi.OrderGetAcceptedCreditCardTypes(r.Context(), countryCode, currencyCode, opts)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"types": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getOrderGetAvailableCoupons(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainPkStr := q.Get("mainPk")
	mainOrdersPKStr := q.Get("mainOrdersPK")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	mainOrdersPK64, _ := strconv.ParseInt(mainOrdersPKStr, 10, 32)
	res, h, err := app.cartV1.OrderApi.OrderGetAvailableCoupons(r.Context(), int32(mainPk64), int32(mainOrdersPK64))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if err := app.writeJSON(w, h.StatusCode, envelope{"coupons": res}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) getOrderGetBonusCreditLimit(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetBraspagAntifraudScirpts(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetCustomsPaymentWarnings(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetDonationOrders(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetInAuthConfig(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderGetInstallments(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetMainOrders(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetManualPayTypes(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderGetMerchantRedirectModel(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderGetMerchantRedirectModelStep2(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetMinimumAllowedPayment(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOpenPayReceiptLink(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrSearchRMAs(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderDetailsForConfirmation(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderDisclosure(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderId(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderInfo(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderItems(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderProducts(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderShipmentProducts(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderShipmentsTotal(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderShipping(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderShippingDetails(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderTaxAmount(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrderTypes(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetOrders(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetPayUBankTransferDetails(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetPayUReceiptLink(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetPaymentInfo(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetPaymentOrderProcessingFees(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetPayments(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetPayments_1(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetRMADetail(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetRMAReasons(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetRMAsProducts(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetRMAsStatus(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderGetRetailCartByRemarketingTrackingId(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetRimanDollarsBalance(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetSalesCampaignByCampaignCode(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetSalesCampaignByCampaignCode_2(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetSalesCampaignByCampaignCode_3(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetSavedCreditCard(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetSiftConfig(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetTapPayConfig(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetTokenExConfig(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetWalletDiscount(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderGetWalletLimit(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderInitInstallments(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderIsAllowPersonalUseFun(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderIsCreditCardRedirectCheckout(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderIsInstallmentsExtensionValid(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderIsIxoPayActiveForGateway(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderIsLacoreFraudSettings(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderIsMercadoPagoOrder(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderIsNiceGtwEnabled(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderIsQRCodeOrderStatusUpdated(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderIsRoutedToZiraat(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderNewPayment(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderOrderHasPackage(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderProcessOrderResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderProcessOrderResponsePost(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderProcessPayPalOrder(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderProcessPayPalOrderResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderSaveInstallments(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putOrderSaveTokenExConfig(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderSendSummury(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderSetInvoiceInsuranceType(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putOrderSetOrderShipping(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putOrderSetOrderTax(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderSetOrderType(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putOrderSetShippingOnOrder(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderSubmitOrder(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putOrderUpdateImportOrderType(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putOrderUpdateOrderMainType(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putOrderUpdateOrderShippingAddress(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putOrderUpdatePayment(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderUpdatePersonalUseFlag(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putOrderUpdateRMA(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderUseCrossBorderCharges(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderUseSavedCreditCard(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderValidateCreditCard(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postOrderValidateDiscountCode(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderVerifyDiscountCode(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) getOrderVerifyUserOrders(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}

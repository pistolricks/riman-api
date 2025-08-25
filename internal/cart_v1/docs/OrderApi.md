# \OrderApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderAddCCPaymentToOrder**](OrderApi.md#OrderAddCCPaymentToOrder) | **Post** /api/v1/order/payment/creditcard | Adds the Credit Card payment information to an order in process
[**OrderAddLogOrderStatus**](OrderApi.md#OrderAddLogOrderStatus) | **Post** /api/v1/order/log | Used to add a status update log to an order.
[**OrderAddMainRetailInfo**](OrderApi.md#OrderAddMainRetailInfo) | **Post** /api/v1/order/retail/info | Used to set retail tracking information for retails orders once they have been completed.
[**OrderAddProcessingFees**](OrderApi.md#OrderAddProcessingFees) | **Post** /api/v1/order/addProcessingFees | Adds processing fees to an order
[**OrderAddProduct**](OrderApi.md#OrderAddProduct) | **Post** /api/v1/order/product | Remove items from an existing, in process order. It will remove all existing items, if any, and add the items provided to the order
[**OrderApplyDiscountCode**](OrderApi.md#OrderApplyDiscountCode) | **Post** /api/v1/orders/{mainOrdersPK}/apply-discount-code | Used to apply a discount code to an order for a distributor.
[**OrderApplyRimanDollarsPayment**](OrderApi.md#OrderApplyRimanDollarsPayment) | **Post** /api/v1/order/{mainOrdersPK}/apply-riman-dollars | Used to apply a Riman dollars to an order for a distributor.
[**OrderCancelMerchantOrder**](OrderApi.md#OrderCancelMerchantOrder) | **Post** /api/v1/order/merchant/cancel | Used to cancelation PayPal Order.
[**OrderCheckOrderPaidStatus**](OrderApi.md#OrderCheckOrderPaidStatus) | **Get** /api/v1/order/{orderId}/paid | 
[**OrderCheckOrderPaymentStarted**](OrderApi.md#OrderCheckOrderPaymentStarted) | **Get** /api/v1/order/{orderId}/payment-started | 
[**OrderClearMerchantPendingStatus**](OrderApi.md#OrderClearMerchantPendingStatus) | **Get** /api/v1/order/{orderId}/clearMerchantPendingStatus | 
[**OrderConvertToLocalCurrency**](OrderApi.md#OrderConvertToLocalCurrency) | **Get** /api/v1/order/convert-to-local-currency/{mainOrdersFK}/{amount} | Returns converted amount to local currency
[**OrderCreateAutoshipOrder**](OrderApi.md#OrderCreateAutoshipOrder) | **Post** /api/v1/orders/create-autoship-order | Runs the manual process to create an autoship order for the currently logged in customer.
[**OrderCreateOrder**](OrderApi.md#OrderCreateOrder) | **Post** /api/v1/order | Create a new order with basic information
[**OrderCreateOrderClimbersClub**](OrderApi.md#OrderCreateOrderClimbersClub) | **Post** /api/v1/order/climbers-club | Finalize climbers club sign up. Creates order and takes money from cliffs wallet. Returns success or failure
[**OrderCreatePrepaidOrder**](OrderApi.md#OrderCreatePrepaidOrder) | **Post** /api/v1/orders/create-prepaid-order | Runs the manual process to create an prepaid order for the currently logged in customer.
[**OrderCreateRMA**](OrderApi.md#OrderCreateRMA) | **Post** /api/v1/RMA | Create RMA
[**OrderDeleteOrder**](OrderApi.md#OrderDeleteOrder) | **Get** /api/v1/order/delete/{orderId} | This method is no longer allowed please do not use
[**OrderGet3Ds2**](OrderApi.md#OrderGet3Ds2) | **Post** /api/v1/order/3ds2 | Get redirect Url for 3D 2 Secure payment
[**OrderGet3DsTokenUrl**](OrderApi.md#OrderGet3DsTokenUrl) | **Post** /api/v1/order/3dstokenurl | Get token url/model for token processing
[**OrderGet3DsUrl**](OrderApi.md#OrderGet3DsUrl) | **Post** /api/v1/order/3dsurl | Get redirect Url for 3D Secure payment
[**OrderGetAcceptedCreditCardTypes**](OrderApi.md#OrderGetAcceptedCreditCardTypes) | **Get** /api/v1/order/payment/creditcard/types/{countryCode}/{currencyCode} | Returns a list of accepted credit card types for the specified country code and currency code.
[**OrderGetAvailableCoupons**](OrderApi.md#OrderGetAvailableCoupons) | **Get** /api/v1/order/get-available-coupons/{mainPk}/{mainOrdersPK} | 
[**OrderGetBonusCreditLimit**](OrderApi.md#OrderGetBonusCreditLimit) | **Get** /api/v1/order/getBonusTokenOrderLimit/{mainOrdersPK} | 
[**OrderGetBraspagAntifraudScirpts**](OrderApi.md#OrderGetBraspagAntifraudScirpts) | **Get** /api/v1/order/get-braspagantifraud-scripts/{identifier} | Used to get InAuth configuration.
[**OrderGetCustomsPaymentWarnings**](OrderApi.md#OrderGetCustomsPaymentWarnings) | **Get** /api/v1/payment/customs-warning | Returns messages related to customs payment warning for the warehouse country
[**OrderGetDonationOrders**](OrderApi.md#OrderGetDonationOrders) | **Get** /api/v1/order/donationOrders | To get information about Donations
[**OrderGetInAuthConfig**](OrderApi.md#OrderGetInAuthConfig) | **Get** /api/v1/order/get-inauth-config | Used to get InAuth configuration.
[**OrderGetInstallments**](OrderApi.md#OrderGetInstallments) | **Post** /api/v1/order/getInstallments | 
[**OrderGetMainOrders**](OrderApi.md#OrderGetMainOrders) | **Get** /api/v1/orders | 
[**OrderGetManualPayTypes**](OrderApi.md#OrderGetManualPayTypes) | **Get** /api/v1/orders/payments/manual | Retrieves available manual paytypes
[**OrderGetMerchantRedirectModel**](OrderApi.md#OrderGetMerchantRedirectModel) | **Post** /api/v1/order/merchant/getRedirectModel | Return merchant redirect model with required parameters for UI
[**OrderGetMerchantRedirectModelStep2**](OrderApi.md#OrderGetMerchantRedirectModelStep2) | **Post** /api/v1/order/merchant/getRedirectModelStep2 | Get script\\iframe for redirect step 2
[**OrderGetMinimumAllowedPayment**](OrderApi.md#OrderGetMinimumAllowedPayment) | **Get** /api/v1/order/getMinimumAllowedPayment/{mainOrdersPK} | 
[**OrderGetOpenPayReceiptLink**](OrderApi.md#OrderGetOpenPayReceiptLink) | **Get** /api/v1/order/getOpenPayReceiptLink/{mainOrdersPK} | 
[**OrderGetOrSearchRMAs**](OrderApi.md#OrderGetOrSearchRMAs) | **Get** /api/v1/RMA | To get items from RMA table order OrderByDescending Issue_Date
[**OrderGetOrderDetailsForConfirmation**](OrderApi.md#OrderGetOrderDetailsForConfirmation) | **Get** /api/v1/order/{orderString}/confirmation-details | Get Order Details for confirmation Page.
[**OrderGetOrderDisclosure**](OrderApi.md#OrderGetOrderDisclosure) | **Get** /api/v1/order/disclosure/{orderId} | 
[**OrderGetOrderId**](OrderApi.md#OrderGetOrderId) | **Get** /api/v1/orders/{mainOrdersPK} | To get single order
[**OrderGetOrderInfo**](OrderApi.md#OrderGetOrderInfo) | **Get** /api/v1/order/details/{mainOrdersFK} | 
[**OrderGetOrderItems**](OrderApi.md#OrderGetOrderItems) | **Get** /api/v1/orders/{mainOrdersPK}/items | To get main order items
[**OrderGetOrderProducts**](OrderApi.md#OrderGetOrderProducts) | **Get** /api/v1/orders/{mainOrdersPK}/products | To get all subitems from main order by id for RMA
[**OrderGetOrderShipmentProducts**](OrderApi.md#OrderGetOrderShipmentProducts) | **Get** /api/v1/orders/{mainOrdersPK}/shipment-products | To get shipment data of orders products
[**OrderGetOrderShipmentsTotal**](OrderApi.md#OrderGetOrderShipmentsTotal) | **Get** /api/v1/orders/{mainOrdersPK}/shipments-total | To get shipments total cost
[**OrderGetOrderShipping**](OrderApi.md#OrderGetOrderShipping) | **Get** /api/v1/order/shipping/{mainOrdersFK}/{shippingChartTypeFK} | Gets the shipping information for an order.
[**OrderGetOrderShippingDetails**](OrderApi.md#OrderGetOrderShippingDetails) | **Get** /api/v1/orders/{mainOrdersPK}/tracking | To get shipping Tracking information of an order
[**OrderGetOrderTaxAmount**](OrderApi.md#OrderGetOrderTaxAmount) | **Get** /api/v1/order/taxes/{mainOrdersFK} | Gets the tax information for an order.
[**OrderGetOrderTypes**](OrderApi.md#OrderGetOrderTypes) | **Get** /api/v1/orders/types | To get order types
[**OrderGetOrders**](OrderApi.md#OrderGetOrders) | **Get** /api/v1/order | Gets order history.
[**OrderGetPayUBankTransferDetails**](OrderApi.md#OrderGetPayUBankTransferDetails) | **Get** /api/v1/order/getPayUBankTransferDetails/{mainOrdersPK} | 
[**OrderGetPayUReceiptLink**](OrderApi.md#OrderGetPayUReceiptLink) | **Get** /api/v1/order/getPayUReceiptLink/{mainOrdersPK} | 
[**OrderGetPaymentInfo**](OrderApi.md#OrderGetPaymentInfo) | **Get** /api/v1/orders/{mainOrdersPK}/payments/info | Get payment info messages for UI
[**OrderGetPaymentMethods**](OrderApi.md#OrderGetPaymentMethods) | **Get** /api/v1/order/payment/{orderId} | Returns a list of payment methods for an order.
[**OrderGetPaymentOrderProcessingFees**](OrderApi.md#OrderGetPaymentOrderProcessingFees) | **Get** /api/v1/order/processingFees | Returns order processing fees by country code
[**OrderGetPayments**](OrderApi.md#OrderGetPayments) | **Get** /api/v1/orders/{mainOrdersPK}/payments | Get payments on order
[**OrderGetPayments_0**](OrderApi.md#OrderGetPayments_0) | **Get** /api/v1/payments | 
[**OrderGetRMADetail**](OrderApi.md#OrderGetRMADetail) | **Get** /api/v1/RMA/{RMANumber}/products | Get RMA Detail
[**OrderGetRMAReasons**](OrderApi.md#OrderGetRMAReasons) | **Get** /api/v1/RMA/reasons | Get GetRMA_Reasons
[**OrderGetRMAsProducts**](OrderApi.md#OrderGetRMAsProducts) | **Get** /api/v1/RMA/RMAs/Products | Get All RMAs and RMAproducts For Order by OrderFK
[**OrderGetRMAsStatus**](OrderApi.md#OrderGetRMAsStatus) | **Get** /api/v1/RMA/status | Get RMAs_Status
[**OrderGetRetailCartByRemarketingTrackingId**](OrderApi.md#OrderGetRetailCartByRemarketingTrackingId) | **Post** /api/v1/order/tracking/{TrackingId} | Used to load the retail cart info based on the re-marketing tracking id.
[**OrderGetRimanDollarsBalance**](OrderApi.md#OrderGetRimanDollarsBalance) | **Get** /api/v1/order/getRimanDollarsBalance/{mainPk}/{mainOrdersFk} | 
[**OrderGetSalesCampaignByCampaignCode**](OrderApi.md#OrderGetSalesCampaignByCampaignCode) | **Get** /api/v1/order/campaigns/info/{campaignCode} | Gets the sales campaign information for a particular campaign.
[**OrderGetSalesCampaignByCampaignCode_0**](OrderApi.md#OrderGetSalesCampaignByCampaignCode_0) | **Get** /api/v1/order/campaigns/info/{productBrandFK}/{countryCode} | Returns sales campaign info based on the campaign code.
[**OrderGetSalesCampaignByCampaignCode_1**](OrderApi.md#OrderGetSalesCampaignByCampaignCode_1) | **Get** /api/v1/order/campaigns/info/{productBrandFK}/{countryCode}/{campaignCode} | Returns sales campaign info based on the campaign code.
[**OrderGetSavedCreditCard**](OrderApi.md#OrderGetSavedCreditCard) | **Get** /api/v1/order/payment/get-card/{sequencePK} | Gets MainCreditCard record based on MainCreditCardsFK
[**OrderGetSiftConfig**](OrderApi.md#OrderGetSiftConfig) | **Get** /api/v1/order/get-sift-config | Used to get Sift configuration.
[**OrderGetTapPayConfig**](OrderApi.md#OrderGetTapPayConfig) | **Get** /api/v1/order/get-TapPay-config | Used to get InAuth configuration.
[**OrderGetTokenExConfig**](OrderApi.md#OrderGetTokenExConfig) | **Get** /api/v1/order/get-ixo-pay-config/{fromProfile} | Returns config for IxoPay tokenEx iframe
[**OrderGetWalletDiscount**](OrderApi.md#OrderGetWalletDiscount) | **Get** /api/v1/order/wallet/discount | Check if Wallet discount available
[**OrderGetWalletLimit**](OrderApi.md#OrderGetWalletLimit) | **Get** /api/v1/order/getWalletOrderLimit/{mainOrdersPK} | 
[**OrderInitInstallments**](OrderApi.md#OrderInitInstallments) | **Post** /api/v1/order/initInstallments | 
[**OrderIsAllowPersonalUseFun**](OrderApi.md#OrderIsAllowPersonalUseFun) | **Get** /api/v1/order/personal-use | To get information about Personal Use status
[**OrderIsCreditCardRedirectCheckout**](OrderApi.md#OrderIsCreditCardRedirectCheckout) | **Get** /api/v1/order/isCreditCardRedirectCheckout/{mainOrderPk} | Check if CC payment is redirect checkout payment
[**OrderIsInstallmentsExtensionValid**](OrderApi.md#OrderIsInstallmentsExtensionValid) | **Post** /api/v1/order/isInstallmentsExtensionValid | 
[**OrderIsIxoPayActiveForGateway**](OrderApi.md#OrderIsIxoPayActiveForGateway) | **Get** /api/v1/order/IsIxoPayActiveForGateway/{mainOrderTypeFk}/{countryCode}/{currencyCode} | Returns config for IxoPay tokenEx iframe
[**OrderIsLacoreFraudSettings**](OrderApi.md#OrderIsLacoreFraudSettings) | **Get** /api/v1/order/isLacoreFraudSettings/{mainOrderTypeFk}/{countryCode}/{currencyCode} | To get Lacore routing and AntiFraud settings
[**OrderIsMercadoPagoOrder**](OrderApi.md#OrderIsMercadoPagoOrder) | **Get** /api/v1/order/isMercadoPagoOrder/{mainOrderTypeFk}/{countryCode} | 
[**OrderIsNiceGtwEnabled**](OrderApi.md#OrderIsNiceGtwEnabled) | **Get** /api/v1/order/isNiceGtwEnabled/{mainOrderTypeFk} | 
[**OrderIsQRCodeOrderStatusUpdated**](OrderApi.md#OrderIsQRCodeOrderStatusUpdated) | **Get** /api/v1/order/isQRCodeOrderStatusUpdated | 
[**OrderIsRoutedToZiraat**](OrderApi.md#OrderIsRoutedToZiraat) | **Get** /api/v1/order/isRoutedToZiraat/{mainOrderTypeFk}/{countryCode} | 
[**OrderNewPayment**](OrderApi.md#OrderNewPayment) | **Post** /api/v1/orders/{mainOrdersPK}/payments | Adds and processes payments to order
[**OrderOrderHasPackage**](OrderApi.md#OrderOrderHasPackage) | **Get** /api/v1/order/{orderIdStr}/has-package | Checks to see if an order contains a package product
[**OrderProcessOrderResponse**](OrderApi.md#OrderProcessOrderResponse) | **Get** /api/v1/order/process-response | Used to process response from 3rd party redirects
[**OrderProcessOrderResponsePost**](OrderApi.md#OrderProcessOrderResponsePost) | **Post** /api/v1/order/process-response | 
[**OrderProcessPayPalOrder**](OrderApi.md#OrderProcessPayPalOrder) | **Post** /api/v1/order/paypal/get-process-url | Used to get redirect url to PayPal.
[**OrderProcessPayPalOrderResponse**](OrderApi.md#OrderProcessPayPalOrderResponse) | **Post** /api/v1/order/merchant/process-response | Used to process response from PayPal.
[**OrderSaveInstallments**](OrderApi.md#OrderSaveInstallments) | **Post** /api/v1/order/saveInstallments | 
[**OrderSaveTokenExConfig**](OrderApi.md#OrderSaveTokenExConfig) | **Put** /api/v1/order/save-ixo-pay-token/{mainOrdersFK} | Returns config for IxoPay tokenEx iframe
[**OrderSendSummury**](OrderApi.md#OrderSendSummury) | **Post** /api/v1/RMA/send-summary | Send summary to the distributor
[**OrderSetInvoiceInsuranceType**](OrderApi.md#OrderSetInvoiceInsuranceType) | **Post** /api/v1/order/set-invoice-insurance-type | Sets invoice insurance type
[**OrderSetOrderShipping**](OrderApi.md#OrderSetOrderShipping) | **Put** /api/v1/order/shipping/{mainOrdersFK}/{shippingChartTypeFK} | Sets the shipping information for an order.
[**OrderSetOrderTax**](OrderApi.md#OrderSetOrderTax) | **Put** /api/v1/order/taxes/{mainOrdersFK} | Gets and sets the tax information for an order.
[**OrderSetOrderType**](OrderApi.md#OrderSetOrderType) | **Post** /api/v1/order/update-type | Updates the order type based on cart information. Will not update paid orders.
[**OrderSetShippingOnOrder**](OrderApi.md#OrderSetShippingOnOrder) | **Put** /api/v1/order/set-shipping | 
[**OrderSubmitOrder**](OrderApi.md#OrderSubmitOrder) | **Post** /api/v1/order/complete | Finalize an order in process
[**OrderUpdateImportOrderType**](OrderApi.md#OrderUpdateImportOrderType) | **Put** /api/v1/order/import-order/{mainOrderPk}/{importOrder} | Updates ImportOrder of the specific order by unique identifier
[**OrderUpdateOrderMainType**](OrderApi.md#OrderUpdateOrderMainType) | **Put** /api/v1/order/main-type | 
[**OrderUpdateOrderShippingAddress**](OrderApi.md#OrderUpdateOrderShippingAddress) | **Put** /api/v1/order/shipping/address/{mainOrdersFK} | Updates the shipping address informaiton for an order.
[**OrderUpdatePayment**](OrderApi.md#OrderUpdatePayment) | **Put** /api/v1/orders/{mainOrdersPK}/payments/{paymentMethodPK} | Updates payments on order
[**OrderUpdatePersonalUseFlag**](OrderApi.md#OrderUpdatePersonalUseFlag) | **Post** /api/v1/orders/{mainOrdersPK}/update-personal-use | Used to update the personal use only flag on main orders.
[**OrderUpdateRMA**](OrderApi.md#OrderUpdateRMA) | **Put** /api/v1/RMA | Update RMA
[**OrderUseCrossBorderCharges**](OrderApi.md#OrderUseCrossBorderCharges) | **Get** /api/v1/order/useCrossBorderCharges/{shipTo}/{shipFrom} | 
[**OrderUseSavedCreditCard**](OrderApi.md#OrderUseSavedCreditCard) | **Post** /api/v1/order/payment/set-card | Updates MainCreditCardsFK column in MainOrders Table
[**OrderValidateCreditCard**](OrderApi.md#OrderValidateCreditCard) | **Post** /api/v1/order/creditcard/validate | Validate Credit Card Number
[**OrderValidateDiscountCode**](OrderApi.md#OrderValidateDiscountCode) | **Post** /api/v1/orders/{mainOrdersPK}/validate | Used to validate an order information before it is submitted.
[**OrderVerifyDiscountCode**](OrderApi.md#OrderVerifyDiscountCode) | **Get** /api/v1/orders/verify-discount-code | Used to verify a discount code for a logged in distributor.
[**OrderVerifyUserOrders**](OrderApi.md#OrderVerifyUserOrders) | **Get** /api/v1/order/{mainpk}/has-completed | 


# **OrderAddCCPaymentToOrder**
> interface{} OrderAddCCPaymentToOrder(ctx, model)
Adds the Credit Card payment information to an order in process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**AddCcPaymentFormModel**](AddCcPaymentFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderAddLogOrderStatus**
> interface{} OrderAddLogOrderStatus(ctx, model)
Used to add a status update log to an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**AddLogOrderStatusFormModel**](AddLogOrderStatusFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderAddMainRetailInfo**
> interface{} OrderAddMainRetailInfo(ctx, model)
Used to set retail tracking information for retails orders once they have been completed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**AddMainRetailInfoFormModel**](AddMainRetailInfoFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderAddProcessingFees**
> float64 OrderAddProcessingFees(ctx, model)
Adds processing fees to an order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**AddProcessingFeesFormModel**](AddProcessingFeesFormModel.md)|  | 

### Return type

**float64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderAddProduct**
> interface{} OrderAddProduct(ctx, model)
Remove items from an existing, in process order. It will remove all existing items, if any, and add the items provided to the order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**AddProductFormModel**](AddProductFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderApplyDiscountCode**
> interface{} OrderApplyDiscountCode(ctx, mainOrdersPK, model)
Used to apply a discount code to an order for a distributor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**| Main Orders PK | 
  **model** | [**ApplyOrderDiscountFormModel**](ApplyOrderDiscountFormModel.md)| Form Model | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderApplyRimanDollarsPayment**
> PaymentProcessingResponse OrderApplyRimanDollarsPayment(ctx, mainOrdersPK, payment)
Used to apply a Riman dollars to an order for a distributor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 
  **payment** | [**RimanDollarsPaymentModel**](RimanDollarsPaymentModel.md)|  | 

### Return type

[**PaymentProcessingResponse**](PaymentProcessingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCancelMerchantOrder**
> interface{} OrderCancelMerchantOrder(ctx, model)
Used to cancelation PayPal Order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**OrderResponseModel**](OrderResponseModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCheckOrderPaidStatus**
> interface{} OrderCheckOrderPaidStatus(ctx, orderId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCheckOrderPaymentStarted**
> interface{} OrderCheckOrderPaymentStarted(ctx, orderId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderClearMerchantPendingStatus**
> interface{} OrderClearMerchantPendingStatus(ctx, orderId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderConvertToLocalCurrency**
> interface{} OrderConvertToLocalCurrency(ctx, mainOrdersFK, amount)
Returns converted amount to local currency

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFK** | **int32**|  | 
  **amount** | **float64**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCreateAutoshipOrder**
> AutoshipOrderInfoViewModel OrderCreateAutoshipOrder(ctx, )
Runs the manual process to create an autoship order for the currently logged in customer.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AutoshipOrderInfoViewModel**](AutoshipOrderInfoViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCreateOrder**
> interface{} OrderCreateOrder(ctx, model)
Create a new order with basic information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**NewOrderFormModel**](NewOrderFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCreateOrderClimbersClub**
> interface{} OrderCreateOrderClimbersClub(ctx, )
Finalize climbers club sign up. Creates order and takes money from cliffs wallet. Returns success or failure

### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCreatePrepaidOrder**
> AutoshipOrderInfoViewModel OrderCreatePrepaidOrder(ctx, products)
Runs the manual process to create an prepaid order for the currently logged in customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **products** | [**[]CartItemFormModel**](CartItemFormModel.md)|  | 

### Return type

[**AutoshipOrderInfoViewModel**](AutoshipOrderInfoViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCreateRMA**
> interface{} OrderCreateRMA(ctx, newRMA)
Create RMA

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newRMA** | [**RmaFormModel**](RmaFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderDeleteOrder**
> interface{} OrderDeleteOrder(ctx, orderId)
This method is no longer allowed please do not use

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int32**| Unique identifier of the order | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGet3Ds2**
> interface{} OrderGet3Ds2(ctx, model)
Get redirect Url for 3D 2 Secure payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**OrderComplete3Ds2Model**](OrderComplete3Ds2Model.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGet3DsTokenUrl**
> interface{} OrderGet3DsTokenUrl(ctx, model)
Get token url/model for token processing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**OrderCompleteTokenModel**](OrderCompleteTokenModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGet3DsUrl**
> interface{} OrderGet3DsUrl(ctx, model)
Get redirect Url for 3D Secure payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**OrderCompleteModel**](OrderCompleteModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetAcceptedCreditCardTypes**
> interface{} OrderGetAcceptedCreditCardTypes(ctx, countryCode, currencyCode, optional)
Returns a list of accepted credit card types for the specified country code and currency code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**| Example: US | 
  **currencyCode** | **string**| Example: USD | 
 **optional** | ***OrderApiOrderGetAcceptedCreditCardTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetAcceptedCreditCardTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mainOrderTypeFK** | **optional.Int32**| Main Order Type PK for the request | 
 **mainOrderPk** | **optional.Int32**| Main Order PK for the request | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetAvailableCoupons**
> []AvailableCouponModel OrderGetAvailableCoupons(ctx, mainPk, mainOrdersPK)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 
  **mainOrdersPK** | **int32**|  | 

### Return type

[**[]AvailableCouponModel**](AvailableCouponModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetBonusCreditLimit**
> SignupTokenOrderLimit OrderGetBonusCreditLimit(ctx, mainOrdersPK)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

[**SignupTokenOrderLimit**](SignupTokenOrderLimit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetBraspagAntifraudScirpts**
> BraspagAntifraudScirptModel OrderGetBraspagAntifraudScirpts(ctx, identifier)
Used to get InAuth configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**|  | 

### Return type

[**BraspagAntifraudScirptModel**](BraspagAntifraudScirptModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetCustomsPaymentWarnings**
> []CustomsPaymentWarningModel OrderGetCustomsPaymentWarnings(ctx, optional)
Returns messages related to customs payment warning for the warehouse country

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderGetCustomsPaymentWarningsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetCustomsPaymentWarningsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **optional.String**|  | 
 **culture** | **optional.String**|  | 

### Return type

[**[]CustomsPaymentWarningModel**](CustomsPaymentWarningModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetDonationOrders**
> FullDonationReport OrderGetDonationOrders(ctx, mainFk)
To get information about Donations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainFk** | **int32**|  | 

### Return type

[**FullDonationReport**](FullDonationReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetInAuthConfig**
> InAuthConfigModel OrderGetInAuthConfig(ctx, )
Used to get InAuth configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InAuthConfigModel**](InAuthConfigModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetInstallments**
> MerchantInstallmentsConfigModel OrderGetInstallments(ctx, model)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**MerchantInstallmentsFormModel**](MerchantInstallmentsFormModel.md)|  | 

### Return type

[**MerchantInstallmentsConfigModel**](MerchantInstallmentsConfigModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetMainOrders**
> MainOrderSimpleViewModel OrderGetMainOrders(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderGetMainOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetMainOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainId** | **optional.String**|  | 
 **mainSiteUrl** | **optional.String**|  | 
 **orderNumber** | **optional.String**|  | 
 **shipmentNumber** | **optional.String**|  | 
 **trackingNumber** | **optional.String**|  | 
 **getCustomerOrders** | **optional.Bool**|  | 
 **getAbandonedOrders** | **optional.Bool**|  | 
 **getEnrollerOrders** | **optional.Bool**|  | 
 **getDownlineOrders** | **optional.Bool**|  | 
 **paidStatus** | **optional.Bool**|  | 
 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 
 **orderType** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **productPKs** | **optional.String**|  | 
 **isRefunded** | **optional.Bool**|  | 
 **countryCode** | **optional.String**|  | 
 **salesCampaignPK** | **optional.Int32**|  | 
 **shippingStatus** | **optional.String**|  | 
 **shouldIncludeShippingDetail** | **optional.Bool**|  | 
 **shouldIncludeOrderItems** | **optional.Bool**|  | 
 **shouldIncludePayments** | **optional.Bool**|  | 
 **shouldCheckIfPrepaid** | **optional.Bool**|  | 
 **orderLevel** | **optional.String**|  | 
 **getDataToExport** | **optional.Bool**|  | 
 **customerOrdersMainId** | **optional.Int32**|  | 
 **weChatOrderNumber** | **optional.String**|  | 
 **memberID** | **optional.String**|  | 

### Return type

[**MainOrderSimpleViewModel**](MainOrderSimpleViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetManualPayTypes**
> []ManualPaymentTypeModel OrderGetManualPayTypes(ctx, )
Retrieves available manual paytypes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ManualPaymentTypeModel**](ManualPaymentTypeModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetMerchantRedirectModel**
> MerchantRedirectViewModel OrderGetMerchantRedirectModel(ctx, model)
Return merchant redirect model with required parameters for UI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**MerchantRedirectFormModel**](MerchantRedirectFormModel.md)|  | 

### Return type

[**MerchantRedirectViewModel**](MerchantRedirectViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetMerchantRedirectModelStep2**
> MerchantRedirectModelStep2 OrderGetMerchantRedirectModelStep2(ctx, model)
Get script\\iframe for redirect step 2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**MerchantRedirectFormModelStep2**](MerchantRedirectFormModelStep2.md)|  | 

### Return type

[**MerchantRedirectModelStep2**](MerchantRedirectModelStep2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetMinimumAllowedPayment**
> float64 OrderGetMinimumAllowedPayment(ctx, mainOrdersPK)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

**float64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOpenPayReceiptLink**
> string OrderGetOpenPayReceiptLink(ctx, mainOrdersPK)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrSearchRMAs**
> RmaView OrderGetOrSearchRMAs(ctx, optional)
To get items from RMA table order OrderByDescending Issue_Date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderGetOrSearchRMAsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetOrSearchRMAsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryRmaNumber** | **optional.Int32**|  | 
 **queryOrderNumber** | **optional.Int32**|  | 
 **queryCustomerSiteurl** | **optional.String**|  | 
 **queryCustomerName** | **optional.String**|  | 
 **queryIssuedBy** | **optional.String**|  | 
 **queryIssuedDate** | **optional.Time**|  | 
 **queryJnsDepartment** | **optional.String**|  | 
 **queryStatusId** | **optional.String**|  | 
 **queryRecivedDate** | **optional.Time**|  | 
 **queryRefundDate** | **optional.Time**|  | 
 **queryOffset** | **optional.Int32**|  | 
 **queryLimit** | **optional.Int32**|  | 
 **queryOrderBy** | **optional.String**|  | 

### Return type

[**RmaView**](RMAView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderDetailsForConfirmation**
> OrderConfirmationDetailViewModel OrderGetOrderDetailsForConfirmation(ctx, orderString)
Get Order Details for confirmation Page.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderString** | **string**| Order Id string | 

### Return type

[**OrderConfirmationDetailViewModel**](OrderConfirmationDetailViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderDisclosure**
> interface{} OrderGetOrderDisclosure(ctx, orderId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderId**
> MainOrderModel OrderGetOrderId(ctx, mainOrdersPK)
To get single order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

[**MainOrderModel**](MainOrderModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderInfo**
> interface{} OrderGetOrderInfo(ctx, mainOrdersFK)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFK** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderItems**
> []MainOrderItemModel OrderGetOrderItems(ctx, mainOrdersPK)
To get main order items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

[**[]MainOrderItemModel**](MainOrderItemModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderProducts**
> interface{} OrderGetOrderProducts(ctx, mainOrdersPK)
To get all subitems from main order by id for RMA

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderShipmentProducts**
> interface{} OrderGetOrderShipmentProducts(ctx, mainOrdersPK)
To get shipment data of orders products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderShipmentsTotal**
> []OrderShipmentsTotalModel OrderGetOrderShipmentsTotal(ctx, mainOrdersPK)
To get shipments total cost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

[**[]OrderShipmentsTotalModel**](OrderShipmentsTotalModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderShipping**
> interface{} OrderGetOrderShipping(ctx, mainOrdersFK, shippingChartTypeFK)
Gets the shipping information for an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFK** | **int32**| OrderNumber | 
  **shippingChartTypeFK** | **int32**| Shipping Key | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderShippingDetails**
> []MainOrderTrackingModel OrderGetOrderShippingDetails(ctx, mainOrdersPK)
To get shipping Tracking information of an order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

[**[]MainOrderTrackingModel**](MainOrderTrackingModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderTaxAmount**
> interface{} OrderGetOrderTaxAmount(ctx, mainOrdersFK)
Gets the tax information for an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFK** | **int32**| OrderNumber | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderTypes**
> MainOrderTypeModel OrderGetOrderTypes(ctx, )
To get order types

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MainOrderTypeModel**](MainOrderTypeModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrders**
> interface{} OrderGetOrders(ctx, entityFk, isMain)
Gets order history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityFk** | **int32**|  | 
  **isMain** | **bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetPayUBankTransferDetails**
> []PayUBankTransferDetails OrderGetPayUBankTransferDetails(ctx, mainOrdersPK)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

[**[]PayUBankTransferDetails**](PayUBankTransferDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetPayUReceiptLink**
> string OrderGetPayUReceiptLink(ctx, mainOrdersPK)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetPaymentInfo**
> []PaymentInfoResponseModel OrderGetPaymentInfo(ctx, mainOrdersPK)
Get payment info messages for UI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

[**[]PaymentInfoResponseModel**](PaymentInfoResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetPaymentMethods**
> PaymentMethodsViewModel OrderGetPaymentMethods(ctx, orderId)
Returns a list of payment methods for an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| Order Number | 

### Return type

[**PaymentMethodsViewModel**](PaymentMethodsViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetPaymentOrderProcessingFees**
> PaymentOrderProcessingFeesModel OrderGetPaymentOrderProcessingFees(ctx, optional)
Returns order processing fees by country code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderGetPaymentOrderProcessingFeesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetPaymentOrderProcessingFeesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **optional.String**|  | 
 **culture** | **optional.String**|  | 

### Return type

[**PaymentOrderProcessingFeesModel**](PaymentOrderProcessingFeesModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetPayments**
> []OrderPayments OrderGetPayments(ctx, mainOrdersPK)
Get payments on order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

[**[]OrderPayments**](OrderPayments.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetPayments_0**
> PaymentSimpleViewModel OrderGetPayments_0(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderGetPayments_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetPayments_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paidStatusFK** | **optional.Int32**|  | 
 **paymentStartDate** | **optional.Time**|  | 
 **paymentEndDate** | **optional.Time**|  | 
 **orderStartDate** | **optional.Time**|  | 
 **orderEndDate** | **optional.Time**|  | 
 **paymentMethodCode** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **page** | **optional.Int32**|  | 

### Return type

[**PaymentSimpleViewModel**](PaymentSimpleViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetRMADetail**
> interface{} OrderGetRMADetail(ctx, rMANumber)
Get RMA Detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rMANumber** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetRMAReasons**
> interface{} OrderGetRMAReasons(ctx, )
Get GetRMA_Reasons

### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetRMAsProducts**
> RmaView OrderGetRMAsProducts(ctx, orderFK)
Get All RMAs and RMAproducts For Order by OrderFK

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderFK** | **int32**|  | 

### Return type

[**RmaView**](RMAView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetRMAsStatus**
> interface{} OrderGetRMAsStatus(ctx, )
Get RMAs_Status

### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetRetailCartByRemarketingTrackingId**
> OrderInfoModel OrderGetRetailCartByRemarketingTrackingId(ctx, trackingId)
Used to load the retail cart info based on the re-marketing tracking id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackingId** | **string**| GUID: Email Tracking Id | 

### Return type

[**OrderInfoModel**](OrderInfoModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetRimanDollarsBalance**
> float64 OrderGetRimanDollarsBalance(ctx, mainPk, mainOrdersFk)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 
  **mainOrdersFk** | **int32**|  | 

### Return type

**float64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetSalesCampaignByCampaignCode**
> SalesCampaignModel OrderGetSalesCampaignByCampaignCode(ctx, campaignCode)
Gets the sales campaign information for a particular campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignCode** | **string**| Campaign Code to Lookup | 

### Return type

[**SalesCampaignModel**](SalesCampaignModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetSalesCampaignByCampaignCode_0**
> SalesCampaignModel OrderGetSalesCampaignByCampaignCode_0(ctx, productBrandFK, countryCode, optional)
Returns sales campaign info based on the campaign code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productBrandFK** | **int32**| Product Brand PK | 
  **countryCode** | **string**| Country Code | 
 **optional** | ***OrderApiOrderGetSalesCampaignByCampaignCode_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetSalesCampaignByCampaignCode_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cartType** | **optional.String**|  | 
 **culture** | **optional.String**|  | 
 **siteUrl** | **optional.String**|  | 

### Return type

[**SalesCampaignModel**](SalesCampaignModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetSalesCampaignByCampaignCode_1**
> SalesCampaignModel OrderGetSalesCampaignByCampaignCode_1(ctx, productBrandFK, countryCode, campaignCode, optional)
Returns sales campaign info based on the campaign code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productBrandFK** | **int32**| Product Brand PK | 
  **countryCode** | **string**| Country Code | 
  **campaignCode** | **string**| Campaign Code | 
 **optional** | ***OrderApiOrderGetSalesCampaignByCampaignCode_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetSalesCampaignByCampaignCode_3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cartType** | **optional.String**|  | 
 **culture** | **optional.String**|  | 
 **siteUrl** | **optional.String**|  | 

### Return type

[**SalesCampaignModel**](SalesCampaignModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetSavedCreditCard**
> interface{} OrderGetSavedCreditCard(ctx, sequencePK)
Gets MainCreditCard record based on MainCreditCardsFK

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sequencePK** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetSiftConfig**
> SiftConfigModel OrderGetSiftConfig(ctx, )
Used to get Sift configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SiftConfigModel**](SiftConfigModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetTapPayConfig**
> InAuthConfigModel OrderGetTapPayConfig(ctx, )
Used to get InAuth configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InAuthConfigModel**](InAuthConfigModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetTokenExConfig**
> interface{} OrderGetTokenExConfig(ctx, fromProfile)
Returns config for IxoPay tokenEx iframe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fromProfile** | **bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetWalletDiscount**
> WalletDiscountModel OrderGetWalletDiscount(ctx, optional)
Check if Wallet discount available

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderGetWalletDiscountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetWalletDiscountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainOrdersFK** | **optional.Int32**|  | 
 **countryCode** | **optional.String**|  | 

### Return type

[**WalletDiscountModel**](WalletDiscountModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetWalletLimit**
> WalletOrderLimit OrderGetWalletLimit(ctx, mainOrdersPK)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 

### Return type

[**WalletOrderLimit**](WalletOrderLimit.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderInitInstallments**
> MerchantInstallmentsConfigModel OrderInitInstallments(ctx, model)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**MerchantInstallmentsFormModel**](MerchantInstallmentsFormModel.md)|  | 

### Return type

[**MerchantInstallmentsConfigModel**](MerchantInstallmentsConfigModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderIsAllowPersonalUseFun**
> PersonalUseModel OrderIsAllowPersonalUseFun(ctx, optional)
To get information about Personal Use status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderIsAllowPersonalUseFunOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderIsAllowPersonalUseFunOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainOrdersPK** | **optional.Int32**|  | 
 **cartKey** | **optional.String**|  | 

### Return type

[**PersonalUseModel**](PersonalUseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderIsCreditCardRedirectCheckout**
> IsCreditCardRedirectCheckoutModel OrderIsCreditCardRedirectCheckout(ctx, mainOrderPk)
Check if CC payment is redirect checkout payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrderPk** | **int32**|  | 

### Return type

[**IsCreditCardRedirectCheckoutModel**](IsCreditCardRedirectCheckoutModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderIsInstallmentsExtensionValid**
> bool OrderIsInstallmentsExtensionValid(ctx, model)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**MerchantInstallmentsFormModel**](MerchantInstallmentsFormModel.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderIsIxoPayActiveForGateway**
> interface{} OrderIsIxoPayActiveForGateway(ctx, mainOrderTypeFk, countryCode, currencyCode)
Returns config for IxoPay tokenEx iframe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrderTypeFk** | **int32**|  | 
  **countryCode** | **string**|  | 
  **currencyCode** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderIsLacoreFraudSettings**
> string OrderIsLacoreFraudSettings(ctx, mainOrderTypeFk, countryCode, currencyCode)
To get Lacore routing and AntiFraud settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrderTypeFk** | **int32**|  | 
  **countryCode** | **string**|  | 
  **currencyCode** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderIsMercadoPagoOrder**
> bool OrderIsMercadoPagoOrder(ctx, mainOrderTypeFk, countryCode)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrderTypeFk** | **int32**|  | 
  **countryCode** | **string**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderIsNiceGtwEnabled**
> bool OrderIsNiceGtwEnabled(ctx, mainOrderTypeFk)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrderTypeFk** | **int32**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderIsQRCodeOrderStatusUpdated**
> bool OrderIsQRCodeOrderStatusUpdated(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderIsQRCodeOrderStatusUpdatedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderIsQRCodeOrderStatusUpdatedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainOrdersFK** | **optional.Int32**|  | 
 **paymentMethodCode** | **optional.String**|  | 
 **paymentReference** | **optional.String**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderIsRoutedToZiraat**
> bool OrderIsRoutedToZiraat(ctx, mainOrderTypeFk, countryCode)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrderTypeFk** | **int32**|  | 
  **countryCode** | **string**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderNewPayment**
> PaymentProcessingResponse OrderNewPayment(ctx, mainOrdersPK, payment)
Adds and processes payments to order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 
  **payment** | [**PaymentFormModel**](PaymentFormModel.md)|  | 

### Return type

[**PaymentProcessingResponse**](PaymentProcessingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderOrderHasPackage**
> bool OrderOrderHasPackage(ctx, orderIdStr)
Checks to see if an order contains a package product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderIdStr** | **string**| Order Number | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderProcessOrderResponse**
> interface{} OrderProcessOrderResponse(ctx, )
Used to process response from 3rd party redirects

### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderProcessOrderResponsePost**
> interface{} OrderProcessOrderResponsePost(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderProcessPayPalOrder**
> interface{} OrderProcessPayPalOrder(ctx, model)
Used to get redirect url to PayPal.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**OrderProcessModel**](OrderProcessModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderProcessPayPalOrderResponse**
> PaymentProcessingResponse OrderProcessPayPalOrderResponse(ctx, model)
Used to process response from PayPal.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**OrderResponseModel**](OrderResponseModel.md)|  | 

### Return type

[**PaymentProcessingResponse**](PaymentProcessingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSaveInstallments**
> MerchantInstallmentsConfigModel OrderSaveInstallments(ctx, model)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**MerchantInstallmentsFormModel**](MerchantInstallmentsFormModel.md)|  | 

### Return type

[**MerchantInstallmentsConfigModel**](MerchantInstallmentsConfigModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSaveTokenExConfig**
> interface{} OrderSaveTokenExConfig(ctx, model, mainOrdersFK)
Returns config for IxoPay tokenEx iframe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**IxoPayTokenModel**](IxoPayTokenModel.md)|  | 
  **mainOrdersFK** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSendSummury**
> interface{} OrderSendSummury(ctx, rmaInfo)
Send summary to the distributor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rmaInfo** | [**RmaSendSummaryInfoModel**](RmaSendSummaryInfoModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSetInvoiceInsuranceType**
> interface{} OrderSetInvoiceInsuranceType(ctx, model)
Sets invoice insurance type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**InvoiceInsuranceTypeModel**](InvoiceInsuranceTypeModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSetOrderShipping**
> interface{} OrderSetOrderShipping(ctx, mainOrdersFK, shippingChartTypeFK)
Sets the shipping information for an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFK** | **int32**| OrderNumber | 
  **shippingChartTypeFK** | **int32**| Shipping Key | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSetOrderTax**
> interface{} OrderSetOrderTax(ctx, mainOrdersFK)
Gets and sets the tax information for an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFK** | **int32**| OrderNumber | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSetOrderType**
> UpdateOrderTypeViewModel OrderSetOrderType(ctx, model)
Updates the order type based on cart information. Will not update paid orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**UpdateOrderTypeFormModel**](UpdateOrderTypeFormModel.md)| SetOrderTypeFormModel | 

### Return type

[**UpdateOrderTypeViewModel**](UpdateOrderTypeViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSetShippingOnOrder**
> interface{} OrderSetShippingOnOrder(ctx, model)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**SetShippingFormModel**](SetShippingFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSubmitOrder**
> interface{} OrderSubmitOrder(ctx, model)
Finalize an order in process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**OrderCompleteModel**](OrderCompleteModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUpdateImportOrderType**
> interface{} OrderUpdateImportOrderType(ctx, mainOrderPk, importOrder)
Updates ImportOrder of the specific order by unique identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrderPk** | **int32**|  | 
  **importOrder** | **bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUpdateOrderMainType**
> interface{} OrderUpdateOrderMainType(ctx, model)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**OrderMainTypeFormModel**](OrderMainTypeFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUpdateOrderShippingAddress**
> interface{} OrderUpdateOrderShippingAddress(ctx, mainOrdersFK, model)
Updates the shipping address informaiton for an order.

Required Fields: FirstName, LastName, Address1, City, State, PostalCode, Country <param name=\"mainOrdersFK\">OrderNumber</param><param name=\"model\">Address Information</param>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFK** | **int32**|  | 
  **model** | [**UpdateOrderShipAddressFormModel**](UpdateOrderShipAddressFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUpdatePayment**
> PaymentProcessingResponse OrderUpdatePayment(ctx, mainOrdersPK, paymentMethodPK, payment)
Updates payments on order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**|  | 
  **paymentMethodPK** | **int32**|  | 
  **payment** | [**PaymentUpdateFormModel**](PaymentUpdateFormModel.md)|  | 

### Return type

[**PaymentProcessingResponse**](PaymentProcessingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUpdatePersonalUseFlag**
> interface{} OrderUpdatePersonalUseFlag(ctx, mainOrdersPK, model)
Used to update the personal use only flag on main orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **string**| Main Orders PK | 
  **model** | [**UpdatePersonalUseFormModel**](UpdatePersonalUseFormModel.md)| Form Model | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUpdateRMA**
> interface{} OrderUpdateRMA(ctx, rmaUpdateQueryModel)
Update RMA

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rmaUpdateQueryModel** | [**RmaUpdateQueryModel**](RmaUpdateQueryModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUseCrossBorderCharges**
> interface{} OrderUseCrossBorderCharges(ctx, shipTo, shipFrom)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shipTo** | **string**|  | 
  **shipFrom** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUseSavedCreditCard**
> interface{} OrderUseSavedCreditCard(ctx, model)
Updates MainCreditCardsFK column in MainOrders Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**SavedCcModel**](SavedCcModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderValidateCreditCard**
> ValidateCreditCardResponse OrderValidateCreditCard(ctx, model)
Validate Credit Card Number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**ValidateCreditCardModel**](ValidateCreditCardModel.md)|  | 

### Return type

[**ValidateCreditCardResponse**](ValidateCreditCardResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderValidateDiscountCode**
> ValidateOrderViewModel OrderValidateDiscountCode(ctx, mainOrdersPK)
Used to validate an order information before it is submitted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **int32**| Main Orders PK | 

### Return type

[**ValidateOrderViewModel**](ValidateOrderViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderVerifyDiscountCode**
> VerifyDiscountCodeReturnModel OrderVerifyDiscountCode(ctx, optional)
Used to verify a discount code for a logged in distributor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderVerifyDiscountCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderVerifyDiscountCodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discountCode** | **optional.String**|  | 
 **mainOrderTypeFk** | **optional.Int32**|  | 
 **mainOrderPk** | **optional.Int32**|  | 

### Return type

[**VerifyDiscountCodeReturnModel**](VerifyDiscountCodeReturnModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderVerifyUserOrders**
> interface{} OrderVerifyUserOrders(ctx, mainpk)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainpk** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


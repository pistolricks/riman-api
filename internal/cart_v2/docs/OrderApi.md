# \OrderApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderAddCCPaymentToOrder**](OrderApi.md#OrderAddCCPaymentToOrder) | **Post** /api/v2/order/payment/creditcard | Adds the Credit Card payment information to an order in process
[**OrderAddProduct**](OrderApi.md#OrderAddProduct) | **Post** /api/v2/order/product | Adds a product to an order.
[**OrderApplyInterativeDiscount**](OrderApi.md#OrderApplyInterativeDiscount) | **Post** /api/v2/order/apply-iterative-discount | Used to auto apply main discount config discount if applicable
[**OrderCreateOrder**](OrderApi.md#OrderCreateOrder) | **Post** /api/v2/order | Create a new order with basic information
[**OrderGetOrderDetails**](OrderApi.md#OrderGetOrderDetails) | **Get** /api/v2/orders/{MainOrdersPK}/details | Retrieves detailed information for a specific order by its primary key.
[**OrderGetOrderTaxAmountString**](OrderApi.md#OrderGetOrderTaxAmountString) | **Get** /api/v2/order/taxes/{mainOrdersFk} | Gets the tax information for an order.
[**OrderGetOrderType**](OrderApi.md#OrderGetOrderType) | **Get** /api/v2/order/type | Retrieves the order type and whether it is an event based on the provided query model.
[**OrderSetShippingOnAutoShipOrder**](OrderApi.md#OrderSetShippingOnAutoShipOrder) | **Put** /api/v2/order/set-shipping-autoship/{mainOrdersPK} | Sets the shipping info on an autoship order.
[**OrderSetShippingOnOrder**](OrderApi.md#OrderSetShippingOnOrder) | **Put** /api/v2/order/set-shipping | Sets the shipping info on an order.
[**OrderUpdateOrderItemsFromCart**](OrderApi.md#OrderUpdateOrderItemsFromCart) | **Put** /api/v2/order/items/add-from-cart/{mainOrdersFk}/{cartKey} | Used to move the items in the user&#39;s cart to an order.
[**OrderUpdateOrderShippingAddress**](OrderApi.md#OrderUpdateOrderShippingAddress) | **Put** /api/v2/order/shipping/address/{mainOrdersFK} | Updates the shipping address on an order record.
[**OrderValidateRetailSignup**](OrderApi.md#OrderValidateRetailSignup) | **Get** /api/v2/order/validate-retail-signup | Use to check if a retail customer has purchased an upgrade to distributor package and can continue to signup.


# **OrderAddCCPaymentToOrder**
> interface{} OrderAddCCPaymentToOrder(ctx, model)
Adds the Credit Card payment information to an order in process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**AddCcPaymentAndBillingFormModel**](AddCcPaymentAndBillingFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderAddProduct**
> interface{} OrderAddProduct(ctx, model)
Adds a product to an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**AddProductFormModelv2**](AddProductFormModelv2.md)| AddProductFormModelv2 model object | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderApplyInterativeDiscount**
> interface{} OrderApplyInterativeDiscount(ctx, model)
Used to auto apply main discount config discount if applicable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**ApplyIterativeDiscountFormModel**](ApplyIterativeDiscountFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderCreateOrder**
> CreateOrderViewModel OrderCreateOrder(ctx, model)
Create a new order with basic information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**NewOrderFormModel**](NewOrderFormModel.md)|  | 

### Return type

[**CreateOrderViewModel**](CreateOrderViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderDetails**
> interface{} OrderGetOrderDetails(ctx, mainOrdersPK, optional)
Retrieves detailed information for a specific order by its primary key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersPK** | **string**|  | 
 **optional** | ***OrderApiOrderGetOrderDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetOrderDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelMainOrdersPK** | **optional.Int32**|  | 
 **modelEmail** | **optional.String**|  | 
 **modelCulture** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderTaxAmountString**
> TaxModel OrderGetOrderTaxAmountString(ctx, mainOrdersFk)
Gets the tax information for an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFk** | **int32**| OrderNumber | 

### Return type

[**TaxModel**](TaxModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderGetOrderType**
> interface{} OrderGetOrderType(ctx, optional)
Retrieves the order type and whether it is an event based on the provided query model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderGetOrderTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderGetOrderTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelMainOrdersPK** | **optional.Int32**|  | 
 **modelEmail** | **optional.String**|  | 
 **modelCulture** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSetShippingOnAutoShipOrder**
> interface{} OrderSetShippingOnAutoShipOrder(ctx, mainOrdersPK)
Sets the shipping info on an autoship order.

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

# **OrderSetShippingOnOrder**
> interface{} OrderSetShippingOnOrder(ctx, model)
Sets the shipping info on an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**SetShippingFormModel**](SetShippingFormModel.md)| SetShippingFormModel model object. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUpdateOrderItemsFromCart**
> interface{} OrderUpdateOrderItemsFromCart(ctx, mainOrdersFk, cartKey)
Used to move the items in the user's cart to an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFk** | **int32**| Main Orders PK | 
  **cartKey** | [**string**](.md)| MainCart Unique Identifier | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderUpdateOrderShippingAddress**
> interface{} OrderUpdateOrderShippingAddress(ctx, mainOrdersFK, model)
Updates the shipping address on an order record.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFK** | **int32**| Main Order PK | 
  **model** | [**UpdateOrderShipAddressFormModel**](UpdateOrderShipAddressFormModel.md)| UpdateOrderShipAddressFormModel model | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderValidateRetailSignup**
> CheckRetailUpgrageViewModel OrderValidateRetailSignup(ctx, optional)
Use to check if a retail customer has purchased an upgrade to distributor package and can continue to signup.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrderApiOrderValidateRetailSignupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrderApiOrderValidateRetailSignupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainFK** | **optional.Int32**|  | 

### Return type

[**CheckRetailUpgrageViewModel**](CheckRetailUpgrageViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


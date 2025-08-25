# \ShippingApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShippingGetShippingMethods**](ShippingApi.md#ShippingGetShippingMethods) | **Get** /api/v1/shipping/methods/{mainOrdersFK}/{countryCode} | Grabs a list of shipping methods for a given order.  {mainOrdersFK}{countryCode}
[**ShippingGetShippingSettings**](ShippingApi.md#ShippingGetShippingSettings) | **Get** /api/v1/shipping/chart-settings/{mainOrdersFk}/{shippingChartTypeFk} | Returns shipping methods for an order
[**ShippingIsAddressValid**](ShippingApi.md#ShippingIsAddressValid) | **Post** /api/v1/shipping/address/validate | Calls the google address validation and returns a status and collection of similar addresses


# **ShippingGetShippingMethods**
> interface{} ShippingGetShippingMethods(ctx, mainOrdersFK, countryCode)
Grabs a list of shipping methods for a given order.  {mainOrdersFK}{countryCode}

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFK** | **int32**|  | 
  **countryCode** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShippingGetShippingSettings**
> interface{} ShippingGetShippingSettings(ctx, mainOrdersFk, shippingChartTypeFk)
Returns shipping methods for an order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainOrdersFk** | **int32**|  | 
  **shippingChartTypeFk** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShippingIsAddressValid**
> AddressValidationModel ShippingIsAddressValid(ctx, address)
Calls the google address validation and returns a status and collection of similar addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | [**AddressModel**](AddressModel.md)|  | 

### Return type

[**AddressValidationModel**](AddressValidationModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# \ShoppingAddressApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShoppingAddressAddBillingAddress**](ShoppingAddressApi.md#ShoppingAddressAddBillingAddress) | **Post** /api/v1/shopping/{cartKey}/address/billing | Adds a billing address record to the shopping cart
[**ShoppingAddressAddMailingAddress**](ShoppingAddressApi.md#ShoppingAddressAddMailingAddress) | **Post** /api/v1/shopping/{cartKey}/address/mailing | Adds a mailing address record to the shopping cart
[**ShoppingAddressAddShippingAddress**](ShoppingAddressApi.md#ShoppingAddressAddShippingAddress) | **Post** /api/v1/shopping/{cartKey}/address/shipping | Adds a shipping address record to the shopping cart
[**ShoppingAddressPatchBillingAddress**](ShoppingAddressApi.md#ShoppingAddressPatchBillingAddress) | **Patch** /api/v1/shopping/{cartKey}/address/billing | Patches a billing address record on the shopping cart
[**ShoppingAddressPatchMailingAddress**](ShoppingAddressApi.md#ShoppingAddressPatchMailingAddress) | **Patch** /api/v1/shopping/{cartKey}/address/mailing | Patches a mailing address record to the shopping cart
[**ShoppingAddressPatchShippingAddress**](ShoppingAddressApi.md#ShoppingAddressPatchShippingAddress) | **Patch** /api/v1/shopping/{cartKey}/address/shipping | Patches a Shipping address record on the shopping cart
[**ShoppingAddressRemoveBillingAddress**](ShoppingAddressApi.md#ShoppingAddressRemoveBillingAddress) | **Delete** /api/v1/shopping/{cartKey}/address/billing | Removes a billing address record on the shopping cart
[**ShoppingAddressRemoveMailingAddress**](ShoppingAddressApi.md#ShoppingAddressRemoveMailingAddress) | **Delete** /api/v1/shopping/{cartKey}/address/mailing | Removes a mailing address record on the shopping cart
[**ShoppingAddressRemoveShippingAddress**](ShoppingAddressApi.md#ShoppingAddressRemoveShippingAddress) | **Delete** /api/v1/shopping/{cartKey}/address/shipping | Removes a shipping address record on the shopping cart
[**ShoppingAddressUpdateBillingAddress**](ShoppingAddressApi.md#ShoppingAddressUpdateBillingAddress) | **Put** /api/v1/shopping/{cartKey}/address/billing | Updates a billing address record on the shopping cart
[**ShoppingAddressUpdateMailingAddress**](ShoppingAddressApi.md#ShoppingAddressUpdateMailingAddress) | **Put** /api/v1/shopping/{cartKey}/address/mailing | Updates a mailing address record on the shopping cart
[**ShoppingAddressUpdateShippingAddress**](ShoppingAddressApi.md#ShoppingAddressUpdateShippingAddress) | **Put** /api/v1/shopping/{cartKey}/address/shipping | Updates a Shipping address record on the shopping cart


# **ShoppingAddressAddBillingAddress**
> CartAddressViewModel ShoppingAddressAddBillingAddress(ctx, cartKey, model)
Adds a billing address record to the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
  **model** | [**CartAddressFormModel**](CartAddressFormModel.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressAddMailingAddress**
> CartAddressViewModel ShoppingAddressAddMailingAddress(ctx, cartKey, model)
Adds a mailing address record to the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
  **model** | [**CartAddressFormModel**](CartAddressFormModel.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressAddShippingAddress**
> CartAddressViewModel ShoppingAddressAddShippingAddress(ctx, cartKey, model)
Adds a shipping address record to the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
  **model** | [**CartAddressFormModel**](CartAddressFormModel.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressPatchBillingAddress**
> CartAddressViewModel ShoppingAddressPatchBillingAddress(ctx, cartKey, model)
Patches a billing address record on the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
  **model** | [**CartAddressViewModel**](CartAddressViewModel.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressPatchMailingAddress**
> CartAddressViewModel ShoppingAddressPatchMailingAddress(ctx, cartKey, model)
Patches a mailing address record to the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
  **model** | [**CartAddressViewModel**](CartAddressViewModel.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressPatchShippingAddress**
> CartAddressViewModel ShoppingAddressPatchShippingAddress(ctx, cartKey, model)
Patches a Shipping address record on the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
  **model** | [**CartAddressViewModel**](CartAddressViewModel.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressRemoveBillingAddress**
> CartAddressViewModel ShoppingAddressRemoveBillingAddress(ctx, cartKey)
Removes a billing address record on the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressRemoveMailingAddress**
> CartAddressViewModel ShoppingAddressRemoveMailingAddress(ctx, cartKey)
Removes a mailing address record on the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressRemoveShippingAddress**
> CartAddressViewModel ShoppingAddressRemoveShippingAddress(ctx, cartKey)
Removes a shipping address record on the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressUpdateBillingAddress**
> CartAddressViewModel ShoppingAddressUpdateBillingAddress(ctx, cartKey, model)
Updates a billing address record on the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
  **model** | [**CartAddressFormModel**](CartAddressFormModel.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressUpdateMailingAddress**
> CartAddressViewModel ShoppingAddressUpdateMailingAddress(ctx, cartKey, model)
Updates a mailing address record on the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
  **model** | [**CartAddressFormModel**](CartAddressFormModel.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingAddressUpdateShippingAddress**
> CartAddressViewModel ShoppingAddressUpdateShippingAddress(ctx, cartKey, model)
Updates a Shipping address record on the shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
  **model** | [**CartAddressFormModel**](CartAddressFormModel.md)|  | 

### Return type

[**CartAddressViewModel**](CartAddressViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# \ProductApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductGetProductConfiguration**](ProductApi.md#ProductGetProductConfiguration) | **Post** /api/v1/product/configuration | Use for products that have &#39;flavors&#39; to an individual item
[**ProductGetProductDetail**](ProductApi.md#ProductGetProductDetail) | **Post** /api/v1/product/detail | Get the detail information for products
[**ProductGetProductRegulatoryDetails**](ProductApi.md#ProductGetProductRegulatoryDetails) | **Get** /api/v1/product/regulatory-details | Returns product details for Regulatory
[**ProductGetRafikiProduct**](ProductApi.md#ProductGetRafikiProduct) | **Get** /api/v1/product/rafiki | Get the details for the dontion product called Rafiki
[**ProductGetRetailProducts**](ProductApi.md#ProductGetRetailProducts) | **Post** /api/v1/product/retail | Gets all retail products metadata for a given siteurl, culturename, and country.
[**ProductGetSamples**](ProductApi.md#ProductGetSamples) | **Get** /api/v1/product/samples/{siteUrl} | 
[**ProductGetWholeSaleProducts**](ProductApi.md#ProductGetWholeSaleProducts) | **Post** /api/v1/product/wholesale | Gets all wholesale products metadata for a given siteurl, culturename, and country.
[**ProductSubmitNvShadeRequest**](ProductApi.md#ProductSubmitNvShadeRequest) | **Post** /api/v1/product/nv/shade-request | Submits a shade form request via email.


# **ProductGetProductConfiguration**
> interface{} ProductGetProductConfiguration(ctx, model)
Use for products that have 'flavors' to an individual item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**ProductConfigurationFormModel**](ProductConfigurationFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetProductDetail**
> interface{} ProductGetProductDetail(ctx, model)
Get the detail information for products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**ProductDetailFormModel**](ProductDetailFormModel.md)| Multiple arrays that are used to look up products | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetProductRegulatoryDetails**
> ProductRegulatoryDetailsViewModel ProductGetProductRegulatoryDetails(ctx, optional)
Returns product details for Regulatory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetProductRegulatoryDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetProductRegulatoryDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sku** | **optional.String**|  | 
 **culture** | **optional.String**|  | 

### Return type

[**ProductRegulatoryDetailsViewModel**](ProductRegulatoryDetailsViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetRafikiProduct**
> interface{} ProductGetRafikiProduct(ctx, )
Get the details for the dontion product called Rafiki

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

# **ProductGetRetailProducts**
> interface{} ProductGetRetailProducts(ctx, model)
Gets all retail products metadata for a given siteurl, culturename, and country.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**ProductFormModel**](ProductFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetSamples**
> interface{} ProductGetSamples(ctx, siteUrl)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetWholeSaleProducts**
> interface{} ProductGetWholeSaleProducts(ctx, model)
Gets all wholesale products metadata for a given siteurl, culturename, and country.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**ProductFormModel**](ProductFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductSubmitNvShadeRequest**
> interface{} ProductSubmitNvShadeRequest(ctx, model)
Submits a shade form request via email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**ShadeRequestModel**](ShadeRequestModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


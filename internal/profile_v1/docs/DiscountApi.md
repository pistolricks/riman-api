# \DiscountApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscountDiscountExtend**](DiscountApi.md#DiscountDiscountExtend) | **Put** /api/v1/discount | extend discount expiration
[**DiscountDiscountGenerate**](DiscountApi.md#DiscountDiscountGenerate) | **Post** /api/v1/discount | generate discounts


# **DiscountDiscountExtend**
> interface{} DiscountDiscountExtend(ctx, discount, optional)
extend discount expiration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **discount** | [**JeunesseProfileCoreDiscountModelsDiscountModel**](JeunesseProfileCoreDiscountModelsDiscountModel.md)|  | 
 **optional** | ***DiscountApiDiscountDiscountExtendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiscountApiDiscountDiscountExtendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscountDiscountGenerate**
> interface{} DiscountDiscountGenerate(ctx, discount, optional)
generate discounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **discount** | [**JeunesseProfileCoreDiscountModelsDiscountModel**](JeunesseProfileCoreDiscountModelsDiscountModel.md)|  | 
 **optional** | ***DiscountApiDiscountDiscountGenerateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiscountApiDiscountDiscountGenerateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


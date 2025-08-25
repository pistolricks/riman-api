# \ShippingApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShippingGet**](ShippingApi.md#ShippingGet) | **Get** /api/v2/shipping/shipping/options | Returns a list of shipping options available.


# **ShippingGet**
> interface{} ShippingGet(ctx, optional)
Returns a list of shipping options available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShippingApiShippingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShippingApiShippingGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartKey** | [**optional.Interface of string**](.md)|  | 
 **address1** | **optional.String**|  | 
 **address2** | **optional.String**|  | 
 **address3** | **optional.String**|  | 
 **country** | **optional.String**|  | 
 **state** | **optional.String**|  | 
 **zip** | **optional.String**|  | 
 **city** | **optional.String**|  | 
 **mainPk** | **optional.Int32**|  | 
 **isPOBox** | **optional.Bool**|  | 
 **mainOrderPK** | **optional.Int32**|  | 
 **isSDProfile** | **optional.Bool**|  | 
 **isAdmin** | **optional.Bool**|  | 
 **isEvent** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


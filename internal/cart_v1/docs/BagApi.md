# \BagApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BagCreate**](BagApi.md#BagCreate) | **Post** /api/v1/bags | Create a new bag
[**BagDelete**](BagApi.md#BagDelete) | **Delete** /api/v1/bags/{bagId} | Delete a bag by its id
[**BagGetAll**](BagApi.md#BagGetAll) | **Get** /api/v1/bags | Get all bags based on the query model
[**BagGetById**](BagApi.md#BagGetById) | **Get** /api/v1/bags/{bagId} | Get a bag by its id


# **BagCreate**
> interface{} BagCreate(ctx, bagFormModel)
Create a new bag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bagFormModel** | [**BagFormModel**](BagFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BagDelete**
> interface{} BagDelete(ctx, bagId)
Delete a bag by its id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bagId** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BagGetAll**
> interface{} BagGetAll(ctx, optional)
Get all bags based on the query model

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BagApiBagGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BagApiBagGetAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bagQueryModelBagPK** | **optional.Int32**|  | 
 **bagQueryModelCustomerPK** | **optional.Int32**|  | 
 **bagQueryModelCartType** | **optional.String**|  | 
 **bagQueryModelCountryCode** | **optional.String**|  | 
 **bagQueryModelCulture** | **optional.String**|  | 
 **bagQueryModelSiteUrl** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BagGetById**
> interface{} BagGetById(ctx, bagId)
Get a bag by its id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bagId** | **int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


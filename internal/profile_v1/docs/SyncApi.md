# \SyncApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncGetDataToSync**](SyncApi.md#SyncGetDataToSync) | **Get** /api/v1/sync/{entityName} | 
[**SyncGetPushOrders**](SyncApi.md#SyncGetPushOrders) | **Get** /api/v1/sync/getPushOrders | 
[**SyncGetQueryStatusOrders**](SyncApi.md#SyncGetQueryStatusOrders) | **Get** /api/v1/sync/getQueryStatusOrders | 
[**SyncProcessPushResults**](SyncApi.md#SyncProcessPushResults) | **Post** /api/v1/sync/processPushResults/{shippingOrdersFK} | 
[**SyncProcessQueryStatusResults**](SyncApi.md#SyncProcessQueryStatusResults) | **Post** /api/v1/sync/processQueryStatusResults/{queryStatusID} | 
[**SyncSetSyncStatus**](SyncApi.md#SyncSetSyncStatus) | **Post** /api/v1/sync/set-status | 
[**SyncSync**](SyncApi.md#SyncSync) | **Post** /api/v1/sync | 


# **SyncGetDataToSync**
> string SyncGetDataToSync(ctx, entityName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityName** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncGetPushOrders**
> JeunesseProfileCoreSynchronizationModelsPushOrderData SyncGetPushOrders(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JeunesseProfileCoreSynchronizationModelsPushOrderData**](Jeunesse.Profile.Core.Synchronization.Models.PushOrderData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncGetQueryStatusOrders**
> JeunesseProfileCoreSynchronizationModelsQueryOrderStatusData SyncGetQueryStatusOrders(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JeunesseProfileCoreSynchronizationModelsQueryOrderStatusData**](Jeunesse.Profile.Core.Synchronization.Models.QueryOrderStatusData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncProcessPushResults**
> bool SyncProcessPushResults(ctx, shippingOrdersFK, result)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shippingOrdersFK** | **int32**|  | 
  **result** | [**JeunesseProfileCoreSynchronizationModelsPushOrdersResult**](JeunesseProfileCoreSynchronizationModelsPushOrdersResult.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncProcessQueryStatusResults**
> bool SyncProcessQueryStatusResults(ctx, queryStatusID, result)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryStatusID** | **int32**|  | 
  **result** | [**JeunesseProfileCoreSynchronizationModelsQueryOrderResultData**](JeunesseProfileCoreSynchronizationModelsQueryOrderResultData.md)|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncSetSyncStatus**
> string SyncSetSyncStatus(ctx, model)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreSynchronizationModelsSyncDataFormModel**](JeunesseProfileCoreSynchronizationModelsSyncDataFormModel.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncSync**
> string SyncSync(ctx, model)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreSynchronizationModelsSyncDataFormModel**](JeunesseProfileCoreSynchronizationModelsSyncDataFormModel.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


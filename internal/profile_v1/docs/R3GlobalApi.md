# \R3GlobalApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**R3GlobalGetUser2x2Matrix**](R3GlobalApi.md#R3GlobalGetUser2x2Matrix) | **Get** /api/v1/r3global/matrix | Get the 2x2 Matrix data for the authenticated user.
[**R3GlobalReset2x2MatrixTimer**](R3GlobalApi.md#R3GlobalReset2x2MatrixTimer) | **Put** /api/v1/r3global/matrix/{mainpk}/reset | Get the 2x2 Matrix data for the authenticated user.


# **R3GlobalGetUser2x2Matrix**
> []JeunesseProfileCoreR3GlobalModelsMatrixNodeModel R3GlobalGetUser2x2Matrix(ctx, optional)
Get the 2x2 Matrix data for the authenticated user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***R3GlobalApiR3GlobalGetUser2x2MatrixOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a R3GlobalApiR3GlobalGetUser2x2MatrixOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **requestChildLegs** | **optional.Int32**|  | 
 **requestOutputDepth** | **optional.Int32**|  | 
 **requestShouldOverrideFullLegsCheck** | **optional.Bool**|  | 

### Return type

[**[]JeunesseProfileCoreR3GlobalModelsMatrixNodeModel**](Jeunesse.Profile.Core.R3Global.Models.MatrixNodeModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **R3GlobalReset2x2MatrixTimer**
> interface{} R3GlobalReset2x2MatrixTimer(ctx, mainpk, optional)
Get the 2x2 Matrix data for the authenticated user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainpk** | **int32**| The mainPK of the user whose timer is to be reset | 
 **optional** | ***R3GlobalApiR3GlobalReset2x2MatrixTimerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a R3GlobalApiR3GlobalReset2x2MatrixTimerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


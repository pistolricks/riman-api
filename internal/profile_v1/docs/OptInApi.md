# \OptInApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptInGetOptInSettings**](OptInApi.md#OptInGetOptInSettings) | **Get** /api/v1/optin/{email}/{culture} | Returns a list of opt in subscription items for an email.
[**OptInUnsubscribe**](OptInApi.md#OptInUnsubscribe) | **Delete** /api/v1/optin/{email} | Removes an email from all subscriptions
[**OptInUpdateOptInSettings**](OptInApi.md#OptInUpdateOptInSettings) | **Put** /api/v1/optin/{email} | Updates an email&#39;s opt in subscription settings.


# **OptInGetOptInSettings**
> []JeunesseProfileCoreOptInModelsOptInItemViewModel OptInGetOptInSettings(ctx, email, culture, optional)
Returns a list of opt in subscription items for an email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**|  | 
  **culture** | **string**|  | 
 **optional** | ***OptInApiOptInGetOptInSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptInApiOptInGetOptInSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreOptInModelsOptInItemViewModel**](Jeunesse.Profile.Core.OptIn.Models.OptInItemViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptInUnsubscribe**
> interface{} OptInUnsubscribe(ctx, email, optional)
Removes an email from all subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**|  | 
 **optional** | ***OptInApiOptInUnsubscribeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptInApiOptInUnsubscribeOpts struct

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

# **OptInUpdateOptInSettings**
> []JeunesseProfileCoreOptInModelsOptInItemViewModel OptInUpdateOptInSettings(ctx, email, model, optional)
Updates an email's opt in subscription settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**|  | 
  **model** | [**JeunesseProfileCoreOptInModelsOptInFormModel**](JeunesseProfileCoreOptInModelsOptInFormModel.md)|  | 
 **optional** | ***OptInApiOptInUpdateOptInSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptInApiOptInUpdateOptInSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreOptInModelsOptInItemViewModel**](Jeunesse.Profile.Core.OptIn.Models.OptInItemViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


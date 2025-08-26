# \MainExtensionApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MainExtensionGetItalyCardData**](MainExtensionApi.md#MainExtensionGetItalyCardData) | **Get** /api/v1/main-extension/it-card | Get card data for Italy user


# **MainExtensionGetItalyCardData**
> JeunesseProfileCoreMainExtensionModelsMainItExtensionModel MainExtensionGetItalyCardData(ctx, optional)
Get card data for Italy user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainExtensionApiMainExtensionGetItalyCardDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainExtensionApiMainExtensionGetItalyCardDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreMainExtensionModelsMainItExtensionModel**](Jeunesse.Profile.Core.MainExtension.Models.MainITExtensionModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


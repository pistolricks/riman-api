# \ServerReportsApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServerReportsGetPdf**](ServerReportsApi.md#ServerReportsGetPdf) | **Get** /api/v1/server-reports/report-pdf/{reportId} | Test endpoint for getting generated pdf from reportServer


# **ServerReportsGetPdf**
> interface{} ServerReportsGetPdf(ctx, reportId, optional)
Test endpoint for getting generated pdf from reportServer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int32**|  | 
 **optional** | ***ServerReportsApiServerReportsGetPdfOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerReportsApiServerReportsGetPdfOpts struct

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


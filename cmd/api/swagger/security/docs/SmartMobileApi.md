# \SmartMobileApi

All URIs are relative to *https://security-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartMobileLogin**](SmartMobileApi.md#SmartMobileLogin) | **Post** /api/v2/smartmobile/authenticate | Login for SmartMobile
[**SmartMobileValidate**](SmartMobileApi.md#SmartMobileValidate) | **Post** /api/v2/smartmobile/authenticate-token | Validates a Joffice token or JWT Token


# **SmartMobileLogin**
> interface{} SmartMobileLogin(ctx, model)
Login for SmartMobile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**SmartMobileLoginRequest**](SmartMobileLoginRequest.md)| Model for request | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartMobileValidate**
> interface{} SmartMobileValidate(ctx, request)
Validates a Joffice token or JWT Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**SmartMobileValidateTokenRequest**](SmartMobileValidateTokenRequest.md)| The request model with either JWTToken set or JOffice Token set | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


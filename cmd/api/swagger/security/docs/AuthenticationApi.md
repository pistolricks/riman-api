# \AuthenticationApi

All URIs are relative to *https://security-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticationAuthenticateByJOfficeToken**](AuthenticationApi.md#AuthenticationAuthenticateByJOfficeToken) | **Get** /api/v2/token/joffice-token/{jofficeToken} | Authenticate using the joffice token.
[**AuthenticationAuthenticateUser**](AuthenticationApi.md#AuthenticationAuthenticateUser) | **Post** /api/v2/token | Used to authenticate a member.
[**AuthenticationAuthenticateUserToWallet**](AuthenticationApi.md#AuthenticationAuthenticateUserToWallet) | **Post** /api/v2/wallet/login | Perform wallet login
[**AuthenticationAuthenticateUserToWallet_0**](AuthenticationApi.md#AuthenticationAuthenticateUserToWallet_0) | **Post** /api/v2/wallet/logout | Perform wallet logout
[**AuthenticationAuthenticateUserV2**](AuthenticationApi.md#AuthenticationAuthenticateUserV2) | **Post** /api/v2/CheckAttemptsAndLogin | Similar  api call to /token but include check for security questions and captcha, have to create new one, cause return model is different from /token. Maybe in a feature   we have to deprecate  /token and use /checkAttemptsAndLogin everywhere
[**AuthenticationCreateReportsLogin**](AuthenticationApi.md#AuthenticationCreateReportsLogin) | **Post** /api/v2/reports/login/create | Used to create a report login account
[**AuthenticationGetTokenByTrackingId**](AuthenticationApi.md#AuthenticationGetTokenByTrackingId) | **Post** /api/v2/token/tracking/{TrackingId} | Returns auth token for use in an abandoned order checkout.
[**AuthenticationLogout**](AuthenticationApi.md#AuthenticationLogout) | **Post** /api/v2/token/logout | Logs a valid member token out and ends their session.
[**AuthenticationReIssue**](AuthenticationApi.md#AuthenticationReIssue) | **Post** /api/v2/token/reissue | Used to reissue an existing token that is not yet expired.
[**AuthenticationReportsLogin**](AuthenticationApi.md#AuthenticationReportsLogin) | **Post** /api/v2/reports/login | Used to authenticate users for report access
[**AuthenticationValidate**](AuthenticationApi.md#AuthenticationValidate) | **Post** /api/v2/token/validate | Used to verify an existing token that is not yet expired. Does not reissue new token
[**AuthenticationValidatePassword**](AuthenticationApi.md#AuthenticationValidatePassword) | **Post** /api/v2/password/validate/{password} | Validates if password follows the business rules for a password. Case, Digit, Special charactesr etc.


# **AuthenticationAuthenticateByJOfficeToken**
> string AuthenticationAuthenticateByJOfficeToken(ctx, jofficeToken)
Authenticate using the joffice token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jofficeToken** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationAuthenticateUser**
> string AuthenticationAuthenticateUser(ctx, request)
Used to authenticate a member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**TokenRequestModel**](TokenRequestModel.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationAuthenticateUserToWallet**
> string AuthenticationAuthenticateUserToWallet(ctx, request)
Perform wallet login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**WalletTokenRequestModel**](WalletTokenRequestModel.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationAuthenticateUserToWallet_0**
> bool AuthenticationAuthenticateUserToWallet_0(ctx, )
Perform wallet logout

### Required Parameters
This endpoint does not need any parameter.

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationAuthenticateUserV2**
> string AuthenticationAuthenticateUserV2(ctx, request)
Similar  api call to /token but include check for security questions and captcha, have to create new one, cause return model is different from /token. Maybe in a feature   we have to deprecate  /token and use /checkAttemptsAndLogin everywhere

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**TokenRequestModel**](TokenRequestModel.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationCreateReportsLogin**
> string AuthenticationCreateReportsLogin(ctx, request)
Used to create a report login account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**ReportLoginFormModel**](ReportLoginFormModel.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationGetTokenByTrackingId**
> string AuthenticationGetTokenByTrackingId(ctx, trackingId)
Returns auth token for use in an abandoned order checkout.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **trackingId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationLogout**
> interface{} AuthenticationLogout(ctx, )
Logs a valid member token out and ends their session.

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

# **AuthenticationReIssue**
> string AuthenticationReIssue(ctx, )
Used to reissue an existing token that is not yet expired.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationReportsLogin**
> string AuthenticationReportsLogin(ctx, request)
Used to authenticate users for report access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**TokenRequestModel**](TokenRequestModel.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationValidate**
> bool AuthenticationValidate(ctx, )
Used to verify an existing token that is not yet expired. Does not reissue new token

### Required Parameters
This endpoint does not need any parameter.

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthenticationValidatePassword**
> PasswordRulesModel AuthenticationValidatePassword(ctx, password)
Validates if password follows the business rules for a password. Case, Digit, Special charactesr etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **password** | **string**|  | 

### Return type

[**PasswordRulesModel**](PasswordRulesModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# \WalletMemberApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletMemberCheckAttempts**](WalletMemberApi.md#WalletMemberCheckAttempts) | **Get** /api/v1/wallet-member/check-attempts | Check actual state of wallet login process
[**WalletMemberCreateWalletAccount**](WalletMemberApi.md#WalletMemberCreateWalletAccount) | **Post** /api/v1/wallet-member/create | Create walet account
[**WalletMemberGetQuestionsList**](WalletMemberApi.md#WalletMemberGetQuestionsList) | **Get** /api/v1/wallet-member/questions | Get questions list
[**WalletMemberGetResetCode**](WalletMemberApi.md#WalletMemberGetResetCode) | **Post** /api/v1/wallet-member/reset-code | Generate reset code and send to email\\sms
[**WalletMemberGetSecurityQuestions**](WalletMemberApi.md#WalletMemberGetSecurityQuestions) | **Get** /api/v1/wallet-member/question | Get question for reset password
[**WalletMemberResetPassword**](WalletMemberApi.md#WalletMemberResetPassword) | **Post** /api/v1/wallet-member/reset-password | Reset password by resetcode
[**WalletMemberSetQuestions**](WalletMemberApi.md#WalletMemberSetQuestions) | **Post** /api/v1/wallet-member/questions | Set security questions and answers
[**WalletMemberUpdatePassword**](WalletMemberApi.md#WalletMemberUpdatePassword) | **Post** /api/v1/wallet-member/update-password | Reset password by resetcode


# **WalletMemberCheckAttempts**
> JeunesseProfileCoreWalletModelsCheckAttemptsResponse WalletMemberCheckAttempts(ctx, optional)
Check actual state of wallet login process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletMemberApiWalletMemberCheckAttemptsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletMemberApiWalletMemberCheckAttemptsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsCheckAttemptsResponse**](Jeunesse.Profile.Core.Wallet.Models.CheckAttemptsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletMemberCreateWalletAccount**
> bool WalletMemberCreateWalletAccount(ctx, model, optional)
Create walet account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsCreateAccountModel**](JeunesseProfileCoreWalletModelsCreateAccountModel.md)|  | 
 **optional** | ***WalletMemberApiWalletMemberCreateWalletAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletMemberApiWalletMemberCreateWalletAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletMemberGetQuestionsList**
> []JeunesseProfileCoreWalletModelsQuestionAnswer WalletMemberGetQuestionsList(ctx, optional)
Get questions list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletMemberApiWalletMemberGetQuestionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletMemberApiWalletMemberGetQuestionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreWalletModelsQuestionAnswer**](Jeunesse.Profile.Core.Wallet.Models.QuestionAnswer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletMemberGetResetCode**
> bool WalletMemberGetResetCode(ctx, model, optional)
Generate reset code and send to email\\sms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsGenerateResetCodeFormModel**](JeunesseProfileCoreWalletModelsGenerateResetCodeFormModel.md)|  | 
 **optional** | ***WalletMemberApiWalletMemberGetResetCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletMemberApiWalletMemberGetResetCodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletMemberGetSecurityQuestions**
> JeunesseProfileCoreWalletModelsCheckAttemptsResponse WalletMemberGetSecurityQuestions(ctx, optional)
Get question for reset password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletMemberApiWalletMemberGetSecurityQuestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletMemberApiWalletMemberGetSecurityQuestionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsCheckAttemptsResponse**](Jeunesse.Profile.Core.Wallet.Models.CheckAttemptsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletMemberResetPassword**
> string WalletMemberResetPassword(ctx, model, optional)
Reset password by resetcode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsResetPasswordFormModel**](JeunesseProfileCoreWalletModelsResetPasswordFormModel.md)|  | 
 **optional** | ***WalletMemberApiWalletMemberResetPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletMemberApiWalletMemberResetPasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletMemberSetQuestions**
> bool WalletMemberSetQuestions(ctx, model, optional)
Set security questions and answers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsQuestionsFormModel**](JeunesseProfileCoreWalletModelsQuestionsFormModel.md)|  | 
 **optional** | ***WalletMemberApiWalletMemberSetQuestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletMemberApiWalletMemberSetQuestionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletMemberUpdatePassword**
> bool WalletMemberUpdatePassword(ctx, model, optional)
Reset password by resetcode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsUpdatePasswordFormModel**](JeunesseProfileCoreWalletModelsUpdatePasswordFormModel.md)|  | 
 **optional** | ***WalletMemberApiWalletMemberUpdatePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletMemberApiWalletMemberUpdatePasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


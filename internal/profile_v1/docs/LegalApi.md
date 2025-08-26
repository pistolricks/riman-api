# \LegalApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LegalCheckAgreements**](LegalApi.md#LegalCheckAgreements) | **Get** /api/v1/legal/check-agreements | Checks if user needs to agree to current legal documents
[**LegalGetLegalDocument**](LegalApi.md#LegalGetLegalDocument) | **Get** /api/v1/legal/documents | Returns the legal documents required for a specific country and language
[**LegalLogAgreedToDocumentsByGuid**](LegalApi.md#LegalLogAgreedToDocumentsByGuid) | **Post** /api/v1/legal/log-agreements-by-guid | Logs documents that users have agreed to.
[**LegalLogAgreedToDocumentsByMainPk**](LegalApi.md#LegalLogAgreedToDocumentsByMainPk) | **Post** /api/v1/legal/log-agreements-by-mainpk | Logs agreement documents by logged in users mainPk


# **LegalCheckAgreements**
> int32 LegalCheckAgreements(ctx, mainPk, mainTypeFk, country, langId, optional)
Checks if user needs to agree to current legal documents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 
  **mainTypeFk** | **int32**|  | 
  **country** | **string**|  | 
  **langId** | **int32**|  | 
 **optional** | ***LegalApiLegalCheckAgreementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LegalApiLegalCheckAgreementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LegalGetLegalDocument**
> interface{} LegalGetLegalDocument(ctx, countryCode, langID, optional)
Returns the legal documents required for a specific country and language

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **langID** | **int32**|  | 
 **optional** | ***LegalApiLegalGetLegalDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LegalApiLegalGetLegalDocumentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromPage** | **optional.String**|  | [default to ]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LegalLogAgreedToDocumentsByGuid**
> interface{} LegalLogAgreedToDocumentsByGuid(ctx, guid, usage, optional)
Logs documents that users have agreed to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **guid** | **string**|  | 
  **usage** | **string**|  | 
 **optional** | ***LegalApiLegalLogAgreedToDocumentsByGuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LegalApiLegalLogAgreedToDocumentsByGuidOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subUsage** | **optional.String**|  | 
 **mainTypeFk** | **optional.Int32**|  | [default to 1]
 **selectedAll** | **optional.Bool**|  | [default to true]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LegalLogAgreedToDocumentsByMainPk**
> interface{} LegalLogAgreedToDocumentsByMainPk(ctx, from, purpose, optional)
Logs agreement documents by logged in users mainPk

\"from\" is the page where the call is coming from  \"purpose\" is an optional field for why they are agreeing. Currently only \"wcAgree\" does anything different

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **from** | **string**|  | 
  **purpose** | **string**|  | 
 **optional** | ***LegalApiLegalLogAgreedToDocumentsByMainPkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LegalApiLegalLogAgreedToDocumentsByMainPkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mainTypeFk** | **optional.Int32**|  | [default to 0]
 **selectedAll** | **optional.Bool**|  | [default to true]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# \SmartMobileApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartMobileGetMessage**](SmartMobileApi.md#SmartMobileGetMessage) | **Get** /api/v1/smartmobile/mailbox/{folderID}/user-messages | Get mailbox messages for the user.
[**SmartMobileGetReportParameters**](SmartMobileApi.md#SmartMobileGetReportParameters) | **Get** /api/v1/smartmobile/reports/{reportId}/parameters | Get report parameters.
[**SmartMobileGetReports**](SmartMobileApi.md#SmartMobileGetReports) | **Get** /api/v1/smartmobile/reports | Gets user&#39;s reports.
[**SmartMobileMoveMainUserToSmartMobile**](SmartMobileApi.md#SmartMobileMoveMainUserToSmartMobile) | **Get** /api/v1/smartmobile/move-user/{mainPk}/{siteUrl} | Move Main user to SmartMobile.
[**SmartMobileRegisterDeviceToken**](SmartMobileApi.md#SmartMobileRegisterDeviceToken) | **Post** /api/v1/smartmobile/registerDeviceToken | Register Mobile Application that will receive Push Notifications.
[**SmartMobileUpdateAllMessages**](SmartMobileApi.md#SmartMobileUpdateAllMessages) | **Put** /api/v1/smartmobile/mailbox/{folderID}/user-messages | Updates messages to mark as read or deleted.
[**SmartMobileUpdateMessage**](SmartMobileApi.md#SmartMobileUpdateMessage) | **Put** /api/v1/smartmobile/mailbox/{folderID}/user-messages/{messageID} | Update the message to mark as read or deleted.


# **SmartMobileGetMessage**
> interface{} SmartMobileGetMessage(ctx, folderID, optional)
Get mailbox messages for the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**| The ID of the user&#39;s folder. | 
 **optional** | ***SmartMobileApiSmartMobileGetMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartMobileApiSmartMobileGetMessageOpts struct

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

# **SmartMobileGetReportParameters**
> interface{} SmartMobileGetReportParameters(ctx, reportId, optional)
Get report parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int32**|  | 
 **optional** | ***SmartMobileApiSmartMobileGetReportParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartMobileApiSmartMobileGetReportParametersOpts struct

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

# **SmartMobileGetReports**
> interface{} SmartMobileGetReports(ctx, optional)
Gets user's reports.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmartMobileApiSmartMobileGetReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartMobileApiSmartMobileGetReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **requestCompanyPK** | **optional.Int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartMobileMoveMainUserToSmartMobile**
> interface{} SmartMobileMoveMainUserToSmartMobile(ctx, mainPk, siteUrl, optional)
Move Main user to SmartMobile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**| Main Users MainPK | 
  **siteUrl** | **string**| Main Users SitreUrl (Optional). | 
 **optional** | ***SmartMobileApiSmartMobileMoveMainUserToSmartMobileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartMobileApiSmartMobileMoveMainUserToSmartMobileOpts struct

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

# **SmartMobileRegisterDeviceToken**
> interface{} SmartMobileRegisterDeviceToken(ctx, request, optional)
Register Mobile Application that will receive Push Notifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**JeunesseProfileCoreSmartMobileRegisterDeviceTokenRequestModel**](JeunesseProfileCoreSmartMobileRegisterDeviceTokenRequestModel.md)|  | 
 **optional** | ***SmartMobileApiSmartMobileRegisterDeviceTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartMobileApiSmartMobileRegisterDeviceTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartMobileUpdateAllMessages**
> interface{} SmartMobileUpdateAllMessages(ctx, folderID, requestModel, optional)
Updates messages to mark as read or deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**| The ID of the user&#39;s folder. | 
  **requestModel** | [**JeunesseProfileCoreSmartMobileUpdateMessageRequest**](JeunesseProfileCoreSmartMobileUpdateMessageRequest.md)| Message Request Model to be updated. [fromBody] | 
 **optional** | ***SmartMobileApiSmartMobileUpdateAllMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartMobileApiSmartMobileUpdateAllMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartMobileUpdateMessage**
> interface{} SmartMobileUpdateMessage(ctx, folderID, messageID, requestModel, optional)
Update the message to mark as read or deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**| The ID of the user&#39;s folder. | 
  **messageID** | **int32**| The ID of the user&#39;s message. | 
  **requestModel** | [**JeunesseProfileCoreSmartMobileUpdateMessageRequest**](JeunesseProfileCoreSmartMobileUpdateMessageRequest.md)| Message Request Model to be updated. [fromBody] | 
 **optional** | ***SmartMobileApiSmartMobileUpdateMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartMobileApiSmartMobileUpdateMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


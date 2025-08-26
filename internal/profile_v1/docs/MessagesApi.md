# \MessagesApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MessagesCreateFolder**](MessagesApi.md#MessagesCreateFolder) | **Post** /api/v1/messages/mailbox/folders | Create new folder  &lt;param name&#x3D;\&quot;query\&quot;&gt;Name and Parent Folder ID.&lt;/param&gt;
[**MessagesDeleteAllMessages**](MessagesApi.md#MessagesDeleteAllMessages) | **Delete** /api/v1/messages/mailbox/{folderID}/delete-all | Update the message to mark as read
[**MessagesDeleteAllUserAlerts**](MessagesApi.md#MessagesDeleteAllUserAlerts) | **Delete** /api/v1/messages/alerts/delete-all | Delete ALL Alerts  ----------------  Wrapper to be consistent with SM API.   There is currently no use case to call it from Message Center UI
[**MessagesDeleteFolder**](MessagesApi.md#MessagesDeleteFolder) | **Delete** /api/v1/messages/mailbox/{folderID} | Delete folder  &lt;param name&#x3D;\&quot;folderID\&quot;&gt;The ID of the user&#39;s folder.&lt;/param&gt;
[**MessagesDeleteUserAlerts**](MessagesApi.md#MessagesDeleteUserAlerts) | **Delete** /api/v1/messages/alerts/delete | Delete selected Alert(s)
[**MessagesGetAlertTypes**](MessagesApi.md#MessagesGetAlertTypes) | **Get** /api/v1/messages/alerts/types | Get unread alerts count
[**MessagesGetFolders**](MessagesApi.md#MessagesGetFolders) | **Get** /api/v1/messages/mailbox/folders | Get list of folders with unread messages count
[**MessagesGetMessage**](MessagesApi.md#MessagesGetMessage) | **Get** /api/v1/messages/mailbox/{folderID}/user-messages/{messageID} | Get mailbox message
[**MessagesGetMessages**](MessagesApi.md#MessagesGetMessages) | **Get** /api/v1/messages/mailbox/{folderID}/user-messages | Get mailbox messages for the user.
[**MessagesGetNews**](MessagesApi.md#MessagesGetNews) | **Get** /api/v1/messages/news | Get smart news feed
[**MessagesGetNewsCategories**](MessagesApi.md#MessagesGetNewsCategories) | **Get** /api/v1/messages/news/categories | Get news categories
[**MessagesGetNewsDetails**](MessagesApi.md#MessagesGetNewsDetails) | **Get** /api/v1/messages/news/{articlePK} | Get news article
[**MessagesGetUnreadAlertsCount**](MessagesApi.md#MessagesGetUnreadAlertsCount) | **Get** /api/v1/messages/alerts/count | Get unread alerts count
[**MessagesGetUnreadMessagesCount**](MessagesApi.md#MessagesGetUnreadMessagesCount) | **Get** /api/v1/messages/mailbox/count | Get unread messages count
[**MessagesGetUserAlerts**](MessagesApi.md#MessagesGetUserAlerts) | **Get** /api/v1/messages/alerts | Get user Alerts
[**MessagesMarkAllUserAlertsAsRead**](MessagesApi.md#MessagesMarkAllUserAlertsAsRead) | **Post** /api/v1/messages/alerts/mark-all-read | Mark ALL Alerts as read
[**MessagesMarkUserAlertsAsRead**](MessagesApi.md#MessagesMarkUserAlertsAsRead) | **Post** /api/v1/messages/alerts/mark-as-read | Mark selected Alert(s) as read
[**MessagesMoveMessages**](MessagesApi.md#MessagesMoveMessages) | **Put** /api/v1/messages/mailbox/{folderID}/move-messages | Move messages to other folder.
[**MessagesRenameeFolder**](MessagesApi.md#MessagesRenameeFolder) | **Put** /api/v1/messages/mailbox/{folderID} | Rename folder  &lt;param name&#x3D;\&quot;folderID\&quot;&gt;The ID of the user&#39;s folder.&lt;/param&gt;&lt;param name&#x3D;\&quot;query\&quot;&gt;New folder name&lt;/param&gt;
[**MessagesUpdateAllMessages**](MessagesApi.md#MessagesUpdateAllMessages) | **Put** /api/v1/messages/mailbox/{folderID}/mark-all-read | Update the message to mark as read
[**MessagesUpdateMessages**](MessagesApi.md#MessagesUpdateMessages) | **Put** /api/v1/messages/mailbox/{folderID}/user-messages | Updates messages to mark as read or deleted.


# **MessagesCreateFolder**
> interface{} MessagesCreateFolder(ctx, query, optional)
Create new folder  <param name=\"query\">Name and Parent Folder ID.</param>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | [**JeunesseProfileCoreMessagesModelsFolderQueryModel**](JeunesseProfileCoreMessagesModelsFolderQueryModel.md)|  | 
 **optional** | ***MessagesApiMessagesCreateFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesCreateFolderOpts struct

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

# **MessagesDeleteAllMessages**
> interface{} MessagesDeleteAllMessages(ctx, folderID, optional)
Update the message to mark as read

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**| The ID of the user&#39;s folder. | 
 **optional** | ***MessagesApiMessagesDeleteAllMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesDeleteAllMessagesOpts struct

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

# **MessagesDeleteAllUserAlerts**
> interface{} MessagesDeleteAllUserAlerts(ctx, optional)
Delete ALL Alerts  ----------------  Wrapper to be consistent with SM API.   There is currently no use case to call it from Message Center UI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesDeleteAllUserAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesDeleteAllUserAlertsOpts struct

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

# **MessagesDeleteFolder**
> interface{} MessagesDeleteFolder(ctx, folderID, optional)
Delete folder  <param name=\"folderID\">The ID of the user's folder.</param>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**|  | 
 **optional** | ***MessagesApiMessagesDeleteFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesDeleteFolderOpts struct

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

# **MessagesDeleteUserAlerts**
> interface{} MessagesDeleteUserAlerts(ctx, optional)
Delete selected Alert(s)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesDeleteUserAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesDeleteUserAlertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **queryMainFK** | **optional.Int32**|  | 
 **queryDeviceCultureName** | **optional.String**|  | 
 **queryAlertIDs** | **optional.String**|  | 
 **queryAlertPKs** | [**optional.Interface of []int32**](int32.md)|  | 
 **queryIsRead** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MessagesGetAlertTypes**
> interface{} MessagesGetAlertTypes(ctx, optional)
Get unread alerts count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesGetAlertTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetAlertTypesOpts struct

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

# **MessagesGetFolders**
> interface{} MessagesGetFolders(ctx, optional)
Get list of folders with unread messages count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesGetFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetFoldersOpts struct

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

# **MessagesGetMessage**
> interface{} MessagesGetMessage(ctx, folderID, messageID, optional)
Get mailbox message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**| The ID of the user&#39;s folder. | 
  **messageID** | **int32**| The ID of the message to get. | 
 **optional** | ***MessagesApiMessagesGetMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetMessageOpts struct

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

# **MessagesGetMessages**
> interface{} MessagesGetMessages(ctx, folderID, optional)
Get mailbox messages for the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**| The ID of the user&#39;s folder. | 
 **optional** | ***MessagesApiMessagesGetMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **queryUsername** | **optional.String**|  | 
 **queryIsFavorite** | **optional.Bool**|  | 
 **queryPageNumber** | **optional.Int32**|  | 
 **queryPageSize** | **optional.Int32**|  | 
 **querySearchTerm** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MessagesGetNews**
> interface{} MessagesGetNews(ctx, optional)
Get smart news feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesGetNewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetNewsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **username** | **optional.String**|  | 
 **lastArticleFK** | **optional.Int32**|  | 
 **length** | **optional.Int32**|  | 
 **categories** | **optional.String**|  | 
 **categoryKeys** | [**optional.Interface of []string**](string.md)|  | 
 **mainFK** | **optional.Int32**|  | 
 **deviceCultureName** | **optional.String**|  | 
 **title** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MessagesGetNewsCategories**
> interface{} MessagesGetNewsCategories(ctx, optional)
Get news categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesGetNewsCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetNewsCategoriesOpts struct

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

# **MessagesGetNewsDetails**
> interface{} MessagesGetNewsDetails(ctx, articlePK, optional)
Get news article

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **articlePK** | **int32**|  | 
 **optional** | ***MessagesApiMessagesGetNewsDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetNewsDetailsOpts struct

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

# **MessagesGetUnreadAlertsCount**
> interface{} MessagesGetUnreadAlertsCount(ctx, optional)
Get unread alerts count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesGetUnreadAlertsCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetUnreadAlertsCountOpts struct

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

# **MessagesGetUnreadMessagesCount**
> interface{} MessagesGetUnreadMessagesCount(ctx, optional)
Get unread messages count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesGetUnreadMessagesCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetUnreadMessagesCountOpts struct

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

# **MessagesGetUserAlerts**
> interface{} MessagesGetUserAlerts(ctx, optional)
Get user Alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesGetUserAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesGetUserAlertsOpts struct

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

# **MessagesMarkAllUserAlertsAsRead**
> interface{} MessagesMarkAllUserAlertsAsRead(ctx, optional)
Mark ALL Alerts as read

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiMessagesMarkAllUserAlertsAsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesMarkAllUserAlertsAsReadOpts struct

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

# **MessagesMarkUserAlertsAsRead**
> interface{} MessagesMarkUserAlertsAsRead(ctx, query, optional)
Mark selected Alert(s) as read

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | [**JeunesseProfileCoreMessagesModelsAlertsQueryModel**](JeunesseProfileCoreMessagesModelsAlertsQueryModel.md)|  | 
 **optional** | ***MessagesApiMessagesMarkUserAlertsAsReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesMarkUserAlertsAsReadOpts struct

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

# **MessagesMoveMessages**
> interface{} MessagesMoveMessages(ctx, folderID, query, optional)
Move messages to other folder.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**| The ID of the target user&#39;s folder. | 
  **query** | [**JeunesseProfileCoreMessagesModelsMailUpdateQueryModel**](JeunesseProfileCoreMessagesModelsMailUpdateQueryModel.md)| Message Request Model to be updated. [fromBody] | 
 **optional** | ***MessagesApiMessagesMoveMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesMoveMessagesOpts struct

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

# **MessagesRenameeFolder**
> interface{} MessagesRenameeFolder(ctx, folderID, query, optional)
Rename folder  <param name=\"folderID\">The ID of the user's folder.</param><param name=\"query\">New folder name</param>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**|  | 
  **query** | [**JeunesseProfileCoreMessagesModelsFolderQueryModel**](JeunesseProfileCoreMessagesModelsFolderQueryModel.md)|  | 
 **optional** | ***MessagesApiMessagesRenameeFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesRenameeFolderOpts struct

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

# **MessagesUpdateAllMessages**
> interface{} MessagesUpdateAllMessages(ctx, folderID, optional)
Update the message to mark as read

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**| The ID of the user&#39;s folder. | 
 **optional** | ***MessagesApiMessagesUpdateAllMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesUpdateAllMessagesOpts struct

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

# **MessagesUpdateMessages**
> interface{} MessagesUpdateMessages(ctx, folderID, query, optional)
Updates messages to mark as read or deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderID** | **int32**| The ID of the user&#39;s folder. | 
  **query** | [**JeunesseProfileCoreMessagesModelsMailUpdateQueryModel**](JeunesseProfileCoreMessagesModelsMailUpdateQueryModel.md)| Message Request Model to be updated. [fromBody] | 
 **optional** | ***MessagesApiMessagesUpdateMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiMessagesUpdateMessagesOpts struct

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


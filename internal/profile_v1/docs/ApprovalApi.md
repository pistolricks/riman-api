# \ApprovalApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApprovalEditTicket**](ApprovalApi.md#ApprovalEditTicket) | **Put** /api/v1/approvals | Edit a Approval Task.
[**ApprovalGetApprovalTasks**](ApprovalApi.md#ApprovalGetApprovalTasks) | **Get** /api/v1/approvals | Gets a list of Approval Tasks.


# **ApprovalEditTicket**
> bool ApprovalEditTicket(ctx, query, optional)
Edit a Approval Task.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | [**JeunesseProfileCoreWalletModelsApprovalsApprovalTaskUpdateQueryModel**](JeunesseProfileCoreWalletModelsApprovalsApprovalTaskUpdateQueryModel.md)|  | 
 **optional** | ***ApprovalApiApprovalEditTicketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApprovalApiApprovalEditTicketOpts struct

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

# **ApprovalGetApprovalTasks**
> JeunesseProfileCoreWalletModelsApprovalsApprovalTasksViewModel ApprovalGetApprovalTasks(ctx, optional)
Gets a list of Approval Tasks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApprovalApiApprovalGetApprovalTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApprovalApiApprovalGetApprovalTasksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **distributorSiteUrl** | **optional.String**|  | 
 **reqSiteUrl** | **optional.String**|  | 
 **typeFK** | **optional.Int32**|  | 
 **statusFKs** | **optional.String**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageNumber** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 

### Return type

[**JeunesseProfileCoreWalletModelsApprovalsApprovalTasksViewModel**](Jeunesse.Profile.Core.Wallet.Models.Approvals.ApprovalTasksViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


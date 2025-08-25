# \LoyaltyPointsApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoyaltyPointsGetRanks**](LoyaltyPointsApi.md#LoyaltyPointsGetRanks) | **Get** /api/v1/loyalty-points/ranks | Gets the list of ranks for the loyalty points preferred customers.
[**LoyaltyPointsGetStatus**](LoyaltyPointsApi.md#LoyaltyPointsGetStatus) | **Get** /api/v1/loyalty-points/{mainPk} | Gets the status of Loyalty Points for a preferred customer.
[**LoyaltyPointsGetTransactions**](LoyaltyPointsApi.md#LoyaltyPointsGetTransactions) | **Get** /api/v1/loyalty-points/{mainPk}/transactions | Gets the transactions of Loyalty Points for a preferred customer ordered by effective date.


# **LoyaltyPointsGetRanks**
> []RewardPointsCustomerRankViewModel LoyaltyPointsGetRanks(ctx, )
Gets the list of ranks for the loyalty points preferred customers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]RewardPointsCustomerRankViewModel**](RewardPointsCustomerRankViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoyaltyPointsGetStatus**
> RewardPointsCustomerStatusViewModel LoyaltyPointsGetStatus(ctx, mainPk)
Gets the status of Loyalty Points for a preferred customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 

### Return type

[**RewardPointsCustomerStatusViewModel**](RewardPointsCustomerStatusViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoyaltyPointsGetTransactions**
> RewardPointsCustomerTransactionsViewModel LoyaltyPointsGetTransactions(ctx, mainPk, optional)
Gets the transactions of Loyalty Points for a preferred customer ordered by effective date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 
 **optional** | ***LoyaltyPointsApiLoyaltyPointsGetTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoyaltyPointsApiLoyaltyPointsGetTransactionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**RewardPointsCustomerTransactionsViewModel**](RewardPointsCustomerTransactionsViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


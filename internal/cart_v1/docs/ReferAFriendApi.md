# \ReferAFriendApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReferAFriendCreateReferrals**](ReferAFriendApi.md#ReferAFriendCreateReferrals) | **Post** /api/v1/rewards/refer-a-friend/{mainPk} | Save a list of referrals for a preferred customer.
[**ReferAFriendGetReferralsAsync**](ReferAFriendApi.md#ReferAFriendGetReferralsAsync) | **Get** /api/v1/rewards/refer-a-friend | Gets the list of referrals for a preferred customer.
[**ReferAFriendGetSettings**](ReferAFriendApi.md#ReferAFriendGetSettings) | **Get** /api/v1/rewards/refer-a-friend/settings | Gets the refer a friend config.
[**ReferAFriendGetStatus**](ReferAFriendApi.md#ReferAFriendGetStatus) | **Get** /api/v1/rewards/refer-a-friend/status | Gets the customer status for referrals.


# **ReferAFriendCreateReferrals**
> []CustomerReferralReturnModel ReferAFriendCreateReferrals(ctx, mainPk, referrals)
Save a list of referrals for a preferred customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 
  **referrals** | [**CustomerReferralFormModel**](CustomerReferralFormModel.md)|  | 

### Return type

[**[]CustomerReferralReturnModel**](CustomerReferralReturnModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReferAFriendGetReferralsAsync**
> CustomerReferralViewModel ReferAFriendGetReferralsAsync(ctx, optional)
Gets the list of referrals for a preferred customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferAFriendApiReferAFriendGetReferralsAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferAFriendApiReferAFriendGetReferralsAsyncOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerMainPk** | **optional.String**|  | 
 **referralId** | **optional.String**|  | 
 **referralEmail** | **optional.String**|  | 
 **referralFirstName** | **optional.String**|  | 
 **referralLastName** | **optional.String**|  | 
 **referralStatus** | **optional.String**|  | 
 **referralBeforeDate** | **optional.Time**|  | 
 **referralAfterDate** | **optional.Time**|  | 
 **earnedPointsBeforeDate** | **optional.Time**|  | 
 **earnedPointsAfterDate** | **optional.Time**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 
 **skipCount** | **optional.Bool**|  | 

### Return type

[**CustomerReferralViewModel**](CustomerReferralViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReferAFriendGetSettings**
> CustomerReferralSettingsSimpleModel ReferAFriendGetSettings(ctx, )
Gets the refer a friend config.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CustomerReferralSettingsSimpleModel**](CustomerReferralSettingsSimpleModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReferAFriendGetStatus**
> []CustomerReferralStatusModel ReferAFriendGetStatus(ctx, )
Gets the customer status for referrals.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CustomerReferralStatusModel**](CustomerReferralStatusModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


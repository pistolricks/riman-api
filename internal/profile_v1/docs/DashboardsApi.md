# \DashboardsApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardsCheckIfShouldShowAgree**](DashboardsApi.md#DashboardsCheckIfShouldShowAgree) | **Get** /api/v1/dashboards/show-agree/{siteUrl} | Get main user details
[**DashboardsDashbortdPromoWinnersModel**](DashboardsApi.md#DashboardsDashbortdPromoWinnersModel) | **Get** /api/v1/dashboards/promo-winners/{promotionId} | Get list of winners for specific promotion
[**DashboardsGetBpSpSnapshot**](DashboardsApi.md#DashboardsGetBpSpSnapshot) | **Get** /api/v1/dashboards/bp-sp-snapshot | Get data for Personal snapshot
[**DashboardsGetCommisions**](DashboardsApi.md#DashboardsGetCommisions) | **Get** /api/v1/dashboards/commissions | Get Commissions data
[**DashboardsGetCommissionsTotal**](DashboardsApi.md#DashboardsGetCommissionsTotal) | **Get** /api/v1/dashboards/commissions-total | Returns commission total for specified time range.
[**DashboardsGetCustCommTier**](DashboardsApi.md#DashboardsGetCustCommTier) | **Get** /api/v1/dashboards/commissions-tier/{siteUrl} | Get Commissions Tier data
[**DashboardsGetCustomers**](DashboardsApi.md#DashboardsGetCustomers) | **Get** /api/v1/dashboards/customers/{siteUrl}/{year} | Get referred customers
[**DashboardsGetDashboardPromoUserData**](DashboardsApi.md#DashboardsGetDashboardPromoUserData) | **Get** /api/v1/dashboards/promo-user-data/{promotionId} | Get user data for specific promotion
[**DashboardsGetDashboardPromotionId**](DashboardsApi.md#DashboardsGetDashboardPromotionId) | **Get** /api/v1/dashboards/promo-id | Get dashboard promotion report id
[**DashboardsGetDistMassiveActionBonusPromoData**](DashboardsApi.md#DashboardsGetDistMassiveActionBonusPromoData) | **Get** /api/v1/dashboards/massive-action-bonus/{promotionId} | Get data for Massive Action Bonus Promo
[**DashboardsGetEnrolleeSalesDetails**](DashboardsApi.md#DashboardsGetEnrolleeSalesDetails) | **Get** /api/v1/dashboards/enrollee-sales/{siteUrl} | Get Enrollee sales details
[**DashboardsGetEventTicketsSnapshots**](DashboardsApi.md#DashboardsGetEventTicketsSnapshots) | **Get** /api/v1/dashboards/event-tickets-snapshot | Get data for event tickets snapshot
[**DashboardsGetFlushPointsDaysLeft**](DashboardsApi.md#DashboardsGetFlushPointsDaysLeft) | **Get** /api/v1/dashboards/flush-points-days-left | Get flush points days left count to show it on UI
[**DashboardsGetMonthlyProductData**](DashboardsApi.md#DashboardsGetMonthlyProductData) | **Get** /api/v1/dashboards/monthly-product-data | Returns a list of monthly datapoints containing product totals for the requested date range in the requested language.
[**DashboardsGetMonthlySalesData**](DashboardsApi.md#DashboardsGetMonthlySalesData) | **Get** /api/v1/dashboards/monthly-sales-data | Returns a list of monthly datapoints containing sales totals for the requested date range.
[**DashboardsGetNewContactCounts**](DashboardsApi.md#DashboardsGetNewContactCounts) | **Get** /api/v1/dashboards/signups | Returns a list of monthly datapoints containing new contact type counts.
[**DashboardsGetNewReps**](DashboardsApi.md#DashboardsGetNewReps) | **Get** /api/v1/dashboards/new-reps | Returns a list of monthly datapoints containing new signup rep counts.
[**DashboardsGetPersonalInfo**](DashboardsApi.md#DashboardsGetPersonalInfo) | **Get** /api/v1/dashboards/personal-info | Get data for Personal Info
[**DashboardsGetPersonalSnapshot**](DashboardsApi.md#DashboardsGetPersonalSnapshot) | **Get** /api/v1/dashboards/personal-snapshot | Get data for Personal snapshot
[**DashboardsGetSmartNewsFeed**](DashboardsApi.md#DashboardsGetSmartNewsFeed) | **Get** /api/v1/dashboards/smartNews | Get smart news feed
[**DashboardsGetTeamDetails**](DashboardsApi.md#DashboardsGetTeamDetails) | **Get** /api/v1/dashboards/team | Get Team details
[**DashboardsGetTopProducers**](DashboardsApi.md#DashboardsGetTopProducers) | **Get** /api/v1/dashboards/top-producers | Get Top producers
[**DashboardsGetTopSalesContact**](DashboardsApi.md#DashboardsGetTopSalesContact) | **Get** /api/v1/dashboards/top-sales-contact | Returns top customer for specified date range.
[**DashboardsGetVolRSBCommisions**](DashboardsApi.md#DashboardsGetVolRSBCommisions) | **Get** /api/v1/dashboards/volume-rsb-commissions/{siteUrl}/{year} | Get Customer Commissions data
[**DashboardsGetVolRSBSales**](DashboardsApi.md#DashboardsGetVolRSBSales) | **Get** /api/v1/dashboards/volume-rsb-sales/{siteUrl}/{year} | Get Customer Sales data


# **DashboardsCheckIfShouldShowAgree**
> bool DashboardsCheckIfShouldShowAgree(ctx, siteUrl, optional)
Get main user details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
 **optional** | ***DashboardsApiDashboardsCheckIfShouldShowAgreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsCheckIfShouldShowAgreeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsDashbortdPromoWinnersModel**
> []JeunesseProfileCoreDashboardsModelsDashbortdPromoWinnerModel DashboardsDashbortdPromoWinnersModel(ctx, promotionId, optional)
Get list of winners for specific promotion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **promotionId** | **int32**|  | 
 **optional** | ***DashboardsApiDashboardsDashbortdPromoWinnersModelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsDashbortdPromoWinnersModelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreDashboardsModelsDashbortdPromoWinnerModel**](Jeunesse.Profile.Core.Dashboards.Models.DashbortdPromoWinnerModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetBpSpSnapshot**
> []JeunesseProfileCoreDashboardsModelsBpSpSnapshotModel DashboardsGetBpSpSnapshot(ctx, optional)
Get data for Personal snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetBpSpSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetBpSpSnapshotOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreDashboardsModelsBpSpSnapshotModel**](Jeunesse.Profile.Core.Dashboards.Models.BpSpSnapshotModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetCommisions**
> JeunesseProfileCoreDashboardsModelsCommissionModel DashboardsGetCommisions(ctx, optional)
Get Commissions data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetCommisionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetCommisionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreDashboardsModelsCommissionModel**](Jeunesse.Profile.Core.Dashboards.Models.CommissionModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetCommissionsTotal**
> JeunesseProfileCoreDashboardsModelsGetCommissionsTotalApiResponseModel DashboardsGetCommissionsTotal(ctx, optional)
Returns commission total for specified time range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetCommissionsTotalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetCommissionsTotalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 

### Return type

[**JeunesseProfileCoreDashboardsModelsGetCommissionsTotalApiResponseModel**](Jeunesse.Profile.Core.Dashboards.Models.GetCommissionsTotalApiResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetCustCommTier**
> JeunesseProfileCoreDashboardsModelsCustomerCommissionsTierModel DashboardsGetCustCommTier(ctx, siteUrl, optional)
Get Commissions Tier data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
 **optional** | ***DashboardsApiDashboardsGetCustCommTierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetCustCommTierOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreDashboardsModelsCustomerCommissionsTierModel**](Jeunesse.Profile.Core.Dashboards.Models.CustomerCommissionsTierModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetCustomers**
> []JeunesseProfileCoreDashboardsModelsCustomerSalesCommissionsModel DashboardsGetCustomers(ctx, siteUrl, year, optional)
Get referred customers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
  **year** | **int32**|  | 
 **optional** | ***DashboardsApiDashboardsGetCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreDashboardsModelsCustomerSalesCommissionsModel**](Jeunesse.Profile.Core.Dashboards.Models.CustomerSalesCommissionsModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetDashboardPromoUserData**
> JeunesseProfileCoreDashboardsModelsDashboardPromoUserDataModel DashboardsGetDashboardPromoUserData(ctx, promotionId, optional)
Get user data for specific promotion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **promotionId** | **int32**|  | 
 **optional** | ***DashboardsApiDashboardsGetDashboardPromoUserDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetDashboardPromoUserDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreDashboardsModelsDashboardPromoUserDataModel**](Jeunesse.Profile.Core.Dashboards.Models.DashboardPromoUserDataModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetDashboardPromotionId**
> int32 DashboardsGetDashboardPromotionId(ctx, optional)
Get dashboard promotion report id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetDashboardPromotionIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetDashboardPromotionIdOpts struct

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

# **DashboardsGetDistMassiveActionBonusPromoData**
> []JeunesseProfileCoreDashboardsModelsMassiveBonusPromoModel DashboardsGetDistMassiveActionBonusPromoData(ctx, promotionId, optional)
Get data for Massive Action Bonus Promo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **promotionId** | **int32**|  | 
 **optional** | ***DashboardsApiDashboardsGetDistMassiveActionBonusPromoDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetDistMassiveActionBonusPromoDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreDashboardsModelsMassiveBonusPromoModel**](Jeunesse.Profile.Core.Dashboards.Models.MassiveBonusPromoModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetEnrolleeSalesDetails**
> []JeunesseProfileCoreDashboardsModelsEnrolleeSalesModel DashboardsGetEnrolleeSalesDetails(ctx, siteUrl, optional)
Get Enrollee sales details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
 **optional** | ***DashboardsApiDashboardsGetEnrolleeSalesDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetEnrolleeSalesDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreDashboardsModelsEnrolleeSalesModel**](Jeunesse.Profile.Core.Dashboards.Models.EnrolleeSalesModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetEventTicketsSnapshots**
> []JeunesseProfileCoreDashboardsModelsEventTicketsSnapshot DashboardsGetEventTicketsSnapshots(ctx, eventId, optional)
Get data for event tickets snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | **int32**|  | 
 **optional** | ***DashboardsApiDashboardsGetEventTicketsSnapshotsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetEventTicketsSnapshotsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchDay** | **optional.String**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreDashboardsModelsEventTicketsSnapshot**](Jeunesse.Profile.Core.Dashboards.Models.EventTicketsSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetFlushPointsDaysLeft**
> int32 DashboardsGetFlushPointsDaysLeft(ctx, optional)
Get flush points days left count to show it on UI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetFlushPointsDaysLeftOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetFlushPointsDaysLeftOpts struct

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

# **DashboardsGetMonthlyProductData**
> []JeunesseProfileCoreDashboardsModelsSmartMobileGetMonthlyProductDataApiResponseModel DashboardsGetMonthlyProductData(ctx, optional)
Returns a list of monthly datapoints containing product totals for the requested date range in the requested language.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetMonthlyProductDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetMonthlyProductDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 

### Return type

[**[]JeunesseProfileCoreDashboardsModelsSmartMobileGetMonthlyProductDataApiResponseModel**](Jeunesse.Profile.Core.Dashboards.Models.SmartMobile.GetMonthlyProductDataApiResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetMonthlySalesData**
> []JeunesseProfileCoreDashboardsModelsSmartMobileGetMonthlySalesDataApiResponseModel DashboardsGetMonthlySalesData(ctx, optional)
Returns a list of monthly datapoints containing sales totals for the requested date range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetMonthlySalesDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetMonthlySalesDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 

### Return type

[**[]JeunesseProfileCoreDashboardsModelsSmartMobileGetMonthlySalesDataApiResponseModel**](Jeunesse.Profile.Core.Dashboards.Models.SmartMobile.GetMonthlySalesDataApiResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetNewContactCounts**
> []JeunesseProfileCoreDashboardsModelsSmartMobileGetNewContactCountsApiResponseModel DashboardsGetNewContactCounts(ctx, optional)
Returns a list of monthly datapoints containing new contact type counts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetNewContactCountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetNewContactCountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 

### Return type

[**[]JeunesseProfileCoreDashboardsModelsSmartMobileGetNewContactCountsApiResponseModel**](Jeunesse.Profile.Core.Dashboards.Models.SmartMobile.GetNewContactCountsApiResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetNewReps**
> []JeunesseProfileCoreUsersModelsSmartMobileContactModel DashboardsGetNewReps(ctx, optional)
Returns a list of monthly datapoints containing new signup rep counts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetNewRepsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetNewRepsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 

### Return type

[**[]JeunesseProfileCoreUsersModelsSmartMobileContactModel**](Jeunesse.Profile.Core.Users.Models.SmartMobile.ContactModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetPersonalInfo**
> JeunesseProfileCoreDashboardsModelsPersonalInfoModel DashboardsGetPersonalInfo(ctx, optional)
Get data for Personal Info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetPersonalInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetPersonalInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreDashboardsModelsPersonalInfoModel**](Jeunesse.Profile.Core.Dashboards.Models.PersonalInfoModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetPersonalSnapshot**
> JeunesseProfileCoreDashboardsModelsPersonalSnapshotModel DashboardsGetPersonalSnapshot(ctx, optional)
Get data for Personal snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetPersonalSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetPersonalSnapshotOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreDashboardsModelsPersonalSnapshotModel**](Jeunesse.Profile.Core.Dashboards.Models.PersonalSnapshotModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetSmartNewsFeed**
> interface{} DashboardsGetSmartNewsFeed(ctx, optional)
Get smart news feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetSmartNewsFeedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetSmartNewsFeedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **siteUrl** | **optional.String**|  | 
 **mainFK** | **optional.Int32**|  | 
 **length** | **optional.Int32**|  | 
 **usePinPriority** | **optional.Bool**|  | 
 **country** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetTeamDetails**
> []JeunesseProfileCoreDashboardsModelsTeamMembersModel DashboardsGetTeamDetails(ctx, optional)
Get Team details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetTeamDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetTeamDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreDashboardsModelsTeamMembersModel**](Jeunesse.Profile.Core.Dashboards.Models.TeamMembersModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetTopProducers**
> JeunesseProfileCoreDashboardsModelsTopProducersModel DashboardsGetTopProducers(ctx, optional)
Get Top producers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetTopProducersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetTopProducersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreDashboardsModelsTopProducersModel**](Jeunesse.Profile.Core.Dashboards.Models.TopProducersModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetTopSalesContact**
> JeunesseProfileCoreUsersModelsSmartMobileContactModel DashboardsGetTopSalesContact(ctx, optional)
Returns top customer for specified date range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DashboardsApiDashboardsGetTopSalesContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetTopSalesContactOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 

### Return type

[**JeunesseProfileCoreUsersModelsSmartMobileContactModel**](Jeunesse.Profile.Core.Users.Models.SmartMobile.ContactModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetVolRSBCommisions**
> []JeunesseProfileCoreDashboardsModelsCustomerSalesCommissionsModel DashboardsGetVolRSBCommisions(ctx, siteUrl, year, optional)
Get Customer Commissions data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
  **year** | **int32**|  | 
 **optional** | ***DashboardsApiDashboardsGetVolRSBCommisionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetVolRSBCommisionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreDashboardsModelsCustomerSalesCommissionsModel**](Jeunesse.Profile.Core.Dashboards.Models.CustomerSalesCommissionsModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DashboardsGetVolRSBSales**
> []JeunesseProfileCoreDashboardsModelsCustomerSalesCommissionsModel DashboardsGetVolRSBSales(ctx, siteUrl, year, optional)
Get Customer Sales data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
  **year** | **int32**|  | 
 **optional** | ***DashboardsApiDashboardsGetVolRSBSalesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DashboardsApiDashboardsGetVolRSBSalesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreDashboardsModelsCustomerSalesCommissionsModel**](Jeunesse.Profile.Core.Dashboards.Models.CustomerSalesCommissionsModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


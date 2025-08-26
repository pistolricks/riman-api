# \RepSiteApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepSiteAddLead**](RepSiteApi.md#RepSiteAddLead) | **Post** /api/v1/repsite/social/addlead | This is a pass through end point that redirects the data to the smart mobile api.  It manages a couple of different calls to that api and was determined to be better handled  Server side rather than client
[**RepSiteGetAdWords**](RepSiteApi.md#RepSiteGetAdWords) | **Get** /api/v1/repsite/gawords/{alias} | 
[**RepSiteGetDistributorDiscountCodes**](RepSiteApi.md#RepSiteGetDistributorDiscountCodes) | **Get** /api/v1/repsite/discount-codes/{alias} | Returns active discount codes for a specific distributor and their downline
[**RepSiteGetImpressumInfo**](RepSiteApi.md#RepSiteGetImpressumInfo) | **Get** /api/v1/repsite/impressum/{country}/{alias} | Get replicated site impressum information, this is used by Germany and Austria
[**RepSiteGetRepSiteMainPk**](RepSiteApi.md#RepSiteGetRepSiteMainPk) | **Get** /api/v1/repsite/mainpk/{alias} | Gets the unique identifier based on the alias or siturl for a distributor
[**RepSiteGetRepsiteInfo**](RepSiteApi.md#RepSiteGetRepsiteInfo) | **Get** /api/v1/repsite | Gets the sponsor&#39;s replicated site information for the currently logged in user.
[**RepSiteGetRepsiteInfo_0**](RepSiteApi.md#RepSiteGetRepsiteInfo_0) | **Get** /api/v1/repsite/{alias} | Get replicated site information, this is primarily used by the jfront for header navigation links
[**RepSiteInvalidateRepSiteInfo**](RepSiteApi.md#RepSiteInvalidateRepSiteInfo) | **Get** /api/v1/repsite/invalidate | 
[**RepSiteSendReferrerEmail**](RepSiteApi.md#RepSiteSendReferrerEmail) | **Post** /api/v1/repsite/sample/refer | 
[**RepSiteValidateSiteUrl**](RepSiteApi.md#RepSiteValidateSiteUrl) | **Post** /api/v1/repsite/url/validate/{siteUrl} | Validates if requested siteUrl is available and valid


# **RepSiteAddLead**
> interface{} RepSiteAddLead(ctx, lead, optional)
This is a pass through end point that redirects the data to the smart mobile api.  It manages a couple of different calls to that api and was determined to be better handled  Server side rather than client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lead** | [**JeunesseProfileCoreSocialModelsNewLeadFormModel**](JeunesseProfileCoreSocialModelsNewLeadFormModel.md)| Object with identifying properties of the lead | 
 **optional** | ***RepSiteApiRepSiteAddLeadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteAddLeadOpts struct

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

# **RepSiteGetAdWords**
> interface{} RepSiteGetAdWords(ctx, alias, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
 **optional** | ***RepSiteApiRepSiteGetAdWordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteGetAdWordsOpts struct

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

# **RepSiteGetDistributorDiscountCodes**
> []JeunesseProfileCoreRepSiteModelsDiscountCodeViewModel RepSiteGetDistributorDiscountCodes(ctx, alias, optional)
Returns active discount codes for a specific distributor and their downline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
 **optional** | ***RepSiteApiRepSiteGetDistributorDiscountCodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteGetDistributorDiscountCodesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **cartType** | **optional.String**|  | 
 **countryCode** | **optional.String**|  | 

### Return type

[**[]JeunesseProfileCoreRepSiteModelsDiscountCodeViewModel**](Jeunesse.Profile.Core.RepSite.Models.DiscountCodeViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepSiteGetImpressumInfo**
> interface{} RepSiteGetImpressumInfo(ctx, country, alias, optional)
Get replicated site impressum information, this is used by Germany and Austria

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**|  | 
  **alias** | **string**| The siteurl of the distributor | 
 **optional** | ***RepSiteApiRepSiteGetImpressumInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteGetImpressumInfoOpts struct

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

# **RepSiteGetRepSiteMainPk**
> string RepSiteGetRepSiteMainPk(ctx, alias, optional)
Gets the unique identifier based on the alias or siturl for a distributor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| the unique name a distributor has defined for their customers and downline | 
 **optional** | ***RepSiteApiRepSiteGetRepSiteMainPkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteGetRepSiteMainPkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepSiteGetRepsiteInfo**
> JeunesseProfileCoreRepSiteModelsReplicateSiteViewModel RepSiteGetRepsiteInfo(ctx, optional)
Gets the sponsor's replicated site information for the currently logged in user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepSiteApiRepSiteGetRepsiteInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteGetRepsiteInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreRepSiteModelsReplicateSiteViewModel**](Jeunesse.Profile.Core.RepSite.Models.ReplicateSiteViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepSiteGetRepsiteInfo_0**
> JeunesseProfileCoreRepSiteModelsReplicateSiteViewModel RepSiteGetRepsiteInfo_0(ctx, alias, optional)
Get replicated site information, this is primarily used by the jfront for header navigation links

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| The siteurl of the distributor | 
 **optional** | ***RepSiteApiRepSiteGetRepsiteInfo_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteGetRepsiteInfo_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreRepSiteModelsReplicateSiteViewModel**](Jeunesse.Profile.Core.RepSite.Models.ReplicateSiteViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepSiteInvalidateRepSiteInfo**
> interface{} RepSiteInvalidateRepSiteInfo(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepSiteApiRepSiteInvalidateRepSiteInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteInvalidateRepSiteInfoOpts struct

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

# **RepSiteSendReferrerEmail**
> interface{} RepSiteSendReferrerEmail(ctx, model, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreRepSiteModelsSampleReferrerFormModel**](JeunesseProfileCoreRepSiteModelsSampleReferrerFormModel.md)|  | 
 **optional** | ***RepSiteApiRepSiteSendReferrerEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteSendReferrerEmailOpts struct

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

# **RepSiteValidateSiteUrl**
> interface{} RepSiteValidateSiteUrl(ctx, siteUrl, optional)
Validates if requested siteUrl is available and valid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
 **optional** | ***RepSiteApiRepSiteValidateSiteUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepSiteApiRepSiteValidateSiteUrlOpts struct

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


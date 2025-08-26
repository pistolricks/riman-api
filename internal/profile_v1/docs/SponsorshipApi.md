# \SponsorshipApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SponsorshipGetSponsorSiteUrl**](SponsorshipApi.md#SponsorshipGetSponsorSiteUrl) | **Get** /api/v1/sponsorship/sponsor/{siteUrl} | Get the sponsor siteUrl
[**SponsorshipGetSponsorshipTotal**](SponsorshipApi.md#SponsorshipGetSponsorshipTotal) | **Get** /api/v1/sponsorship | Get the sponsorship report data
[**SponsorshipTreeInDownline**](SponsorshipApi.md#SponsorshipTreeInDownline) | **Get** /api/v1/sponsorship/inDownline/{mainSiteUrl}/{downlineSiteUrl} | Check is user downline


# **SponsorshipGetSponsorSiteUrl**
> interface{} SponsorshipGetSponsorSiteUrl(ctx, siteUrl, optional)
Get the sponsor siteUrl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
 **optional** | ***SponsorshipApiSponsorshipGetSponsorSiteUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SponsorshipApiSponsorshipGetSponsorSiteUrlOpts struct

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

# **SponsorshipGetSponsorshipTotal**
> interface{} SponsorshipGetSponsorshipTotal(ctx, optional)
Get the sponsorship report data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SponsorshipApiSponsorshipGetSponsorshipTotalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SponsorshipApiSponsorshipGetSponsorshipTotalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **mainSiteUrl** | **optional.String**|  | 
 **siteURL** | **optional.String**|  | 
 **fName** | **optional.String**|  | 
 **lName** | **optional.String**|  | 
 **displayName** | **optional.String**|  | 
 **isCountedByPlatinum** | **optional.Bool**|  | 
 **statuses** | **optional.String**|  | 
 **highestRanks** | **optional.String**|  | 
 **searchMainPK** | **optional.Int32**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SponsorshipTreeInDownline**
> interface{} SponsorshipTreeInDownline(ctx, mainSiteUrl, downlineSiteUrl, optional)
Check is user downline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainSiteUrl** | **string**|  | 
  **downlineSiteUrl** | **string**|  | 
 **optional** | ***SponsorshipApiSponsorshipTreeInDownlineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SponsorshipApiSponsorshipTreeInDownlineOpts struct

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


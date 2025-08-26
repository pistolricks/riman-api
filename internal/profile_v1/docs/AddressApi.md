# \AddressApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressAddressAutocomplete**](AddressApi.md#AddressAddressAutocomplete) | **Get** /api/v1/address/auto-complete | gets a list of potential address matches. Uses long/lat to narrow focus if provided.
[**AddressAddressGetDetails**](AddressApi.md#AddressAddressGetDetails) | **Post** /api/v1/address/validate | Gets detailed address of provided address or placeId  Use placeId instead of address if possible to reduce cost
[**AddressAddressValidate**](AddressApi.md#AddressAddressValidate) | **Get** /api/v1/address/validate | Validate address by dynamic library


# **AddressAddressAutocomplete**
> []JeunesseProfileCoreAddressAddressModelView AddressAddressAutocomplete(ctx, addressString, country, optional)
gets a list of potential address matches. Uses long/lat to narrow focus if provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressString** | **string**|  | 
  **country** | **string**|  | 
 **optional** | ***AddressApiAddressAddressAutocompleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressApiAddressAddressAutocompleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **latitude** | **optional.String**|  | [default to ]
 **longitude** | **optional.String**|  | [default to ]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreAddressAddressModelView**](Jeunesse.Profile.Core.Address.AddressModelView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressAddressGetDetails**
> JeunesseProfileCoreAddressAddressModelView AddressAddressGetDetails(ctx, optional)
Gets detailed address of provided address or placeId  Use placeId instead of address if possible to reduce cost

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddressApiAddressAddressGetDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressApiAddressAddressGetDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **placeId** | **optional.String**|  | [default to ]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreAddressAddressModelView**](Jeunesse.Profile.Core.Address.AddressModelView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressAddressValidate**
> JeunesseProfileCoreAddressModelsAddressValidationViewModel AddressAddressValidate(ctx, optional)
Validate address by dynamic library

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddressApiAddressAddressValidateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressApiAddressAddressValidateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **selectedShippingMethod** | **optional.Int32**|  | 
 **isAutoComplete** | **optional.Bool**|  | 
 **description** | **optional.String**|  | 
 **address1** | **optional.String**|  | 
 **city** | **optional.String**|  | 
 **state** | **optional.String**|  | 
 **zipCode** | **optional.String**|  | 
 **country** | **optional.String**|  | 
 **placeId** | **optional.String**|  | 

### Return type

[**JeunesseProfileCoreAddressModelsAddressValidationViewModel**](Jeunesse.Profile.Core.Address.Models.AddressValidationViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# \LocalizationApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalizationComunaCheck**](LocalizationApi.md#LocalizationComunaCheck) | **Get** /api/v1/localization/comuna/{stateOrProvince} | Get List of Comuna Codes
[**LocalizationGetAllCountries**](LocalizationApi.md#LocalizationGetAllCountries) | **Get** /api/v1/localization/all-countries | Get list of all countries
[**LocalizationGetAreasForCity**](LocalizationApi.md#LocalizationGetAreasForCity) | **Get** /api/v1/localization/areas/{cityId}/{countryCode} | 
[**LocalizationGetBillingCountries**](LocalizationApi.md#LocalizationGetBillingCountries) | **Get** /api/v1/localization/billing/countries | Returns a list of countries to be used with a billing address.
[**LocalizationGetCitiesForCountry**](LocalizationApi.md#LocalizationGetCitiesForCountry) | **Get** /api/v1/localization/cities/{countryCode}/{isoCode} | 
[**LocalizationGetCountries**](LocalizationApi.md#LocalizationGetCountries) | **Get** /api/v1/localization/countries | Get list of countries
[**LocalizationGetCountryGroups**](LocalizationApi.md#LocalizationGetCountryGroups) | **Get** /api/v1/localization/country-groups | Get list of country groups
[**LocalizationGetCountryZipFormats**](LocalizationApi.md#LocalizationGetCountryZipFormats) | **Get** /api/v1/localization/get-country-zip-formats | 
[**LocalizationGetEuGatewayCountries**](LocalizationApi.md#LocalizationGetEuGatewayCountries) | **Get** /api/v1/localization/eu-gateway-countries | 
[**LocalizationGetGatewayAddress**](LocalizationApi.md#LocalizationGetGatewayAddress) | **Get** /api/v1/localization/gateway-address | 
[**LocalizationGetIdTypes**](LocalizationApi.md#LocalizationGetIdTypes) | **Get** /api/v1/localization/id | 
[**LocalizationGetMinEnrollmentAge**](LocalizationApi.md#LocalizationGetMinEnrollmentAge) | **Get** /api/v1/localization/min-age | 
[**LocalizationGetRegions**](LocalizationApi.md#LocalizationGetRegions) | **Get** /api/v1/localization/countries/{countryCode} | Returns a list of regions for a specified country.
[**LocalizationGetShippingCountries**](LocalizationApi.md#LocalizationGetShippingCountries) | **Get** /api/v1/localization/shipping/countries/{countryCode} | Based on the country requesting, will provide a listing of available countries that can be shipped to
[**LocalizationGetStreetsForArea**](LocalizationApi.md#LocalizationGetStreetsForArea) | **Get** /api/v1/localization/streets/{areaId} | Turkey where it grabs street values
[**LocalizationGetTimezones**](LocalizationApi.md#LocalizationGetTimezones) | **Get** /api/v1/localization/timezones | 
[**LocalizationInvalidateLocalization**](LocalizationApi.md#LocalizationInvalidateLocalization) | **Get** /api/v1/localization/invalidate | 
[**LocalizationLookupJPAddressByPostalCode**](LocalizationApi.md#LocalizationLookupJPAddressByPostalCode) | **Get** /api/v1/localization/jp-address-lookup | Looks for adress suggestions for JP based on provide postal code value


# **LocalizationComunaCheck**
> interface{} LocalizationComunaCheck(ctx, stateOrProvince, optional)
Get List of Comuna Codes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stateOrProvince** | **string**| State or Porvince for searching | 
 **optional** | ***LocalizationApiLocalizationComunaCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationComunaCheckOpts struct

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

# **LocalizationGetAllCountries**
> []JeunesseProfileCoreLocalizationModelsCountrySimpleModel LocalizationGetAllCountries(ctx, optional)
Get list of all countries

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalizationApiLocalizationGetAllCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetAllCountriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreLocalizationModelsCountrySimpleModel**](Jeunesse.Profile.Core.Localization.Models.CountrySimpleModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalizationGetAreasForCity**
> interface{} LocalizationGetAreasForCity(ctx, cityId, countryCode, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cityId** | **string**|  | 
  **countryCode** | **string**|  | 
 **optional** | ***LocalizationApiLocalizationGetAreasForCityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetAreasForCityOpts struct

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

# **LocalizationGetBillingCountries**
> []JeunesseProfileCoreLocalizationModelsBillingCountriesViewModel LocalizationGetBillingCountries(ctx, optional)
Returns a list of countries to be used with a billing address.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalizationApiLocalizationGetBillingCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetBillingCountriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreLocalizationModelsBillingCountriesViewModel**](Jeunesse.Profile.Core.Localization.Models.BillingCountriesViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalizationGetCitiesForCountry**
> interface{} LocalizationGetCitiesForCountry(ctx, countryCode, isoCode, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **isoCode** | **string**|  | 
 **optional** | ***LocalizationApiLocalizationGetCitiesForCountryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetCitiesForCountryOpts struct

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

# **LocalizationGetCountries**
> interface{} LocalizationGetCountries(ctx, optional)
Get list of countries

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalizationApiLocalizationGetCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetCountriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **countryGroups** | **optional.String**|  | 
 **singleCountryCode** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalizationGetCountryGroups**
> interface{} LocalizationGetCountryGroups(ctx, optional)
Get list of country groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalizationApiLocalizationGetCountryGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetCountryGroupsOpts struct

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

# **LocalizationGetCountryZipFormats**
> interface{} LocalizationGetCountryZipFormats(ctx, cultureCode, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cultureCode** | **string**|  | 
 **optional** | ***LocalizationApiLocalizationGetCountryZipFormatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetCountryZipFormatsOpts struct

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

# **LocalizationGetEuGatewayCountries**
> interface{} LocalizationGetEuGatewayCountries(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalizationApiLocalizationGetEuGatewayCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetEuGatewayCountriesOpts struct

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

# **LocalizationGetGatewayAddress**
> interface{} LocalizationGetGatewayAddress(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalizationApiLocalizationGetGatewayAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetGatewayAddressOpts struct

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

# **LocalizationGetIdTypes**
> []JeunesseProfileCoreLocalizationModelsIdTypesModel LocalizationGetIdTypes(ctx, country, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**|  | 
 **optional** | ***LocalizationApiLocalizationGetIdTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetIdTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **langId** | **optional.Int32**|  | [default to 1]
 **pageName** | **optional.String**|  | [default to addnew]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreLocalizationModelsIdTypesModel**](Jeunesse.Profile.Core.Localization.Models.IdTypesModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalizationGetMinEnrollmentAge**
> int32 LocalizationGetMinEnrollmentAge(ctx, country, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**|  | 
 **optional** | ***LocalizationApiLocalizationGetMinEnrollmentAgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetMinEnrollmentAgeOpts struct

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

# **LocalizationGetRegions**
> interface{} LocalizationGetRegions(ctx, countryCode, optional)
Returns a list of regions for a specified country.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**| Two Letter ISO Country code. Example: US for United States | 
 **optional** | ***LocalizationApiLocalizationGetRegionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetRegionsOpts struct

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

# **LocalizationGetShippingCountries**
> interface{} LocalizationGetShippingCountries(ctx, countryCode, optional)
Based on the country requesting, will provide a listing of available countries that can be shipped to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**| The initiating country | 
 **optional** | ***LocalizationApiLocalizationGetShippingCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetShippingCountriesOpts struct

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

# **LocalizationGetStreetsForArea**
> interface{} LocalizationGetStreetsForArea(ctx, areaId, optional)
Turkey where it grabs street values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **areaId** | **int32**| The initiating CountryAreas | 
 **optional** | ***LocalizationApiLocalizationGetStreetsForAreaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetStreetsForAreaOpts struct

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

# **LocalizationGetTimezones**
> []JeunesseProfileCoreLocalizationModelsTimeZoneViewModel LocalizationGetTimezones(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalizationApiLocalizationGetTimezonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationGetTimezonesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreLocalizationModelsTimeZoneViewModel**](Jeunesse.Profile.Core.Localization.Models.TimeZoneViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalizationInvalidateLocalization**
> interface{} LocalizationInvalidateLocalization(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalizationApiLocalizationInvalidateLocalizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationInvalidateLocalizationOpts struct

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

# **LocalizationLookupJPAddressByPostalCode**
> []JeunesseProfileCoreLocalizationModelsJpAddressLookupViewModel LocalizationLookupJPAddressByPostalCode(ctx, optional)
Looks for adress suggestions for JP based on provide postal code value

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalizationApiLocalizationLookupJPAddressByPostalCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalizationApiLocalizationLookupJPAddressByPostalCodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **postalCode** | **optional.String**|  | 

### Return type

[**[]JeunesseProfileCoreLocalizationModelsJpAddressLookupViewModel**](Jeunesse.Profile.Core.Localization.Models.JPAddressLookupViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


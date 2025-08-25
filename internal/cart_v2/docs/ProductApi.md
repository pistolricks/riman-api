# \ProductApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductGetBestSellers**](ProductApi.md#ProductGetBestSellers) | **Get** /api/v2/products/best-sellers | Returns a list of top 5 best selling products.
[**ProductGetBestsellers**](ProductApi.md#ProductGetBestsellers) | **Get** /api/v2/products/bestseller | Returns a list of best selling products.
[**ProductGetBrandSpotlight**](ProductApi.md#ProductGetBrandSpotlight) | **Get** /api/v2/products/brand/best-sellers | Gets a list of products that have been selected to spotlight
[**ProductGetCAPHeadersWithCVRanges**](ProductApi.md#ProductGetCAPHeadersWithCVRanges) | **Get** /api/v2/products/cap-headers/{country}/{langId} | Returns a headers for CAP with CV ranges
[**ProductGetCategories**](ProductApi.md#ProductGetCategories) | **Get** /api/v2/products/categories | Get a listing of all available categories
[**ProductGetCategoryById**](ProductApi.md#ProductGetCategoryById) | **Get** /api/v2/products/categories/{categoryId} | Get specific category information
[**ProductGetMayLoveProducts**](ProductApi.md#ProductGetMayLoveProducts) | **Get** /api/v2/products/may-love | Returns a list of may love products
[**ProductGetProductBrands**](ProductApi.md#ProductGetProductBrands) | **Get** /api/v2/products/brands | Extract all product brands
[**ProductGetProductById**](ProductApi.md#ProductGetProductById) | **Get** /api/v2/products/{productFk}/{cartType}/{culture} | Get specific product information
[**ProductGetProductById_0**](ProductApi.md#ProductGetProductById_0) | **Get** /api/v2/products/detail | Get specific product information
[**ProductGetProductFunctions**](ProductApi.md#ProductGetProductFunctions) | **Get** /api/v2/products/functions | Returns a list of functions/purposes for available products
[**ProductGetProducts**](ProductApi.md#ProductGetProducts) | **Get** /api/v2/products | Get full listing of all products
[**ProductGetRemainingProductQuantities**](ProductApi.md#ProductGetRemainingProductQuantities) | **Get** /api/v2/products/remaining-products | Get remaining product quantities
[**ProductGetRemainingSignUpProductQuantities**](ProductApi.md#ProductGetRemainingSignUpProductQuantities) | **Get** /api/v2/products/remaining-signup-products | Get remaining sign up product quantities
[**ProductGetRitualProducts**](ProductApi.md#ProductGetRitualProducts) | **Get** /api/v2/products/ritual | Returns a list of ritual products
[**ProductGetRituals**](ProductApi.md#ProductGetRituals) | **Get** /api/v2/products/all-rituals | Returns a list of ritual products.
[**ProductGetSubCategories**](ProductApi.md#ProductGetSubCategories) | **Get** /api/v2/products/sub-categories | Get a listing of all available sub categories
[**ProductInvalidateCache**](ProductApi.md#ProductInvalidateCache) | **Get** /api/v2/products/invalidate-cache | Invalidates the local cache
[**ProductSearchProduct**](ProductApi.md#ProductSearchProduct) | **Get** /api/v2/products/search | Search products by keyword


# **ProductGetBestSellers**
> []BestSellersViewModel ProductGetBestSellers(ctx, optional)
Returns a list of top 5 best selling products.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetBestSellersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetBestSellersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartKey** | [**optional.Interface of string**](.md)|  | 
 **productsToShow** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **cartType** | **optional.String**|  | 
 **countryCode** | **optional.String**|  | 
 **culture** | **optional.String**|  | 

### Return type

[**[]BestSellersViewModel**](BestSellersViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetBestsellers**
> []FullProductViewModel ProductGetBestsellers(ctx, optional)
Returns a list of best selling products.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetBestsellersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetBestsellersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCartKey** | [**optional.Interface of string**](.md)|  | 
 **modelProductsToShow** | **optional.String**|  | 
 **modelLimit** | **optional.Int32**|  | 
 **modelCartType** | **optional.String**|  | 
 **modelCountryCode** | **optional.String**|  | 
 **modelCulture** | **optional.String**|  | 

### Return type

[**[]FullProductViewModel**](FullProductViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetBrandSpotlight**
> []BrandViewModel ProductGetBrandSpotlight(ctx, optional)
Gets a list of products that have been selected to spotlight

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetBrandSpotlightOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetBrandSpotlightOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartKey** | [**optional.Interface of string**](.md)|  | 
 **productsToShow** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **cartType** | **optional.String**|  | 
 **countryCode** | **optional.String**|  | 
 **culture** | **optional.String**|  | 

### Return type

[**[]BrandViewModel**](BrandViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetCAPHeadersWithCVRanges**
> []CapHeaderModel ProductGetCAPHeadersWithCVRanges(ctx, country, langId)
Returns a headers for CAP with CV ranges

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country | 
  **langId** | **int32**| LangId | 

### Return type

[**[]CapHeaderModel**](CAPHeaderModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetCategories**
> []CategoryModel ProductGetCategories(ctx, optional)
Get a listing of all available categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetCategoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartType** | **optional.String**|  | 
 **countryCode** | **optional.String**|  | 
 **culture** | **optional.String**|  | 

### Return type

[**[]CategoryModel**](CategoryModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetCategoryById**
> CategoryModel ProductGetCategoryById(ctx, categoryId)
Get specific category information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryId** | **int32**| Product Category PK | 

### Return type

[**CategoryModel**](CategoryModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetMayLoveProducts**
> []SuggestProductModel ProductGetMayLoveProducts(ctx, suggestProductQueryProductPK, optional)
Returns a list of may love products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **suggestProductQueryProductPK** | **int32**|  | 
 **optional** | ***ProductApiProductGetMayLoveProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetMayLoveProductsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suggestProductQueryCartType** | **optional.String**|  | 
 **suggestProductQueryCountryCode** | **optional.String**|  | 
 **suggestProductQueryLanguagePK** | **optional.Int32**|  | 

### Return type

[**[]SuggestProductModel**](SuggestProductModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetProductBrands**
> []ProductBrand ProductGetProductBrands(ctx, )
Extract all product brands

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProductBrand**](ProductBrand.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetProductById**
> FullProductViewModel ProductGetProductById(ctx, productFk, cartType, culture)
Get specific product information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productFk** | **int32**| Product PK | 
  **cartType** | **string**| Cart Type | 
  **culture** | **string**| Culture | 

### Return type

[**FullProductViewModel**](FullProductViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetProductById_0**
> FullProductViewModel ProductGetProductById_0(ctx, optional)
Get specific product information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetProductById_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetProductById_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productPK** | **optional.Int32**|  | 
 **countryCode** | **optional.String**|  | 
 **culture** | **optional.String**|  | 
 **cartType** | **optional.String**|  | 
 **mainFK** | **optional.Int32**|  | 
 **repSiteUrl** | **optional.String**|  | 
 **languageId** | **optional.Int32**|  | 

### Return type

[**FullProductViewModel**](FullProductViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetProductFunctions**
> []BrandViewModel ProductGetProductFunctions(ctx, optional)
Returns a list of functions/purposes for available products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetProductFunctionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetProductFunctionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartType** | **optional.String**|  | 
 **countryCode** | **optional.String**|  | 
 **culture** | **optional.String**|  | 

### Return type

[**[]BrandViewModel**](BrandViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetProducts**
> []FullProductViewModel ProductGetProducts(ctx, optional)
Get full listing of all products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetProductsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartType** | **optional.String**|  | 
 **repSiteUrl** | **optional.String**|  | 
 **mainOrderTypeFk** | **optional.Int32**|  | 
 **mainTypeFk** | **optional.Int32**|  | 
 **countryCode** | **optional.String**|  | 
 **eventId** | **optional.Int32**|  | 
 **categoryId** | **optional.Int32**|  | 
 **otherCategoryId** | **optional.Int32**|  | 
 **culture** | **optional.String**|  | 
 **brandId** | **optional.Int32**|  | 
 **menuId** | **optional.Int32**|  | 
 **functionId** | **optional.Int32**|  | 
 **lineId** | **optional.Int32**|  | 
 **take** | **optional.Int32**|  | 
 **skip** | **optional.Int32**|  | 
 **searchTerm** | **optional.String**|  | 
 **includePromos** | **optional.Bool**|  | 
 **isRedemption** | **optional.Bool**|  | 
 **mainPK** | **optional.Int32**|  | 
 **productPK** | **optional.Int32**|  | 
 **sku** | **optional.String**|  | 
 **priceListPK** | **optional.Int32**|  | 
 **isASEligible** | **optional.Bool**|  | 
 **isPPASEligible** | **optional.Bool**|  | 
 **isAdminEligible** | **optional.Bool**|  | 
 **isSignupEligible** | **optional.Bool**|  | 
 **isSignupPackage** | **optional.Bool**|  | 
 **isStarterKit** | **optional.Bool**|  | 
 **isCreateAPack** | **optional.Bool**|  | 
 **isImportEligible** | **optional.Bool**|  | 
 **isPayCard** | **optional.Bool**|  | 
 **isCart** | **optional.Bool**|  | 
 **currencyCode** | **optional.String**|  | 

### Return type

[**[]FullProductViewModel**](FullProductViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetRemainingProductQuantities**
> []RemainingProductQuantitiesModel ProductGetRemainingProductQuantities(ctx, optional)
Get remaining product quantities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetRemainingProductQuantitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetRemainingProductQuantitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelProductIds** | **optional.String**|  | 
 **modelMainPK** | **optional.Int32**|  | 

### Return type

[**[]RemainingProductQuantitiesModel**](RemainingProductQuantitiesModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetRemainingSignUpProductQuantities**
> []RemainingProductQuantitiesModel ProductGetRemainingSignUpProductQuantities(ctx, optional)
Get remaining sign up product quantities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetRemainingSignUpProductQuantitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetRemainingSignUpProductQuantitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelProductIds** | **optional.String**|  | 
 **modelMainPK** | **optional.Int32**|  | 

### Return type

[**[]RemainingProductQuantitiesModel**](RemainingProductQuantitiesModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetRitualProducts**
> []SuggestProductModel ProductGetRitualProducts(ctx, suggestProductQueryProductPK, optional)
Returns a list of ritual products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **suggestProductQueryProductPK** | **int32**|  | 
 **optional** | ***ProductApiProductGetRitualProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetRitualProductsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suggestProductQueryCartType** | **optional.String**|  | 
 **suggestProductQueryCountryCode** | **optional.String**|  | 
 **suggestProductQueryLanguagePK** | **optional.Int32**|  | 

### Return type

[**[]SuggestProductModel**](SuggestProductModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetRituals**
> []FullProductViewModel ProductGetRituals(ctx, optional)
Returns a list of ritual products.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetRitualsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetRitualsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCartType** | **optional.String**|  | 
 **modelCountryCode** | **optional.String**|  | 
 **modelCulture** | **optional.String**|  | 

### Return type

[**[]FullProductViewModel**](FullProductViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductGetSubCategories**
> []SubMenuBlockModel ProductGetSubCategories(ctx, optional)
Get a listing of all available sub categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductGetSubCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductGetSubCategoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartType** | **optional.String**|  | 
 **countryCode** | **optional.String**|  | 
 **culture** | **optional.String**|  | 

### Return type

[**[]SubMenuBlockModel**](SubMenuBlockModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductInvalidateCache**
> interface{} ProductInvalidateCache(ctx, )
Invalidates the local cache

### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductSearchProduct**
> ProductSearchViewModel ProductSearchProduct(ctx, optional)
Search products by keyword

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductApiProductSearchProductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductApiProductSearchProductOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **optional.String**|  | 
 **culture** | **optional.String**|  | 
 **cartType** | **optional.String**|  | 
 **searchTerm** | **optional.String**|  | 

### Return type

[**ProductSearchViewModel**](ProductSearchViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


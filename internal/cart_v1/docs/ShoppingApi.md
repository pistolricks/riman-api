# \ShoppingApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShoppingApplyDiscount**](ShoppingApi.md#ShoppingApplyDiscount) | **Put** /api/v1/shopping/apply-discount | Applies either a SalesCampign promo to a cart or MainDiscountConfig to an order.
[**ShoppingClearReferrer**](ShoppingApi.md#ShoppingClearReferrer) | **Delete** /api/v1/shopping/{id}/clear-referrer | Clears out the set referrer information on the cart. ** Use patch call going forward **
[**ShoppingCreate**](ShoppingApi.md#ShoppingCreate) | **Post** /api/v1/shopping | Creates a new shopping cart object.
[**ShoppingCreateMainOrder**](ShoppingApi.md#ShoppingCreateMainOrder) | **Post** /api/v1/shopping/{cartKey}/create-order | Used to create a ready to process order record from the main cart information.
[**ShoppingGetCalculatedTaxes**](ShoppingApi.md#ShoppingGetCalculatedTaxes) | **Get** /api/v1/shopping/{cartKey}/estimated-taxes | Returns the estimate tax for the cart.
[**ShoppingGetCampaigns**](ShoppingApi.md#ShoppingGetCampaigns) | **Get** /api/v1/shopping/campaigns | List campaigns
[**ShoppingGetCart**](ShoppingApi.md#ShoppingGetCart) | **Get** /api/v1/shopping/{id} | Gets a cart object by the cart key value.
[**ShoppingGetCheckoutMessage**](ShoppingApi.md#ShoppingGetCheckoutMessage) | **Get** /api/v1/shopping/get-checkout-messages | Returns a list of checkout messages/warnings for checkout
[**ShoppingGetQualifiedSalesCampaignInfo**](ShoppingApi.md#ShoppingGetQualifiedSalesCampaignInfo) | **Post** /api/v1/shopping/campaigns/qualified | Checks to see if the shopping cart has qualified for any sales campaigns
[**ShoppingGetRank**](ShoppingApi.md#ShoppingGetRank) | **Get** /api/v1/shopping/{id}/rank | Get Rank, calculated by items in cart, that user will achieve on purchase
[**ShoppingGetShippingOptions**](ShoppingApi.md#ShoppingGetShippingOptions) | **Get** /api/v1/shopping/{cartKey}/shipping-options | Returns shipping options and pricing available for the cart.
[**ShoppingPatch**](ShoppingApi.md#ShoppingPatch) | **Patch** /api/v1/shopping/{id} | Used for partial updates on a cart object item.
[**ShoppingRemoveCart**](ShoppingApi.md#ShoppingRemoveCart) | **Delete** /api/v1/shopping/{id} | Deletes a cart object.
[**ShoppingSetShipping**](ShoppingApi.md#ShoppingSetShipping) | **Put** /api/v1/shopping/{cartKey}/set-shipping | Sets the shipping info on the cart selected by the user.
[**ShoppingUpdate**](ShoppingApi.md#ShoppingUpdate) | **Put** /api/v1/shopping/{id} | Updates a cart object info.
[**ShoppingValidateCart**](ShoppingApi.md#ShoppingValidateCart) | **Get** /api/v1/shopping/{cartKey}/validate | Call to validate cart information


# **ShoppingApplyDiscount**
> ShoppingApplyDiscount(ctx, model)
Applies either a SalesCampign promo to a cart or MainDiscountConfig to an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**CartDiscountFormModel**](CartDiscountFormModel.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingClearReferrer**
> CartViewModel ShoppingClearReferrer(ctx, id)
Clears out the set referrer information on the cart. ** Use patch call going forward **

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 

### Return type

[**CartViewModel**](CartViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingCreate**
> CartFormModel ShoppingCreate(ctx, model)
Creates a new shopping cart object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**CartFormModel**](CartFormModel.md)| CartFormModel object | 

### Return type

[**CartFormModel**](CartFormModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingCreateMainOrder**
> int32 ShoppingCreateMainOrder(ctx, cartKey)
Used to create a ready to process order record from the main cart information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingGetCalculatedTaxes**
> TaxModel ShoppingGetCalculatedTaxes(ctx, cartKey)
Returns the estimate tax for the cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)| Unique identifier of the cart. | 

### Return type

[**TaxModel**](TaxModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingGetCampaigns**
> []CampaignResponseModel ShoppingGetCampaigns(ctx, optional)
List campaigns

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShoppingApiShoppingGetCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShoppingApiShoppingGetCampaignsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **culture** | **optional.String**|  | 

### Return type

[**[]CampaignResponseModel**](CampaignResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingGetCart**
> CartViewModel ShoppingGetCart(ctx, id)
Gets a cart object by the cart key value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart. | 

### Return type

[**CartViewModel**](CartViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingGetCheckoutMessage**
> []CheckoutMessageViewModel ShoppingGetCheckoutMessage(ctx, optional)
Returns a list of checkout messages/warnings for checkout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ShoppingApiShoppingGetCheckoutMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShoppingApiShoppingGetCheckoutMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **optional.String**|  | 
 **languageFK** | **optional.Int32**|  | 
 **cartKey** | [**optional.Interface of string**](.md)|  | 

### Return type

[**[]CheckoutMessageViewModel**](CheckoutMessageViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingGetQualifiedSalesCampaignInfo**
> GetQualifiedSalesCampaignViewModel ShoppingGetQualifiedSalesCampaignInfo(ctx, model)
Checks to see if the shopping cart has qualified for any sales campaigns

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**GetQualifiedSalesCampaignRequestModel**](GetQualifiedSalesCampaignRequestModel.md)|  | 

### Return type

[**GetQualifiedSalesCampaignViewModel**](GetQualifiedSalesCampaignViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingGetRank**
> GetRankResponseModel ShoppingGetRank(ctx, id)
Get Rank, calculated by items in cart, that user will achieve on purchase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart. | 

### Return type

[**GetRankResponseModel**](GetRankResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingGetShippingOptions**
> ShippingMethodsViewModel ShoppingGetShippingOptions(ctx, cartKey)
Returns shipping options and pricing available for the cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)| Unique identifier of the cart. | 

### Return type

[**ShippingMethodsViewModel**](ShippingMethodsViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingPatch**
> CartViewModel ShoppingPatch(ctx, id, model)
Used for partial updates on a cart object item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 
  **model** | [**CartPatchModel**](CartPatchModel.md)| CartFormModel object | 

### Return type

[**CartViewModel**](CartViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingRemoveCart**
> interface{} ShoppingRemoveCart(ctx, id)
Deletes a cart object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingSetShipping**
> interface{} ShoppingSetShipping(ctx, cartKey, model)
Sets the shipping info on the cart selected by the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)| Unique identifier of the cart. | 
  **model** | [**CartSetShippingFormModel**](CartSetShippingFormModel.md)| Model info for the shipping chosen by the user. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingUpdate**
> CartFormModel ShoppingUpdate(ctx, id, model)
Updates a cart object info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 
  **model** | [**CartFormModel**](CartFormModel.md)| CartFormModel object | 

### Return type

[**CartFormModel**](CartFormModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingValidateCart**
> ValidateCartViewModel ShoppingValidateCart(ctx, cartKey, optional)
Call to validate cart information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cartKey** | [**string**](.md)|  | 
 **optional** | ***ShoppingApiShoppingValidateCartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShoppingApiShoppingValidateCartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderNumber** | **optional.String**|  | 

### Return type

[**ValidateCartViewModel**](ValidateCartViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# \ShoppingItemsApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShoppingItemsAddAffiliateItems**](ShoppingItemsApi.md#ShoppingItemsAddAffiliateItems) | **Post** /api/v1/shopping/{id}/affiliate-items | Used to quickly validate and add items to a cart for affiliate deep linking.
[**ShoppingItemsAddItem**](ShoppingItemsApi.md#ShoppingItemsAddItem) | **Post** /api/v1/shopping/{id}/items | Adds a new product item into a shopping cart.
[**ShoppingItemsAddItems**](ShoppingItemsApi.md#ShoppingItemsAddItems) | **Post** /api/v1/shopping/{id}/items/list | Adds a list of products to the shopping cart.
[**ShoppingItemsAddQuickCartItems**](ShoppingItemsApi.md#ShoppingItemsAddQuickCartItems) | **Post** /api/v1/shopping/{id}/quick-checkout | Used to quickly add items to a cart for fast checkout
[**ShoppingItemsGetCartItems**](ShoppingItemsApi.md#ShoppingItemsGetCartItems) | **Get** /api/v1/shopping/{id}/items | Returns a list of items in a cart.
[**ShoppingItemsPatchItem**](ShoppingItemsApi.md#ShoppingItemsPatchItem) | **Patch** /api/v1/shopping/{id}/items/{itemId} | Patches an existing item in a shopping cart
[**ShoppingItemsRemoveCartItems**](ShoppingItemsApi.md#ShoppingItemsRemoveCartItems) | **Delete** /api/v1/shopping/{id}/items | Used to remove all cart items from a cart.
[**ShoppingItemsRemoveCartItemsList**](ShoppingItemsApi.md#ShoppingItemsRemoveCartItemsList) | **Delete** /api/v1/shopping/{id}/items/list | Used to remove a list of items from a cart.
[**ShoppingItemsRemoveItem**](ShoppingItemsApi.md#ShoppingItemsRemoveItem) | **Delete** /api/v1/shopping/{id}/items/{itemId} | Remove an item from a shopping cart
[**ShoppingItemsUpdateItem**](ShoppingItemsApi.md#ShoppingItemsUpdateItem) | **Put** /api/v1/shopping/{id}/items/{itemId} | Updates an existing item in a shopping cart
[**ShoppingItemsValidateCartItems**](ShoppingItemsApi.md#ShoppingItemsValidateCartItems) | **Get** /api/v1/shopping/{id}/items/validate | Returns a validation message in case items set is prohibited.


# **ShoppingItemsAddAffiliateItems**
> AffiliateCartResponseModel ShoppingItemsAddAffiliateItems(ctx, id, itemsList)
Used to quickly validate and add items to a cart for affiliate deep linking.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| CartKey | 
  **itemsList** | [**CartItemAffiliateModel**](CartItemAffiliateModel.md)| Products | 

### Return type

[**AffiliateCartResponseModel**](AffiliateCartResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingItemsAddItem**
> CartItemViewModel ShoppingItemsAddItem(ctx, id, model)
Adds a new product item into a shopping cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 
  **model** | [**CartItemFormModel**](CartItemFormModel.md)| CartItemModel object | 

### Return type

[**CartItemViewModel**](CartItemViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingItemsAddItems**
> []CartItemViewModel ShoppingItemsAddItems(ctx, id, models)
Adds a list of products to the shopping cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 
  **models** | [**[]CartItemFormModel**](CartItemFormModel.md)| array of CartItemModel object | 

### Return type

[**[]CartItemViewModel**](CartItemViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingItemsAddQuickCartItems**
> AffiliateCartResponseModel ShoppingItemsAddQuickCartItems(ctx, id, itemsList)
Used to quickly add items to a cart for fast checkout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 
  **itemsList** | [**CartItemAffiliateModel**](CartItemAffiliateModel.md)|  | 

### Return type

[**AffiliateCartResponseModel**](AffiliateCartResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingItemsGetCartItems**
> []CartItemViewModel ShoppingItemsGetCartItems(ctx, id)
Returns a list of items in a cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 

### Return type

[**[]CartItemViewModel**](CartItemViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingItemsPatchItem**
> CartItemViewModel ShoppingItemsPatchItem(ctx, id, itemId, model)
Patches an existing item in a shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 
  **itemId** | **int32**| Product PK | 
  **model** | [**CartItemPatchModel**](CartItemPatchModel.md)| CartItemModel object | 

### Return type

[**CartItemViewModel**](CartItemViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingItemsRemoveCartItems**
> interface{} ShoppingItemsRemoveCartItems(ctx, id)
Used to remove all cart items from a cart.

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

# **ShoppingItemsRemoveCartItemsList**
> interface{} ShoppingItemsRemoveCartItemsList(ctx, id, items)
Used to remove a list of items from a cart.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 
  **items** | [**[]int32**](int32.md)| MainCartItemsPk array | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingItemsRemoveItem**
> interface{} ShoppingItemsRemoveItem(ctx, id, itemId)
Remove an item from a shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 
  **itemId** | **int32**| Product PK | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingItemsUpdateItem**
> CartItemViewModel ShoppingItemsUpdateItem(ctx, id, itemId, model)
Updates an existing item in a shopping cart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 
  **itemId** | **int32**| Product PK | 
  **model** | [**CartItemFormModel**](CartItemFormModel.md)| CartItemModel object | 

### Return type

[**CartItemViewModel**](CartItemViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShoppingItemsValidateCartItems**
> []ValidateCartItemsViewModel ShoppingItemsValidateCartItems(ctx, id)
Returns a validation message in case items set is prohibited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Unique identifier of the cart | 

### Return type

[**[]ValidateCartItemsViewModel**](ValidateCartItemsViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


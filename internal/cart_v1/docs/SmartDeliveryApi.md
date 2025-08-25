# \SmartDeliveryApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartDeliveryAddKidsDonation**](SmartDeliveryApi.md#SmartDeliveryAddKidsDonation) | **Put** /api/v1/smartdelivery/items/addkidsdonation | Update smartdelivery item
[**SmartDeliveryCreateSettings**](SmartDeliveryApi.md#SmartDeliveryCreateSettings) | **Post** /api/v1/smartdelivery | Create new SD profile without items
[**SmartDeliveryDeactivateAutoship**](SmartDeliveryApi.md#SmartDeliveryDeactivateAutoship) | **Post** /api/v1/smartdelivery/deactivate | Deactivate user&#39;s autoship
[**SmartDeliveryGetItems**](SmartDeliveryApi.md#SmartDeliveryGetItems) | **Get** /api/v1/smartdelivery/items | Get List of smartdelivery items
[**SmartDeliveryGetPB**](SmartDeliveryApi.md#SmartDeliveryGetPB) | **Get** /api/v1/smartdelivery/prepaid | Get List of prepaid orders
[**SmartDeliveryGetPayments**](SmartDeliveryApi.md#SmartDeliveryGetPayments) | **Get** /api/v1/smartdelivery/payment | Get smartdelivery available payments
[**SmartDeliveryGetSettings**](SmartDeliveryApi.md#SmartDeliveryGetSettings) | **Get** /api/v1/smartdelivery | Get smartdelivery settings
[**SmartDeliveryInsertItem**](SmartDeliveryApi.md#SmartDeliveryInsertItem) | **Post** /api/v1/smartdelivery/items | Insert a smartdelivery item
[**SmartDeliveryInsertItemList**](SmartDeliveryApi.md#SmartDeliveryInsertItemList) | **Post** /api/v1/smartdelivery/items/list | Insert a list smartdelivery item
[**SmartDeliveryRemoveItem**](SmartDeliveryApi.md#SmartDeliveryRemoveItem) | **Delete** /api/v1/smartdelivery/items/{autoshipitempk} | Delete smartdelivery item
[**SmartDeliveryReversePA**](SmartDeliveryApi.md#SmartDeliveryReversePA) | **Put** /api/v1/smartdelivery/prepaid/reverse/{order} | Reverse PA Orders for particular PBOrder for make it refundable
[**SmartDeliverySkipAutoshipMonth**](SmartDeliveryApi.md#SmartDeliverySkipAutoshipMonth) | **Post** /api/v1/smartdelivery/skip | Skip month of user&#39;s autoship
[**SmartDeliveryUpdateItem**](SmartDeliveryApi.md#SmartDeliveryUpdateItem) | **Put** /api/v1/smartdelivery/items/{autoshipitempk} | Update smartdelivery item
[**SmartDeliveryUpdatePBDate**](SmartDeliveryApi.md#SmartDeliveryUpdatePBDate) | **Put** /api/v1/smartdelivery/prepaid/date/{order} | Update StartDate for particular PBOrder and all active following PBOrders
[**SmartDeliveryUpdatePBStatus**](SmartDeliveryApi.md#SmartDeliveryUpdatePBStatus) | **Put** /api/v1/smartdelivery/prepaid/status/{order} | Update status for PBOrder
[**SmartDeliveryUpdateSettings**](SmartDeliveryApi.md#SmartDeliveryUpdateSettings) | **Put** /api/v1/smartdelivery | Update smartdelivery settings


# **SmartDeliveryAddKidsDonation**
> SmartDeliveryItemModel SmartDeliveryAddKidsDonation(ctx, )
Update smartdelivery item

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SmartDeliveryItemModel**](SmartDeliveryItemModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryCreateSettings**
> SmartDeliveryViewModel SmartDeliveryCreateSettings(ctx, )
Create new SD profile without items

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SmartDeliveryViewModel**](SmartDeliveryViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryDeactivateAutoship**
> interface{} SmartDeliveryDeactivateAutoship(ctx, )
Deactivate user's autoship

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

# **SmartDeliveryGetItems**
> []SmartDeliveryItemModel SmartDeliveryGetItems(ctx, )
Get List of smartdelivery items

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SmartDeliveryItemModel**](SmartDeliveryItemModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryGetPB**
> PrepaidViewModel SmartDeliveryGetPB(ctx, optional)
Get List of prepaid orders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SmartDeliveryApiSmartDeliveryGetPBOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SmartDeliveryApiSmartDeliveryGetPBOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainSiteUrl** | **optional.String**|  | 
 **orderNumber** | **optional.String**|  | 
 **startDate** | **optional.Time**|  | 
 **endDate** | **optional.Time**|  | 
 **prepaidStatus** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**PrepaidViewModel**](PrepaidViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryGetPayments**
> []PaymentMethodViewModel SmartDeliveryGetPayments(ctx, )
Get smartdelivery available payments

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PaymentMethodViewModel**](PaymentMethodViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryGetSettings**
> SmartDeliveryViewModel SmartDeliveryGetSettings(ctx, )
Get smartdelivery settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SmartDeliveryViewModel**](SmartDeliveryViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryInsertItem**
> SmartDeliveryItemModel SmartDeliveryInsertItem(ctx, model)
Insert a smartdelivery item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**SmartDeliveryItemFormModel**](SmartDeliveryItemFormModel.md)| form model of item | 

### Return type

[**SmartDeliveryItemModel**](SmartDeliveryItemModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryInsertItemList**
> []SmartDeliveryItemModel SmartDeliveryInsertItemList(ctx, model)
Insert a list smartdelivery item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**[]SmartDeliveryItemFormModel**](SmartDeliveryItemFormModel.md)| form model of item | 

### Return type

[**[]SmartDeliveryItemModel**](SmartDeliveryItemModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryRemoveItem**
> interface{} SmartDeliveryRemoveItem(ctx, autoshipitempk)
Delete smartdelivery item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **autoshipitempk** | **int32**| id of the smartdelivery item | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryReversePA**
> PrepaidPbOrderModel SmartDeliveryReversePA(ctx, order, pborder)
Reverse PA Orders for particular PBOrder for make it refundable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **order** | **int32**|  | 
  **pborder** | [**PrepaidPbOrderModel**](PrepaidPbOrderModel.md)|  | 

### Return type

[**PrepaidPbOrderModel**](PrepaidPBOrderModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliverySkipAutoshipMonth**
> interface{} SmartDeliverySkipAutoshipMonth(ctx, )
Skip month of user's autoship

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

# **SmartDeliveryUpdateItem**
> SmartDeliveryItemModel SmartDeliveryUpdateItem(ctx, autoshipitempk, model)
Update smartdelivery item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **autoshipitempk** | **int32**| id of the smartdelivery item | 
  **model** | [**SmartDeliveryItemFormModel**](SmartDeliveryItemFormModel.md)| form model of item | 

### Return type

[**SmartDeliveryItemModel**](SmartDeliveryItemModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryUpdatePBDate**
> PrepaidPbOrderModel SmartDeliveryUpdatePBDate(ctx, order, pborder)
Update StartDate for particular PBOrder and all active following PBOrders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **order** | **int32**|  | 
  **pborder** | [**PrepaidPbOrderModel**](PrepaidPbOrderModel.md)|  | 

### Return type

[**PrepaidPbOrderModel**](PrepaidPBOrderModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryUpdatePBStatus**
> PrepaidPbOrderModel SmartDeliveryUpdatePBStatus(ctx, order, pborder)
Update status for PBOrder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **order** | **int32**|  | 
  **pborder** | [**PrepaidPbOrderModel**](PrepaidPbOrderModel.md)|  | 

### Return type

[**PrepaidPbOrderModel**](PrepaidPBOrderModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmartDeliveryUpdateSettings**
> SmartDeliveryViewModel SmartDeliveryUpdateSettings(ctx, formmodel)
Update smartdelivery settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **formmodel** | [**SmartDeliveryFormModel**](SmartDeliveryFormModel.md)|  | 

### Return type

[**SmartDeliveryViewModel**](SmartDeliveryViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


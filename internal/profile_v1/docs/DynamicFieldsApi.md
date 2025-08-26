# \DynamicFieldsApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicFieldsGetDataSourceValuesForInputName**](DynamicFieldsApi.md#DynamicFieldsGetDataSourceValuesForInputName) | **Post** /api/v1/form/data-source | Used to pull data source values for a &lt;select&gt;&lt;/select&gt; type input. Simply pass the   name of the input and expect a list of SelectListItem values back.   i.e. { name: string, value: string } An example of this would be pulling a   list of States for the passed in country. Pass the countryCode and the   inputName of \&quot;state\&quot; and get the data needed. In this case the dependent fields  property would be empty.  A slightly more complicated example of when we&#39;d use this would be for   populating the &lt;select&gt;&lt;/select&gt; list of shipping methods. Because the shipping   methods require the state, country, zip (maybe city too), we&#39;d pass their names and   values in the dependent fields property to get the shipping methods.
[**DynamicFieldsGetDynamicFieldData**](DynamicFieldsApi.md#DynamicFieldsGetDynamicFieldData) | **Get** /api/v1/dynamic-fields/get-data | Call to get existing user data for the sent array of fields  It checks against the UpdateSchema JSON to find where to  retrieve from for the key(s) provided
[**DynamicFieldsGetJsonFormData**](DynamicFieldsApi.md#DynamicFieldsGetJsonFormData) | **Get** /api/v1/form | Retrieves the Json form data for the specified country and form type
[**DynamicFieldsGetValueForInputName**](DynamicFieldsApi.md#DynamicFieldsGetValueForInputName) | **Post** /api/v1/form/input-value | Used to pull an actual value for an input based on other inputs entered. This would   always be tied to a populate-type rule associated with the form. An example would   be for the Brazil address form where filling in the postal code will trigger this   API for the Neighborhood, City, etc inputs to auto-populate with values.
[**DynamicFieldsMigrateDynamicSignupData**](DynamicFieldsApi.md#DynamicFieldsMigrateDynamicSignupData) | **Post** /api/v1/dynamic-fields/migrate | Migrates data for signup users to the relevant tables. Uses  the UpdateSchema field to determine where the data should be migrated to.  After copying over it removes the record from the signup dynamic table.
[**DynamicFieldsOptionsSearchPost**](DynamicFieldsApi.md#DynamicFieldsOptionsSearchPost) | **Post** /api/v1/dynamic-fields/options-search | Call to get options list for dropdowns source  Works with parameters
[**DynamicFieldsSetDynamicFieldData**](DynamicFieldsApi.md#DynamicFieldsSetDynamicFieldData) | **Post** /api/v1/dynamic-fields/set-data | Call to update user data for dynamic fields. Checks against the UpdateSchema JSON  to determine where to store data by the fields provided
[**DynamicFieldsValidateFormInput**](DynamicFieldsApi.md#DynamicFieldsValidateFormInput) | **Post** /api/v1/form/validate/{countryCode} | Used to validate the value of a single input. Designed for inputs that would   require API calls to validate. An example would be when we use the username   field when checking out a new customer with SmartDelivery. We use an API to   determine if the username is available.
[**DynamicFieldsValidation**](DynamicFieldsApi.md#DynamicFieldsValidation) | **Post** /api/v1/dynamic-fields/validation | Does validation for to check for duplicates and/or run a  sproc or function to validate the format. Checks against the  UpdateSchema JSON to get info on if it is checking duplicates  or if it needs to build a query for the JSON


# **DynamicFieldsGetDataSourceValuesForInputName**
> []SystemWebWebPagesHtmlSelectListItem DynamicFieldsGetDataSourceValuesForInputName(ctx, queryModel)
Used to pull data source values for a <select></select> type input. Simply pass the   name of the input and expect a list of SelectListItem values back.   i.e. { name: string, value: string } An example of this would be pulling a   list of States for the passed in country. Pass the countryCode and the   inputName of \"state\" and get the data needed. In this case the dependent fields  property would be empty.  A slightly more complicated example of when we'd use this would be for   populating the <select></select> list of shipping methods. Because the shipping   methods require the state, country, zip (maybe city too), we'd pass their names and   values in the dependent fields property to get the shipping methods.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryModel** | [**JeunesseProfileCoreDynamicFieldsModelsPopulateInputQueryModel**](JeunesseProfileCoreDynamicFieldsModelsPopulateInputQueryModel.md)| model of data to retrieve | 

### Return type

[**[]SystemWebWebPagesHtmlSelectListItem**](System.Web.WebPages.Html.SelectListItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicFieldsGetDynamicFieldData**
> []JeunesseProfileCoreDynamicFieldsModelsDynamicValuesViewModel DynamicFieldsGetDynamicFieldData(ctx, optional)
Call to get existing user data for the sent array of fields  It checks against the UpdateSchema JSON to find where to  retrieve from for the key(s) provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DynamicFieldsApiDynamicFieldsGetDynamicFieldDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicFieldsApiDynamicFieldsGetDynamicFieldDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **optional.String**|  | 
 **schemaType** | **optional.Int32**|  | 
 **signupKey** | **optional.String**|  | 
 **mainPk** | **optional.Int32**|  | 
 **cartKey** | [**optional.Interface of string**](.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)|  | 
 **inputName** | **optional.String**|  | 
 **fieldValues** | [**optional.Interface of []interface{}**](interface{}.md)|  | 

### Return type

[**[]JeunesseProfileCoreDynamicFieldsModelsDynamicValuesViewModel**](Jeunesse.Profile.Core.DynamicFields.Models.DynamicValuesViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicFieldsGetJsonFormData**
> []JeunesseProfileCoreDynamicFieldsModelsDynamicFormModel DynamicFieldsGetJsonFormData(ctx, optional)
Retrieves the Json form data for the specified country and form type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DynamicFieldsApiDynamicFieldsGetJsonFormDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicFieldsApiDynamicFieldsGetJsonFormDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **optional.String**|  | 
 **formType** | **optional.Int32**|  | 

### Return type

[**[]JeunesseProfileCoreDynamicFieldsModelsDynamicFormModel**](Jeunesse.Profile.Core.DynamicFields.Models.DynamicFormModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicFieldsGetValueForInputName**
> []JeunesseProfileCoreDynamicFieldsModelsFormInputModel DynamicFieldsGetValueForInputName(ctx, queryModel)
Used to pull an actual value for an input based on other inputs entered. This would   always be tied to a populate-type rule associated with the form. An example would   be for the Brazil address form where filling in the postal code will trigger this   API for the Neighborhood, City, etc inputs to auto-populate with values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **queryModel** | [**JeunesseProfileCoreDynamicFieldsModelsPopulateValuesQueryModel**](JeunesseProfileCoreDynamicFieldsModelsPopulateValuesQueryModel.md)| model of data to retrieve | 

### Return type

[**[]JeunesseProfileCoreDynamicFieldsModelsFormInputModel**](Jeunesse.Profile.Core.DynamicFields.Models.FormInputModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicFieldsMigrateDynamicSignupData**
> interface{} DynamicFieldsMigrateDynamicSignupData(ctx, signupKey, mainPk, country)
Migrates data for signup users to the relevant tables. Uses  the UpdateSchema field to determine where the data should be migrated to.  After copying over it removes the record from the signup dynamic table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **signupKey** | **string**|  | 
  **mainPk** | **int32**|  | 
  **country** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicFieldsOptionsSearchPost**
> interface{} DynamicFieldsOptionsSearchPost(ctx, model)
Call to get options list for dropdowns source  Works with parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreDynamicFieldsModelsDynamicValuesGetFormModel**](JeunesseProfileCoreDynamicFieldsModelsDynamicValuesGetFormModel.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicFieldsSetDynamicFieldData**
> []JeunesseProfileCoreDynamicFieldsModelsDynamicValuesViewModel DynamicFieldsSetDynamicFieldData(ctx, model)
Call to update user data for dynamic fields. Checks against the UpdateSchema JSON  to determine where to store data by the fields provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreDynamicFieldsModelsDynamicValuesPostFormModel**](JeunesseProfileCoreDynamicFieldsModelsDynamicValuesPostFormModel.md)|  | 

### Return type

[**[]JeunesseProfileCoreDynamicFieldsModelsDynamicValuesViewModel**](Jeunesse.Profile.Core.DynamicFields.Models.DynamicValuesViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicFieldsValidateFormInput**
> []JeunesseProfileCoreDynamicFieldsModelsInputValidatedModel DynamicFieldsValidateFormInput(ctx, countryCode, formInput)
Used to validate the value of a single input. Designed for inputs that would   require API calls to validate. An example would be when we use the username   field when checking out a new customer with SmartDelivery. We use an API to   determine if the username is available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**| country to get data for | 
  **formInput** | [**JeunesseProfileCoreDynamicFieldsModelsFormInputModel**](JeunesseProfileCoreDynamicFieldsModelsFormInputModel.md)| the input to validate with the name and value | 

### Return type

[**[]JeunesseProfileCoreDynamicFieldsModelsInputValidatedModel**](Jeunesse.Profile.Core.DynamicFields.Models.InputValidatedModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicFieldsValidation**
> []JeunesseProfileCoreDynamicFieldsModelsDynamicValuesViewModel DynamicFieldsValidation(ctx, model)
Does validation for to check for duplicates and/or run a  sproc or function to validate the format. Checks against the  UpdateSchema JSON to get info on if it is checking duplicates  or if it needs to build a query for the JSON

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreDynamicFieldsModelsDynamicValuesPostFormModel**](JeunesseProfileCoreDynamicFieldsModelsDynamicValuesPostFormModel.md)|  | 

### Return type

[**[]JeunesseProfileCoreDynamicFieldsModelsDynamicValuesViewModel**](Jeunesse.Profile.Core.DynamicFields.Models.DynamicValuesViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


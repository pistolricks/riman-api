# \TrainingApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrainingGetTrainingRecordsForUser**](TrainingApi.md#TrainingGetTrainingRecordsForUser) | **Get** /api/v1/training/records | Returns the list to training records available from Moodle.
[**TrainingRecordTrainingAction**](TrainingApi.md#TrainingRecordTrainingAction) | **Post** /api/v1/training/record | Records a training log for a user
[**TrainingSetupTrainingRecordsForUser**](TrainingApi.md#TrainingSetupTrainingRecordsForUser) | **Post** /api/v1/training/setup | Sets up a training record for a user and returns the training url.


# **TrainingGetTrainingRecordsForUser**
> []JeunesseProfileCoreTrainingModelsTrainingRecordViewModel TrainingGetTrainingRecordsForUser(ctx, optional)
Returns the list to training records available from Moodle.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TrainingApiTrainingGetTrainingRecordsForUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrainingApiTrainingGetTrainingRecordsForUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreTrainingModelsTrainingRecordViewModel**](Jeunesse.Profile.Core.Training.Models.TrainingRecordViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrainingRecordTrainingAction**
> string TrainingRecordTrainingAction(ctx, user, events, date, optional)
Records a training log for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **user** | **int32**|  | 
  **events** | **string**|  | 
  **date** | **time.Time**|  | 
 **optional** | ***TrainingApiTrainingRecordTrainingActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrainingApiTrainingRecordTrainingActionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **course** | **optional.String**|  | [default to ]
 **chapter** | **optional.String**|  | [default to ]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrainingSetupTrainingRecordsForUser**
> JeunesseProfileCoreTrainingModelsTrainingSetupViewModel TrainingSetupTrainingRecordsForUser(ctx, model, optional)
Sets up a training record for a user and returns the training url.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreTrainingModelsTrainingSetupQueryModel**](JeunesseProfileCoreTrainingModelsTrainingSetupQueryModel.md)|  | 
 **optional** | ***TrainingApiTrainingSetupTrainingRecordsForUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrainingApiTrainingSetupTrainingRecordsForUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreTrainingModelsTrainingSetupViewModel**](Jeunesse.Profile.Core.Training.Models.TrainingSetupViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


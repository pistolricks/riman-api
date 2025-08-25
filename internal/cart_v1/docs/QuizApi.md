# \QuizApi

All URIs are relative to *https://cart-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuizGetProductByQuizAnswers**](QuizApi.md#QuizGetProductByQuizAnswers) | **Post** /api/v1/quiz | Returns the product based on the quiz answers
[**QuizGetQuiz**](QuizApi.md#QuizGetQuiz) | **Get** /api/v1/quiz | Returns the quiz questions and answers


# **QuizGetProductByQuizAnswers**
> FullProductViewModel QuizGetProductByQuizAnswers(ctx, quizAnswer, optional)
Returns the product based on the quiz answers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **quizAnswer** | [**QuizAnswerFormModel**](QuizAnswerFormModel.md)| Answer IDs | 
 **optional** | ***QuizApiQuizGetProductByQuizAnswersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuizApiQuizGetProductByQuizAnswersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countryCode** | **optional.String**|  | [default to ]

### Return type

[**FullProductViewModel**](FullProductViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuizGetQuiz**
> QuizModel QuizGetQuiz(ctx, optional)
Returns the quiz questions and answers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuizApiQuizGetQuizOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuizApiQuizGetQuizOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **optional.String**|  | [default to ]

### Return type

[**QuizModel**](QuizModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


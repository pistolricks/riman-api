# \MainUsersApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MainUsersAddUserAward**](MainUsersApi.md#MainUsersAddUserAward) | **Post** /api/v1/users/award | 
[**MainUsersAllowUserChangeUsername**](MainUsersApi.md#MainUsersAllowUserChangeUsername) | **Get** /api/v1/users/allow-change-username | Check is user can change username
[**MainUsersCanUserChangeSponsor**](MainUsersApi.md#MainUsersCanUserChangeSponsor) | **Get** /api/v1/users/can-change-sponsor | Check is user can change his sponsor
[**MainUsersChangeUserSponsor**](MainUsersApi.md#MainUsersChangeUserSponsor) | **Post** /api/v1/users/change-user-sponsor | Change user sponsor data
[**MainUsersChangeUsername**](MainUsersApi.md#MainUsersChangeUsername) | **Post** /api/v1/users/change-username | 
[**MainUsersCheckUserDownline**](MainUsersApi.md#MainUsersCheckUserDownline) | **Get** /api/v1/users/{MainPk}/downline/{DownlineMainPk} | Checks a user Downline for a specific user
[**MainUsersCheckUserEnrollerTree**](MainUsersApi.md#MainUsersCheckUserEnrollerTree) | **Get** /api/v1/users/{MainPk}/enroller-tree/{EnrollerTreeMainPk} | Checks a user Enroller Tree for a specific user
[**MainUsersCreateSignupUser**](MainUsersApi.md#MainUsersCreateSignupUser) | **Post** /api/v1/users/signup | Creates a new signup user
[**MainUsersGetAwardsTypes**](MainUsersApi.md#MainUsersGetAwardsTypes) | **Get** /api/v1/users/award-types | 
[**MainUsersGetESignatureInfo**](MainUsersApi.md#MainUsersGetESignatureInfo) | **Get** /api/v1/users/e-signature | Returns the E-Signature info for the provided user.
[**MainUsersGetMainUserSponsorPlacementView**](MainUsersApi.md#MainUsersGetMainUserSponsorPlacementView) | **Get** /api/v1/users/sponsor-placement | Returns sponsor placement display text
[**MainUsersGetRedirectRule**](MainUsersApi.md#MainUsersGetRedirectRule) | **Get** /api/v1/users/get-redirect-rule | Returns the redirection rule of the user.
[**MainUsersGetReferrerInfo**](MainUsersApi.md#MainUsersGetReferrerInfo) | **Get** /api/v1/users/referrer | Anonymous; gets limited referer info
[**MainUsersGetRenewalStatusInfoAsync**](MainUsersApi.md#MainUsersGetRenewalStatusInfoAsync) | **Get** /api/v1/users/get-renewal-status-info | Returns the renewal info of the user.
[**MainUsersGetSignupUser**](MainUsersApi.md#MainUsersGetSignupUser) | **Get** /api/v1/users/signup | gets a signup user
[**MainUsersGetUser**](MainUsersApi.md#MainUsersGetUser) | **Get** /api/v1/users/{MainPK} | Returns a user by MainPk
[**MainUsersGetUserAwards**](MainUsersApi.md#MainUsersGetUserAwards) | **Get** /api/v1/users/awards/{MainPK} | 
[**MainUsersGetUserBySiteUrl**](MainUsersApi.md#MainUsersGetUserBySiteUrl) | **Get** /api/v1/users/siteUrl/{SiteUrl} | Returns a user by siteUrl
[**MainUsersGetUserHighestRank**](MainUsersApi.md#MainUsersGetUserHighestRank) | **Get** /api/v1/users/user-highest-rank | Get User highest rank for talkDesk system
[**MainUsersGetUserMarketingInfo**](MainUsersApi.md#MainUsersGetUserMarketingInfo) | **Get** /api/v1/users/marketing/{MainPK} | 
[**MainUsersGetUserRecognitionInfo**](MainUsersApi.md#MainUsersGetUserRecognitionInfo) | **Get** /api/v1/users/recognition/{MainPK} | 
[**MainUsersGetUsers**](MainUsersApi.md#MainUsersGetUsers) | **Get** /api/v1/users | Returns a list of users from the Main Table
[**MainUsersIsNewSponsorValid**](MainUsersApi.md#MainUsersIsNewSponsorValid) | **Get** /api/v1/users/is-sponsor-valid/{sponsorSiteUrl} | Check is sponsor is valid for user
[**MainUsersRemindUsername**](MainUsersApi.md#MainUsersRemindUsername) | **Post** /api/v1/users/remind-username | Remind username for specific user.
[**MainUsersResetPassword**](MainUsersApi.md#MainUsersResetPassword) | **Post** /api/v1/users/reset-password | Resets password for specific user.
[**MainUsersSetESignatureInfo**](MainUsersApi.md#MainUsersSetESignatureInfo) | **Post** /api/v1/users/e-signature | Sets the E-Signature info for a user.
[**MainUsersSyncCustomers**](MainUsersApi.md#MainUsersSyncCustomers) | **Get** /api/v1/users/customers/sync | Sync specific user
[**MainUsersUpdateMarketingInfo**](MainUsersApi.md#MainUsersUpdateMarketingInfo) | **Put** /api/v1/users/marketing | Updates user marketing info
[**MainUsersUpdatePassword**](MainUsersApi.md#MainUsersUpdatePassword) | **Post** /api/v1/users/update-password | Reset password by resetcode
[**MainUsersUpdateProfile**](MainUsersApi.md#MainUsersUpdateProfile) | **Put** /api/v1/users | Updates the profile info for the logged in user.
[**MainUsersUpdateSignupUser**](MainUsersApi.md#MainUsersUpdateSignupUser) | **Put** /api/v1/users/signup | updates a signup user
[**MainUsersUpdateSoD**](MainUsersApi.md#MainUsersUpdateSoD) | **Put** /api/v1/users/update-success-on-demand/{MainPk}/{NewStatus} | Updates Success on Demand status
[**MainUsersUpdateUserRecognitionInfo**](MainUsersApi.md#MainUsersUpdateUserRecognitionInfo) | **Put** /api/v1/users/recognition | Updates user recognition info
[**MainUsersValidateEmail**](MainUsersApi.md#MainUsersValidateEmail) | **Post** /api/v1/users/validate-email | Validates whether the given email is unique across all member types (customers/planners) and listed countries.


# **MainUsersAddUserAward**
> []JeunesseBackDataPocoAwardsAward MainUsersAddUserAward(ctx, award, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **award** | [**JeunesseBackDataPocoAwardsAward**](JeunesseBackDataPocoAwardsAward.md)|  | 
 **optional** | ***MainUsersApiMainUsersAddUserAwardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersAddUserAwardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseBackDataPocoAwardsAward**](Jeunesse.Back.Data.Poco.Awards.Award.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersAllowUserChangeUsername**
> bool MainUsersAllowUserChangeUsername(ctx, optional)
Check is user can change username

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersAllowUserChangeUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersAllowUserChangeUsernameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersCanUserChangeSponsor**
> bool MainUsersCanUserChangeSponsor(ctx, optional)
Check is user can change his sponsor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersCanUserChangeSponsorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersCanUserChangeSponsorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersChangeUserSponsor**
> bool MainUsersChangeUserSponsor(ctx, model, optional)
Change user sponsor data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreUsersModelsUserSponsorChangeModel**](JeunesseProfileCoreUsersModelsUserSponsorChangeModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersChangeUserSponsorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersChangeUserSponsorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersChangeUsername**
> bool MainUsersChangeUsername(ctx, newUsername, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newUsername** | **string**|  | 
 **optional** | ***MainUsersApiMainUsersChangeUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersChangeUsernameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersCheckUserDownline**
> JeunesseProfileCoreUsersModelsMainUserDownlineCheckModel MainUsersCheckUserDownline(ctx, mainPk, downlineMainPk, optional)
Checks a user Downline for a specific user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 
  **downlineMainPk** | **int32**|  | 
 **optional** | ***MainUsersApiMainUsersCheckUserDownlineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersCheckUserDownlineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsMainUserDownlineCheckModel**](Jeunesse.Profile.Core.Users.Models.MainUserDownlineCheckModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersCheckUserEnrollerTree**
> JeunesseProfileCoreUsersModelsMainUserEnrollerTreeModel MainUsersCheckUserEnrollerTree(ctx, mainPk, enrollerTreeMainPk, optional)
Checks a user Enroller Tree for a specific user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 
  **enrollerTreeMainPk** | **int32**|  | 
 **optional** | ***MainUsersApiMainUsersCheckUserEnrollerTreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersCheckUserEnrollerTreeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsMainUserEnrollerTreeModel**](Jeunesse.Profile.Core.Users.Models.MainUserEnrollerTreeModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersCreateSignupUser**
> JeunesseProfileCoreUsersModelsSignupUserViewModel MainUsersCreateSignupUser(ctx, model, optional)
Creates a new signup user

minimum populated fields are country, languageFk, referrerSiteUrl, maintypeFk, source, lastUrl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreUsersModelsSignupUserFormModel**](JeunesseProfileCoreUsersModelsSignupUserFormModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersCreateSignupUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersCreateSignupUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsSignupUserViewModel**](Jeunesse.Profile.Core.Users.Models.SignupUserViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetAwardsTypes**
> []JeunesseBackDataPocoAwardsAwardType MainUsersGetAwardsTypes(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersGetAwardsTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetAwardsTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseBackDataPocoAwardsAwardType**](Jeunesse.Back.Data.Poco.Awards.AwardType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetESignatureInfo**
> JeunesseProfileCoreUsersModelsESignatureInfoViewModel MainUsersGetESignatureInfo(ctx, optional)
Returns the E-Signature info for the provided user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersGetESignatureInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetESignatureInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **mainPK** | **optional.Int32**|  | 

### Return type

[**JeunesseProfileCoreUsersModelsESignatureInfoViewModel**](Jeunesse.Profile.Core.Users.Models.ESignatureInfoViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetMainUserSponsorPlacementView**
> JeunesseProfileCoreUsersModelsMainUserSponsorPlacementViewModel MainUsersGetMainUserSponsorPlacementView(ctx, optional)
Returns sponsor placement display text

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersGetMainUserSponsorPlacementViewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetMainUserSponsorPlacementViewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **mainFK** | **optional.Int32**|  | 
 **languageFk** | **optional.Int32**|  | 

### Return type

[**JeunesseProfileCoreUsersModelsMainUserSponsorPlacementViewModel**](Jeunesse.Profile.Core.Users.Models.MainUserSponsorPlacementViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetRedirectRule**
> JeunesseProfileCoreUsersModelsRedirectRule MainUsersGetRedirectRule(ctx, optional)
Returns the redirection rule of the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersGetRedirectRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetRedirectRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **mainPK** | **optional.Int32**|  | 
 **country** | **optional.String**|  | 
 **sourceURL** | **optional.String**|  | 

### Return type

[**JeunesseProfileCoreUsersModelsRedirectRule**](Jeunesse.Profile.Core.Users.Models.RedirectRule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetReferrerInfo**
> JeunesseProfileCoreUsersModelsSignupUserViewModel MainUsersGetReferrerInfo(ctx, siteUrl, optional)
Anonymous; gets limited referer info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
 **optional** | ***MainUsersApiMainUsersGetReferrerInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetReferrerInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkFlow** | **optional.Bool**|  | [default to false]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsSignupUserViewModel**](Jeunesse.Profile.Core.Users.Models.SignupUserViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetRenewalStatusInfoAsync**
> JeunesseProfileCoreUsersModelsRenewalStatusInfo MainUsersGetRenewalStatusInfoAsync(ctx, optional)
Returns the renewal info of the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersGetRenewalStatusInfoAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetRenewalStatusInfoAsyncOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **mainPK** | **optional.Int32**|  | 

### Return type

[**JeunesseProfileCoreUsersModelsRenewalStatusInfo**](Jeunesse.Profile.Core.Users.Models.RenewalStatusInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetSignupUser**
> JeunesseProfileCoreUsersModelsSignupUserViewModel MainUsersGetSignupUser(ctx, optional)
gets a signup user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersGetSignupUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetSignupUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guid** | **optional.String**|  | [default to ]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsSignupUserViewModel**](Jeunesse.Profile.Core.Users.Models.SignupUserViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetUser**
> JeunesseProfileCoreUsersModelsMainUserModel MainUsersGetUser(ctx, mainPK, optional)
Returns a user by MainPk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPK** | **int32**|  | 
 **optional** | ***MainUsersApiMainUsersGetUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsMainUserModel**](Jeunesse.Profile.Core.Users.Models.MainUserModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetUserAwards**
> []JeunesseBackDataPocoAwardsAward MainUsersGetUserAwards(ctx, mainPK, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPK** | **int32**|  | 
 **optional** | ***MainUsersApiMainUsersGetUserAwardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetUserAwardsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseBackDataPocoAwardsAward**](Jeunesse.Back.Data.Poco.Awards.Award.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetUserBySiteUrl**
> JeunesseProfileCoreUsersModelsMainUserModel MainUsersGetUserBySiteUrl(ctx, siteUrl, optional)
Returns a user by siteUrl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteUrl** | **string**|  | 
 **optional** | ***MainUsersApiMainUsersGetUserBySiteUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetUserBySiteUrlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsMainUserModel**](Jeunesse.Profile.Core.Users.Models.MainUserModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetUserHighestRank**
> interface{} MainUsersGetUserHighestRank(ctx, optional)
Get User highest rank for talkDesk system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersGetUserHighestRankOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetUserHighestRankOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **mainPK** | **optional.Int32**|  | 
 **key** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetUserMarketingInfo**
> JeunesseProfileCoreUsersModelsUserMarketingModel MainUsersGetUserMarketingInfo(ctx, mainPK, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPK** | **int32**|  | 
 **optional** | ***MainUsersApiMainUsersGetUserMarketingInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetUserMarketingInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsUserMarketingModel**](Jeunesse.Profile.Core.Users.Models.UserMarketingModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetUserRecognitionInfo**
> JeunesseProfileCoreUsersModelsUserRecognitionModel MainUsersGetUserRecognitionInfo(ctx, mainPK, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPK** | **int32**|  | 
 **optional** | ***MainUsersApiMainUsersGetUserRecognitionInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetUserRecognitionInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsUserRecognitionModel**](Jeunesse.Profile.Core.Users.Models.UserRecognitionModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersGetUsers**
> JeunesseProfileCoreUsersModelsMainUsersSimpleViewModel MainUsersGetUsers(ctx, optional)
Returns a list of users from the Main Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersGetUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **search** | **optional.String**|  | 
 **exactSearch** | **optional.Bool**|  | 
 **siteUrl** | **optional.String**|  | 
 **mainPK** | **optional.Int32**|  | 
 **email** | **optional.String**|  | 
 **firstName** | **optional.String**|  | 
 **lastName** | **optional.String**|  | 
 **company** | **optional.String**|  | 
 **phone** | **optional.String**|  | 
 **sponsorSiteUrl** | **optional.String**|  | 
 **orderNumber** | **optional.Int32**|  | 
 **skipCount** | **optional.Bool**|  | 
 **mainType** | **optional.String**|  | 
 **orderBy** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**JeunesseProfileCoreUsersModelsMainUsersSimpleViewModel**](Jeunesse.Profile.Core.Users.Models.MainUsersSimpleViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersIsNewSponsorValid**
> bool MainUsersIsNewSponsorValid(ctx, sponsorSiteUrl, optional)
Check is sponsor is valid for user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sponsorSiteUrl** | **string**|  | 
 **optional** | ***MainUsersApiMainUsersIsNewSponsorValidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersIsNewSponsorValidOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersRemindUsername**
> string MainUsersRemindUsername(ctx, model, optional)
Remind username for specific user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreUsersModelsRemindUsernameModel**](JeunesseProfileCoreUsersModelsRemindUsernameModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersRemindUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersRemindUsernameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersResetPassword**
> string MainUsersResetPassword(ctx, model, optional)
Resets password for specific user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreUsersModelsResetPasswordModel**](JeunesseProfileCoreUsersModelsResetPasswordModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersResetPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersResetPasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersSetESignatureInfo**
> JeunesseProfileCoreUsersModelsESignatureInfoViewModel MainUsersSetESignatureInfo(ctx, model, optional)
Sets the E-Signature info for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreUsersModelsESignatureInfoPostModel**](JeunesseProfileCoreUsersModelsESignatureInfoPostModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersSetESignatureInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersSetESignatureInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsESignatureInfoViewModel**](Jeunesse.Profile.Core.Users.Models.ESignatureInfoViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersSyncCustomers**
> JeunesseProfileCoreUsersModelsSmartMobileContactsResponseModel MainUsersSyncCustomers(ctx, optional)
Sync specific user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MainUsersApiMainUsersSyncCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersSyncCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastSyncedDate** | **optional.Time**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsSmartMobileContactsResponseModel**](Jeunesse.Profile.Core.Users.Models.SmartMobile.ContactsResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersUpdateMarketingInfo**
> JeunesseProfileCoreUsersModelsUserMarketingModel MainUsersUpdateMarketingInfo(ctx, model, optional)
Updates user marketing info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreUsersModelsUserMarketingFormModel**](JeunesseProfileCoreUsersModelsUserMarketingFormModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersUpdateMarketingInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersUpdateMarketingInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsUserMarketingModel**](Jeunesse.Profile.Core.Users.Models.UserMarketingModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersUpdatePassword**
> bool MainUsersUpdatePassword(ctx, model, optional)
Reset password by resetcode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreUsersModelsUpdatePasswordModel**](JeunesseProfileCoreUsersModelsUpdatePasswordModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersUpdatePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersUpdatePasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersUpdateProfile**
> JeunesseProfileCoreUsersModelsMainUserModel MainUsersUpdateProfile(ctx, model, optional)
Updates the profile info for the logged in user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsMemberFormModel**](JeunesseProfileCoreMemberModelsMemberFormModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersUpdateProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersUpdateProfileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsMainUserModel**](Jeunesse.Profile.Core.Users.Models.MainUserModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersUpdateSignupUser**
> JeunesseProfileCoreUsersModelsSignupUserViewModel MainUsersUpdateSignupUser(ctx, model, optional)
updates a signup user

minimum populated fields are country, languageFk, referrerSiteUrl, maintypeFk, source, lastUrl

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreUsersModelsSignupUserFormModel**](JeunesseProfileCoreUsersModelsSignupUserFormModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersUpdateSignupUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersUpdateSignupUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsSignupUserViewModel**](Jeunesse.Profile.Core.Users.Models.SignupUserViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersUpdateSoD**
> bool MainUsersUpdateSoD(ctx, mainPk, newStatus, optional)
Updates Success on Demand status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 
  **newStatus** | **bool**|  | 
 **optional** | ***MainUsersApiMainUsersUpdateSoDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersUpdateSoDOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersUpdateUserRecognitionInfo**
> JeunesseProfileCoreUsersModelsUserRecognitionFormModel MainUsersUpdateUserRecognitionInfo(ctx, model, optional)
Updates user recognition info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreUsersModelsUserRecognitionFormModel**](JeunesseProfileCoreUsersModelsUserRecognitionFormModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersUpdateUserRecognitionInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersUpdateUserRecognitionInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreUsersModelsUserRecognitionFormModel**](Jeunesse.Profile.Core.Users.Models.UserRecognitionFormModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MainUsersValidateEmail**
> bool MainUsersValidateEmail(ctx, request, optional)
Validates whether the given email is unique across all member types (customers/planners) and listed countries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**JeunesseProfileCoreUsersModelsEmailValidationQueryModel**](JeunesseProfileCoreUsersModelsEmailValidationQueryModel.md)|  | 
 **optional** | ***MainUsersApiMainUsersValidateEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MainUsersApiMainUsersValidateEmailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


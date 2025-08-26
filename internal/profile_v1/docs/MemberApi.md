# \MemberApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberBonusCreditsGoingToExpire**](MemberApi.md#MemberBonusCreditsGoingToExpire) | **Get** /api/v1/member/bonus-credits-going-to-expire/{daysCount} | Check if user has bonus credits going to expire
[**MemberConvertAnonymousMemberToCustomer**](MemberApi.md#MemberConvertAnonymousMemberToCustomer) | **Post** /api/v1/member/convert | Used to convert an anonymous member account to a customer log in account.
[**MemberCreateMember**](MemberApi.md#MemberCreateMember) | **Post** /api/v1/member | Creates a new Jeunesse member
[**MemberCreateMemberFromEventGuest**](MemberApi.md#MemberCreateMemberFromEventGuest) | **Post** /api/v1/member-from-event-guest | Sets up a new member record for event guest checkout
[**MemberCreatePaymentMethod**](MemberApi.md#MemberCreatePaymentMethod) | **Post** /api/v1/payment | Creates a new payment method and returns an updated list of payment details
[**MemberCreateTicket**](MemberApi.md#MemberCreateTicket) | **Post** /api/v1/member/create-ticket | Allow distributors from India create a compliance ticket.
[**MemberGetAddressLookUp**](MemberApi.md#MemberGetAddressLookUp) | **Get** /api/v1/member/zipcode/{countryCode}/{zipCode} | Grabs Brazil address values: Address2, Address3, City, and State
[**MemberGetBillingAddress**](MemberApi.md#MemberGetBillingAddress) | **Get** /api/v1/member/address/billing | Returns the saved shipping address for a member
[**MemberGetHighestRealRank**](MemberApi.md#MemberGetHighestRealRank) | **Get** /api/v1/highestrank | To get highest rank of user.
[**MemberGetMemberInfo**](MemberApi.md#MemberGetMemberInfo) | **Get** /api/v1/member/{mainpk} | Will lookup a customer by their primary key, if found, will return the key to their referrer
[**MemberGetPaymentDetails**](MemberApi.md#MemberGetPaymentDetails) | **Get** /api/v1/payment | Gets a list of payment details
[**MemberGetReferrer**](MemberApi.md#MemberGetReferrer) | **Get** /api/v1/member/referrer | Returns the referrer information for the currently logged in user,  or if the siteurl of a distributor is passed in it will return the referrer info  if for the siteurl if they are under the logged in distributors downline.
[**MemberGetShippingAddress**](MemberApi.md#MemberGetShippingAddress) | **Get** /api/v1/member/address/shipping | A function to return the saved shipping address for a member
[**MemberGetTWUnifiedDonationOrgOptions**](MemberApi.md#MemberGetTWUnifiedDonationOrgOptions) | **Get** /api/v1/member/unified-donation-options | Returns a list of donation type options for Taiwan Unified settings
[**MemberGetTWUnifiedInfo**](MemberApi.md#MemberGetTWUnifiedInfo) | **Get** /api/v1/member/unified-profile | Returns the unified info options including the set values for the main user.
[**MemberGetUpgradeLink**](MemberApi.md#MemberGetUpgradeLink) | **Get** /api/v1/member/upgrade-link/{mainPk}/{code2} | 
[**MemberGetUserUnreadCounts**](MemberApi.md#MemberGetUserUnreadCounts) | **Get** /api/v1/member/unread-counts | Collect all unread counts to display in badges
[**MemberInsertUpdateTWUnifiedInfo**](MemberApi.md#MemberInsertUpdateTWUnifiedInfo) | **Post** /api/v1/member/unified-profile | Inserts or updates the Unified Invoice profile settings for a TW distributor.
[**MemberIsIxoPayActiveForGateway**](MemberApi.md#MemberIsIxoPayActiveForGateway) | **Get** /api/v1/member/IsIxoPayActiveForGateway | Returns config for IxoPay tokenEx iframe
[**MemberPayPalAgreementAvailable**](MemberApi.md#MemberPayPalAgreementAvailable) | **Get** /api/v1/payment/paypal-agreement | Checks to see if PayPal Agreement available for user country
[**MemberRemovePaymentMethod**](MemberApi.md#MemberRemovePaymentMethod) | **Put** /api/v1/payment/{sequencePk}/{useOrder} | deletes a single payment detail for a user and returns an updated list of payment details
[**MemberRequiresCvv**](MemberApi.md#MemberRequiresCvv) | **Get** /api/v1/payment/requires-cvv | Checks to see if user will need cvv field in recurring cc info
[**MemberSendImsEmailToUser**](MemberApi.md#MemberSendImsEmailToUser) | **Get** /api/v1/member/ims/send-email/{siteurl} | Utilized by IMS, this will send an email to the users downline if all criteria is met
[**MemberSendSmsMessage**](MemberApi.md#MemberSendSmsMessage) | **Get** /api/v1/member/sms | 
[**MemberSetEmailAndPassword**](MemberApi.md#MemberSetEmailAndPassword) | **Post** /api/v1/member/set-email-and-password | Used to set the email and password for a guest account after checkout.
[**MemberSetMemberPriceList**](MemberApi.md#MemberSetMemberPriceList) | **Post** /api/v1/set-price-list | Sets the pricelistPK on an existing member&#39;s account
[**MemberSetReferrer**](MemberApi.md#MemberSetReferrer) | **Put** /api/v1/member/referrer | 
[**MemberUpdateBillingAddress**](MemberApi.md#MemberUpdateBillingAddress) | **Put** /api/v1/member/address/billing | Updates a member&#39;s billing address.
[**MemberUpdateIsShippingAddress**](MemberApi.md#MemberUpdateIsShippingAddress) | **Put** /api/v1/member/update-isUseShippingAddress | Update isUseShippingAddress in Main Table
[**MemberUpdatePassword**](MemberApi.md#MemberUpdatePassword) | **Put** /api/v1/password | Updates the users password
[**MemberUpdatePaymentDetails**](MemberApi.md#MemberUpdatePaymentDetails) | **Put** /api/v1/payment | updates a single payment detail for a user and returns an updated list of payment details
[**MemberUpdatePaymentOrder**](MemberApi.md#MemberUpdatePaymentOrder) | **Put** /api/v1/payment/order | updates a single payment order
[**MemberUpdateShippingAddress**](MemberApi.md#MemberUpdateShippingAddress) | **Put** /api/v1/member/address/shipping | Updates a member&#39;s shipping address.
[**MemberUpdateShowWelcome**](MemberApi.md#MemberUpdateShowWelcome) | **Put** /api/v1/member/update-show-welcome | 
[**MemberUpdateTOSStatus**](MemberApi.md#MemberUpdateTOSStatus) | **Post** /api/v1/member/tos | Updates a user&#39;s terms of service status.
[**MemberUpgrade**](MemberApi.md#MemberUpgrade) | **Get** /api/v1/member/upgrade | Update affiliate to Distributor
[**MemberUserAlertCheck**](MemberApi.md#MemberUserAlertCheck) | **Get** /api/v1/member/alertCheck | 
[**MemberValidateSiteUrl**](MemberApi.md#MemberValidateSiteUrl) | **Post** /api/v1/member/gov-id/validate | Validates if government ID is valid for the given country
[**MemberVerifyCpf**](MemberApi.md#MemberVerifyCpf) | **Get** /api/v1/member/verify-cpf | 


# **MemberBonusCreditsGoingToExpire**
> interface{} MemberBonusCreditsGoingToExpire(ctx, daysCount, optional)
Check if user has bonus credits going to expire

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **daysCount** | **int32**|  | 
 **optional** | ***MemberApiMemberBonusCreditsGoingToExpireOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberBonusCreditsGoingToExpireOpts struct

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

# **MemberConvertAnonymousMemberToCustomer**
> interface{} MemberConvertAnonymousMemberToCustomer(ctx, model, optional)
Used to convert an anonymous member account to a customer log in account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsConvertAnonymousMemberFormModel**](JeunesseProfileCoreMemberModelsConvertAnonymousMemberFormModel.md)|  | 
 **optional** | ***MemberApiMemberConvertAnonymousMemberToCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberConvertAnonymousMemberToCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberCreateMember**
> int32 MemberCreateMember(ctx, model, optional)
Creates a new Jeunesse member

Required Fields:  Please use the V2 version of this call!  SiteUrl, FirstName, LastName, Address1, City, State, PostalCode, Country, ShipFirstName, ShipLastName, ShipAddress1, ShipCity, ShipState, ShipPostalCode, ShipCountry, PhoneNumber, Email, Password, Referrer, MainType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberMemberModel**](JeunesseProfileCoreMemberMemberModel.md)|  | 
 **optional** | ***MemberApiMemberCreateMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberCreateMemberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberCreateMemberFromEventGuest**
> JeunesseProfileCoreMemberModelsMemberFromEventGuestViewModel MemberCreateMemberFromEventGuest(ctx, model, optional)
Sets up a new member record for event guest checkout

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileApiViewModelsAddMemberFromEventGuestFormModel**](JeunesseProfileApiViewModelsAddMemberFromEventGuestFormModel.md)|  | 
 **optional** | ***MemberApiMemberCreateMemberFromEventGuestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberCreateMemberFromEventGuestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreMemberModelsMemberFromEventGuestViewModel**](Jeunesse.Profile.Core.Member.Models.MemberFromEventGuestViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberCreatePaymentMethod**
> JeunesseProfileCoreMemberModelsPaymentViewModel MemberCreatePaymentMethod(ctx, model, optional)
Creates a new payment method and returns an updated list of payment details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsPaymentFormModel**](JeunesseProfileCoreMemberModelsPaymentFormModel.md)|  | 
 **optional** | ***MemberApiMemberCreatePaymentMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberCreatePaymentMethodOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreMemberModelsPaymentViewModel**](Jeunesse.Profile.Core.Member.Models.PaymentViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberCreateTicket**
> interface{} MemberCreateTicket(ctx, model, optional)
Allow distributors from India create a compliance ticket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsComplianceFormModel**](JeunesseProfileCoreMemberModelsComplianceFormModel.md)|  | 
 **optional** | ***MemberApiMemberCreateTicketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberCreateTicketOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberGetAddressLookUp**
> interface{} MemberGetAddressLookUp(ctx, countryCode, zipCode, optional)
Grabs Brazil address values: Address2, Address3, City, and State

Require Fields  countryCode and zipCode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**|  | 
  **zipCode** | **string**|  | 
 **optional** | ***MemberApiMemberGetAddressLookUpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetAddressLookUpOpts struct

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

# **MemberGetBillingAddress**
> interface{} MemberGetBillingAddress(ctx, optional)
Returns the saved shipping address for a member

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberGetBillingAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetBillingAddressOpts struct

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

# **MemberGetHighestRealRank**
> int32 MemberGetHighestRealRank(ctx, optional)
To get highest rank of user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberGetHighestRealRankOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetHighestRealRankOpts struct

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

# **MemberGetMemberInfo**
> JeunesseProfileCoreMemberCustomerModel MemberGetMemberInfo(ctx, mainpk, optional)
Will lookup a customer by their primary key, if found, will return the key to their referrer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainpk** | **int32**|  | 
 **optional** | ***MemberApiMemberGetMemberInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetMemberInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreMemberCustomerModel**](Jeunesse.Profile.Core.Member.CustomerModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberGetPaymentDetails**
> []JeunesseProfileCoreMemberModelsPaymentViewModel MemberGetPaymentDetails(ctx, optional)
Gets a list of payment details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberGetPaymentDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetPaymentDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreMemberModelsPaymentViewModel**](Jeunesse.Profile.Core.Member.Models.PaymentViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberGetReferrer**
> interface{} MemberGetReferrer(ctx, optional)
Returns the referrer information for the currently logged in user,  or if the siteurl of a distributor is passed in it will return the referrer info  if for the siteurl if they are under the logged in distributors downline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberGetReferrerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetReferrerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteUrl** | **optional.String**| Optional SiteURL | [default to ]
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberGetShippingAddress**
> interface{} MemberGetShippingAddress(ctx, optional)
A function to return the saved shipping address for a member

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberGetShippingAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetShippingAddressOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **isOTG** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberGetTWUnifiedDonationOrgOptions**
> JeunesseProfileCoreMemberModelsTwUnifiedTwUnifiedDonationOrgOptions MemberGetTWUnifiedDonationOrgOptions(ctx, optional)
Returns a list of donation type options for Taiwan Unified settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberGetTWUnifiedDonationOrgOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetTWUnifiedDonationOrgOptionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreMemberModelsTwUnifiedTwUnifiedDonationOrgOptions**](Jeunesse.Profile.Core.Member.Models.TWUnified.TWUnifiedDonationOrgOptions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberGetTWUnifiedInfo**
> []JeunesseProfileCoreMemberModelsTwUnifiedTwUnifiedInfo MemberGetTWUnifiedInfo(ctx, optional)
Returns the unified info options including the set values for the main user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberGetTWUnifiedInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetTWUnifiedInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **mainFK** | **optional.Int32**|  | 

### Return type

[**[]JeunesseProfileCoreMemberModelsTwUnifiedTwUnifiedInfo**](Jeunesse.Profile.Core.Member.Models.TWUnified.TWUnifiedInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberGetUpgradeLink**
> interface{} MemberGetUpgradeLink(ctx, mainPk, code2, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainPk** | **int32**|  | 
  **code2** | **string**|  | 
 **optional** | ***MemberApiMemberGetUpgradeLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetUpgradeLinkOpts struct

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

# **MemberGetUserUnreadCounts**
> interface{} MemberGetUserUnreadCounts(ctx, optional)
Collect all unread counts to display in badges

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberGetUserUnreadCountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberGetUserUnreadCountsOpts struct

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

# **MemberInsertUpdateTWUnifiedInfo**
> JeunesseProfileCoreMemberModelsTwUnifiedTwUnifiedInfo MemberInsertUpdateTWUnifiedInfo(ctx, request, optional)
Inserts or updates the Unified Invoice profile settings for a TW distributor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**JeunesseProfileCoreMemberModelsTwUnifiedTwUnifiedInfo**](JeunesseProfileCoreMemberModelsTwUnifiedTwUnifiedInfo.md)|  | 
 **optional** | ***MemberApiMemberInsertUpdateTWUnifiedInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberInsertUpdateTWUnifiedInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreMemberModelsTwUnifiedTwUnifiedInfo**](Jeunesse.Profile.Core.Member.Models.TWUnified.TWUnifiedInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberIsIxoPayActiveForGateway**
> interface{} MemberIsIxoPayActiveForGateway(ctx, optional)
Returns config for IxoPay tokenEx iframe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberIsIxoPayActiveForGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberIsIxoPayActiveForGatewayOpts struct

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

# **MemberPayPalAgreementAvailable**
> bool MemberPayPalAgreementAvailable(ctx, optional)
Checks to see if PayPal Agreement available for user country

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberPayPalAgreementAvailableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberPayPalAgreementAvailableOpts struct

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

# **MemberRemovePaymentMethod**
> JeunesseProfileCoreMemberModelsPaymentViewModel MemberRemovePaymentMethod(ctx, sequencePk, useOrder, optional)
deletes a single payment detail for a user and returns an updated list of payment details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sequencePk** | **int32**|  | 
  **useOrder** | **int32**|  | 
 **optional** | ***MemberApiMemberRemovePaymentMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberRemovePaymentMethodOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreMemberModelsPaymentViewModel**](Jeunesse.Profile.Core.Member.Models.PaymentViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberRequiresCvv**
> bool MemberRequiresCvv(ctx, optional)
Checks to see if user will need cvv field in recurring cc info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberRequiresCvvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberRequiresCvvOpts struct

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

# **MemberSendImsEmailToUser**
> interface{} MemberSendImsEmailToUser(ctx, siteurl, optional)
Utilized by IMS, this will send an email to the users downline if all criteria is met

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteurl** | **string**|  | 
 **optional** | ***MemberApiMemberSendImsEmailToUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberSendImsEmailToUserOpts struct

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

# **MemberSendSmsMessage**
> interface{} MemberSendSmsMessage(ctx, phoneNumber, message, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phoneNumber** | **string**|  | 
  **message** | **string**|  | 
 **optional** | ***MemberApiMemberSendSmsMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberSendSmsMessageOpts struct

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

# **MemberSetEmailAndPassword**
> bool MemberSetEmailAndPassword(ctx, model, optional)
Used to set the email and password for a guest account after checkout.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsUpgradeRetailFormModel**](JeunesseProfileCoreMemberModelsUpgradeRetailFormModel.md)| Post Model | 
 **optional** | ***MemberApiMemberSetEmailAndPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberSetEmailAndPasswordOpts struct

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

# **MemberSetMemberPriceList**
> interface{} MemberSetMemberPriceList(ctx, model, optional)
Sets the pricelistPK on an existing member's account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsSetMemberPriceListFormModel**](JeunesseProfileCoreMemberModelsSetMemberPriceListFormModel.md)|  | 
 **optional** | ***MemberApiMemberSetMemberPriceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberSetMemberPriceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberSetReferrer**
> interface{} MemberSetReferrer(ctx, model, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsSetReferrerFormModel**](JeunesseProfileCoreMemberModelsSetReferrerFormModel.md)|  | 
 **optional** | ***MemberApiMemberSetReferrerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberSetReferrerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberUpdateBillingAddress**
> bool MemberUpdateBillingAddress(ctx, model, optional)
Updates a member's billing address.

Required Fields:  FirstName, LastName, Address1, City, State, PostalCode, Country  billing address should only be updated in connection with a credit card (when updating credit card info)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberAddressModel**](JeunesseProfileCoreMemberAddressModel.md)|  | 
 **optional** | ***MemberApiMemberUpdateBillingAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUpdateBillingAddressOpts struct

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

# **MemberUpdateIsShippingAddress**
> bool MemberUpdateIsShippingAddress(ctx, model, optional)
Update isUseShippingAddress in Main Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsUpdateIsShippingAddressModel**](JeunesseProfileCoreMemberModelsUpdateIsShippingAddressModel.md)|  | 
 **optional** | ***MemberApiMemberUpdateIsShippingAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUpdateIsShippingAddressOpts struct

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

# **MemberUpdatePassword**
> MemberUpdatePassword(ctx, model, optional)
Updates the users password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsPasswordFormModel**](JeunesseProfileCoreMemberModelsPasswordFormModel.md)|  | 
 **optional** | ***MemberApiMemberUpdatePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUpdatePasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberUpdatePaymentDetails**
> JeunesseProfileCoreMemberModelsCreatePaymentMethodModel MemberUpdatePaymentDetails(ctx, model, optional)
updates a single payment detail for a user and returns an updated list of payment details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsPaymentFormModel**](JeunesseProfileCoreMemberModelsPaymentFormModel.md)|  | 
 **optional** | ***MemberApiMemberUpdatePaymentDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUpdatePaymentDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreMemberModelsCreatePaymentMethodModel**](Jeunesse.Profile.Core.Member.Models.CreatePaymentMethodModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberUpdatePaymentOrder**
> bool MemberUpdatePaymentOrder(ctx, model, optional)
updates a single payment order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberModelsPaymentOrderModel**](JeunesseProfileCoreMemberModelsPaymentOrderModel.md)|  | 
 **optional** | ***MemberApiMemberUpdatePaymentOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUpdatePaymentOrderOpts struct

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

# **MemberUpdateShippingAddress**
> bool MemberUpdateShippingAddress(ctx, model, optional)
Updates a member's shipping address.

Required Fields:  FirstName, LastName, Address1, City, State, PostalCode, Country

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberAddressModel**](JeunesseProfileCoreMemberAddressModel.md)|  | 
 **optional** | ***MemberApiMemberUpdateShippingAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUpdateShippingAddressOpts struct

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

# **MemberUpdateShowWelcome**
> bool MemberUpdateShowWelcome(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberUpdateShowWelcomeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUpdateShowWelcomeOpts struct

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

# **MemberUpdateTOSStatus**
> interface{} MemberUpdateTOSStatus(ctx, optional)
Updates a user's terms of service status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberUpdateTOSStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUpdateTOSStatusOpts struct

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

# **MemberUpgrade**
> bool MemberUpgrade(ctx, state, mainType, mainFk, agreeTerms, optional)
Update affiliate to Distributor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **state** | **string**|  | 
  **mainType** | **int32**|  | 
  **mainFk** | **int32**|  | 
  **agreeTerms** | **bool**|  | 
 **optional** | ***MemberApiMemberUpgradeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUpgradeOpts struct

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

# **MemberUserAlertCheck**
> interface{} MemberUserAlertCheck(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MemberApiMemberUserAlertCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberUserAlertCheckOpts struct

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

# **MemberValidateSiteUrl**
> interface{} MemberValidateSiteUrl(ctx, model, optional)
Validates if government ID is valid for the given country

CN_Ssn = 1,  TW_ResidentID = 2,  TH_ID = 5,  US_Ssn = 6,  MX_Ssn = 7,  VN_ResidentID = 8,  PH_Ssn = 11,  ID_Ssn = 12,  MY_Ssn = 13,  SG_Ssm = 14,  CA_Sin = 15,  US_Ein = 16,  TW_Arc = 17,  TW_UnifiedID = 18,  KR_Ssn = 19,  VN_PassportNo = 20,  HK_IdCard = 22,  HK_PassportNo = 23,  HK_BR =24,  ID_PassportNo = 25,  ID_DrivingLicense = 26,  ID_IdentityNumber = 27,  DO_Ssn = 28,  CX_Cpf = 29,  TR_NationalID = 31,  KW_Ssn = 32,  KW_Ein = 33,  US_Fin = 34,  FR_Ssn = 35,  FR_Siret = 36,  IL_IdNumber = 37,  IT_Ssn = 38,  IT_IdCard = 39,  IT_Passport = 40,  IT_DriversLicense = 41,  PE_Ssn = 42,  PA_Ssn = 43,  PE_Dni = 44,  PE_CE = 45,  PE_Rui = 46,  CO_CitizenID = 47,  CO_ForeignerID = 48,  CO_Passport = 49,  IL_CompanyNumber = 50,  CA_BusinessNumber = 51,  CR_CitizenID = 52,  CR_ForeignerID = 53

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreMemberGovIdModel**](JeunesseProfileCoreMemberGovIdModel.md)|  | 
 **optional** | ***MemberApiMemberValidateSiteUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberValidateSiteUrlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MemberVerifyCpf**
> interface{} MemberVerifyCpf(ctx, cpf, mainPk, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cpf** | **string**|  | 
  **mainPk** | **int32**|  | 
 **optional** | ***MemberApiMemberVerifyCpfOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MemberApiMemberVerifyCpfOpts struct

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


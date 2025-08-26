# \WalletApi

All URIs are relative to *https://profile-api.riman.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletAddFundsGetRedirectUrl**](WalletApi.md#WalletAddFundsGetRedirectUrl) | **Post** /api/v1/wallet/add-funds | 
[**WalletAdminUpdatePhotoIdData**](WalletApi.md#WalletAdminUpdatePhotoIdData) | **Post** /api/v1/wallet/admin-photoId | 
[**WalletCreateFundRequest**](WalletApi.md#WalletCreateFundRequest) | **Post** /api/v1/wallet/funds-request | 
[**WalletCreateSignupToken**](WalletApi.md#WalletCreateSignupToken) | **Post** /api/v1/wallet/signup-tokens | 
[**WalletDeactivateBankAccount**](WalletApi.md#WalletDeactivateBankAccount) | **Post** /api/v1/wallet/eft/deactivate | 
[**WalletDeleteFundRequest**](WalletApi.md#WalletDeleteFundRequest) | **Delete** /api/v1/wallet/funds-request/{fundRequestId} | 
[**WalletDeleteSignupToken**](WalletApi.md#WalletDeleteSignupToken) | **Delete** /api/v1/wallet/signup-tokens/{tokenId} | 
[**WalletGetAccountCurrencies**](WalletApi.md#WalletGetAccountCurrencies) | **Get** /api/v1/wallet/eft/account-currencies/{bankCountry} | 
[**WalletGetAccountDetails**](WalletApi.md#WalletGetAccountDetails) | **Get** /api/v1/wallet/account | 
[**WalletGetAccountHistory**](WalletApi.md#WalletGetAccountHistory) | **Get** /api/v1/wallet/account-history | 
[**WalletGetBankCountries**](WalletApi.md#WalletGetBankCountries) | **Get** /api/v1/wallet/eft/bank-countries/{isLocalEFT} | 
[**WalletGetEFTLocalStatus**](WalletApi.md#WalletGetEFTLocalStatus) | **Get** /api/v1/wallet/eft/local/status | 
[**WalletGetEFTRequiredDocuments**](WalletApi.md#WalletGetEFTRequiredDocuments) | **Get** /api/v1/wallet/eft/local/required-documents | 
[**WalletGetEWalletLoginForm**](WalletApi.md#WalletGetEWalletLoginForm) | **Post** /api/v1/wallet/ewallet-login-form | 
[**WalletGetFundRequestCheckRuns**](WalletApi.md#WalletGetFundRequestCheckRuns) | **Get** /api/v1/wallet/fund-request-checkruns | 
[**WalletGetFundRequestInvoiceDetails**](WalletApi.md#WalletGetFundRequestInvoiceDetails) | **Get** /api/v1/wallet/fund-request-invoice-details | 
[**WalletGetFundsRequest**](WalletApi.md#WalletGetFundsRequest) | **Get** /api/v1/wallet/funds-request | 
[**WalletGetFundsRequestSettings**](WalletApi.md#WalletGetFundsRequestSettings) | **Get** /api/v1/wallet/funds-request/settings | 
[**WalletGetIPayoutCardsBin**](WalletApi.md#WalletGetIPayoutCardsBin) | **Get** /api/v1/wallet/paycard/ipayoutbin | 
[**WalletGetIdentificationTypes**](WalletApi.md#WalletGetIdentificationTypes) | **Get** /api/v1/wallet/identification-types | 
[**WalletGetPayCardReplaceState**](WalletApi.md#WalletGetPayCardReplaceState) | **Get** /api/v1/wallet/paycard/replace | 
[**WalletGetPayCardState**](WalletApi.md#WalletGetPayCardState) | **Get** /api/v1/wallet/paycard/state | 
[**WalletGetPayCardValidation**](WalletApi.md#WalletGetPayCardValidation) | **Post** /api/v1/wallet/paycard/validate | 
[**WalletGetPersonalInfo**](WalletApi.md#WalletGetPersonalInfo) | **Get** /api/v1/wallet/eft/personal-info | 
[**WalletGetPersonalInfoLocal**](WalletApi.md#WalletGetPersonalInfoLocal) | **Get** /api/v1/wallet/eft/local/personal-info | 
[**WalletGetPhotoIdData**](WalletApi.md#WalletGetPhotoIdData) | **Get** /api/v1/wallet/photoId | 
[**WalletGetSignupTokens**](WalletApi.md#WalletGetSignupTokens) | **Get** /api/v1/wallet/signup-tokens | 
[**WalletGetUserBankBook**](WalletApi.md#WalletGetUserBankBook) | **Get** /api/v1/wallet/bank-book | 
[**WalletGetUserBankDocument**](WalletApi.md#WalletGetUserBankDocument) | **Get** /api/v1/wallet/bank-document | 
[**WalletGetUserPaymentMethods**](WalletApi.md#WalletGetUserPaymentMethods) | **Get** /api/v1/wallet/payment-methods | 
[**WalletGetUserPayoutMethods**](WalletApi.md#WalletGetUserPayoutMethods) | **Get** /api/v1/wallet/payout-methods/{activeOnly} | 
[**WalletGetWalletBalance**](WalletApi.md#WalletGetWalletBalance) | **Get** /api/v1/wallet/account-history/wallet-balance | 
[**WalletGetWalletSettings**](WalletApi.md#WalletGetWalletSettings) | **Get** /api/v1/wallet/settings | 
[**WalletRegisterPaycard**](WalletApi.md#WalletRegisterPaycard) | **Post** /api/v1/wallet/paycard | 
[**WalletReplacePayCard**](WalletApi.md#WalletReplacePayCard) | **Post** /api/v1/wallet/paycard/replace | 
[**WalletResetWalletPassword**](WalletApi.md#WalletResetWalletPassword) | **Post** /api/v1/wallet/reset-password | 
[**WalletSaveEFTBankInfo**](WalletApi.md#WalletSaveEFTBankInfo) | **Post** /api/v1/wallet/eft/bank-info | 
[**WalletSaveEFTLocalBankInfo**](WalletApi.md#WalletSaveEFTLocalBankInfo) | **Post** /api/v1/wallet/eft/local/bank-info | 
[**WalletUpdatePhotoIdData**](WalletApi.md#WalletUpdatePhotoIdData) | **Post** /api/v1/wallet/photoId | 
[**WalletUpdateUserBankBook**](WalletApi.md#WalletUpdateUserBankBook) | **Post** /api/v1/wallet/bank-book | 
[**WalletUpdateUserBankDocument**](WalletApi.md#WalletUpdateUserBankDocument) | **Post** /api/v1/wallet/bank-document | 
[**WalletUpdateUserFundRequestInvoice**](WalletApi.md#WalletUpdateUserFundRequestInvoice) | **Post** /api/v1/wallet/fund-request-update-invoice | 
[**WalletUpdateUserPayoutMethods**](WalletApi.md#WalletUpdateUserPayoutMethods) | **Post** /api/v1/wallet/payout-methods | 


# **WalletAddFundsGetRedirectUrl**
> string WalletAddFundsGetRedirectUrl(ctx, form, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **form** | [**JeunesseProfileCoreWalletModelsAddFundsFormModel**](JeunesseProfileCoreWalletModelsAddFundsFormModel.md)|  | 
 **optional** | ***WalletApiWalletAddFundsGetRedirectUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletAddFundsGetRedirectUrlOpts struct

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

# **WalletAdminUpdatePhotoIdData**
> string WalletAdminUpdatePhotoIdData(ctx, model, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsPhotoIdAdminUpdatePhotoIdFormModel**](JeunesseProfileCoreWalletModelsPhotoIdAdminUpdatePhotoIdFormModel.md)|  | 
 **optional** | ***WalletApiWalletAdminUpdatePhotoIdDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletAdminUpdatePhotoIdDataOpts struct

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

# **WalletCreateFundRequest**
> bool WalletCreateFundRequest(ctx, formModel, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **formModel** | [**JeunesseProfileCoreWalletModelsFundRequestFundRequestFormModel**](JeunesseProfileCoreWalletModelsFundRequestFundRequestFormModel.md)|  | 
 **optional** | ***WalletApiWalletCreateFundRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletCreateFundRequestOpts struct

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

# **WalletCreateSignupToken**
> JeunesseProfileCoreWalletModelsSignupTokensGetSignupTokensReportResult WalletCreateSignupToken(ctx, commandFormModel, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commandFormModel** | [**JeunesseProfileCoreWalletModelsSignupTokensSignupTokenFormModel**](JeunesseProfileCoreWalletModelsSignupTokensSignupTokenFormModel.md)|  | 
 **optional** | ***WalletApiWalletCreateSignupTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletCreateSignupTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsSignupTokensGetSignupTokensReportResult**](Jeunesse.Profile.Core.Wallet.Models.SignupTokens.GetSignupTokensReportResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletDeactivateBankAccount**
> JeunesseProfileCoreWalletModelsEftEftResponseModel WalletDeactivateBankAccount(ctx, mainId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainId** | **int32**|  | 
 **optional** | ***WalletApiWalletDeactivateBankAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletDeactivateBankAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsEftEftResponseModel**](Jeunesse.Profile.Core.Wallet.Models.EFT.EFTResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletDeleteFundRequest**
> bool WalletDeleteFundRequest(ctx, fundRequestId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundRequestId** | **int32**|  | 
 **optional** | ***WalletApiWalletDeleteFundRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletDeleteFundRequestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mainId** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletDeleteSignupToken**
> JeunesseProfileCoreWalletModelsSignupTokensGetSignupTokensReportResult WalletDeleteSignupToken(ctx, tokenId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **int32**|  | 
 **optional** | ***WalletApiWalletDeleteSignupTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletDeleteSignupTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mainId** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsSignupTokensGetSignupTokensReportResult**](Jeunesse.Profile.Core.Wallet.Models.SignupTokens.GetSignupTokensReportResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetAccountCurrencies**
> []JeunesseProfileCoreWalletModelsEftAccountCurrencyModel WalletGetAccountCurrencies(ctx, bankCountry, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankCountry** | **string**|  | 
 **optional** | ***WalletApiWalletGetAccountCurrenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetAccountCurrenciesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreWalletModelsEftAccountCurrencyModel**](Jeunesse.Profile.Core.Wallet.Models.EFT.AccountCurrencyModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetAccountDetails**
> JeunesseProfileCoreWalletModelsAccountDetailsModel WalletGetAccountDetails(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetAccountDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetAccountDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainId** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsAccountDetailsModel**](Jeunesse.Profile.Core.Wallet.Models.AccountDetailsModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetAccountHistory**
> []JeunesseProfileCoreWalletModelsFundRequestFundRequestModel WalletGetAccountHistory(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetAccountHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetAccountHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **queryIsSequential** | **optional.Bool**|  | 
 **queryUserName** | **optional.String**|  | 
 **queryPageSize** | **optional.Int32**|  | 
 **queryPageNumber** | **optional.Int32**|  | 
 **queryOrderBy** | **optional.String**|  | 

### Return type

[**[]JeunesseProfileCoreWalletModelsFundRequestFundRequestModel**](Jeunesse.Profile.Core.Wallet.Models.FundRequest.FundRequestModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetBankCountries**
> []JeunesseProfileCoreWalletModelsEftBankCountryModel WalletGetBankCountries(ctx, isLocalEFT, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **isLocalEFT** | **bool**|  | 
 **optional** | ***WalletApiWalletGetBankCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetBankCountriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mainFk** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreWalletModelsEftBankCountryModel**](Jeunesse.Profile.Core.Wallet.Models.EFT.BankCountryModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetEFTLocalStatus**
> JeunesseProfileCoreWalletModelsEftEftLocalStatusModel WalletGetEFTLocalStatus(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetEFTLocalStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetEFTLocalStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainFk** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsEftEftLocalStatusModel**](Jeunesse.Profile.Core.Wallet.Models.EFT.EFTLocalStatusModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetEFTRequiredDocuments**
> []JeunesseProfileCoreWalletModelsEftDocumentModel WalletGetEFTRequiredDocuments(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetEFTRequiredDocumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetEFTRequiredDocumentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainFk** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreWalletModelsEftDocumentModel**](Jeunesse.Profile.Core.Wallet.Models.EFT.DocumentModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetEWalletLoginForm**
> JeunesseProfileCoreWalletModelsMerchantIPayoutModel WalletGetEWalletLoginForm(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetEWalletLoginFormOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetEWalletLoginFormOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsMerchantIPayoutModel**](Jeunesse.Profile.Core.Wallet.Models.MerchantIPayoutModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetFundRequestCheckRuns**
> []JeunesseProfileCoreWalletModelsFundRequestInvoiceFundRequestCheckRunModel WalletGetFundRequestCheckRuns(ctx, mainId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainId** | **int32**|  | 
 **optional** | ***WalletApiWalletGetFundRequestCheckRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetFundRequestCheckRunsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreWalletModelsFundRequestInvoiceFundRequestCheckRunModel**](Jeunesse.Profile.Core.Wallet.Models.FundRequestInvoice.FundRequestCheckRunModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetFundRequestInvoiceDetails**
> JeunesseProfileCoreWalletModelsFundRequestInvoiceFundRequestInvoiceModel WalletGetFundRequestInvoiceDetails(ctx, mainId, withholdRecordId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainId** | **int32**|  | 
  **withholdRecordId** | **int32**|  | 
 **optional** | ***WalletApiWalletGetFundRequestInvoiceDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetFundRequestInvoiceDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsFundRequestInvoiceFundRequestInvoiceModel**](Jeunesse.Profile.Core.Wallet.Models.FundRequestInvoice.FundRequestInvoiceModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetFundsRequest**
> []JeunesseProfileCoreWalletModelsFundRequestFundRequestModel WalletGetFundsRequest(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetFundsRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetFundsRequestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **queryMainId** | **optional.Int32**|  | 
 **queryPageSize** | **optional.Int32**|  | 
 **queryPageNumber** | **optional.Int32**|  | 
 **queryOrderBy** | **optional.String**|  | 

### Return type

[**[]JeunesseProfileCoreWalletModelsFundRequestFundRequestModel**](Jeunesse.Profile.Core.Wallet.Models.FundRequest.FundRequestModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetFundsRequestSettings**
> JeunesseProfileCoreWalletModelsFundRequestFundRequestSettingsModel WalletGetFundsRequestSettings(ctx, userName, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userName** | **string**|  | 
 **optional** | ***WalletApiWalletGetFundsRequestSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetFundsRequestSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsFundRequestFundRequestSettingsModel**](Jeunesse.Profile.Core.Wallet.Models.FundRequest.FundRequestSettingsModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetIPayoutCardsBin**
> []string WalletGetIPayoutCardsBin(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetIPayoutCardsBinOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetIPayoutCardsBinOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetIdentificationTypes**
> []JeunesseProfileCoreWalletModelsIdentificationTypeModel WalletGetIdentificationTypes(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetIdentificationTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetIdentificationTypesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **optional.String**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreWalletModelsIdentificationTypeModel**](Jeunesse.Profile.Core.Wallet.Models.IdentificationTypeModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetPayCardReplaceState**
> JeunesseProfileCoreWalletModelsPayCardRegisterAllowed WalletGetPayCardReplaceState(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetPayCardReplaceStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetPayCardReplaceStateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsPayCardRegisterAllowed**](Jeunesse.Profile.Core.Wallet.Models.PayCardRegisterAllowed.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetPayCardState**
> JeunesseProfileCoreWalletModelsPayCardRegisterAllowed WalletGetPayCardState(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetPayCardStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetPayCardStateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsPayCardRegisterAllowed**](Jeunesse.Profile.Core.Wallet.Models.PayCardRegisterAllowed.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetPayCardValidation**
> JeunesseProfileCoreWalletModelsPayCardRegisterAllowed WalletGetPayCardValidation(ctx, model, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsPayCardModel**](JeunesseProfileCoreWalletModelsPayCardModel.md)|  | 
 **optional** | ***WalletApiWalletGetPayCardValidationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetPayCardValidationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsPayCardRegisterAllowed**](Jeunesse.Profile.Core.Wallet.Models.PayCardRegisterAllowed.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetPersonalInfo**
> JeunesseProfileCoreWalletModelsAccountDetailsModel WalletGetPersonalInfo(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetPersonalInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetPersonalInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainFk** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsAccountDetailsModel**](Jeunesse.Profile.Core.Wallet.Models.AccountDetailsModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetPersonalInfoLocal**
> JeunesseProfileCoreWalletModelsAccountDetailsModel WalletGetPersonalInfoLocal(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetPersonalInfoLocalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetPersonalInfoLocalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainFk** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsAccountDetailsModel**](Jeunesse.Profile.Core.Wallet.Models.AccountDetailsModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetPhotoIdData**
> JeunesseProfileCoreWalletModelsPhotoIdModel WalletGetPhotoIdData(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetPhotoIdDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetPhotoIdDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainId** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsPhotoIdModel**](Jeunesse.Profile.Core.Wallet.Models.PhotoIdModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetSignupTokens**
> JeunesseProfileCoreWalletModelsSignupTokensSignupTokenViewModel WalletGetSignupTokens(ctx, tokenStatus, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenStatus** | **string**|  | 
 **optional** | ***WalletApiWalletGetSignupTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetSignupTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **mainId** | **optional.Int32**|  | 
 **tokenName** | **optional.String**|  | 
 **tokenStartDate** | **optional.Time**|  | 
 **tokenEndDate** | **optional.Time**|  | 
 **useMainOrdersFK** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **pageNumber** | **optional.Int32**|  | 
 **orderBy** | **optional.String**|  | 

### Return type

[**JeunesseProfileCoreWalletModelsSignupTokensSignupTokenViewModel**](Jeunesse.Profile.Core.Wallet.Models.SignupTokens.SignupTokenViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetUserBankBook**
> JeunesseProfileCoreWalletModelsBankBookBankBookFormModel WalletGetUserBankBook(ctx, mainId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainId** | **int32**|  | 
 **optional** | ***WalletApiWalletGetUserBankBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetUserBankBookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsBankBookBankBookFormModel**](Jeunesse.Profile.Core.Wallet.Models.BankBook.BankBookFormModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetUserBankDocument**
> JeunesseProfileCoreWalletModelsBankDocumentBankDocumentModel WalletGetUserBankDocument(ctx, mainId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainId** | **int32**|  | 
 **optional** | ***WalletApiWalletGetUserBankDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetUserBankDocumentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsBankDocumentBankDocumentModel**](Jeunesse.Profile.Core.Wallet.Models.BankDocument.BankDocumentModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetUserPaymentMethods**
> []JeunesseProfileCoreWalletModelsPaymentMethodModel WalletGetUserPaymentMethods(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetUserPaymentMethodsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetUserPaymentMethodsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreWalletModelsPaymentMethodModel**](Jeunesse.Profile.Core.Wallet.Models.PaymentMethodModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetUserPayoutMethods**
> []JeunesseProfileCoreWalletModelsPayoutMethodModel WalletGetUserPayoutMethods(ctx, activeOnly, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activeOnly** | **bool**|  | 
 **optional** | ***WalletApiWalletGetUserPayoutMethodsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetUserPayoutMethodsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mainId** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**[]JeunesseProfileCoreWalletModelsPayoutMethodModel**](Jeunesse.Profile.Core.Wallet.Models.PayoutMethodModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetWalletBalance**
> []JeunesseProfileCoreWalletModelsFundRequestFundRequestModel WalletGetWalletBalance(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetWalletBalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetWalletBalanceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]
 **queryUserName** | **optional.String**|  | 

### Return type

[**[]JeunesseProfileCoreWalletModelsFundRequestFundRequestModel**](Jeunesse.Profile.Core.Wallet.Models.FundRequest.FundRequestModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletGetWalletSettings**
> JeunesseProfileCoreWalletModelsWalletHomeSettings WalletGetWalletSettings(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletGetWalletSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletGetWalletSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsWalletHomeSettings**](Jeunesse.Profile.Core.Wallet.Models.WalletHomeSettings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletRegisterPaycard**
> JeunesseProfileCoreWalletModelsRegisterPayardViewModel WalletRegisterPaycard(ctx, model, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsRegisterPayardFormModel**](JeunesseProfileCoreWalletModelsRegisterPayardFormModel.md)|  | 
 **optional** | ***WalletApiWalletRegisterPaycardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletRegisterPaycardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsRegisterPayardViewModel**](Jeunesse.Profile.Core.Wallet.Models.RegisterPayСardViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletReplacePayCard**
> JeunesseProfileCoreWalletModelsPayCardRegisterAllowed WalletReplacePayCard(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletApiWalletReplacePayCardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletReplacePayCardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsPayCardRegisterAllowed**](Jeunesse.Profile.Core.Wallet.Models.PayCardRegisterAllowed.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletResetWalletPassword**
> string WalletResetWalletPassword(ctx, mainId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mainId** | **int32**|  | 
 **optional** | ***WalletApiWalletResetWalletPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletResetWalletPasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletSaveEFTBankInfo**
> bool WalletSaveEFTBankInfo(ctx, model, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsEftEftFormModel**](JeunesseProfileCoreWalletModelsEftEftFormModel.md)|  | 
 **optional** | ***WalletApiWalletSaveEFTBankInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletSaveEFTBankInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mainFk** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletSaveEFTLocalBankInfo**
> JeunesseProfileCoreWalletModelsEftEftResponseModel WalletSaveEFTLocalBankInfo(ctx, model, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsEftEftLocalFormModel**](JeunesseProfileCoreWalletModelsEftEftLocalFormModel.md)|  | 
 **optional** | ***WalletApiWalletSaveEFTLocalBankInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletSaveEFTLocalBankInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mainFk** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsEftEftResponseModel**](Jeunesse.Profile.Core.Wallet.Models.EFT.EFTResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletUpdatePhotoIdData**
> JeunesseProfileCoreWalletModelsPhotoIdModel WalletUpdatePhotoIdData(ctx, model, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | [**JeunesseProfileCoreWalletModelsPhotoIdUpdatePhotoIdFormModel**](JeunesseProfileCoreWalletModelsPhotoIdUpdatePhotoIdFormModel.md)|  | 
 **optional** | ***WalletApiWalletUpdatePhotoIdDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletUpdatePhotoIdDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

[**JeunesseProfileCoreWalletModelsPhotoIdModel**](Jeunesse.Profile.Core.Wallet.Models.PhotoIdModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletUpdateUserBankBook**
> string WalletUpdateUserBankBook(ctx, bankBook, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankBook** | [**JeunesseProfileCoreWalletModelsBankBookBankBookFormModel**](JeunesseProfileCoreWalletModelsBankBookBankBookFormModel.md)|  | 
 **optional** | ***WalletApiWalletUpdateUserBankBookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletUpdateUserBankBookOpts struct

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

# **WalletUpdateUserBankDocument**
> string WalletUpdateUserBankDocument(ctx, bankDocument, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankDocument** | [**JeunesseProfileCoreWalletModelsBankDocumentBankDocumentFormModel**](JeunesseProfileCoreWalletModelsBankDocumentBankDocumentFormModel.md)|  | 
 **optional** | ***WalletApiWalletUpdateUserBankDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletUpdateUserBankDocumentOpts struct

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

# **WalletUpdateUserFundRequestInvoice**
> string WalletUpdateUserFundRequestInvoice(ctx, fundRequestInvoice, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fundRequestInvoice** | [**JeunesseProfileCoreWalletModelsFundRequestInvoiceFundRequestInvoiceFormModel**](JeunesseProfileCoreWalletModelsFundRequestInvoiceFundRequestInvoiceFormModel.md)|  | 
 **optional** | ***WalletApiWalletUpdateUserFundRequestInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletUpdateUserFundRequestInvoiceOpts struct

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

# **WalletUpdateUserPayoutMethods**
> interface{} WalletUpdateUserPayoutMethods(ctx, methods, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **methods** | [**[]JeunesseProfileCoreWalletModelsPayoutMethodModel**](Jeunesse.Profile.Core.Wallet.Models.PayoutMethodModel.md)|  | 
 **optional** | ***WalletApiWalletUpdateUserPayoutMethodsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletApiWalletUpdateUserPayoutMethodsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mainId** | **optional.Int32**|  | 
 **apiVersion** | **optional.String**| The requested API version | [default to 1.0]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/x-www-form-urlencoded
 - **Accept**: application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


package main

import (
	"net/http"
)

func (app *application) getWallet_GetWalletSettings(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getWallet_GetPayCardState(w http.ResponseWriter, r *http.Request)         {}
func (app *application) getWallet_GetPayCardReplaceState(w http.ResponseWriter, r *http.Request)  {}
func (app *application) postWallet_ReplacePayCard(w http.ResponseWriter, r *http.Request)         {}
func (app *application) getWallet_GetUserPaymentMethods(w http.ResponseWriter, r *http.Request)   {}
func (app *application) getWallet_GetIPayoutCardsBin(w http.ResponseWriter, r *http.Request)      {}
func (app *application) getWallet_GetAccountDetails(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getWallet_GetPhotoIdData(w http.ResponseWriter, r *http.Request)          {}
func (app *application) postWallet_UpdatePhotoIdData(w http.ResponseWriter, r *http.Request)      {}
func (app *application) getWallet_GetUserBankBook(w http.ResponseWriter, r *http.Request)         {}
func (app *application) postWallet_UpdateUserBankBook(w http.ResponseWriter, r *http.Request)     {}
func (app *application) getWallet_GetUserBankDocument(w http.ResponseWriter, r *http.Request)     {}
func (app *application) postWallet_UpdateUserBankDocument(w http.ResponseWriter, r *http.Request) {}
func (app *application) getWallet_GetEFTLocalStatus(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getWallet_GetPersonalInfo(w http.ResponseWriter, r *http.Request)         {}
func (app *application) getWallet_GetAccountCurrencies(w http.ResponseWriter, r *http.Request)    {}
func (app *application) getWallet_GetFundRequestCheckRuns(w http.ResponseWriter, r *http.Request) {}
func (app *application) getWallet_GetPersonalInfoLocal(w http.ResponseWriter, r *http.Request)    {}
func (app *application) getWallet_GetFundsRequestSettings(w http.ResponseWriter, r *http.Request) {}
func (app *application) getWallet_GetUserPayoutMethods(w http.ResponseWriter, r *http.Request)    {}
func (app *application) getWallet_GetEFTRequiredDocuments(w http.ResponseWriter, r *http.Request) {}
func (app *application) getWallet_GetIdentificationTypes(w http.ResponseWriter, r *http.Request)  {}
func (app *application) getWallet_GetBankCountries(w http.ResponseWriter, r *http.Request)        {}
func (app *application) getWallet_GetWalletBalance(w http.ResponseWriter, r *http.Request)        {}
func (app *application) getWallet_GetFundRequestInvoiceDetails(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getWallet_GetFundsRequest(w http.ResponseWriter, r *http.Request)         {}
func (app *application) postWallet_CreateFundRequest(w http.ResponseWriter, r *http.Request)      {}
func (app *application) getWallet_GetAccountHistory(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getWallet_GetSignupTokens(w http.ResponseWriter, r *http.Request)         {}
func (app *application) postWallet_CreateSignupToken(w http.ResponseWriter, r *http.Request)      {}
func (app *application) postWallet_RegisterPaycard(w http.ResponseWriter, r *http.Request)        {}
func (app *application) postWallet_AddFundsGetRedirectUrl(w http.ResponseWriter, r *http.Request) {}
func (app *application) postWallet_AdminUpdatePhotoIdData(w http.ResponseWriter, r *http.Request) {}
func (app *application) postWallet_GetPayCardValidation(w http.ResponseWriter, r *http.Request)   {}
func (app *application) postWallet_GetEWalletLoginForm(w http.ResponseWriter, r *http.Request)    {}
func (app *application) postWallet_UpdateUserFundRequestInvoice(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postWallet_SaveEFTBankInfo(w http.ResponseWriter, r *http.Request)         {}
func (app *application) postWallet_DeactivateBankAccount(w http.ResponseWriter, r *http.Request)   {}
func (app *application) postWallet_UpdateUserPayoutMethods(w http.ResponseWriter, r *http.Request) {}
func (app *application) postWallet_ResetWalletPassword(w http.ResponseWriter, r *http.Request)     {}
func (app *application) postWallet_SaveEFTLocalBankInfo(w http.ResponseWriter, r *http.Request)    {}
func (app *application) deleteWallet_DeleteSignupToken(w http.ResponseWriter, r *http.Request)     {}
func (app *application) deleteWallet_DeleteFundRequest(w http.ResponseWriter, r *http.Request)     {}
func (app *application) getWalletMember_GetSecurityQuestions(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getWalletMember_GetQuestionsList(w http.ResponseWriter, r *http.Request) {}
func (app *application) postWalletMember_SetQuestions(w http.ResponseWriter, r *http.Request)    {}
func (app *application) getWalletMember_CheckAttempts(w http.ResponseWriter, r *http.Request)    {}
func (app *application) postWalletMember_CreateWalletAccount(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postWalletMember_GetResetCode(w http.ResponseWriter, r *http.Request)   {}
func (app *application) postWalletMember_ResetPassword(w http.ResponseWriter, r *http.Request)  {}
func (app *application) postWalletMember_UpdatePassword(w http.ResponseWriter, r *http.Request) {}

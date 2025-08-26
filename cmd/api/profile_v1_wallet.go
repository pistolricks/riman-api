package main

/*
func (app *application) getGetWalletSettings(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetPayCardState(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetPayCardReplaceState(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postReplacePayCard(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetUserPaymentMethods(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetIPayoutCardsBin(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetAccountDetails(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetPhotoIdData(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postUpdatePhotoIdData(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetUserBankBook(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postUpdateUserBankBook(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetUserBankDocument(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postUpdateUserBankDocument(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetEFTLocalStatus(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetPersonalInfo(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetAccountCurrencies(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetFundRequestCheckRuns(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetPersonalInfoLocal(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetFundsRequestSettings(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetUserPayoutMethods(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetEFTRequiredDocuments(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetIdentificationTypes(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetBankCountries(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetWalletBalance(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetFundRequestInvoiceDetails(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getGetFundsRequest(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postCreateFundRequest(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetAccountHistory(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getGetSignupTokens(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postCreateSignupToken(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postRegisterPaycard(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postAddFundsGetRedirectUrl(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postAdminUpdatePhotoIdData(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postGetPayCardValidation(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postGetEWalletLoginForm(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postUpdateUserFundRequestInvoice(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postSaveEFTBankInfo(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postDeactivateBankAccount(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postUpdateUserPayoutMethods(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postResetWalletPassword(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) postSaveEFTLocalBankInfo(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) deleteDeleteSignupToken(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) deleteDeleteFundRequest(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletApi
}
func (app *application) getWalletMemberSecurityQuestions(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletMemberApi
}
func (app *application) getWalletMemberQuestionsList(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletMemberApi
}
func (app *application) postWalletMemberSetQuestions(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletMemberApi
}
func (app *application) getWalletMemberCheckAttempts(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletMemberApi
}
func (app *application) postWalletMemberCreateWalletAccount(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletMemberApi
}
func (app *application) postWalletMemberResetCode(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletMemberApi
}
func (app *application) postWalletMemberResetPassword(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletMemberApi
}
func (app *application) postWalletMemberUpdatePassword(w http.ResponseWriter, r *http.Request) {
	app.profileV1.WalletMemberApi
}
*/

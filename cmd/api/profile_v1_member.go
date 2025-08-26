package main

import (
	"net/http"
)

func RepSite_GetRepsiteInfo()              {}
func RepSite_InvalidateRepSiteInfo()       {}
func RepSite_GetRepSiteMainPk()            {}
func RepSite_GetAdWords()                  {}
func RepSite_GetImpressumInfo()            {}
func RepSite_GetDistributorDiscountCodes() {}
func RepSite_SendReferrerEmail()           {}
func RepSite_AddLead()                     {}
func RepSite_ValidateSiteUrl()             {}

func (app *application) postMember_CreatePaymentMethod(w http.ResponseWriter, r *http.Request)     {}
func (app *application) putMember_UpdatePaymentDetails(w http.ResponseWriter, r *http.Request)     {}
func (app *application) getMember_GetHighestRealRank(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getMember_GetMemberInfo(w http.ResponseWriter, r *http.Request)            {}
func (app *application) getMember_UserAlertCheck(w http.ResponseWriter, r *http.Request)           {}
func (app *application) getMember_GetUserUnreadCounts(w http.ResponseWriter, r *http.Request)      {}
func (app *application) getMember_RequiresCvv(w http.ResponseWriter, r *http.Request)              {}
func (app *application) getMember_GetBillingAddress(w http.ResponseWriter, r *http.Request)        {}
func (app *application) putMember_UpdateBillingAddress(w http.ResponseWriter, r *http.Request)     {}
func (app *application) getMember_PayPalAgreementAvailable(w http.ResponseWriter, r *http.Request) {}
func (app *application) getMember_SendImsEmailToUser(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getMember_IsIxoPayActiveForGateway(w http.ResponseWriter, r *http.Request) {}
func (app *application) getMember_GetTWUnifiedDonationOrgOptions(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getMember_GetReferrer(w http.ResponseWriter, r *http.Request)           {}
func (app *application) putMember_SetReferrer(w http.ResponseWriter, r *http.Request)           {}
func (app *application) getMember_GetUpgradeLink(w http.ResponseWriter, r *http.Request)        {}
func (app *application) getMember_GetShippingAddress(w http.ResponseWriter, r *http.Request)    {}
func (app *application) putMember_UpdateShippingAddress(w http.ResponseWriter, r *http.Request) {}
func (app *application) getMember_GetTWUnifiedInfo(w http.ResponseWriter, r *http.Request)      {}
func (app *application) postMember_InsertUpdateTWUnifiedInfo(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getMember_GetAddressLookUp(w http.ResponseWriter, r *http.Request)          {}
func (app *application) getMember_VerifyCpf(w http.ResponseWriter, r *http.Request)                 {}
func (app *application) getMember_BonusCreditsGoingToExpire(w http.ResponseWriter, r *http.Request) {}
func (app *application) getMember_SendSmsMessage(w http.ResponseWriter, r *http.Request)            {}
func (app *application) getMember_Upgrade(w http.ResponseWriter, r *http.Request)                   {}
func (app *application) putMember_UpdatePassword(w http.ResponseWriter, r *http.Request)            {}
func (app *application) putMember_UpdatePaymentOrder(w http.ResponseWriter, r *http.Request)        {}
func (app *application) putMember_UpdateShowWelcome(w http.ResponseWriter, r *http.Request)         {}
func (app *application) putMember_RemovePaymentMethod(w http.ResponseWriter, r *http.Request)       {}
func (app *application) putMember_UpdateIsShippingAddress(w http.ResponseWriter, r *http.Request)   {}
func (app *application) postMember_CreateMember(w http.ResponseWriter, r *http.Request)             {}
func (app *application) postMember_UpdateTOSStatus(w http.ResponseWriter, r *http.Request)          {}
func (app *application) postMember_ConvertAnonymousMemberToCustomer(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postMember_SetMemberPriceList(w http.ResponseWriter, r *http.Request) {}
func (app *application) postMember_CreateTicket(w http.ResponseWriter, r *http.Request)       {}
func (app *application) postMember_ValidateSiteUrl(w http.ResponseWriter, r *http.Request)    {}
func (app *application) postMember_CreateMemberFromEventGuest(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postMember_SetEmailAndPassword(w http.ResponseWriter, r *http.Request) {}

package main

import (
	"net/http"
)

func (app *application) getMainUsers_GetUser(w http.ResponseWriter, r *http.Request)                {}
func (app *application) getMainUsers_GetAwardsTypes(w http.ResponseWriter, r *http.Request)         {}
func (app *application) getMainUsers_GetUserAwards(w http.ResponseWriter, r *http.Request)          {}
func (app *application) getMainUsers_GetUserBySiteUrl(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getMainUsers_CanUserChangeSponsor(w http.ResponseWriter, r *http.Request)   {}
func (app *application) getMainUsers_GetUserMarketingInfo(w http.ResponseWriter, r *http.Request)   {}
func (app *application) getMainUsers_GetSignupUser(w http.ResponseWriter, r *http.Request)          {}
func (app *application) postMainUsers_CreateSignupUser(w http.ResponseWriter, r *http.Request)      {}
func (app *application) putMainUsers_UpdateSignupUser(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getMainUsers_GetUserRecognitionInfo(w http.ResponseWriter, r *http.Request) {}
func (app *application) getMainUsers_AllowUserChangeUsername(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getMainUsers_GetESignatureInfo(w http.ResponseWriter, r *http.Request)  {}
func (app *application) postMainUsers_SetESignatureInfo(w http.ResponseWriter, r *http.Request) {}
func (app *application) getMainUsers_IsNewSponsorValid(w http.ResponseWriter, r *http.Request)  {}
func (app *application) getMainUsers_CheckUserDownline(w http.ResponseWriter, r *http.Request)  {}
func (app *application) getMainUsers_GetRenewalStatusInfoAsync(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getMainUsers_GetUserHighestRank(w http.ResponseWriter, r *http.Request)    {}
func (app *application) getMainUsers_CheckUserEnrollerTree(w http.ResponseWriter, r *http.Request) {}
func (app *application) getMainUsers_SyncCustomers(w http.ResponseWriter, r *http.Request)         {}
func (app *application) getMainUsers_GetReferrerInfo(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getMainUsers_GetMainUserSponsorPlacementView(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getMainUsers_GetRedirectRule(w http.ResponseWriter, r *http.Request)     {}
func (app *application) getMainUsers_GetUsers(w http.ResponseWriter, r *http.Request)            {}
func (app *application) putMainUsers_UpdateProfile(w http.ResponseWriter, r *http.Request)       {}
func (app *application) putMainUsers_UpdateMarketingInfo(w http.ResponseWriter, r *http.Request) {}
func (app *application) putMainUsers_UpdateUserRecognitionInfo(w http.ResponseWriter, r *http.Request) {
}
func (app *application) putMainUsers_UpdateSoD(w http.ResponseWriter, r *http.Request)          {}
func (app *application) postMainUsers_AddUserAward(w http.ResponseWriter, r *http.Request)      {}
func (app *application) postMainUsers_ResetPassword(w http.ResponseWriter, r *http.Request)     {}
func (app *application) postMainUsers_ValidateEmail(w http.ResponseWriter, r *http.Request)     {}
func (app *application) postMainUsers_RemindUsername(w http.ResponseWriter, r *http.Request)    {}
func (app *application) postMainUsers_UpdatePassword(w http.ResponseWriter, r *http.Request)    {}
func (app *application) postMainUsers_ChangeUserSponsor(w http.ResponseWriter, r *http.Request) {}
func (app *application) postMainUsers_ChangeUsername(w http.ResponseWriter, r *http.Request)    {}
func (app *application) getMember_GetPaymentDetails(w http.ResponseWriter, r *http.Request)     {}

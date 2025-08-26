package main

import (
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	"github.com/pistolricks/riman-api/internal/profile_v1"
)

func (app *application) getGetUser(w http.ResponseWriter, r *http.Request) {
	mainPkStr := r.URL.Query().Get("mainPK")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetUser(r.Context(), int32(mainPk64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"user": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetAwardsTypes(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetAwardsTypes(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"awardTypes": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUserAwards(w http.ResponseWriter, r *http.Request) {
	mainPkStr := r.URL.Query().Get("mainPK")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetUserAwards(r.Context(), int32(mainPk64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"awards": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUserBySiteUrl(w http.ResponseWriter, r *http.Request) {
	siteUrl := r.URL.Query().Get("siteUrl")
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetUserBySiteUrl(r.Context(), siteUrl, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"user": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getCanUserChangeSponsor(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MainUsersApi.MainUsersCanUserChangeSponsor(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"canChangeSponsor": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUserMarketingInfo(w http.ResponseWriter, r *http.Request) {
	mainPkStr := r.URL.Query().Get("mainPK")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetUserMarketingInfo(r.Context(), int32(mainPk64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"marketing": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetSignupUser(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetSignupUser(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"signupUser": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postCreateSignupUser(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreUsersModelsSignupUserFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersCreateSignupUser(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"signupUser": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateSignupUser(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreUsersModelsSignupUserFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersUpdateSignupUser(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"signupUser": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUserRecognitionInfo(w http.ResponseWriter, r *http.Request) {
	mainPkStr := r.URL.Query().Get("mainPK")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetUserRecognitionInfo(r.Context(), int32(mainPk64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"recognition": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getAllowUserChangeUsername(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MainUsersApi.MainUsersAllowUserChangeUsername(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"allowed": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetESignatureInfo(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetESignatureInfo(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"eSignature": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postSetESignatureInfo(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreUsersModelsESignatureInfoPostModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersSetESignatureInfo(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"eSignature": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getIsNewSponsorValid(w http.ResponseWriter, r *http.Request) {
	sponsorSiteUrl := r.URL.Query().Get("sponsorSiteUrl")
	res, h, err := app.profileV1.MainUsersApi.MainUsersIsNewSponsorValid(r.Context(), sponsorSiteUrl, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"valid": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getCheckUserDownline(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainPkStr := q.Get("mainPk")
	downlineMainPkStr := q.Get("downlineMainPk")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	downlineMainPk64, _ := strconv.ParseInt(downlineMainPkStr, 10, 32)
	res, h, err := app.profileV1.MainUsersApi.MainUsersCheckUserDownline(r.Context(), int32(mainPk64), int32(downlineMainPk64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"downline": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetRenewalStatusInfoAsync(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetRenewalStatusInfoAsync(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"renewalStatus": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUserHighestRank(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetUserHighestRank(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"highestRank": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getCheckUserEnrollerTree(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainPkStr := q.Get("mainPk")
	enrollerTreeMainPkStr := q.Get("enrollerTreeMainPk")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	enrollerTreeMainPk64, _ := strconv.ParseInt(enrollerTreeMainPkStr, 10, 32)
	res, h, err := app.profileV1.MainUsersApi.MainUsersCheckUserEnrollerTree(r.Context(), int32(mainPk64), int32(enrollerTreeMainPk64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"enrollerTree": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getSyncCustomers(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MainUsersApi.MainUsersSyncCustomers(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"synced": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetReferrerInfo(w http.ResponseWriter, r *http.Request) {
	siteUrl := r.URL.Query().Get("siteUrl")
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetReferrerInfo(r.Context(), siteUrl, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"referrer": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetMainUserSponsorPlacementView(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetMainUserSponsorPlacementView(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"sponsorPlacement": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetRedirectRule(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &profile_v1.MainUsersApiMainUsersGetRedirectRuleOpts{}
	if v := q.Get("mainPK"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.MainPK = optional.NewInt32(int32(n)) } }
	if v := q.Get("country"); v != "" { opts.Country = optional.NewString(v) }
	if v := q.Get("sourceURL"); v != "" { opts.SourceURL = optional.NewString(v) }
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetRedirectRule(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"redirectRule": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUsers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &profile_v1.MainUsersApiMainUsersGetUsersOpts{}
	if v := q.Get("search"); v != "" { opts.Search = optional.NewString(v) }
	if v := q.Get("exactSearch"); v != "" { opts.ExactSearch = optional.NewBool(v == "true" || v == "1") }
	if v := q.Get("siteUrl"); v != "" { opts.SiteUrl = optional.NewString(v) }
	if v := q.Get("mainPK"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.MainPK = optional.NewInt32(int32(n)) } }
	if v := q.Get("email"); v != "" { opts.Email = optional.NewString(v) }
	if v := q.Get("firstName"); v != "" { opts.FirstName = optional.NewString(v) }
	if v := q.Get("lastName"); v != "" { opts.LastName = optional.NewString(v) }
	if v := q.Get("company"); v != "" { opts.Company = optional.NewString(v) }
	if v := q.Get("phone"); v != "" { opts.Phone = optional.NewString(v) }
	if v := q.Get("sponsorSiteUrl"); v != "" { opts.SponsorSiteUrl = optional.NewString(v) }
	if v := q.Get("orderNumber"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.OrderNumber = optional.NewInt32(int32(n)) } }
	if v := q.Get("skipCount"); v != "" { opts.SkipCount = optional.NewBool(v == "true" || v == "1") }
	if v := q.Get("mainType"); v != "" { opts.MainType = optional.NewString(v) }
	if v := q.Get("orderBy"); v != "" { opts.OrderBy = optional.NewString(v) }
	if v := q.Get("offset"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.Offset = optional.NewInt32(int32(n)) } }
	if v := q.Get("limit"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.Limit = optional.NewInt32(int32(n)) } }
	res, h, err := app.profileV1.MainUsersApi.MainUsersGetUsers(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"users": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateProfile(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsMemberFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersUpdateProfile(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"profile": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateMarketingInfo(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreUsersModelsUserMarketingFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersUpdateMarketingInfo(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"marketing": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateUserRecognitionInfo(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreUsersModelsUserRecognitionFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersUpdateUserRecognitionInfo(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"recognition": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateSoD(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainPkStr := q.Get("mainPk")
	newStatusStr := q.Get("newStatus")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	newStatus := newStatusStr == "true" || newStatusStr == "1"
	res, h, err := app.profileV1.MainUsersApi.MainUsersUpdateSoD(r.Context(), int32(mainPk64), newStatus, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"updated": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postAddUserAward(w http.ResponseWriter, r *http.Request) {
	var award profile_v1.JeunesseBackDataPocoAwardsAward
	if err := app.readJSON(w, r, &award); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersAddUserAward(r.Context(), award, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"awards": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postResetPassword(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreUsersModelsResetPasswordModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersResetPassword(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postValidateEmail(w http.ResponseWriter, r *http.Request) {
	var request profile_v1.JeunesseProfileCoreUsersModelsEmailValidationQueryModel
	if err := app.readJSON(w, r, &request); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersValidateEmail(r.Context(), request, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"valid": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postRemindUsername(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreUsersModelsRemindUsernameModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersRemindUsername(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"username": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postUpdatePassword(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreUsersModelsUpdatePasswordModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersUpdatePassword(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"updated": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postChangeUserSponsor(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreUsersModelsUserSponsorChangeModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MainUsersApi.MainUsersChangeUserSponsor(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"changed": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postChangeUsername(w http.ResponseWriter, r *http.Request) {
	newUsername := r.URL.Query().Get("newUsername")
	res, h, err := app.profileV1.MainUsersApi.MainUsersChangeUsername(r.Context(), newUsername, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"changed": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetPaymentDetails(w http.ResponseWriter, r *http.Request) {
	// This endpoint maps to Member API
	res, h, err := app.profileV1.MemberApi.MemberGetPaymentDetails(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"paymentDetails": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}

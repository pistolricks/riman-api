package main

import (
	"net/http"
	"strconv"

	"github.com/pistolricks/riman-api/internal/profile_v1"
)

func (app *application) postCreatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsPaymentFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberCreatePaymentMethod(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"payment": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdatePaymentDetails(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsPaymentFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberUpdatePaymentDetails(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"payment": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetHighestRealRank(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberGetHighestRealRank(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"rank": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetMemberInfo(w http.ResponseWriter, r *http.Request) {
	mainPkStr := r.URL.Query().Get("mainpk")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	res, h, err := app.profileV1.MemberApi.MemberGetMemberInfo(r.Context(), int32(mainPk64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"member": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getUserAlertCheck(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberUserAlertCheck(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"alerts": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUserUnreadCounts(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberGetUserUnreadCounts(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"unreadCounts": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getRequiresCvv(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberRequiresCvv(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"requiresCvv": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetBillingAddress(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberGetBillingAddress(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"billingAddress": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateBillingAddress(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberAddressModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberUpdateBillingAddress(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"updated": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getPayPalAgreementAvailable(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberPayPalAgreementAvailable(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"available": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getSendImsEmailToUser(w http.ResponseWriter, r *http.Request) {
	siteurl := r.URL.Query().Get("siteurl")
	res, h, err := app.profileV1.MemberApi.MemberSendImsEmailToUser(r.Context(), siteurl, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getIsIxoPayActiveForGateway(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberIsIxoPayActiveForGateway(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"active": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetTWUnifiedDonationOrgOptions(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberGetTWUnifiedDonationOrgOptions(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"orgOptions": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetReferrer(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberGetReferrer(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"referrer": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putSetReferrer(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsSetReferrerFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberSetReferrer(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUpgradeLink(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	mainPkStr := q.Get("mainPk")
	code2 := q.Get("code2")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	res, h, err := app.profileV1.MemberApi.MemberGetUpgradeLink(r.Context(), int32(mainPk64), code2, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"upgradeLink": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetShippingAddress(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberGetShippingAddress(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"shippingAddress": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateShippingAddress(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberAddressModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberUpdateShippingAddress(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"updated": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetTWUnifiedInfo(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberGetTWUnifiedInfo(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"twInfo": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postInsertUpdateTWUnifiedInfo(w http.ResponseWriter, r *http.Request) {
	var request profile_v1.JeunesseProfileCoreMemberModelsTwUnifiedTwUnifiedInfo
	if err := app.readJSON(w, r, &request); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberInsertUpdateTWUnifiedInfo(r.Context(), request, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"twInfo": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetAddressLookUp(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	countryCode := q.Get("countryCode")
	zipCode := q.Get("zipCode")
	res, h, err := app.profileV1.MemberApi.MemberGetAddressLookUp(r.Context(), countryCode, zipCode, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"addressLookup": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getVerifyCpf(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	cpf := q.Get("cpf")
	mainPkStr := q.Get("mainPk")
	mainPk64, _ := strconv.ParseInt(mainPkStr, 10, 32)
	res, h, err := app.profileV1.MemberApi.MemberVerifyCpf(r.Context(), cpf, int32(mainPk64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"verifyCpf": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getBonusCreditsGoingToExpire(w http.ResponseWriter, r *http.Request) {
	daysStr := r.URL.Query().Get("daysCount")
	days64, _ := strconv.ParseInt(daysStr, 10, 32)
	res, h, err := app.profileV1.MemberApi.MemberBonusCreditsGoingToExpire(r.Context(), int32(days64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"bonusCredits": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getSendSmsMessage(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	phone := q.Get("phoneNumber")
	message := q.Get("message")
	res, h, err := app.profileV1.MemberApi.MemberSendSmsMessage(r.Context(), phone, message, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getUpgrade(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	state := q.Get("state")
	mainTypeStr := q.Get("mainType")
	mainFkStr := q.Get("mainFk")
	agreeTerms := q.Get("agreeTerms") == "true" || q.Get("agreeTerms") == "1"
	mainType64, _ := strconv.ParseInt(mainTypeStr, 10, 32)
	mainFk64, _ := strconv.ParseInt(mainFkStr, 10, 32)
	res, h, err := app.profileV1.MemberApi.MemberUpgrade(r.Context(), state, int32(mainType64), int32(mainFk64), agreeTerms, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"upgraded": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdatePassword(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsPasswordFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	resp, err := app.profileV1.MemberApi.MemberUpdatePassword(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, resp.StatusCode, envelope{"result": "ok"}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdatePaymentOrder(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsPaymentOrderModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberUpdatePaymentOrder(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"updated": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateShowWelcome(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberUpdateShowWelcome(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putRemovePaymentMethod(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	sequencePkStr := q.Get("sequencePk")
	useOrderStr := q.Get("useOrder")
	sequencePk64, _ := strconv.ParseInt(sequencePkStr, 10, 32)
	useOrder64, _ := strconv.ParseInt(useOrderStr, 10, 32)
	res, h, err := app.profileV1.MemberApi.MemberRemovePaymentMethod(r.Context(), int32(sequencePk64), int32(useOrder64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"removed": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateIsShippingAddress(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsUpdateIsShippingAddressModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberUpdateIsShippingAddress(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"updated": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postCreateMember(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberMemberModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberCreateMember(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"memberId": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postUpdateTOSStatus(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.MemberApi.MemberUpdateTOSStatus(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postConvertAnonymousMemberToCustomer(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsConvertAnonymousMemberFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberConvertAnonymousMemberToCustomer(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postSetMemberPriceList(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsSetMemberPriceListFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberSetMemberPriceList(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postCreateTicket(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsComplianceFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberCreateTicket(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postValidateSiteUrl(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberGovIdModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberValidateSiteUrl(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postCreateMemberFromEventGuest(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileApiViewModelsAddMemberFromEventGuestFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberCreateMemberFromEventGuest(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"memberFromEventGuest": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postSetEmailAndPassword(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMemberModelsUpgradeRetailFormModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MemberApi.MemberSetEmailAndPassword(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"result": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}

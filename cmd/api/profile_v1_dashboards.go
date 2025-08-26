package main

import (
	"net/http"
	"strconv"
)

func (app *application) getTeamDetails(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetTeamDetails(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"teamDetails": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getDashboardPromotionId(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetDashboardPromotionId(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"promotionId": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getCommissions(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetCommisions(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"commissions": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getPersonalInfo(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetPersonalInfo(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"personalInfo": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getTopProducers(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetTopProducers(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"topProducers": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getBpSpSnapshot(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetBpSpSnapshot(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"bpSpSnapshot": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getPersonalSnapshot(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetPersonalSnapshot(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"personalSnapshot": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getCheckIfShouldShowAgree(w http.ResponseWriter, r *http.Request) {
	siteUrl := r.URL.Query().Get("siteUrl")
	res, h, err := app.profileV1.DashboardsApi.DashboardsCheckIfShouldShowAgree(r.Context(), siteUrl, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"shouldShow": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getFlushPointsDaysLeft(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetFlushPointsDaysLeft(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"daysLeft": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getEnrolleeSalesDetails(w http.ResponseWriter, r *http.Request) {
	siteUrl := r.URL.Query().Get("siteUrl")
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetEnrolleeSalesDetails(r.Context(), siteUrl, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"enrolleeSales": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getCustCommTier(w http.ResponseWriter, r *http.Request) {
	siteUrl := r.URL.Query().Get("siteUrl")
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetCustCommTier(r.Context(), siteUrl, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"custCommTier": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getCustomers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	siteUrl := q.Get("siteUrl")
	yearStr := q.Get("year")
	year64, _ := strconv.ParseInt(yearStr, 10, 32)
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetCustomers(r.Context(), siteUrl, int32(year64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"customers": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getDashbortdPromoWinnersModel(w http.ResponseWriter, r *http.Request) {
	promotionIdStr := r.URL.Query().Get("promotionId")
	promotionId64, _ := strconv.ParseInt(promotionIdStr, 10, 32)
	res, h, err := app.profileV1.DashboardsApi.DashboardsDashbortdPromoWinnersModel(r.Context(), int32(promotionId64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"winners": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getDashboardPromoUserData(w http.ResponseWriter, r *http.Request) {
	promotionIdStr := r.URL.Query().Get("promotionId")
	promotionId64, _ := strconv.ParseInt(promotionIdStr, 10, 32)
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetDashboardPromoUserData(r.Context(), int32(promotionId64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"promoUserData": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getVolRSBSales(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	siteUrl := q.Get("siteUrl")
	yearStr := q.Get("year")
	year64, _ := strconv.ParseInt(yearStr, 10, 32)
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetVolRSBSales(r.Context(), siteUrl, int32(year64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"volRSBSales": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getDisShippingApitMassiveActionBonusPromoData(w http.ResponseWriter, r *http.Request) {
	promotionIdStr := r.URL.Query().Get("promotionId")
	promotionId64, _ := strconv.ParseInt(promotionIdStr, 10, 32)
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetDistMassiveActionBonusPromoData(r.Context(), int32(promotionId64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"massiveBonusPromo": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getVolRSBCommissions(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	siteUrl := q.Get("siteUrl")
	yearStr := q.Get("year")
	year64, _ := strconv.ParseInt(yearStr, 10, 32)
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetVolRSBCommisions(r.Context(), siteUrl, int32(year64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"volRSBCommissions": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getNewContactCounts(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetNewContactCounts(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"newContactCounts": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getNewReps(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetNewReps(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"newReps": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getCommissionsTotal(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetCommissionsTotal(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"commissionsTotal": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getTopSalesContact(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetTopSalesContact(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"topSalesContact": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getMonthlySalesData(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetMonthlySalesData(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"monthlySalesData": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getMonthlyProductData(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetMonthlyProductData(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"monthlyProductData": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getEventTicketsSnapshots(w http.ResponseWriter, r *http.Request) {
	eventIdStr := r.URL.Query().Get("eventId")
	eventId64, _ := strconv.ParseInt(eventIdStr, 10, 32)
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetEventTicketsSnapshots(r.Context(), int32(eventId64), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"eventTicketsSnapshots": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getSmartNewsFeed(w http.ResponseWriter, r *http.Request) {
	res, h, err := app.profileV1.DashboardsApi.DashboardsGetSmartNewsFeed(r.Context(), nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"smartNewsFeed": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}

package main

import (
	"net/http"
)

func (app *application) getAddressAddressAutocomplete(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getAddressAddressValidate(w http.ResponseWriter, r *http.Request)     { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postAddressAddressGetDetails(w http.ResponseWriter, r *http.Request)   { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getApprovalGetApprovalTasks(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) putApprovalEditTicket(w http.ResponseWriter, r *http.Request)       { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) postDiscountDiscountGenerate(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) putDiscountDiscountExtend(w http.ResponseWriter, r *http.Request)    { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getDynamicFieldsGetDynamicFieldData(w http.ResponseWriter, r *http.Request)       { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getDynamicFieldsGetJsonFormData(w http.ResponseWriter, r *http.Request)           { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postDynamicFieldsGetDataSourceValuesForInputName(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postDynamicFieldsGetValueForInputName(w http.ResponseWriter, r *http.Request)           { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postDynamicFieldsSetDynamicFieldData(w http.ResponseWriter, r *http.Request)            { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postDynamicFieldsValidation(w http.ResponseWriter, r *http.Request)                     { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postDynamicFieldsValidateFormInput(w http.ResponseWriter, r *http.Request)             { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postDynamicFieldsOptionsSearchPost(w http.ResponseWriter, r *http.Request)             { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postDynamicFieldsMigrateDynamicSignupData(w http.ResponseWriter, r *http.Request)      { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) postLacoreConnectNotificationProcess(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getLegalGetLegalDocument(w http.ResponseWriter, r *http.Request)      { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getLegalCheckAgreements(w http.ResponseWriter, r *http.Request)       { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postLegalLogAgreedToDocumentsByMainPk(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postLegalLogAgreedToDocumentsByGuid(w http.ResponseWriter, r *http.Request)   { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getMainExtensionGetItalyCardData(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getOptInGetOptInSettings(w http.ResponseWriter, r *http.Request)   { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) putOptInUpdateOptInSettings(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) deleteOptInUnsubscribe(w http.ResponseWriter, r *http.Request)      { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getR3GlobalGetUser2x2Matrix(w http.ResponseWriter, r *http.Request)   { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) putR3GlobalReset2x2MatrixTimer(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getRepSiteGetRepsiteInfo(w http.ResponseWriter, r *http.Request)                 { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getRepSiteInvalidateRepSiteInfo(w http.ResponseWriter, r *http.Request)          { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getRepSiteGetRepSiteMainPk(w http.ResponseWriter, r *http.Request)               { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getRepSiteGetAdWords(w http.ResponseWriter, r *http.Request)                     { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getRepSiteGetImpressumInfo(w http.ResponseWriter, r *http.Request)               { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getRepSiteGetDistributorDiscountCodes(w http.ResponseWriter, r *http.Request)    { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postRepSiteSendReferrerEmail(w http.ResponseWriter, r *http.Request)             { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postRepSiteAddLead(w http.ResponseWriter, r *http.Request)                       { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postRepSiteValidateSiteUrl(w http.ResponseWriter, r *http.Request)               { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getServerReportsGetPdf(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getSmartMobileMoveMainUserToSmartMobile(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getSmartMobileGetReportParameters(w http.ResponseWriter, r *http.Request)       { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getSmartMobileGetReports(w http.ResponseWriter, r *http.Request)                { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getSmartMobileGetMessage(w http.ResponseWriter, r *http.Request)                { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) putSmartMobileUpdateAllMessages(w http.ResponseWriter, r *http.Request)         { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) putSmartMobileUpdateMessage(w http.ResponseWriter, r *http.Request)             { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postSmartMobileRegisterDeviceToken(w http.ResponseWriter, r *http.Request)      { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getSponsorshipGetSponsorSiteUrl(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getSponsorshipTreeInDownline(w http.ResponseWriter, r *http.Request)    { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getSponsorshipGetSponsorshipTotal(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getSyncGetDataToSync(w http.ResponseWriter, r *http.Request)          { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getSyncGetPushOrders(w http.ResponseWriter, r *http.Request)          { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getSyncGetQueryStatusOrders(w http.ResponseWriter, r *http.Request)   { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postSyncSync(w http.ResponseWriter, r *http.Request)                  { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postSyncSetSyncStatus(w http.ResponseWriter, r *http.Request)         { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postSyncProcessPushResults(w http.ResponseWriter, r *http.Request)    { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postSyncProcessQueryStatusResults(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getTestGet(w http.ResponseWriter, r *http.Request)         { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) getTestNugetTester(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

func (app *application) getTrainingGetTrainingRecordsForUser(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postTrainingSetupTrainingRecordsForUser(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }
func (app *application) postTrainingRecordTrainingAction(w http.ResponseWriter, r *http.Request) { app.errorResponse(w, r, http.StatusNotImplemented, "not implemented") }

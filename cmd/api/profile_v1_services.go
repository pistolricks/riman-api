package main

import (
	"net/http"
)

func (app *application) getAddress_AddressAutocomplete(w http.ResponseWriter, r *http.Request) {}
func (app *application) getAddress_AddressValidate(w http.ResponseWriter, r *http.Request)     {}
func (app *application) postAddress_AddressGetDetails(w http.ResponseWriter, r *http.Request)  {}
func (app *application) getApproval_GetApprovalTasks(w http.ResponseWriter, r *http.Request)   {}
func (app *application) putApproval_EditTicket(w http.ResponseWriter, r *http.Request)         {}
func (app *application) postDiscount_DiscountGenerate(w http.ResponseWriter, r *http.Request)  {}
func (app *application) putDiscount_DiscountExtend(w http.ResponseWriter, r *http.Request)     {}
func (app *application) getDynamicFields_GetDynamicFieldData(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getDynamicFields_GetJsonFormData(w http.ResponseWriter, r *http.Request) {}
func (app *application) postDynamicFields_GetDataSourceValuesForInputName(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postDynamicFields_GetValueForInputName(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postDynamicFields_SetDynamicFieldData(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postDynamicFields_Validation(w http.ResponseWriter, r *http.Request)        {}
func (app *application) postDynamicFields_ValidateFormInput(w http.ResponseWriter, r *http.Request) {}
func (app *application) postDynamicFields_OptionsSearchPost(w http.ResponseWriter, r *http.Request) {}
func (app *application) postDynamicFields_MigrateDynamicSignupData(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postLacoreConnect_NotificationProcess(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getLegal_GetLegalDocument(w http.ResponseWriter, r *http.Request) {}
func (app *application) getLegal_CheckAgreements(w http.ResponseWriter, r *http.Request)  {}
func (app *application) postLegal_LogAgreedToDocumentsByMainPk(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postLegal_LogAgreedToDocumentsByGuid(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getMainExtension_GetItalyCardData(w http.ResponseWriter, r *http.Request) {}
func (app *application) getOptIn_GetOptInSettings(w http.ResponseWriter, r *http.Request)         {}
func (app *application) putOptIn_UpdateOptInSettings(w http.ResponseWriter, r *http.Request)      {}
func (app *application) deleteOptIn_Unsubscribe(w http.ResponseWriter, r *http.Request)           {}
func (app *application) getR3Global_GetUser2x2Matrix(w http.ResponseWriter, r *http.Request)      {}
func (app *application) putR3Global_Reset2x2MatrixTimer(w http.ResponseWriter, r *http.Request)   {}
func (app *application) getRepSite_GetRepsiteInfo(w http.ResponseWriter, r *http.Request)         {}
func (app *application) getRepSite_InvalidateRepSiteInfo(w http.ResponseWriter, r *http.Request)  {}
func (app *application) getRepSite_GetRepSiteMainPk(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getRepSite_GetAdWords(w http.ResponseWriter, r *http.Request)             {}
func (app *application) getRepSite_GetImpressumInfo(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getRepSite_GetDistributorDiscountCodes(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postRepSite_SendReferrerEmail(w http.ResponseWriter, r *http.Request) {}
func (app *application) postRepSite_AddLead(w http.ResponseWriter, r *http.Request)           {}
func (app *application) postRepSite_ValidateSiteUrl(w http.ResponseWriter, r *http.Request)   {}
func (app *application) getServerReports_GetPdf(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getSmartMobile_MoveMainUserToSmartMobile(w http.ResponseWriter, r *http.Request) {
}
func (app *application) getSmartMobile_GetReportParameters(w http.ResponseWriter, r *http.Request)  {}
func (app *application) getSmartMobile_GetReports(w http.ResponseWriter, r *http.Request)           {}
func (app *application) getSmartMobile_GetMessage(w http.ResponseWriter, r *http.Request)           {}
func (app *application) putSmartMobile_UpdateAllMessages(w http.ResponseWriter, r *http.Request)    {}
func (app *application) putSmartMobile_UpdateMessage(w http.ResponseWriter, r *http.Request)        {}
func (app *application) postSmartMobile_RegisterDeviceToken(w http.ResponseWriter, r *http.Request) {}
func (app *application) getSponsorship_GetSponsorSiteUrl(w http.ResponseWriter, r *http.Request)    {}
func (app *application) getSponsorship_TreeInDownline(w http.ResponseWriter, r *http.Request)       {}
func (app *application) getSponsorship_GetSponsorshipTotal(w http.ResponseWriter, r *http.Request)  {}
func (app *application) getSync_GetDataToSync(w http.ResponseWriter, r *http.Request)               {}
func (app *application) getSync_GetPushOrders(w http.ResponseWriter, r *http.Request)               {}
func (app *application) getSync_GetQueryStatusOrders(w http.ResponseWriter, r *http.Request)        {}
func (app *application) postSync_Sync(w http.ResponseWriter, r *http.Request)                       {}
func (app *application) postSync_SetSyncStatus(w http.ResponseWriter, r *http.Request)              {}
func (app *application) postSync_ProcessPushResults(w http.ResponseWriter, r *http.Request)         {}
func (app *application) postSync_ProcessQueryStatusResults(w http.ResponseWriter, r *http.Request)  {}
func (app *application) getTest_Get(w http.ResponseWriter, r *http.Request)                         {}
func (app *application) getTest_NugetTester(w http.ResponseWriter, r *http.Request)                 {}
func (app *application) getTraining_GetTrainingRecordsForUser(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postTraining_SetupTrainingRecordsForUser(w http.ResponseWriter, r *http.Request) {
}
func (app *application) postTraining_RecordTrainingAction(w http.ResponseWriter, r *http.Request) {}

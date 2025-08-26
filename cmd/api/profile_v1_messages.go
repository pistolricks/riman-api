package main

import (
	"net/http"
)

func (app *application) getMessages_GetNews(w http.ResponseWriter, r *http.Request)                {}
func (app *application) getMessages_GetUserAlerts(w http.ResponseWriter, r *http.Request)          {}
func (app *application) getMessages_GetUnreadAlertsCount(w http.ResponseWriter, r *http.Request)   {}
func (app *application) getMessages_GetAlertTypes(w http.ResponseWriter, r *http.Request)          {}
func (app *application) getMessages_GetUnreadMessagesCount(w http.ResponseWriter, r *http.Request) {}
func (app *application) getMessages_GetFolders(w http.ResponseWriter, r *http.Request)             {}
func (app *application) postMessages_CreateFolder(w http.ResponseWriter, r *http.Request)          {}
func (app *application) getMessages_GetNewsCategories(w http.ResponseWriter, r *http.Request)      {}
func (app *application) getMessages_GetNewsDetails(w http.ResponseWriter, r *http.Request)         {}
func (app *application) getMessages_GetMessage(w http.ResponseWriter, r *http.Request)             {}
func (app *application) getMessages_GetMessages(w http.ResponseWriter, r *http.Request)            {}
func (app *application) putMessages_UpdateMessages(w http.ResponseWriter, r *http.Request)         {}
func (app *application) putMessages_RenameeFolder(w http.ResponseWriter, r *http.Request)          {}
func (app *application) deleteMessages_DeleteFolder(w http.ResponseWriter, r *http.Request)        {}
func (app *application) putMessages_UpdateAllMessages(w http.ResponseWriter, r *http.Request)      {}
func (app *application) putMessages_MoveMessages(w http.ResponseWriter, r *http.Request)           {}
func (app *application) postMessages_MarkUserAlertsAsRead(w http.ResponseWriter, r *http.Request)  {}
func (app *application) postMessages_MarkAllUserAlertsAsRead(w http.ResponseWriter, r *http.Request) {
}
func (app *application) deleteMessages_DeleteUserAlerts(w http.ResponseWriter, r *http.Request)    {}
func (app *application) deleteMessages_DeleteAllUserAlerts(w http.ResponseWriter, r *http.Request) {}
func (app *application) deleteMessages_DeleteAllMessages(w http.ResponseWriter, r *http.Request)   {}

func LacoreConnect_NotificationProcess() {}

func OptIn_GetOptInSettings()    {}
func OptIn_UpdateOptInSettings() {}
func OptIn_Unsubscribe()         {}

func SmartMobile_MoveMainUserToSmartMobile() {}
func SmartMobile_GetReportParameters()       {}
func SmartMobile_GetReports()                {}
func SmartMobile_GetMessage()                {}
func SmartMobile_UpdateAllMessages()         {}
func SmartMobile_UpdateMessage()             {}
func SmartMobile_RegisterDeviceToken()       {}

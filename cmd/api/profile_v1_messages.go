package main

import (
	"net/http"
	"strconv"

	"github.com/antihax/optional"
	"github.com/pistolricks/riman-api/internal/profile_v1"
)

func (app *application) getGetNews(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	opts := &profile_v1.MessagesApiMessagesGetNewsOpts{}
	if v := q.Get("username"); v != "" { opts.Username = optional.NewString(v) }
	if v := q.Get("lastArticleFK"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.LastArticleFK = optional.NewInt32(int32(n)) } }
	if v := q.Get("length"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.Length = optional.NewInt32(int32(n)) } }
	if v := q.Get("categories"); v != "" { opts.Categories = optional.NewString(v) }
	if v := q.Get("mainFK"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.MainFK = optional.NewInt32(int32(n)) } }
	if v := q.Get("deviceCultureName"); v != "" { opts.DeviceCultureName = optional.NewString(v) }
	if v := q.Get("title"); v != "" { opts.Title = optional.NewString(v) }
	res, h, err := app.profileV1.MessagesApi.MessagesGetNews(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"news": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUserAlerts(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.MessagesApiMessagesGetUserAlertsOpts{}
	res, h, err := app.profileV1.MessagesApi.MessagesGetUserAlerts(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"alerts": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUnreadAlertsCount(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.MessagesApiMessagesGetUnreadAlertsCountOpts{}
	res, h, err := app.profileV1.MessagesApi.MessagesGetUnreadAlertsCount(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"unreadAlerts": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetAlertTypes(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.MessagesApiMessagesGetAlertTypesOpts{}
	res, h, err := app.profileV1.MessagesApi.MessagesGetAlertTypes(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"alertTypes": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetUnreadMessagesCount(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.MessagesApiMessagesGetUnreadMessagesCountOpts{}
	res, h, err := app.profileV1.MessagesApi.MessagesGetUnreadMessagesCount(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"unreadMessages": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetFolders(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.MessagesApiMessagesGetFoldersOpts{}
	res, h, err := app.profileV1.MessagesApi.MessagesGetFolders(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"folders": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) postCreateFolder(w http.ResponseWriter, r *http.Request) {
	var model profile_v1.JeunesseProfileCoreMessagesModelsFolderQueryModel
	if err := app.readJSON(w, r, &model); err != nil { app.badRequestResponse(w, r, err); return }
	res, h, err := app.profileV1.MessagesApi.MessagesCreateFolder(r.Context(), model, nil)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"folder": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetNewsCategories(w http.ResponseWriter, r *http.Request) {
	opts := &profile_v1.MessagesApiMessagesGetNewsCategoriesOpts{}
	res, h, err := app.profileV1.MessagesApi.MessagesGetNewsCategories(r.Context(), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"categories": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetNewsDetails(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	articlePkStr := q.Get("articlePK")
	articlePk64, _ := strconv.ParseInt(articlePkStr, 10, 32)
	opts := &profile_v1.MessagesApiMessagesGetNewsDetailsOpts{}
	res, h, err := app.profileV1.MessagesApi.MessagesGetNewsDetails(r.Context(), int32(articlePk64), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"article": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetMessage(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	folderIdStr := q.Get("folderID")
	messageIdStr := q.Get("messageID")
	folderId64, _ := strconv.ParseInt(folderIdStr, 10, 32)
	messageId64, _ := strconv.ParseInt(messageIdStr, 10, 32)
	opts := &profile_v1.MessagesApiMessagesGetMessageOpts{}
	res, h, err := app.profileV1.MessagesApi.MessagesGetMessage(r.Context(), int32(folderId64), int32(messageId64), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"message": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) getGetMessages(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	folderIdStr := q.Get("folderID")
	folderId64, _ := strconv.ParseInt(folderIdStr, 10, 32)
	opts := &profile_v1.MessagesApiMessagesGetMessagesOpts{}
 if v := q.Get("pageNumber"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.QueryPageNumber = optional.NewInt32(int32(n)) } }
	if v := q.Get("pageSize"); v != "" { if n, err := strconv.ParseInt(v, 10, 32); err == nil { opts.QueryPageSize = optional.NewInt32(int32(n)) } }
	res, h, err := app.profileV1.MessagesApi.MessagesGetMessages(r.Context(), int32(folderId64), opts)
	if err != nil { app.serverErrorResponse(w, r, err); return }
	if err := app.writeJSON(w, h.StatusCode, envelope{"messages": res}, nil); err != nil { app.serverErrorResponse(w, r, err) }
}
func (app *application) putUpdateMessages(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putRenameeFolder(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) deleteDeleteFolder(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putUpdateAllMessages(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) putMoveMessages(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postMarkUserAlertsAsRead(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) postMarkAllUserAlertsAsRead(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) deleteDeleteUserAlerts(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) deleteDeleteAllUserAlerts(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}
func (app *application) deleteDeleteAllMessages(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotImplemented, "not implemented")
}

package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/pistolricks/riman-api/internal/data"
	"github.com/pistolricks/riman-api/internal/security"
	"github.com/pistolricks/riman-api/internal/validator"
)

func (app *application) secureVendorLoginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateUsername(v, input.Username)
	data.ValidatePasswordPlaintext(v, input.Password)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	client, err := app.models.Clients.GetByUsername(input.Username)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidCredentialsResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	match, err := client.Password.MatchesPassword(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.invalidCredentialsResponse(w, r)
		return
	}

	tokenRequest := security.TokenRequestModel{
		UserName: input.Username,
		Password: input.Password,
	}

	token, h, err := app.security.AuthenticationApi.AuthenticationAuthenticateUser(context.Background(), tokenRequest)
	if err != nil {
		return
	}

	err = app.models.Clients.UpdateToken(client.ID, token)
	if err != nil {
		return
	}
	err = app.writeJSON(w, h.StatusCode, envelope{"username": client.Username, "authentication_token": token}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) secureTokenRefreshHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Username string `json:"username"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	client, err := app.models.Clients.GetByUsername(input.Username)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.invalidCredentialsResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	token, h, err := app.security.AuthenticationApi.AuthenticationReIssue(context.Background())
	if err != nil {
		return
	}

	err = app.models.Clients.UpdateToken(client.ID, token)
	if err != nil {
		return
	}
	err = app.writeJSON(w, h.StatusCode, envelope{"username": client.Username, "authentication_token": token}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

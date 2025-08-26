package security

import (
	isec "github.com/pistolricks/riman-api/internal/security"
)

type APIClient = isec.APIClient

type Configuration = isec.Configuration

type AuthenticationApiService = isec.AuthenticationApiService

type TokenRequestModel = isec.TokenRequestModel

type WalletTokenRequestModel = isec.WalletTokenRequestModel

type PasswordRulesModel = isec.PasswordRulesModel

var NewAPIClient = isec.NewAPIClient

var NewConfiguration = isec.NewConfiguration

package data

import (
	id "github.com/pistolricks/riman-api/internal/data"
	iv "github.com/pistolricks/riman-api/internal/validator"
)

// Re-export minimal symbols required by cmd/api/security.go (and compatible with main.go)

type Models = id.Models

var NewModels = id.NewModels

var (
	ErrRecordNotFound = id.ErrRecordNotFound
)

// Re-export validator dependency for function signatures convenience
// Note: we re-use the internal validator type via pkg/validator in the handler, but for direct function forwarding we need the exact type.
type Validator = iv.Validator

func ValidateUsername(v *iv.Validator, username string) { id.ValidateUsername(v, username) }
func ValidatePasswordPlaintext(v *iv.Validator, password string) { id.ValidatePasswordPlaintext(v, password) }

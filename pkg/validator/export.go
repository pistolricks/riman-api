package validator

import (
	"regexp"
	iv "github.com/pistolricks/riman-api/internal/validator"
)

// Re-export the internal validator for use by cmd when running single files.

type Validator = iv.Validator

var New = iv.New

var (
	EmailRX = iv.EmailRX
)

func PermittedValue[T comparable](value T, permittedValues ...T) bool { return iv.PermittedValue(value, permittedValues...) }
func Matches(value string, rx *regexp.Regexp) bool { return iv.Matches(value, rx) }
func Unique[T comparable](values []T) bool { return iv.Unique(values) }

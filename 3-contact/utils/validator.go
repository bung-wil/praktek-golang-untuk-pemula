package utils

import (
	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct memvalidasi struct berdasarkan tag validate
func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}

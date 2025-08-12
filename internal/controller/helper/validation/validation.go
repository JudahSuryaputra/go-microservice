package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Struct(s interface{}) error {
	return validate.Struct(s)
}

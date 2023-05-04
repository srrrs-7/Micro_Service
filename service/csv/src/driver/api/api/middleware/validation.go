package middleware

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func NewValidation() {
	// validation setting
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("customValidation", customValidation)
	}
}

func customValidation(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()

	// value is empty ?
	return fieldValue != ""
}

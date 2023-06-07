package utils

import (
	"gopkg.in/go-playground/validator.v9"
)

func StructValidator(s interface{}) *string {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		validateError := err.Error()
		return &validateError
	}
	return nil
}

package utils

import (
	globalmodels "github.com/shine-bright-team/LAAS/v2/global_models"
	"gopkg.in/go-playground/validator.v9"
)

func StructValidator(s interface{}) *globalmodels.ErrorResponse {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		return &globalmodels.ErrorResponse{
			Type:    "Bad Request",
			Message: err.Error(),
		}
	}
	return nil
}

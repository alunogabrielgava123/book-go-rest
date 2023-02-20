package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func ValidateBody(s interface {}, ctx *gin.Context) error {

	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		return err
	}

	return nil

}

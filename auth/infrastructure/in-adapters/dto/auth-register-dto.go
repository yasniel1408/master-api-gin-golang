package auth_infrastructure_dto

import (
	"github.com/go-playground/validator"
)

type AuthDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (data *AuthDTO) Validate() error {
	v := validator.New()

	err := v.Struct(data)
	if err != nil {
		return err
	}
	return nil
}

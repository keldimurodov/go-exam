package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserDetail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type ResponseMessage struct {
	Content string `json:"content"`
}

type VerifyUserRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type VerifyUserResponse struct {
	UserInfo User `json:"userinfo"`
}

type ResponseUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Code      string `json:"code"`
}

type UserValidationStauts struct {
	Status bool `json:"status"`
}

type CheckUser struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func (rm *UserDetail) Validate() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(&rm.Email, validation.Required, is.Email),
		validation.Field(
			&rm.Password,
			validation.Required,
			validation.Length(8, 30),
			validation.Match(regexp.MustCompile("[a-z]|[A-Z][1-9]")),
		),
	)
}
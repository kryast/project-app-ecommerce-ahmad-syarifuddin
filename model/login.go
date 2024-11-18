package model

type Login struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Phone    string `json:"phone,omitempty" validate:"required_without=Email"`
	Password string `json:"password" validate:"required"`
}

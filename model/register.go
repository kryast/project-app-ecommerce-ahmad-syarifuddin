package model

import "database/sql"

type Register struct {
	ID       int            `json:"id,omitempty"`
	Name     string         `json:"name" validate:"required,min=3,max=100"`
	Email    sql.NullString `json:"email,omitempty" validate:"required,email"`
	Phone    sql.NullString `json:"phone,omitempty" validate:"required_without=Email,phone,excludes=Email"`
	Password string         `json:"password" validate:"required"`
}

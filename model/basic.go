package model

import "time"

type Basic struct {
	ID         uint       `json:"id" validate:"required"`
	Created_at time.Time  `json:"created_at" validate:"required"`
	Updated_at time.Time  `json:"updated_at" validate:"required"`
	Deleted_at *time.Time `json:"deleted_at_at,omitempty" validate:"required"`
}

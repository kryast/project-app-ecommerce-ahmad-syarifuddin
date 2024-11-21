package model

import (
	"database/sql"
	"time"
)

type Product struct {
	ID        int          `json:"id"`
	Name      string       `json:"product_name"`
	Detail    string       `json:"product_detail"`
	Price     float64      `json:"product_price"`
	Category  string       `json:"category_name"`
	PhotoURL  string       `json:"photo_url"`
	FlagNew   string       `json:"flag_new,omitempty"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
}

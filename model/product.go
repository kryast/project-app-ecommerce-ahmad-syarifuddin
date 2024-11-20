package model

import (
	"database/sql"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"product_name"`
	Detail    string  `json:"product_detail"`
	Price     float64 `json:"product_price"`
	Category  string  `json:"category_name"`
	PhotoURL  string  `json:"photo_url"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

package model

type Wishlist struct {
	ID               int     `json:"id"`
	ProductDetailsID int     `json:"product_details_id,omitempty"`
	Name             string  `json:"product_name,omitempty"`
	Detail           string  `json:"product_detail,omitempty"`
	Price            float64 `json:"product_price,omitempty"`
	Category         string  `json:"category_name,omitempty"`
	PhotoURL         string  `json:"photo_url,omitempty"`
}

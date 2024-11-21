package model

type Cart struct {
	ID         int     `json:"id"`
	ProductID  int     `json:"product_details_id,omitempty"`
	Product    string  `json:"product_name"`
	Price      float64 `json:"product_price"`
	Quantity   int     `json:"product_quantity"`
	Subtotal   float64 `json:"product_subtotal"`
	TotalPrice float64 `json:"total_price,omitempty"`
}

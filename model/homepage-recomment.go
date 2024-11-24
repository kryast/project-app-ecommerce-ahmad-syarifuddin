package model

type Recomment struct {
	ID            int    `json:"id"`
	Product       string `json:"title"`
	ProductDetail string `json:"subtitle"`
	PhotoURL      string `json:"photo"`
	ProductID     int    `json:"product_id"`
}

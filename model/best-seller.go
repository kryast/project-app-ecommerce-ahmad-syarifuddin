package model

type ProductBestSeller struct {
	ID        int    `json:"id"`
	Name      string `json:"product_name"`
	TotalSold int    `json:"total_sold"`
}

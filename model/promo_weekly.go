package model

type PromoWeekly struct {
	ID               int    `json:"id"`
	ProductDetailsID int    `json:"product_details_id"`
	StartWeek        string `json:"start_week"`
	EndWeek          string `json:"end_week"`
}

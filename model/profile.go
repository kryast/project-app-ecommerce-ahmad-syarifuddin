package model

type Profile struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   *string `json:"email,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Address *string `json:"address"`
}

package model

type Order struct {
	ID      int `json:"id,omitempty"`
	AuthID  int
	Address string
	Payment string
	Cart    []*Cart
}

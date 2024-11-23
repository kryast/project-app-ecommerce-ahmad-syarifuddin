package model

type Address struct {
	ID          int    `json:"id"`
	AuthID      int    `json:"auth_id"`
	Address     string `json:"address"`
	FlagDefault bool   `json:"flag_default"`
}

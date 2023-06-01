package models

import "time"

type Wallet struct {
	ID              int32     `json:"id"`
	Quantidade      float64   `json:"quantidade"`
	UpdatedAt       time.Time `json:"updated_at"`
	QuantidadeAntes float64   `json:"quantidade_antes"`
}

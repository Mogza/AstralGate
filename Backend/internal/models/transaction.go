package models

import "time"

type Transaction struct {
	Id            int64     `json:"id"`
	WalletId      int64     `json:"wallet_id"`
	ProductId     int64     `json:"product_id"`
	ClientAddress string    `json:"client_address"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	TxHash        string    `json:"tx_hash"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

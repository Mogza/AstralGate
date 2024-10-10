package models

import (
	"time"
)

type Wallet struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	CryptoAddress string    `json:"crypto_address"`
	Balance       float64   `json:"balance"`
	Currency      string    `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

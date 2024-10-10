package models

import "time"

type Products struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UsdPrice    float64   `json:"usd_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

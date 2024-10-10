package models

import "time"

type User struct {
	Id           int64     `json:"id"`
	Username     string    `json:"username"`
	Firstname    string    `json:"first_name" gorm:"column:first_name"`
	Lastname     string    `json:"last_name" gorm:"column:last_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	PhoneNumber  string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Mogza/AstralGate/internal/models"
	"github.com/Mogza/AstralGate/internal/utils"
	"net/http"
)

type RegisterRequest struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	err = h.DB.Where("username = ?", req.Username).First(&existingUser).Error
	if err == nil {
		http.Error(w, "User with this username already exists", http.StatusConflict)
		return
	}
	err = h.DB.Where("email = ?", req.Email).First(&existingUser).Error
	if err == nil {
		http.Error(w, "User with this email already exists", http.StatusConflict)
		return
	}
	err = h.DB.Where("phone_number = ?", req.PhoneNumber).First(&existingUser).Error
	if err == nil {
		http.Error(w, "User with this phone number already exists", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username:     req.Username,
		Firstname:    req.FirstName,
		Lastname:     req.LastName,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Role:         "merchant",
		PhoneNumber:  req.PhoneNumber,
	}

	err = h.DB.Create(&user).Error
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "{\"ok\":\"user successfully created\"}")
	utils.LogFatal(err, "Fprintf failed")
}

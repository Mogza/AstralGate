package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Mogza/AstralGate/internal/models"
	"github.com/Mogza/AstralGate/internal/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// LoginRequest : Requested body for login
type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// RegisterRequest : Requested body for registration
type RegisterRequest struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve targeted user
	var user models.User
	err = h.DB.Where("username = ?", req.Login).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = h.DB.Where("email = ?", req.Login).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Invalid username/email or password", http.StatusUnauthorized)
			return
		}
	}

	// Check password
	if !utils.CheckPasswordHash(user.PasswordHash, req.Password) {
		http.Error(w, "Invalid username/email or password", http.StatusUnauthorized)
		return
	}

	// Generate the token
	tokenString, err := utils.GenerateJWT(user.Id, user.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	if err != nil {
		return
	}
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Search for existing users
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

	// Hash the password
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

	// Create user in database
	err = h.DB.Create(&user).Error
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Create hot wallet for the user
	walletAddress := utils.CreateWallet(req.Password)

	wallet := models.Wallet{
		UserId:        user.Id,
		CryptoAddress: walletAddress,
		Currency:      "POL",
	}

	// Save wallet in database
	err = h.DB.Create(&wallet).Error
	if err != nil {
		http.Error(w, "Failed to wallet for the user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "{\"ok\":\"user successfully created\"}")
	utils.LogFatal(err, "Fprintf failed")
}

func (h Handler) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	// Retrieve all users information
	var User []models.User
	err := h.DB.Order("id").Find(&User).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(User)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetUserMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// Retrieve expected user information
	var user models.User
	err := h.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["user_id"])

	// Retrieve expected user information
	var user models.User
	err := h.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetUserWalletMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// Retrieve user wallets
	var wallets []models.Wallet
	err := h.DB.Where("user_id = ?", userID).Find(&wallets).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(wallets)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetUserProductMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// Retrieve user products
	var products []models.Products
	err := h.DB.Where("user_id = ?", userID).Find(&products).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetUserWalletById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["user_id"])

	// Retrieve expected wallet
	var wallets []models.Wallet
	err := h.DB.Where("user_id = ?", id).Find(&wallets).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(wallets)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetUserProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["user_id"])

	// Retrieve products of expected user
	var products []models.Products
	err := h.DB.Where("user_id = ?", id).Find(&products).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["user_id"])

	// Retrieve expected user
	var user models.User
	err := h.DB.First(&user, id).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode request body
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("No Body")
	}

	// Update user in database
	err = h.DB.Save(&user).Error
	if err != nil {
		http.Error(w, "User not updated", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["user_id"])

	// Delete expected user
	var user models.User
	err := h.DB.Delete(&user, id).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "{\"ok\":\"user successfully deleted\"}")
	utils.LogFatal(err, "Fprintf failed")
}

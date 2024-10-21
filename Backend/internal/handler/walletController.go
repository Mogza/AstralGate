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

func (h Handler) GetAllWallets(w http.ResponseWriter, _ *http.Request) {
	// Retrieve all wallets
	var wallet []models.Wallet
	err := h.DB.Order("id").Find(&wallet).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(wallet)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetWalletById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["wallet_id"])

	// Retrieve expected wallet
	var wallet models.Wallet
	err := h.DB.First(&wallet, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Wallet not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(wallet)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) UpdateWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["wallet_id"])

	// Retrieve expected wallet
	var wallet models.Wallet
	err := h.DB.First(&wallet, id).Error
	if err != nil {
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	}

	// Decode request body
	err = json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		fmt.Println("No Body")
	}

	// Update wallet in database
	err = h.DB.Save(&wallet).Error
	if err != nil {
		http.Error(w, "Wallet not updated", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(wallet)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) UpdateBalance() {
	// Retrieve all wallets where currency == POL
	var wallets []models.Wallet
	err := h.DB.Where("currency = ?", "POL").Find(&wallets).Error
	if err != nil {
		fmt.Println("[UpdateBalance] Error while retrieving all wallets: ", err)
		return
	}

	if len(wallets) == 0 {
		fmt.Println("[UpdateBalance] No wallets found")
		return
	}

	// Update wallets balance
	for _, wallet := range wallets {
		newBalance := utils.GetMaticBalance(wallet.CryptoAddress)
		if newBalance != wallet.Balance {
			wallet.Balance = newBalance

			err = h.DB.Save(wallet).Error
			if err != nil {
				fmt.Println("[UpdateBalance] Error while saving wallet ( ", wallet.CryptoAddress, "): ", err)
			}
		}
	}
}

func (h Handler) DeleteWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["wallet_id"])

	// Delete expected wallet
	var wallet models.Wallet
	err := h.DB.Delete(&wallet, id).Error
	if err != nil {
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "{\"ok\":\"wallet successfully deleted\"}")
	utils.LogFatal(err, "Fprintf failed")
}

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

func (h Handler) GetAllTransactions(w http.ResponseWriter, _ *http.Request) {
	// Retrieve all transactions
	var transactions []models.Transaction
	err := h.DB.Order("id").Find(&transactions).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(transactions)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["transaction_id"])

	// Retrieve expected transaction
	var transaction models.Transaction
	err := h.DB.First(&transaction, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Transaction not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(transaction)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) CreatePOLTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// Decode body request
	var newTx models.Transaction
	err := json.NewDecoder(r.Body).Decode(&newTx)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var wallets models.Wallet
	err = h.DB.Where("user_id = ? AND currency = ?", userID, "POL").Find(&wallets).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	newTx.WalletId = wallets.Id
	newTx.Currency = "POL"
	newTx.Status = "pending"
	// TODO: call to extern API to convert usd price to POL price
	newTx.Amount = 1

	// Create Transaction
	err = h.DB.Create(&newTx).Error
	if err != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintln(w, "{\"merchant_address\": ", wallets.CryptoAddress, "}")
	utils.LogFatal(err, "Fprintf failed")
}

func (h Handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["transaction_id"])

	// Retrieve expected transaction
	var transaction models.Transaction
	err := h.DB.First(&transaction, id).Error
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	// Decode request body
	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		fmt.Println("No Body")
	}

	// Update transaction in database
	err = h.DB.Save(&transaction).Error
	if err != nil {
		http.Error(w, "Wallet not updated", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(transaction)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["transaction_id"])

	// Delete expected transaction
	var transaction models.Transaction
	err := h.DB.Delete(&transaction, id).Error
	if err != nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "{\"ok\":\"transaction successfully deleted\"}")
	utils.LogFatal(err, "Fprintf failed")
}

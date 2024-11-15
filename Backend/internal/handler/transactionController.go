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

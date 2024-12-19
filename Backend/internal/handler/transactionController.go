package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Mogza/AstralGate/internal/models"
	"github.com/Mogza/AstralGate/internal/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PolygonTXResult struct {
	Hash  string `json:"hash"`
	From  string `json:"from"`
	Value string `json:"value"`
}

type PolygonTX struct {
	Result []PolygonTXResult `json:"result"`
}

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

func (h Handler) CheckPaidTransaction() {
	// Retrieve all wallets where currency == POL
	var wallets []models.Wallet
	err := h.DB.Where("currency = ?", "POL").Find(&wallets).Error
	if err != nil {
		fmt.Println("[UpdateBalance] Error while retrieving all wallets: ", err)
		return
	}

	// Count to respect calls per seconds of the api
	count := 0

	for _, wallet := range wallets {
		count += 1

		polygonScanToken := os.Getenv("POLYGONSCAN_TOKEN")

		// Get Wallet onchain TX
		polygonScanUrl := "https://api-amoy.polygonscan.com/api?module=account&action=txlist&address=" + wallet.CryptoAddress + "&tag=latest&apikey=" + polygonScanToken

		// Creating request
		req, err := http.NewRequest("GET", polygonScanUrl, nil)
		if err != nil {
			fmt.Println("checkPaidTransaction : Error while creating new request")
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("chackPaidTransaction : Error while creating sending request")
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				utils.LogFatal(err, "Error while closing the body")
			}
		}(resp.Body)

		body, _ := io.ReadAll(resp.Body)
		var tx PolygonTX
		err = json.Unmarshal(body, &tx)
		utils.LogFatal(err, "Error while unmarshalling response body")

		for _, transaction := range tx.Result {
			// Ensure this is not an already verify tx
			var paidTX []models.Transaction
			err = h.DB.Where("tx_hash = ?", transaction.Hash).First(&paidTX).Error
			if err == nil {
				continue
			}

			txValuePOL, err := strconv.ParseFloat(transaction.Value, 64)
			if err != nil {
				utils.LogFatal(err, "Error while converting value to float")
			}
			txValueWei := txValuePOL / 1e18

			var pendingTX models.Transaction
			err = h.DB.Where("client_address = ? AND amount = ? AND status = ?", transaction.From, txValueWei, "pending").First(&pendingTX).Error
			if err == nil {
				pendingTX.Status = "paid"
				pendingTX.TxHash = transaction.Hash
				h.DB.Save(pendingTX)
			}
		}

		if count >= 5 {
			count = 0
			time.Sleep(1 * time.Second)
		}
	}

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

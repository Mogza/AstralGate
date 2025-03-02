package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Mogza/AstralGate/internal/models"
	"github.com/Mogza/AstralGate/internal/service"
	"github.com/Mogza/AstralGate/internal/utils"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type AlchemyPriceData struct {
	Value string `json:"value"`
}

type AlchemyReturnData struct {
	Prices []AlchemyPriceData `json:"prices"`
}

type AlchemyReturn struct {
	Data []AlchemyReturnData `json:"data"`
}

type RevenueReturn struct {
	Revenue string `json:"revenue"`
}

type UsersOnboardedReturn struct {
	Count int64 `json:"count"`
}

type ItemSoldData struct {
	Name   string `json:"name"`
	Number int64  `json:"number"`
}

type ItemSoldReturn struct {
	Data []ItemSoldData `json:"data"`
}

type ActivityData struct {
	Period string `json:"period"`
	Number int64  `json:"number"`
}

type ActivityReturn struct {
	Data []ActivityData `json:"data"`
}

func (h Handler) GetUserRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// Retrieve expected wallet
	var wallet models.Wallet
	err := h.DB.Where("user_id = ?", userID).First(&wallet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Wallet not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Retrieve wallet balance
	polRevenue := service.GetTotalRevenue(wallet)

	// Convert POL to USD
	alchemyApiUrl := "https://api.g.alchemy.com/prices/v1/tokens/by-symbol?symbols=POL"

	req, err := http.NewRequest("GET", alchemyApiUrl, nil)
	if err != nil {
		fmt.Println("GetUserRevenue : Error while creating new request")
		return
	}

	alchemyToken := os.Getenv("ALCHEMY_TOKEN")
	req.Header.Set("Authorization", "Bearer "+alchemyToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("GetUserRevenue : Error while creating sending request")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			utils.LogFatal(err, "Error while closing the body")
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	var priceData AlchemyReturn
	err = json.Unmarshal(body, &priceData)
	utils.LogFatal(err, "Error while unmarshalling response body")

	priceString := priceData.Data[0].Prices[0].Value
	price, _ := strconv.ParseFloat(priceString, 32)
	totalRevenue := RevenueReturn{
		Revenue: fmt.Sprintf("%.2f", price*polRevenue.Revenue),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(totalRevenue)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetUsersOnboarded(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// Retrieve expected wallet
	var wallet models.Wallet
	err := h.DB.Where("user_id = ?", userID).First(&wallet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Wallet not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	var count int64
	err = h.DB.Model(&models.Transaction{}).Where("wallet_id = ? AND status = ?", wallet.Id, "paid").Count(&count).Error
	utils.LogFatal(err, "Error while retrieving count of transactions")
	usersOnboarded := UsersOnboardedReturn{
		Count: count,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(usersOnboarded)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetItemsSold(w http.ResponseWriter, r *http.Request) {
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
	}

	// Retrieve expected wallet
	var wallet models.Wallet
	err = h.DB.Where("user_id = ?", userID).First(&wallet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Wallet not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	var itemReturn ItemSoldReturn
	for _, item := range products {
		var count int64
		err = h.DB.Model(&models.Transaction{}).Where("product_id = ? AND status = ?", item.Id, "paid").Count(&count).Error
		if err != nil {
			break
		}

		itemReturn.Data = append(itemReturn.Data, ItemSoldData{
			Name:   item.Title,
			Number: count,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(itemReturn)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetActivity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// Retrieve expected wallet
	var wallet models.Wallet
	err := h.DB.Where("user_id = ?", userID).First(&wallet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Wallet not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	var timePeriod []time.Time
	daysToAnalyze := 7
	now := time.Now()

	for i := 0; i < daysToAnalyze; i++ {
		timePeriod = append(timePeriod, now.AddDate(0, 0, -i))
	}

	var activityReturn ActivityReturn
	for _, period := range timePeriod {
		startOfDay := period.Truncate(24 * time.Hour)
		endOfDay := startOfDay.Add(24 * time.Hour)

		var count int64
		err = h.DB.Model(&models.Transaction{}).
			Where("wallet_id = ? AND status = ? AND updated_at >= ? AND updated_at < ?", wallet.Id, "paid", startOfDay, endOfDay).
			Count(&count).Error

		if err != nil {
			break
		}

		activityReturn.Data = append(activityReturn.Data, ActivityData{
			Period: period.Format("02/01"),
			Number: count,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(activityReturn)
	utils.LogFatal(err, "Error while encoding response")
}

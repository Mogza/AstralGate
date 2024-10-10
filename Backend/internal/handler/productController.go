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

type ProductBodyCreation struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UsdPrice    float64 `json:"usd_price"`
}

func (h Handler) GetAllProducts(w http.ResponseWriter, _ *http.Request) {
	var products []models.Products
	err := h.DB.Order("id").Find(&products).Error
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["product_id"])

	var product models.Products
	err := h.DB.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(product)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) CreateProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	var req ProductBodyCreation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var existingProduct models.Products
	err = h.DB.Where("title = ?", req.Title).First(&existingProduct).Error
	if err == nil {
		http.Error(w, "Product with this title already exists", http.StatusConflict)
		return
	}

	product := models.Products{
		UserId:      int64(userID),
		Title:       req.Title,
		Description: req.Description,
		UsdPrice:    req.UsdPrice,
	}

	err = h.DB.Create(&product).Error
	if err != nil {
		http.Error(w, "Failed to product user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "{\"ok\":\"product successfully created\"}")
	utils.LogFatal(err, "Fprintf failed")
}

func (h Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["product_id"])

	var product models.Products
	err := h.DB.First(&product, id).Error
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println("No Body")
	}
	err = h.DB.Save(&product).Error
	if err != nil {
		http.Error(w, "Product not updated", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(product)
	utils.LogFatal(err, "Error while encoding response")
}

func (h Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["product_id"])

	var product models.Products
	err := h.DB.Delete(&product, id).Error
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, "{\"ok\":\"product successfully deleted\"}")
	utils.LogFatal(err, "Fprintf failed")
}

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
)

// ProductBodyCreation : Requested body for product creation
type ProductBodyCreation struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UsdPrice    float64 `json:"usd_price"`
}

func (h Handler) GetAllProducts(w http.ResponseWriter, _ *http.Request) {
	// Retrieve all products
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

	// Retrieve expected product
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

	// Parse the form
	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	// Extract from the form
	file, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		testname, _, err := r.FormFile("title")
		fmt.Println("test", testname)
		http.Error(w, "Failed to get image from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Decode from the form
	var req ProductBodyCreation
	req.Title = r.FormValue("title")
	req.Description = r.FormValue("description")
	req.UsdPrice, err = strconv.ParseFloat(r.FormValue("usd_price"), 64)
	if err != nil {
		http.Error(w, "Invalid price value", http.StatusBadRequest)
		return
	}

	// See if products already exists (by title)
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

	// Create Product
	err = h.DB.Create(&product).Error
	if err != nil {
		http.Error(w, "Failed to product user", http.StatusInternalServerError)
		return
	}

	// Save image
	imagePath := fmt.Sprintf("/app/images/%d.jpg", product.Id)
	out, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
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

	// Retrieve expected product
	var product models.Products
	err := h.DB.First(&product, id).Error
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Decode body request
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println("No Body")
	}

	// Update product in database
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

	// Delete expected product
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

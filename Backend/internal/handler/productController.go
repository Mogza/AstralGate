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
	"path/filepath"
	"strconv"
	"time"
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

	// Parse the form with a larger size limit for images
	err := r.ParseMultipartForm(32 << 20) // 32MB limit
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	// Extract from the form
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to get image from request: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type
	if !isValidImageType(fileHeader.Header.Get("Content-Type")) {
		http.Error(w, "Invalid file type. Only jpg, jpeg, and png are allowed", http.StatusBadRequest)
		return
	}

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

	// Create image directory if it doesn't exist
	imageDirPath := "images"
	if err := os.MkdirAll(imageDirPath, 0755); err != nil {
		http.Error(w, "Failed to create image directory", http.StatusInternalServerError)
		return
	}

	// Generate unique filename using UUID
	fileExt := filepath.Ext(fileHeader.Filename)
	imageFileName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), int64(userID), fileExt)
	imagePath := filepath.Join(imageDirPath, imageFileName)

	// Create Product with image path
	product := models.Products{
		UserId:      int64(userID),
		Title:       req.Title,
		Description: req.Description,
		UsdPrice:    req.UsdPrice,
		ImagePath:   imagePath, // Add this field to your Products model
	}

	// Create Product in DB
	err = h.DB.Create(&product).Error
	if err != nil {
		http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Save image file
	out, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, "Failed to create image file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Failed to save image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "product successfully created",
		"product": product,
	})
	if err != nil {
		utils.LogFatal(err, "Error while encoding response")
	}
}

// Helper function to validate image types
func isValidImageType(contentType string) bool {
	validTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
	}
	return validTypes[contentType]
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

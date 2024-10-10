package routes

import (
	"github.com/Mogza/AstralGate/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/products/{product_id}", h.GetProductById).Methods("GET")
	router.HandleFunc("/products/", h.CreateProducts).Methods("POST")
	router.HandleFunc("/products/{product_id}", h.UpdateProduct).Methods("PUT")
}

func RegisterProductAdminRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/products", h.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{product_id}", h.DeleteProduct).Methods("DELETE")
}

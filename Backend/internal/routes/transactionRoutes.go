package routes

import (
	"github.com/Mogza/AstralGate/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterTransactionRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/transaction/{transaction_id}", h.GetTransactionById).Methods("GET")
	router.HandleFunc("/transaction/POL/", h.CreatePOLTransactions).Methods("POST")
	router.HandleFunc("/transaction/{transaction_id}", h.UpdateTransaction).Methods("PUT")
}

func RegisterTransactionAdminRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/transaction", h.GetAllTransactions).Methods("GET")
	router.HandleFunc("/transaction/{transaction_id}", h.DeleteTransaction).Methods("DELETE")
}

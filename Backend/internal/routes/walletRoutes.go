package routes

import (
	"github.com/Mogza/AstralGate/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterWalletRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/wallets/{wallet_id}", h.GetWalletById).Methods("GET")
	router.HandleFunc("/wallets/{wallet_id}", h.UpdateWallet).Methods("PUT")
}

func RegisterWalletAdminRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/wallets", h.GetAllWallets).Methods("GET")
	router.HandleFunc("/wallets/{wallet_id}", h.DeleteWallet).Methods("DELETE")
}

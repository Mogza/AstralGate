package routes

import (
	"github.com/Mogza/AstralGate/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterDashboardRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/revenue", h.GetUserRevenue).Methods("GET")
	router.HandleFunc("/users", h.GetUsersOnboarded).Methods("GET")
	router.HandleFunc("/items", h.GetItemsSold).Methods("GET")
	router.HandleFunc("/activity", h.GetActivity).Methods("GET")
}

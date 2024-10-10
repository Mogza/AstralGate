package routes

import (
	"github.com/Mogza/AstralGate/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/register/", h.Register).Methods("POST")
	router.HandleFunc("/login/", h.Login).Methods("POST")
}

func RegisterUserRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/users/me", h.GetUserMe).Methods("GET")
	router.HandleFunc("/users/wallets/{user_id}", h.GetUserWallet).Methods("GET")
	router.HandleFunc("/users/{user_id}", h.UpdateUser).Methods("PUT")
}

func RegisterUserAdminRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/users", h.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{user_id}", h.GetUserById).Methods("GET")
	router.HandleFunc("/users/{user_id}", h.DeleteUser).Methods("DELETE")
}

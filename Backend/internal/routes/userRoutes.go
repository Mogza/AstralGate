package routes

import (
	"github.com/Mogza/AstralGate/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router, h handler.Handler) {
	router.HandleFunc("/register/", h.Register).Methods("POST")
	router.HandleFunc("/login/", h.Login).Methods("POST")
}

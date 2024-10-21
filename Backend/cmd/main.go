package main

import (
	"github.com/Mogza/AstralGate/internal/db"
	"github.com/Mogza/AstralGate/internal/handler"
	"github.com/Mogza/AstralGate/internal/middleware"
	"github.com/Mogza/AstralGate/internal/routes"
	"github.com/Mogza/AstralGate/internal/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

func main() {
	// Database  & Handler init
	DB := db.Init()
	h := handler.New(DB)

	// CORS Setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "*"},
		AllowCredentials: true,
	})

	// Router creation
	router := mux.NewRouter()
	registerRoutes(router, h)
	han := c.Handler(router)

	pollingTicker := time.NewTicker(30 * time.Second)
	defer pollingTicker.Stop()

	// Goroutine : Listen and Serve
	go func() {
		log.Println("Listening on port 8080")
		err := http.ListenAndServe(":8080", han)
		utils.LogFatal(err, "Error starting server")
	}()

	// Goroutine : Polling functions
	go func() {
		for {
			select {
			case <-pollingTicker.C:
				h.UpdateBalance()
			}
		}
	}()

	select {}
}

func registerRoutes(router *mux.Router, h handler.Handler) {
	// Login router setup
	routes.RegisterAuthRoutes(router, h)

	// Api router setup
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.JwtMiddleware)
	routes.RegisterUserRoutes(apiRouter, h)
	routes.RegisterWalletRoutes(apiRouter, h)
	routes.RegisterProductRoutes(apiRouter, h)

	// Admin router setup
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AdminMiddleware)
	routes.RegisterUserAdminRoutes(adminRouter, h)
	routes.RegisterWalletAdminRoutes(adminRouter, h)
	routes.RegisterProductAdminRoutes(adminRouter, h)

}

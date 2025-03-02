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

	// Serve static files from the images directory
	router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	registerRoutes(router, h)
	han := c.Handler(router)

	pollingTicker30 := time.NewTicker(30 * time.Second)
	defer pollingTicker30.Stop()
	pollingTicker5 := time.NewTicker(5 * time.Second)
	defer pollingTicker5.Stop()

	// Goroutine : Listen and Serve
	go func() {
		log.Println("Listening on port 8080")
		err := http.ListenAndServe(":8080", han)
		utils.LogFatal(err, "Error starting server")
	}()

	// Goroutine : Polling functions 30sec
	go func() {
		for {
			select {
			case <-pollingTicker30.C:
				h.UpdateBalance()
			}
		}
	}()

	// Goroutine : Polling functions 5sec
	go func() {
		for {
			select {
			case <-pollingTicker5.C:
				h.CheckPaidTransaction()
			}
		}
	}()

	select {}
}

func registerRoutes(router *mux.Router, h handler.Handler) {
	// Login router setup
	routes.RegisterAuthRoutes(router, h)
	routes.RegisterPublicTransactionRoutes(router, h)

	// Api router setup
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.JwtMiddleware)
	apiRouter.Use(middleware.CORS)
	routes.RegisterUserRoutes(apiRouter, h)
	routes.RegisterWalletRoutes(apiRouter, h)
	routes.RegisterProductRoutes(apiRouter, h)
	routes.RegisterTransactionRoutes(apiRouter, h)

	// Api stats router setup
	statsRouter := router.PathPrefix("/stats").Subrouter()
	statsRouter.Use(middleware.JwtMiddleware)
	routes.RegisterDashboardRoutes(statsRouter, h)

	// Admin router setup
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AdminMiddleware)
	routes.RegisterUserAdminRoutes(adminRouter, h)
	routes.RegisterWalletAdminRoutes(adminRouter, h)
	routes.RegisterProductAdminRoutes(adminRouter, h)
	routes.RegisterTransactionAdminRoutes(adminRouter, h)

}

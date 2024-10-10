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
	DB := db.Init()
	h := handler.New(DB)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "*"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()
	registerRoutes(router, h)
	han := c.Handler(router)

	pollingTicker := time.NewTicker(30 * time.Second)
	defer pollingTicker.Stop()

	go func() {
		log.Println("Listening on port 8080")
		err := http.ListenAndServe(":8080", han)
		utils.LogFatal(err, "Error starting server")
	}()

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
	routes.RegisterAuthRoutes(router, h)

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.JwtMiddleware)
	routes.RegisterUserRoutes(apiRouter, h)
	routes.RegisterWalletRoutes(apiRouter, h)

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AdminMiddleware)
	routes.RegisterUserAdminRoutes(adminRouter, h)
	routes.RegisterWalletAdminRoutes(adminRouter, h)

}

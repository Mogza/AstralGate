package main

import (
	"github.com/Mogza/AstralGate/internal/db"
	"github.com/Mogza/AstralGate/internal/handler"
	"github.com/Mogza/AstralGate/internal/routes"
	"github.com/Mogza/AstralGate/internal/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	DB := db.Init()
	h := handler.New(DB)

	//client, err := rpc.Dial("tcp", os.Getenv("POLYGON_AMOY_RPC"))
	//if err != nil {
	//	log.Fatalf("Failed to connect to the Polygon network: %v", err)
	//}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "*"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()
	registerRoutes(router, h)
	han := c.Handler(router)

	go func() {
		log.Println("Listening on port 8080")
		err := http.ListenAndServe(":8080", han)
		utils.LogFatal(err, "Error starting server")
	}()

	select {}
}

func registerRoutes(router *mux.Router, h handler.Handler) {
	routes.RegisterAuthRoutes(router, h)
}

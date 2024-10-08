package main

import (
	"github.com/Mogza/AstralGate/internal/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// DB := db.Init()
	// h := handler.New(DB)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "*"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()
	han := c.Handler(router)

	go func() {
		log.Println("Listening on port 8080")
		err := http.ListenAndServe(":8080", han)
		utils.LogFatal(err, "Error starting server")
	}()

	select {}
}

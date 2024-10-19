package main

import (
	"log"
	"net/http"

	"github.com/thongsoi/testd/database"
	"github.com/thongsoi/testd/internal/service"

	"github.com/thongsoi/testd/internal/repository"

	"github.com/thongsoi/testd/internal/handler"
)

func main() {
	// Initialize the database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer func() {
		if err := database.CloseDB(); err != nil {
			log.Println("Failed to close the database:", err)
		}
	}()
	r := mux.NewRouter()
	repo := &repository.Repository{}
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	r.HandleFunc("/", handler.OrderHandler).Methods("GET")
	r.HandleFunc("/markets", handler.GetMarkets).Methods("POST")
	r.HandleFunc("/submarkets", handler.GetSubmarkets).Methods("POST")
	r.HandleFunc("/submit-order", handler.SubmitOrder).Methods("POST")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

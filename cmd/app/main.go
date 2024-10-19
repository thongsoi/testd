package main

import (
	"log"
	"net/http"

	"github.com/thongsoi/testd/internal/service"

	"github.com/thongsoi/testd/internal/repository"

	"github.com/thongsoi/testd/internal/handler"
)

func main() {
	repo := &repository.Repository{}
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	http.HandleFunc("/markets", handler.GetMarkets)
	http.HandleFunc("/submarkets", handler.GetSubmarkets)
	http.HandleFunc("/submit-order", handler.SubmitOrder)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

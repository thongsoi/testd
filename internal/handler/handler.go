package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thongsoi/testd/internal/model"
	"github.com/thongsoi/testd/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetMarkets(w http.ResponseWriter, r *http.Request) {
	markets := h.service.GetMarkets()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(markets)
}

func (h *Handler) GetSubmarkets(w http.ResponseWriter, r *http.Request) {
	marketIDStr := r.URL.Query().Get("marketID")
	marketID, err := strconv.Atoi(marketIDStr)
	if err != nil {
		http.Error(w, "Invalid market ID", http.StatusBadRequest)
		return
	}

	submarkets := h.service.GetSubmarketsByMarketID(marketID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submarkets)
}

func (h *Handler) SubmitOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.CreateOrder(order)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

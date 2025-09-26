package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"ecommercefinal/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartHandler struct {
	service *services.CartService
}

func NewCartHandler(service *services.CartService) *CartHandler {
	return &CartHandler{service: service}
}

type AddCartRequest struct {
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var req AddCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		http.Error(w, "invalid product_id", http.StatusBadRequest)
		return
	}

	if err := h.service.AddToCart(context.Background(), userID, productID); err != nil {
		http.Error(w, "could not add to cart", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "added to cart"})
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userIDHex := mux.Vars(r)["userId"]
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	cart, err := h.service.GetCart(context.Background(), userID)
	if err != nil {
		http.Error(w, "could not fetch cart", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":  userID.Hex(),
		"products": cart,
	})
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userIDHex := mux.Vars(r)["userId"]
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	if err := h.service.ClearCart(context.Background(), userID); err != nil {
		http.Error(w, "could not clear cart", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "cart cleared"})
}

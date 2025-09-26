package handler

import (
	"context"
	"ecommercefinal/models"
	"ecommercefinal/services"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	User *services.UserService
}

func NewUserHandler(user *services.UserService) *UserHandler {
	return &UserHandler{User: user}
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	u.User.CreateUser(context.TODO(), user)
	json.NewEncoder(w).Encode(user)

}

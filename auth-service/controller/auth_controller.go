package controller

import (
	model2 "auth/model"
	service2 "auth/service"
	"encoding/json"
	"net/http"
)

type AuthController struct {
	authService *service2.AuthService
}

func NewAuthController(authService *service2.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	var user model2.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	if err := c.authService.Signup(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var user model2.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	token, err := c.authService.Login(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(model2.AuthResponse{Token: token})
}

package handlers

import (
	"encoding/json"
	"net/http"
)

type authHandler struct {
	httpclient *http.Client
}

func NewAuthHandler(client *http.Client) *authHandler {
	return &authHandler{
		httpclient: client,
	}
}
func (h *authHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", h.handleLogin)
	mux.HandleFunc("/logout", h.handleLogout)
	mux.HandleFunc("/register", h.handleRegister)
	mux.HandleFunc("/refresh", h.handleRefresh)
	mux.HandleFunc("/forgot-password", h.handleForgotPassword)
	mux.HandleFunc("/reset-password", h.handleResetPassword)
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *authHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	return
}

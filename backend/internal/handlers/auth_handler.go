package handlers

import (
	"net/http"
	"time"

	"github.com/schtvr/ttrpg/backend/pkg"
	"golang.org/x/crypto/bcrypt"
)

type authHandler struct {
	httpclient *http.Client
}

func NewAuthHandler(client *http.Client) *authHandler {
	return &authHandler{
		httpclient: client,
	}
}
func (ah *authHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", ah.handleRegister)
	mux.HandleFunc("/login", ah.handleLogin)
	// mux.HandleFunc("/logout", h.handleLogout)
}

type authRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ah *authHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var req authRequestBody
	if err := pkg.UnmarshalBody(r, &req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	UserRegistration(&User{
		Username:  req.Username,
		Password:  hashPassword(req.Password),
		CreatedAt: time.Now(),
	})
}

func (ah *authHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req authRequestBody
	if err := pkg.UnmarshalBody(r, &req); err != nil {
		http.Error(w, "invalid reqest", http.StatusBadRequest)
	}
	if err := UserLogin(&User{
		Username: req.Username,
		Password: hashPassword(req.Password),
	}); err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}
}

// Hash the password (e.g., during signup)
func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// Compare the password (e.g., during login)
func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

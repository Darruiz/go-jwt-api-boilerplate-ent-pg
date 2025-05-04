package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.UserID == "" || input.Password == "" {
		http.Error(w, "invalid credentials", http.StatusBadRequest)
		return
	}

	if input.Password != os.Getenv("STATIC_PASSWORD") {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	tokens, err := GenerateTokens(input.UserID)
	if err != nil {
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	var input RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.RefreshToken == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := ValidateToken(input.RefreshToken)
	if err != nil || !token.Valid {
		http.Error(w, "invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["typ"] != "refresh" {
		http.Error(w, "invalid refresh token format", http.StatusUnauthorized)
		return
	}

	userId, _ := claims["sub"].(string)
	tokens, err := GenerateTokens(userId)
	if err != nil {
		http.Error(w, "failed to refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)
}

package auth

import (
	"net/http"
	"os"

	"github.com/darruiz/dzfinance-go-api/internal/middleware"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/auth/login", LoginHandler)
	mux.HandleFunc("/auth/refresh", RefreshHandler)
	mux.Handle("/auth/me", middleware.AuthMiddleware([]byte(os.Getenv("JWT_SECRET")))(http.HandlerFunc(meHandler)))
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(string)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"user_id": "` + userId + `"}`))
}
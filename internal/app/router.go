package app

import (
	"net/http"

	"github.com/darruiz/dzfinance-go-api/internal/features/auth"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()
	auth.SetupRoutes(mux)
	return mux
}

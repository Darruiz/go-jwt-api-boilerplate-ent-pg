package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/darruiz/dzfinance-go-api/internal/app"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️ .env não carregado automaticamente, usando variáveis do container")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}
	log.Printf("🚀 API Go rodando em http://localhost:%s", port)
	http.ListenAndServe(":"+port, app.SetupRouter())
}
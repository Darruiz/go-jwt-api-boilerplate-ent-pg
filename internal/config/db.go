package config

import (
	"context"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectEnt() *ent.Client {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ erro conectando ao Ent: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("❌ erro criando schema Ent: %v", err)
	}

	return client
}

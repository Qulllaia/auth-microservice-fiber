package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CreateConnectionString() string {

	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, dbSSLMode)

	return  connStr;
	
}


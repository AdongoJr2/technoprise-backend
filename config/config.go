package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseSSLMode  string
	ServerPort       string
}

// LoadConfig loads configuration from environment variables or .env file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment variables directly.")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("DB_HOST environment variable not set.")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB_PORT environment variable not set.")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB_USER environment variable not set.")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("DB_PASSWORD environment variable not set.")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME environment variable not set.")
	}

	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable" // default value if not set
	}

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		log.Fatal("PORT environment variable not set.")
	}

	return &Config{
		DatabaseHost:     dbHost,
		DatabasePort:     dbPort,
		DatabaseUser:     dbUser,
		DatabasePassword: dbPassword,
		DatabaseName:     dbName,
		DatabaseSSLMode:  dbSSLMode,
		ServerPort:       serverPort,
	}
}

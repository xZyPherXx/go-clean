package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Name string
		Port string
		FL   string
		PK   string
	}
	Log struct {
		LogLevel string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
}

func LoadConfig() *Config {

	if err := godotenv.Load("/config/.env"); err != nil {

		panic(err)
	}

	config := Config{

		App: struct {
			Name string
			Port string
			FL   string
			PK   string
		}{

			Name: os.Getenv("APP_NAME"),
			Port: os.Getenv("APP_PORT"),
			FL:   os.Getenv("FULLCHAIN"),
			PK:   os.Getenv("PRIVKEY"),
		},

		Log: struct{ LogLevel string }{
			LogLevel: os.Getenv("LOG_LEVEL"),
		},

		Database: struct {
			Host     string
			Port     string
			User     string
			Password string
			DBName   string
			SSLMode  string
		}{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
	}

	log.Printf("Config loaded: %v", config)
	return &config
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN  string `env:"DSN" envDefault:"host=localhost user=postgres password=rahat12 dbname=tickets port=9920 sslmode=disable TimeZone=Asia/Shanghai"`
	PORT string `env:"PORT" envDefault:"8080"`
	JWTSecretKey string `env:"	JWT_SECRET_KEY" envDefault:"Sy9frXLrOngAQXcMiuF7yAfmNTUgziBH"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}

	return &Config{
		DSN:  os.Getenv("DSN"),
		PORT: os.Getenv("PORT"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}, nil
}
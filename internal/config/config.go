package config

import (
	"sim-clinic-api/internal/utils"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var getEnv = utils.GetEnv
var parseDuration = utils.ParseDuration

type Config struct {
	AppPort   string
	AppEnv    string
	JWTSecret string
	JWTExpire time.Duration

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBLogLevel string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found, using environment variables")
	}

	return &Config{
		AppPort:   getEnv("APP_PORT", "8080"),
		AppEnv:    getEnv("APP_ENV", "development"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		JWTExpire: parseDuration(getEnv("JWT_EXPIRE", "24h")),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "sim-clinic"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		DBLogLevel: getEnv("DB_LOG_LEVEL", "info"),
	}, nil
}

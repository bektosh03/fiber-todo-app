package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
	"os"
)

type Config struct {
	Environment          string // development, production
	PostgresHost         string
	PostgresPort         int
	PostgresDB           string
	PostgresUser         string
	PostgresPassword     string
	HTTPPort             string
	PathToCasbinConfFile string
	JWTSigningKey        string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return Config{
		Environment:          cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		PostgresHost:         cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:         cast.ToInt(getOrReturnDefault("POSTGRES_PORT", "5432")),
		PostgresDB:           cast.ToString(getOrReturnDefault("POSTGRES_DB", "tododb")),
		PostgresUser:         cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres")),
		PostgresPassword:     cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "12345")),
		HTTPPort:             cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080")),
		PathToCasbinConfFile: cast.ToString(getOrReturnDefault("PATH_TO_CASBIN_CONF_FILE", "config/rbac_model.conf")),
		JWTSigningKey:        cast.ToString(getOrReturnDefault("JWT_SIGNING_KEY", "ASDv12aeAE!@ba{df]6hB}")),
	}
}

func getOrReturnDefault(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}

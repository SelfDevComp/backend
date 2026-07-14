package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigLogger struct {
	Env string
}

type ConfigServer struct {
	Host         string
	Addr         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	Middleware   string
}

type ConfigDatabase struct {
	PostgresURL string
	RedisURL    string
}

type ConfigJWT struct {
	Secret string
}

type ConfigAuth struct {
	ClientId     string
	ClientSecret string
	RedirectURI  string
}

type config struct {
	Server   ConfigServer
	Logger   ConfigLogger
	Database ConfigDatabase
	JWT      ConfigJWT
	Auth     ConfigAuth
}

func NewConfig() (config, error) {
	if err := godotenv.Load(); err != nil {
		return config{}, err
	}

	readTimeout, _ := strconv.Atoi(getEnv("READ_TIMEOUT", "10"))
	writeTimeout, _ := strconv.Atoi(getEnv("WRITE_TIMEOUT", "10"))
	idleTimeout, _ := strconv.Atoi(getEnv("IDLE_TIMEOUT", "60"))

	cfg := config{
		Server: ConfigServer{
			Host:         getEnv("HOST", "localhost"),
			Addr:         getEnv("ADDR", ":8080"),
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
			Middleware:   getEnv("MIDDLEWARE", "http://localhost:80"),
		},
		Logger: ConfigLogger{
			Env: getEnv("ENV", "production"),
		},
		Database: ConfigDatabase{
			PostgresURL: getEnv("POSTGRES_URL", "postgres://admin:admin@db:5432/self-dev"),
			RedisURL: getEnv("REDIS_URL", "redis://:admin@redis:6379/0"),
		},
		JWT: ConfigJWT{
			Secret: getEnv("SECRET", ""),
		},
		Auth: ConfigAuth{
			ClientId:     getEnv("AUTH_CLIENT_ID", ""),
			ClientSecret: getEnv("AUTH_CLIENT_SECRET", ""),
		},
	}

	return cfg, nil
}

func getEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

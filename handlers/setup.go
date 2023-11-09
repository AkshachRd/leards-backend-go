package handlers

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

type EnvVars struct {
	AvatarBasePath string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func NewEnvVars() *EnvVars {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	envVars := EnvVars{}

	envVars.AvatarBasePath = getEnv("AVATAR_BASE_PATH", "./")

	return &envVars
}

type Server struct {
	DB      *gorm.DB
	EnvVars *EnvVars
}

func NewServer(db *gorm.DB) *Server {
	return &Server{DB: db, EnvVars: NewEnvVars()}
}

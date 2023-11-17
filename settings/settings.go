package settings

import (
	"github.com/joho/godotenv"
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

type Settings struct {
	EnvVars *EnvVars
}

var AppSettings = &Settings{}

func Setup() {
	AppSettings.EnvVars = NewEnvVars()
}

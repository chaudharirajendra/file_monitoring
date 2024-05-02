package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// EnvLoader defines the behavior for loading environment variables.
type EnvLoader interface {
	Load() error
}

// GodotEnvLoader implements the EnvLoader interface using the godotenv package.
type GodotEnvLoader struct{}

func (g *GodotEnvLoader) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	TargetDir   string
	StoragePath string
	Concurrency int
}

func NewConfig(loader EnvLoader) (*Config, error) {
	err := loader.Load()
	if err != nil {
		return nil, err
	}

	targetDir := getEnv("TARGET_DIR", "")
	storagePath := getEnv("STORAGE_PATH", "")
	concurrency := getIntEnv("CONCURRENCY", 5)

	return &Config{
		TargetDir:   targetDir,
		StoragePath: storagePath,
		Concurrency: concurrency,
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getIntEnv(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Error converting %s to int: %v\n", key, err)
		return defaultValue
	}
	return value
}

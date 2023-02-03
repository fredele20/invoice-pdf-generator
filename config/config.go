package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)


type UniDocConfig struct {
	Key string
}

func GetUniDocCred() (*UniDocConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	key := os.Getenv("UNICODE_LICENSE_API_KEY")
	if len(key) == 0 {
		return nil, errors.New("unidoc key not found")
	}

	return &UniDocConfig{Key: key}, nil
}
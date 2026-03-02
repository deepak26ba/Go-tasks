package config

import (
	"brevo/pkg/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config() (string, error) {

	var conn models.Connection

	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("Failed loading the environment : %v", err)
	}

	conn.ApiKey = os.Getenv("YOUR_API_KEY")

	return conn.ApiKey, nil
}

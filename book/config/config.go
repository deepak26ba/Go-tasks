package config

import (
	"books/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config() (string, error) {

	var con models.ConnectionString
	var connectionString string

	err := godotenv.Load()
	if err != nil {
		return connectionString, fmt.Errorf("Failed loading the environment : %v", err)
	}

	con.User = os.Getenv("DB_USER")
	con.DBName = os.Getenv("DB_NAME")
	con.Password = os.Getenv("DB_PASSWORD")
	con.SslMode = os.Getenv("DB_SSLMODE")
	con.Port = os.Getenv("DB_PORT")

	connectionString = fmt.Sprintf("user=%s port=%s dbname=%s password=%s sslmode=%s", con.User, con.Port, con.DBName, con.Password, con.SslMode)

	return connectionString, nil
}

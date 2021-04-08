package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var Controller *xorm.Engine

func GetConnectionString() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	var databaseHost string
	if databaseHost = os.Getenv("POSTGRES_HOST"); databaseHost == "" {
		return "", errors.New("failed to parse database host from environment")
	}

	var databasePort string
	if databasePort = os.Getenv("POSTGRES_PORT"); databasePort == "" {
		return "", errors.New("failed to parse database port from environment")
	}

	var userName string
	if userName = os.Getenv("POSTGRES_USER"); userName == "" {
		return "", errors.New("failed to parse database user from environment")
	}

	var userPassword string
	if userPassword = os.Getenv("POSTGRES_PASSWORD"); userPassword == "" {
		return "", errors.New("failed to parse database user password from environment")
	}

	var databaseName string
	if databaseName = os.Getenv("POSTGRES_DB"); databaseName == "" {
		return "", errors.New("failed to parse database name from environment")
	}

	var databaseSSL string
	if databaseSSL = os.Getenv("POSTGRES_SSL"); databaseSSL == "" {
		return "", errors.New("failed to parse ssl mode from environment")
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", databaseHost, databasePort, userName, userPassword, databaseName, databaseSSL), nil
}

func Connect() error {
	connection, err := GetConnectionString()
	if err != nil {
		return err
	}

	Controller, err = xorm.NewEngine("postgres", string(connection))
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	if err := Controller.Close(); err != nil {
		return err
	}
	return nil
}

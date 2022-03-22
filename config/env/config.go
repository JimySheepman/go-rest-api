package env

import (
	"errors"

	"github.com/joho/godotenv"
)

func LoadEnvironmentConfigure(path string) (bool, error) {
	err := godotenv.Load(path)
	if err != nil {
		return false, errors.New("Error loading .env file")
	}
	return true, nil
}

package config

import (
	"errors"
	"fmt"
	"os"
)

var errMissingEnvVariable error = errors.New("environment variable not found")

type Config struct {
	ServerPort string
}

func NewConfig() (Config, error) {
	serverPort, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		return Config{}, fmt.Errorf("%w: SERVER_PORT", errMissingEnvVariable)
	}

	return Config{
		ServerPort: serverPort,
	}, nil
}

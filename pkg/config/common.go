package config

import "errors"

var errMissingEnvVariable error = errors.New("environment variable not found")

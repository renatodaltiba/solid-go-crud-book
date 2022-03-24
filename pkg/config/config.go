package config

import (
	"os"

	"github.com/renatodalitba/books-solid-go-lang/pkg/errors"
	"github.com/sherifabdlnaby/configuro"
)

type Config struct {
	Server struct {
		HTTP *Server `validate:"required"`
	} `validate:"required"`
	Database *Database `validate:"required"`
}

type Database struct {
	Driver   string `validate:"required"`
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Name     string `validate:"required"`
}

type Server struct {
	Host string `validate:"required"`
}

func NewConfig(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.WrapErrorf(err, errors.ErrCodeNotFound, "config file not found")
	}

	loader, err := configuro.NewConfig(
		configuro.WithLoadFromConfigFile(configPath, false),
	)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := loader.Load(config); err != nil {
		return nil, err
	}

	if err := loader.Validate(config); err != nil {
		return nil, err
	}

	return config, nil
}

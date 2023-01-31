package config

import (
	"os"

	err "github.com/Alexander021192/finance_bot/internal/errors"
	"gopkg.in/yaml.v3"
)

const configFile = "data/config.yaml"

const (
	CannotReadConfig = err.ConstError("Cannot Read Config File")
	CannotUnmarshalConfig = err.ConstError("Cannot Unmarshal Config File")
)

type Config struct {
	Token string `yaml:"token"`
}

type Service struct {
	config Config
}

func New() (*Service, error) {
	s := &Service{}

	rawYAML, err := os.ReadFile(configFile)
	if err != nil {
		return nil, CannotReadConfig
	}

	err = yaml.Unmarshal(rawYAML, &s.config)
	if err != nil {
		return nil, CannotUnmarshalConfig
	}
	return s, nil
}

func(s *Service) Token() string {
	return s.config.Token
}
